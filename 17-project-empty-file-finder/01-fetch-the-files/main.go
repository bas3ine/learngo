// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}
}
