package utils

import (
	"fmt"
	"testing"
)

func TestCuuid(t *testing.T) {
	s := GenString()
	fmt.Println(s)
}
