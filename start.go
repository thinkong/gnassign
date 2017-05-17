package gnassign

import (
	"os"
	"bufio"
	"log"
)

type CheckList struct {
	List []string
}

func (cl *CheckList) ReadConf(filename string) (error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cl.List = append(cl.List, scanner.Text())
		log.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return err
	}
	return nil
}