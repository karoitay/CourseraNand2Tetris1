package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: " + os.Args[0] + " <filename1> [<filename2> ...]\n")
	} else {
		for i := 1; i < len(os.Args); i++ {
			compile(os.Args[i])
		}
	}
}

func compile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	symbols := NewSymbolTable()

	FirstPass{}.Compile(NewFileParser(f), symbols)

	if _, err := f.Seek(0, 0); err != nil {
		panic(err)
	}

	SecondPass{strings.Replace(filename, ".asm", ".hack", -1)}.Compile(NewFileParser(f), symbols)
}
