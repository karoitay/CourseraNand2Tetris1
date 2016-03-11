// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

(OPTIMIZE) // copy RAM[0], RAM[1] to m, n such that m <= n
  @R0
  D=M
  @R1
  D=D-M
  @R0GR1
  D;JGE

  (R1GR0)
    @R0
    D=M
    @m
    M=D
    @R1
    D=M
    @n
    M=D
    @COMPUTE
    0;JMP
  (R0GR1)
    @R0
    D=M
    @n
    M=D
    @R1
    D=M
    @m
    M=D

(COMPUTE) // now RAM[0] >= RAM[1]
  @R2
  M=0
  @i
  M=0

  (LOOP)
    @i
    D=M
    @m
    D=M-D
    @END
    D;JEQ

    @n
    D=M
    @R2
    M=M+D

    @i
    M=M+1
    @LOOP
    0;JMP

(END)
  @END
  0;JMP