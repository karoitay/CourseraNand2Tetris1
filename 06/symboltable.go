package main

import (
	"fmt"
)

type SymbolTable map[string]int

func NewSymbolTable() *SymbolTable {
	symbols := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24567,
	}
	st := SymbolTable(symbols)
	return &st
}

func (st *SymbolTable) asMap() *map[string]int {
	m := map[string]int(*st)
	return &m
}

func (st *SymbolTable) Get(key string) (val int, ok bool) {
	val, ok = (*st.asMap())[key]
	return
}

func (st *SymbolTable) Put(key string, address int) {
	fmt.Printf("Put %s: %v\n", key, address)
	(*st.asMap())[key] = address
}
