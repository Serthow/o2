package asm

// Assembler represents a 65816 immediate assembler
type Assembler interface {
	FlagsTracker

	REP(c Flags)
	SEP(c Flags)
	NOP()
	JSR_abs(addr uint16)
	JSL(addr uint32)
	JSL_lhb(lo, hi, bank uint8)
	JML(addr uint32)
	RTS()
	RTL()
	LDA_imm8_b(m uint8)
	LDA_imm16_w(m uint16)
	LDA_imm16_lh(lo, hi uint8)
	STA_long(addr uint32)
}