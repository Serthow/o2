package emulator

import (
	"testing"
)

func TestSystem_CreateEmulator(t *testing.T) {
	// verify our ROM, SRAM, WRAM mappings in the bus:
	tests := []struct {
		name   string
		verify func(t *testing.T, q *System)
	}{
		// ROM:
		{
			name: "ROM bank 00",
			verify: func(t *testing.T, q *System) {
				q.ROM[0x0000] = 0xFE
				if actual, expected := q.Bus.EaRead(0x00_8000), uint8(0xFE); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
				if actual, expected := q.Bus.EaRead(0x80_8000), uint8(0xFE); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "ROM bank 01",
			verify: func(t *testing.T, q *System) {
				q.ROM[0x8000] = 0xFD
				if actual, expected := q.Bus.EaRead(0x01_8000), uint8(0xFD); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
				if actual, expected := q.Bus.EaRead(0x81_8000), uint8(0xFD); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		// SRAM:
		{
			name: "SRAM bank 00",
			verify: func(t *testing.T, q *System) {
				q.SRAM[0x0000] = 0xFC
				if actual, expected := q.Bus.EaRead(0x70_0000), uint8(0xFC); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "SRAM bank 01",
			verify: func(t *testing.T, q *System) {
				q.SRAM[0x8000] = 0xFB
				if actual, expected := q.Bus.EaRead(0x71_0000), uint8(0xFB); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		// WRAM:
		{
			name: "WRAM $7E:0000",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x0000] = 0xFA
				if actual, expected := q.Bus.EaRead(0x7E_0000), uint8(0xFA); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $00:0000",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x0000] = 0xF9
				if actual, expected := q.Bus.EaRead(0x00_0000), uint8(0xF9); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $7E:1FFF",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x1FFF] = 0xF8
				if actual, expected := q.Bus.EaRead(0x7E_1FFF), uint8(0xF8); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $00:1FFF",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x1FFF] = 0xF7
				if actual, expected := q.Bus.EaRead(0x00_1FFF), uint8(0xF7); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $7E:2000",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x2000] = 0xF6
				if actual, expected := q.Bus.EaRead(0x7E_2000), uint8(0xF6); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $7F:2000",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x12000] = 0xF5
				if actual, expected := q.Bus.EaRead(0x7F_2000), uint8(0xF5); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
		{
			name: "WRAM $7F:FFFF",
			verify: func(t *testing.T, q *System) {
				q.WRAM[0x1FFFF] = 0xF4
				if actual, expected := q.Bus.EaRead(0x7F_FFFF), uint8(0xF4); actual != expected {
					t.Errorf("mapping failed, actual = %v, expected = %v", actual, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &System{}
			if err := q.CreateEmulator(); err != nil {
				t.Fatal(err)
			}
			tt.verify(t, q)
		})
	}
}

func TestSystem_SetupPatch(t *testing.T) {
	// verify our ROM, SRAM, WRAM mappings in the bus:
	tests := []struct {
		name   string
		verify func(t *testing.T, q *System)
	}{
		{
			name: "Verify entry",
			verify: func(t *testing.T, q *System) {
				if actual, expected := q.Bus.EaRead(0x00_8000), uint8(0xE2); actual != expected {
					t.Errorf("mapping failed, actual = %02x, expected = %02x", actual, expected)
				}
			},
		},
		{
			name: "Step",
			verify: func(t *testing.T, q *System) {
				q.CPU.Reset()
				cycles, aborted := q.CPU.Step()
				if actual, expected := cycles, 3; actual != expected {
					t.Errorf("Step() failed, cycles actual = %v, expected = %v", actual, expected)
				}
				if actual, expected := aborted, false; actual != expected {
					t.Errorf("Step() failed, cycles actual = %v, expected = %v", actual, expected)
				}
				if actual, expected := q.CPU.PC, uint16(0x8002); actual != expected {
					t.Errorf("Step() failed, PC actual = %04x, expected = %04x", actual, expected)
				}
			},
		},
		{
			name: "Run",
			verify: func(t *testing.T, q *System) {
				q.CPU.Reset()
				for q.CPU.PC != 0x8006 {
					_, _ = q.CPU.Step()
				}
				if actual, expected := q.CPU.PC, uint16(0x8006); actual != expected {
					t.Errorf("Step() failed, PC actual = %04x, expected = %04x", actual, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &System{}
			if err := q.CreateEmulator(); err != nil {
				t.Fatal(err)
			}
			rom, err := MakeTestROM("TEST")
			if err != nil {
				t.Fatal(err)
			}
			copy(q.ROM[:], rom.Contents)
			if err := q.SetupPatch(); err != nil {
				t.Fatal(err)
			}
			tt.verify(t, q)
		})
	}
}
