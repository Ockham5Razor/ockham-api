package model

import (
	"fmt"
	"testing"
)

func TestV2RayConfJSON(t *testing.T) {
	v2RayConf := GenDefaultConfig()
	fmt.Println(v2RayConf.AsJSON())
}
