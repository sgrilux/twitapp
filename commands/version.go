package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	TwitApp.Version = func() {
		content, err := ioutil.ReadFile("VERSION.txt")
		if err != nil {
			log.Fatal(err)
		}

		ver := string(content)
		fmt.Println(ver)
		os.Exit(0)
	}
}
