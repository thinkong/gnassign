package gnassign

import (
	"os"
	"bufio"
	"log"
)

type CheckList struct {
	List            []string // everything else
	SubDomain       []string // stuff like *.gracenote.com
	StartsWith      []string // stuff that starts with
	MatchEverything bool
}

func (cl CheckList) SearchMatch(url string) (bool) {
	if cl.MatchEverything {
		return true
	}
	for _, s := range cl.List {
		log.Println(s)
	}
	return true
}

func (cl *CheckList) AddString(str string) {
	if str[0] == '*' {
		if len(str) == 1 {
			cl.MatchEverything = true
			return
		}
		if str[1] == '.' {
			cl.SubDomain = append(cl.SubDomain, str)
		} else {
			cl.List = append(cl.List, str)
		}
	} else if str[len(str) - 1] == '/' {
		cl.StartsWith = append(cl.StartsWith, str)
	}
}

func (cl *CheckList) ReadConf(filename string) (error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			continue
		}
		cl.AddString(str)
	}

	err = scanner.Err()
	if err != nil {
		return err
	}
	return nil
}