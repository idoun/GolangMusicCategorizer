package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var l = len(os.Args)

	if l > 1 {
		var target = getTargetDir(os.Args)
		fmt.Println(target)

		if len(target) == 0 {
			os.Exit(1)
		}

		var source = checkSlashSuffix(os.Args[1])

		f, err := os.Open(source)
		if err != nil {
			fmt.Println(err)
		}

		files, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			fmt.Println(err)
		}

		var errorCnt = 0
		for _, file := range files {
			var fileName = file.Name()
			rgx, _ := regexp.Compile("[^\\d*][^-]*")

			// pass dir
			if file.IsDir() {

			} else if strings.HasSuffix(strings.ToLower(fileName), "mp3") { // mp3 file
				var artist = strings.TrimSpace(rgx.FindString(fileName))
				// fmt.Println(artist)

				var newFilename = strings.TrimSpace(fileName[strings.Index(fileName, artist):])

				// create target dir
				_ = os.Mkdir(target+artist, os.ModeDir)

				// copy(or move) file with new name
				fmt.Println(source + fileName)
				fmt.Println(target + artist + "/" + newFilename)

				err = copyFile(source+fileName, target+artist+"/"+fileName)

				if err != nil {
					fmt.Println(err)
					errorCnt++
				}
			}
		}
		fmt.Printf("Errors: %v\n", errorCnt)
	} else {
		fmt.Println("No Args!")
	}
}

func getTargetDir(args []string) string {
	var l = len(os.Args)

	if l < 3 {
		// not target defined, choose parent!
		var base = filepath.Base(os.Args[1])
		return strings.TrimSuffix(os.Args[1], base)
	} else if len(os.Args[2]) > 0 { // not empty
		return checkSlashSuffix(os.Args[2])
	}

	return ""
}

func checkSlashSuffix(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
