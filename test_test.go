package test

import (
	"testing"
)

func TestF(t*testing.T){
	c:=f(1,2)
	if c != 3 {
		t.Error("fail")
	}
}
