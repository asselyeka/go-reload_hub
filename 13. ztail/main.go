package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	size := 0
	lenArg := 0
	index := -1

	for i, input := range arguments {
		if input == "-c" {
			index = i
		}
		lenArg++
	}

	if index != -1 {
		if index+1 < lenArg {
			if student.CheckStr(arguments[index+1]) == true {
				size = student.Atoi(arguments[index+1])
			} else {
				fmt.Printf("tail: invalid number of bytes: '%s'\n", arguments[index+1])
				return
			}
		} else {
			fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
		}
	}

	var filesSlice []string
	if index == 0 {
		filesSlice = []string(arguments[2:])
	} else {
		filesSlice = []string(arguments[0:index])
	}

	filesLen := 0
	for range filesSlice {
		filesLen++
	}

	for _, filesN := range filesSlice {

		file, err := os.Open(filesN)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", filesN)
			os.Exit(0)
		}
		defer file.Close()

		if filesLen > 1 {
			fmt.Printf("==> %v <==\n", filesN)
		}

		fileStat, _ := os.Stat(filesN)
		sizeF := fileStat.Size()
		var content []byte
		for i := 0; i < int(sizeF); i++ {
			content = append(content, 0)
		}
		file.Read(content)

		lenF := 0
		for range content {
			lenF++
		}
		if size <= lenF {
			content = content[lenF-size:]
		}
		fmt.Printf("%s", content)
	}
}
