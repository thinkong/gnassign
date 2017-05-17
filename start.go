package gnassign

import (
	"os"
	"bufio"
	"log"
	"net/url"
	"strings"
)

type CheckList struct {
	List            []string // everything else
	SubDomain       []string // stuff like *.gracenote.com
	StartsWith      []string // stuff that starts with
	MatchEverything bool
}

func (cl CheckList) SearchMatch(urlstring string) (bool) {
	if cl.MatchEverything {
		return true
	}
	u, uerr := url.ParseRequestURI(urlstring)
	if uerr != nil {
		//log.Println(uerr.Error())
		//log.Println("http:// missing")
		u, uerr = url.ParseRequestURI("http://" + urlstring)
		if uerr != nil {
			log.Println(uerr.Error())
			return false
		}
	}

	for _, s := range cl.List {
		matched, err := Match(s, urlstring)
		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		if matched {
			return true
		}
	}
	for _, s := range cl.SubDomain {
		if s != u.Hostname() && strings.Contains(u.Hostname(), s) {
			return true
		}
	}
	for _, s := range cl.StartsWith {
		if strings.HasPrefix(urlstring, s) {
			return true
		}
	}
	return false
}

func (cl *CheckList) AddString(str string) {
	if str[0] == '*' {
		if len(str) == 1 {
			cl.MatchEverything = true
			return
		}
		if str[1] == '.' {
			str = str[2:]
			log.Println(str)
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
	cnt := 0
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			continue
		}
		cl.AddString(str)
		cnt++
	}

	err = scanner.Err()
	if err != nil {
		return err
	}
	log.Println("loaded", cnt, "blacklists")
	return nil
}