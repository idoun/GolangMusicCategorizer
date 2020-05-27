package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var l = len(os.Args)

	if l > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}

		files, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {
			var fileName = file.Name()
			if file.IsDir() {

			} else if strings.HasSuffix(strings.ToLower(fileName), "mp3") {
				fmt.Println("found:" + fileName)
			}
		}
	} else {
		fmt.Println("No Args!")
	}
}
