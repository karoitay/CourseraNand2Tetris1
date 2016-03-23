package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: " + os.Args[0] + " <filename>\n")
	} else {
		compile(os.Args[1])
	}
}

func compile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	outfile := strings.Replace(filename, ".asm", ".hack", -1)
	of, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	fp := NewFileParser(f)
	for inst := fp.nextLine(); inst != nil; inst = fp.nextLine() {
		if _, err = of.WriteString(inst.GetCode() + "\n"); err != nil {
			panic(err)
		}
	}
}
