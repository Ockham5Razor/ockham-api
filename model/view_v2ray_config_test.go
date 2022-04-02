package model

import (
	"fmt"
	"testing"
)

func TestV2RayConfJSON(t *testing.T) {
	v2RayConf := GenConfig(10086, 8080, "/some-path/")
	fmt.Println(v2RayConf.AsJSON())
}
