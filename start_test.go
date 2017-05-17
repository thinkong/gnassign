package gnassign

import "testing"

var cl CheckList

func initializeTest() {
	cl = CheckList{}
	cl.List = []string{"*/img/*.jpg",
		"*.gracenote.com",
		"gracenote.com/example/"}
}

func TestCheckList_ReadConf(t *testing.T) {
	initializeTest()
	// Nothing to test...
}