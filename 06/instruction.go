package main

import (
	"fmt"
	"strings"
)

var destinations map[string]string = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var computations map[string]string = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}

var jumps map[string]string = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type Instruction interface {
	GetCode() string
}

type AInstruction struct {
	Source  string
	Value   string
	Address int
}

type CInstruction struct {
	Source string
	Dest   string
	Comp   string
	Jump   string
}

type Label struct {
	Name string
}

func (inst AInstruction) GetCode() string {
	return fmt.Sprintf("%016b", inst.Address)
}

func (inst CInstruction) GetCode() string {
	var code string
	if strings.Index(inst.Comp, "M") > -1 {
		code = "1111" + computations[strings.Replace(inst.Comp, "M", "A", -1)]
	} else {
		code = "1110" + computations[inst.Comp]
	}
	return code + destinations[inst.Dest] + jumps[inst.Jump]
}

func (lbl Label) GetCode() string {
	return ""
}
