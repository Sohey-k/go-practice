package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	var useColor bool
	flag.BoolVar(&useColor, "c", false, "Use color output")
	flag.Parse()

	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if useColor {
			if file.IsDir() {
				color.Yellow(file.Name())
			} else {
				fmt.Println(file.Name())
			}
		} else {
			fmt.Println(file.Name())
		}
	}
}
