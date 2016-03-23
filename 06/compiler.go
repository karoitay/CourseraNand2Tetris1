package main

import (
	"os"
	"strconv"
)

type Compiler interface {
	Compile(fp *FileParser, symbols *SymbolTable)
}

type FirstPass struct{}

func (pass FirstPass) Compile(fp *FileParser, symbols *SymbolTable) {
	nextAddr := 0
	for inst := fp.nextInstruction(); inst != nil; inst = fp.nextInstruction() {
		switch lbl := inst.(type) {
		case Label:
			symbols.Put(lbl.Name, nextAddr)
		default:
			nextAddr = nextAddr + 1
			break
		}
	}
}

type SecondPass struct {
	OutFileName string
}

func (pass SecondPass) Compile(fp *FileParser, symbols *SymbolTable) {
	f, err := os.Create(pass.OutFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nextVarAddress := 16
	for inst := fp.nextInstruction(); inst != nil; inst = fp.nextInstruction() {
		var code string
		switch i := inst.(type) {
		case AInstruction:
			if address, err := strconv.Atoi(i.Value); err == nil {
				i.Address = address
			} else {
				if address, found := symbols.Get(i.Value); !found {
					i.Address = nextVarAddress
					symbols.Put(i.Value, nextVarAddress)
					nextVarAddress = nextVarAddress + 1
				} else {
					i.Address = address
				}
			}
			code = i.GetCode()
			break
		default:
			code = i.GetCode()
			break
		}

		if code != "" {
			if _, err = f.WriteString(code + "\n"); err != nil {
				panic(err)
			}
		}
	}
}
