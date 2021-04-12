package asm

import (
	"bytes"
	"fmt"
	"strings"
)

// Emitter implements Assembler and bytes.Buffer; a 65816 immediate assembler that emits to the buffer
type Emitter struct {
	flagsTracker

	Code *bytes.Buffer
	Text *strings.Builder

	address uint32
	baseSet bool
}

func (a *Emitter) Clone() *Emitter {
	return &Emitter{
		flagsTracker: a.flagsTracker,
		Code:         &bytes.Buffer{},
		Text:         &strings.Builder{},
		address:      a.address,
		baseSet:      a.baseSet,
	}
}

func (a *Emitter) Append(e *Emitter) {
	a.address = e.address
	a.baseSet = e.baseSet
	a.flagsTracker = e.flagsTracker

	_, _ = e.Code.WriteTo(a.Code)
	_, _ = a.Text.WriteString(e.Text.String())
}

func (a *Emitter) SetBase(addr uint32) {
	a.address = addr
	a.baseSet = true
}

func (a *Emitter) GetBase() uint32 {
	return a.address
}

func (a *Emitter) emitBase() {
	if !a.baseSet {
		return
	}

	_, _ = a.Text.WriteString(fmt.Sprintf("base $%06x\n", a.address))
	a.baseSet = false
}

func (a *Emitter) emit1(ins string, d [1]byte) {
	if a.Code != nil {
		_, _ = a.Code.Write(d[:])
	}
	if a.Text != nil {
		a.emitBase()
		// TODO: adjust these format widths
		_, _ = a.Text.WriteString(fmt.Sprintf("    %-5s %-8s ; $%06x  %02x\n", ins, "", a.address, d[0]))
	}
	a.address += 1
}

func (a *Emitter) emit2(ins, argsFormat string, d [2]byte) {
	if a.Code != nil {
		_, _ = a.Code.Write(d[:])
	}
	if a.Text != nil {
		a.emitBase()
		args := fmt.Sprintf(argsFormat, d[1])
		// TODO: adjust these format widths
		_, _ = a.Text.WriteString(fmt.Sprintf("    %-5s %-8s ; $%06x  %02x %02x\n", ins, args, a.address, d[0], d[1]))
	}
	a.address += 2
}

func (a *Emitter) emit3(ins, argsFormat string, d [3]byte) {
	if a.Code != nil {
		_, _ = a.Code.Write(d[:])
	}
	if a.Text != nil {
		a.emitBase()
		args := fmt.Sprintf(argsFormat, d[1], d[2])
		// TODO: adjust these format widths
		_, _ = a.Text.WriteString(fmt.Sprintf("    %-5s %-8s ; $%06x  %02x %02x %02x\n", ins, args, a.address, d[0], d[1], d[2]))
	}
	a.address += 3
}

func (a *Emitter) emit4(ins, argsFormat string, d [4]byte) {
	if a.Code != nil {
		_, _ = a.Code.Write(d[:])
	}
	if a.Text != nil {
		a.emitBase()
		args := fmt.Sprintf(argsFormat, d[1], d[2], d[3])
		// TODO: adjust these format widths
		_, _ = a.Text.WriteString(fmt.Sprintf("    %-5s %-8s ; $%06x  %02x %02x %02x %02x\n", ins, args, a.address, d[0], d[1], d[2], d[3]))
	}
	a.address += 4
}

func imm24(v uint32) (byte, byte, byte) {
	return byte(v), byte(v >> 8), byte(v >> 16)
}

func imm16(v uint16) (byte, byte) {
	return byte(v), byte(v >> 8)
}

func (a *Emitter) Comment(s string) {
	if a.Text != nil {
		_, _ = a.Text.WriteString(fmt.Sprintf("    ; %s\n", s))
	}
}

func (a *Emitter) REP(c Flags) {
	a.AssumeREP(c)
	a.emit2("rep", "#$%02x", [2]byte{0xC2, byte(c)})
}

func (a *Emitter) SEP(c Flags) {
	a.AssumeSEP(c)
	a.emit2("sep", "#$%02x", [2]byte{0xE2, byte(c)})
}

func (a *Emitter) NOP() {
	a.emit1("nop", [1]byte{0xEA})
}

