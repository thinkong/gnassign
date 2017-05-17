package gnassign

import "testing"

func TestMatch(t *testing.T) {
	matched, err := Match("*.img", "hello/gracenote/1.img")
	if err != nil {
		t.Error(err.Error())
	}
	if !matched {
		t.Error("Expected Match1")
	}
	matched, err = Match("*/img/*.jpg", "www.gracenote.com/img/1.jpg")
	if err != nil {
		t.Error(err.Error())
	}
	if !matched {
		t.Error("Expected Match2")
	}
}