package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SignatureHashType string

const (
	HmacSha256 = "HMAC-SHA256"
)

type Signature struct {
	HashType   SignatureHashType
	Credential struct {
		AccessKeyID string
		Action      string
	}
	SignedHeaders []string
	Signature     string
}

var defaultSignedHeaders = []string{HeaderContentType, HeaderHost, HeaderTimestamp}

func (s *Signature) String() string {
	return fmt.Sprintf(
		"%v Credential=%v/%v, SignedHeaders=%v, Signature=%v",
		HmacSha256,
		s.Credential.AccessKeyID,
		s.Credential.Action,
		strings.Join(s.SignedHeaders, ";"),
		s.Signature,
	)
}

func Compare(sig1, sig2 string) bool {
	return subtle.ConstantTimeCompare([]byte(sig1), []byte(sig2)) == 1
}

func DecodeSignature(signatureBody string) (err error, signature Signature) {
	signatureBodySplitType := strings.SplitN(signatureBody, " ", 2)
	if len(signatureBodySplitType) < 2 {
		return errors.New("signature illegal"), signature
	}
	hashTypeStr := signatureBodySplitType[0]
	hashType := SignatureHashType(hashTypeStr)

	content := signatureBodySplitType[1]
	realContent := strings.Replace(content, " ", "", -1)
	contentSplit := strings.Split(realContent, ",")

	contentMap := make(map[string]string)
	for _, contentPart := range contentSplit {
		contentPartSplit := strings.Split(contentPart, "=")
		if len(contentPartSplit) != 2 {
			return errors.New("signature illegal"), signature
		}
		contentMap[contentPartSplit[0]] = contentPartSplit[1]
	}

	// Credential
	credentialSplit := strings.Split(contentMap["Credential"], "/")
	if len(credentialSplit) != 2 {
		return errors.New("signature illegal"), signature
	}
	accessKeyID := credentialSplit[0]
	action := credentialSplit[1]

	signature = Signature{
		HashType: hashType,
		Credential: struct {
			AccessKeyID string
			Action      string
		}{AccessKeyID: accessKeyID, Action: action},
		SignedHeaders: strings.Split(contentMap["SignedHeaders"], ";"),
		Signature:     contentMap["Signature"],
	}
	return nil, signature
}

func CreateSignature(req *http.Request, accessKeyID, secretAccessKey, actionType string, signedHeaders ...string) (Signature, error) {
	// Request Body 摘要
	var bodyDigestBytes []byte
	if req.Body != nil {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return Signature{}, err
		}
		// 限位，如果没有限位动作，那么会被填充 0
		digestPickEnd := 0
		if len(bodyBytes) > 0 && len(bodyBytes) < 128 {
			digestPickEnd = len(bodyBytes)
		} else if len(bodyBytes) >= 128 {
			digestPickEnd = 128
		}
		bodyDigestBytes = bodyBytes[0:digestPickEnd]
	}

	// 如果没有传入，则使用默认头
	if len(signedHeaders) == 0 {
		signedHeaders = defaultSignedHeaders
	}
	// 按位排序
	var signedHeaderValues []string
	for _, header := range signedHeaders {
		var headerVal string
		if header == HeaderHost {
			headerVal = req.Host
		} else {
			headerVal = req.Header.Get(header)
		}
		signedHeaderValues = append(signedHeaderValues, headerVal)
	}
	// 拼接成串
	unsignedHeaderValStr := strings.Join(signedHeaderValues, ";")
	var zeroSpace = append([]byte{}, 0)
	key := []byte(secretAccessKey)
	m := hmac.New(sha256.New, key)
	m.Write([]byte(unsignedHeaderValStr)) // 写入头
	m.Write(zeroSpace)                    // 写入零值分隔符
	m.Write(bodyDigestBytes)              // 写入摘要
	signature := hex.EncodeToString(m.Sum(nil))

	return Signature{
		HashType: HmacSha256,
		Credential: struct {
			AccessKeyID string
			Action      string
		}{AccessKeyID: accessKeyID, Action: actionType},
		SignedHeaders: signedHeaders,
		Signature:     signature,
	}, nil
}
