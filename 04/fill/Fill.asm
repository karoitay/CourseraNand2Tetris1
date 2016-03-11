// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, the
// program clears the screen, i.e. writes "white" in every pixel.

// Put your code here.

(CACHE)
  @prev_val
  M=0

(LOOP)
  @KBD
  D=M;
  @UNPRESSED
  D;JEQ

(PRESSED)
  @val
  M=-1
  @SHOULD_PAINT
  0;JMP

(UNPRESSED)
  @val
  M=0

(SHOULD_PAINT) 
  @val
  D=M
  @prev_val
  D=D+M
  @LOOP
  D;JEQ

(SET_PREV_VAL)
  @val
  D=M
  @prev_val
  M=D

(PAINT)
  @i
  M=0

  (PAINT_LOOP)
    @i
    D=M
    @8192 // screen size
    D=D-A
    @END
    D;JGE

    @SCREEN
    D=A
    @i
    D=D+M
    @screen_add
    M=D

    @val
    D=M
    @screen_add
    A=M
    M=D

    @i
    M=M+1

    @PAINT_LOOP
    0;JMP

(END)
  @LOOP
  0;JMP
