package gnassign

import "testing"

var cl CheckList

func initializeTest() {
	cl = CheckList{}
	cl.List = []string{"*/img/*.jpg"}
	cl.SubDomain = []string{"*.gracenote.com"}
	cl.StartsWith = []string{"gracenote.com/example/"}
}

func TestCheckList_ReadConf(t *testing.T) {
	initializeTest()
	// Nothing to test...
}

func TestCheckList_AddString(t *testing.T) {
	cList := CheckList{}
	cList.AddString("*/img/*.jpg")
	cList.AddString("*.gracenote.com")
	cList.AddString("gracenote.com/example/")

	if len(cList.StartsWith) != 1 || len(cList.SubDomain) != 1 || len(cList.List) != 1 {
		t.Error("check wrong")
	}
}