package snes

import (
	"encoding/binary"
	"encoding/hex"
	"testing"
)

func sampleROM() []byte {
	contents := make([]byte, 0x8000)
	hex.Decode(
		contents[0x7FB0:],
		[]byte("018d2401e2306bffffffffffffffffff"+
			"544845204c4547454e44204f46205a45"+
			"4c4441202020020a03010100f2500daf"+
			"ffffffff2c82ffff2c82c9800080d882"+
			"ffffffff2c822c822c822c820080d882"),
	)
	return contents
}

func TestNewROM(t *testing.T) {
	contents := sampleROM()
	gotR, err := NewROM(contents)
	if err != nil {
		t.Fatal(err)
	}

	// check:
	if gotR.Header.MakerCode != 0x8D01 {
		t.Fatal("MakerCode")
	}
	if gotR.Header.GameCode != 0x30E20124 {
		t.Fatal("GameCode")
	}
	if gotR.NativeVectors.NMI != 0x80c9 {
		t.Fatal("NativeVectors.NMI")
	}
}

func TestROM_BusReader(t *testing.T) {
	contents := sampleROM()
	rom, err := NewROM(contents)
	if err != nil {
		t.Fatal(err)
	}

	r := rom.BusReader(0x00FFEA)
	p := uint16(0)
	err = binary.Read(r, binary.LittleEndian, &p)
	if err != nil {
		t.Fatal(err)
	}
	if p != 0x80c9 {
		t.Fatal("expected NMI vector at $FFEA")
	}
}
