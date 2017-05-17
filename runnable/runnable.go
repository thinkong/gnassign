package main

import (
	"github.com/thinkong/gnassign"
	"log"
)

func main() {
	blackList := gnassign.CheckList{}
	err := blackList.ReadConf("blacklist.conf")

	if err != nil {
		log.Fatalln(err)
	}
}