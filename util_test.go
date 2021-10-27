package main

import (
	"fmt"
	"gol-c/util"
	"testing"
)

func TestCuuid(t *testing.T) {
	s := util.GenString()
	fmt.Println(s)
}

func TestJwt(t *testing.T)  {
	fmt.Println(util.GenToken("dave.smith", "idXsw2fasaxalrgpc"))
}