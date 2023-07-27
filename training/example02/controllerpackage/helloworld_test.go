package controllerpackage

import (
	"controllerpackage"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("appleboy")
	if hello != "Hi, appleboy" {
		t.Errorf("Testing Failed")
	}
	hello = HelloWorld("appleboy ")
	if hello != "Hi, appleboy" {
		t.Errorf("Testing fail")
	}
}