func (a *Emitter) JSR_abs(addr uint16) {
	var d [3]byte
	d[0] = 0x20
	d[1], d[2] = imm16(addr)
	a.emit3("jsr", "$%02[2]x%02[1]x", d)
}

func (a *Emitter) JSL(addr uint32) {
	var d [4]byte
	d[0] = 0x22
	d[1], d[2], d[3] = imm24(addr)
	a.emit4("jsl", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) JSL_lhb(lo, hi, bank uint8) {
	var d [4]byte
	d[0] = 0x22
	d[1], d[2], d[3] = lo, hi, bank
	a.emit4("jsl", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) JML(addr uint32) {
	var d [4]byte
	d[0] = 0x5C
	d[1], d[2], d[3] = imm24(addr)
	a.emit4("jml", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) RTS() {
	a.emit1("rts", [1]byte{0x60})
}

func (a *Emitter) RTL() {
	a.emit1("rtl", [1]byte{0x6B})
}

func (a *Emitter) LDA_imm8_b(m uint8) {
	if a.IsM16bit() {
		panic(fmt.Errorf("asm: LDA_imm8_b called but 'm' flag is 16-bit; call SEP(0x20) or AssumeSEP(0x20) first"))
	}
	var d [2]byte
	d[0] = 0xA9
	d[1] = m
	a.emit2("lda.b", "#$%02x", d)
}

func (a *Emitter) LDA_imm16_w(m uint16) {
	if !a.IsM16bit() {
		panic(fmt.Errorf("asm: LDA_imm16_w called but 'm' flag is 8-bit; call REP(0x20) or AssumeREP(0x20) first"))
	}
	var d [3]byte
	d[0] = 0xA9
	d[1], d[2] = imm16(m)
	a.emit3("lda.w", "#$%02[2]x%02[1]x", d)
}

func (a *Emitter) LDA_imm16_lh(lo, hi uint8) {
	if !a.IsM16bit() {
		panic(fmt.Errorf("asm: LDA_imm16_lh called but 'm' flag is 8-bit; call REP(0x20) or AssumeREP(0x20) first"))
	}
	var d [3]byte
	d[0] = 0xA9
	d[1], d[2] = lo, hi
	a.emit3("lda.w", "#$%02[2]x%02[1]x", d)
}

func (a *Emitter) LDA_long(addr uint32) {
	var d [4]byte
	d[0] = 0xAF
	d[1], d[2], d[3] = imm24(addr)
	a.emit4("lda.l", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) STA_long(addr uint32) {
	var d [4]byte
	d[0] = 0x8F
	d[1], d[2], d[3] = imm24(addr)
	a.emit4("sta.l", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) STA_abs(addr uint16) {
	var d [3]byte
	d[0] = 0x8D
	d[1], d[2] = imm16(addr)
	a.emit3("sta.w", "$%02[2]x%02[1]x", d)
}

func (a *Emitter) STA_dp(addr uint8) {
	var d [2]byte
	d[0] = 0x85
	d[1] = addr
	a.emit2("sta.b", "$%02[1]x", d)
}

func (a *Emitter) ORA_long(addr uint32) {
	var d [4]byte
	d[0] = 0x0F
	d[1], d[2], d[3] = imm24(addr)
	a.emit4("ora.l", "$%02[3]x%02[2]x%02[1]x", d)
}

func (a *Emitter) CMP_imm8_b(m uint8) {
	if a.IsM16bit() {
		panic(fmt.Errorf("asm: CMP_imm8_b called but 'm' flag is 16-bit; call SEP(0x20) or AssumeSEP(0x20) first"))
	}
	var d [2]byte
	d[0] = 0xC9
	d[1] = m
	a.emit2("cmp.b", "#$%02x", d)
}

func (a *Emitter) BNE(m int8) {
	var d [2]byte
	d[0] = 0xD0
	d[1] = uint8(m)
	a.emit2("bne", "$%02x", d)
}

func (a *Emitter) ADC_imm8_b(m uint8) {
	if a.IsM16bit() {
		panic(fmt.Errorf("asm: ADC_imm8_b called but 'm' flag is 16-bit; call SEP(0x20) or AssumeSEP(0x20) first"))
	}
	var d [2]byte
	d[0] = 0x69
	d[1] = m
	a.emit2("adc.b", "#$%02x", d)
}
