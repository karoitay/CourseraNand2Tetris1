package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type FileParser struct {
	scanner *bufio.Scanner
}

func NewFileParser(f *os.File) *FileParser {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	return &FileParser{scanner}
}

func (fp FileParser) nextLine() Instruction {
	for fp.scanner.Scan() {
		text := strings.TrimSpace(fp.scanner.Text())
		commentIndex := strings.Index(text, "//")
		if commentIndex >= 0 {
			text = strings.TrimSpace(text[0:commentIndex])
		}
		if len(text) > 0 {
			return parseInstruction(text)
		}
	}
	return nil
}

func parseInstruction(text string) Instruction {
	if strings.HasPrefix(text, "@") {
		if address, err := strconv.Atoi(text[1:]); err != nil {
			panic(err)
		} else {
			return AInstruction{address}
		}
	}
	if strings.HasPrefix(text, "(") {
		panic("Need to handle labels: " + text)
	}
	inst := CInstruction{}
	split := strings.Split(text, ";")
	if len(split) == 2 {
		inst.Jump = split[1]
	}
	split = strings.Split(split[0], "=")
	if len(split) == 1 {
		inst.Comp = split[0]
	} else {
		inst.Dest = split[0]
		inst.Comp = split[1]
	}
	return inst
}
