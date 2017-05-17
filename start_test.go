package gnassign

import "testing"

var cl CheckList

func init() {
	cl = CheckList{}
	cl.AddString("*/img/*.jpg")
	cl.AddString("*.gracenote.com")
	cl.AddString("gracenote.com/example/")
}

func TestCheckList_ReadConf(t *testing.T) {
	//initializeTest()
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
	cList.AddString("*")
	if !cList.MatchEverything {
		t.Error("matcheverything must be true")
	}
}

func TestCheckList_SearchMatch(t *testing.T) {
	//initializeTest()
	matched := cl.SearchMatch("www.gracenote.com/img/1.jpg")
	if !matched {
		t.Error("supposed to match1")
	}
	matched = cl.SearchMatch("gracenote.com/abc.png")
	if matched {
		t.Error("supposed to not match")
	}
	matched = cl.SearchMatch("gracenote.com/example/abc.png")
	if !matched {
		t.Error("supposed to match")
	}
}

func BenchmarkCheckList_SearchMatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matched := cl.SearchMatch("www.gracenote.com/img/1.jpg")
		if !matched {
			b.Error("supposed to match1")
		}
	}
}

func BenchmarkCheckList_SearchMatch2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matched := cl.SearchMatch("gracenote.com/abc.png")
		if matched {
			b.Error("supposed to not match")
		}
	}
}