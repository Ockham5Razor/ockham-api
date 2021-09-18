package v1

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	s := randomCode()
	fmt.Println(s)
}
