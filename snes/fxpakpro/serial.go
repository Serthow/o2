package fxpakpro

import (
	"fmt"
	"go.bug.st/serial"
	"o2/snes"
)

func sendSerial(f serial.Port, buf []byte) error {
	sent := 0
	for sent < len(buf) {
		n, e := f.Write(buf[sent:])
		if e != nil {
			return e
		}
		sent += n
	}
	return nil
}

func recvSerial(f serial.Port, rsp []byte, expected int) error {
	o := 0
	for o < expected {
		n, err := f.Read(rsp[o:expected])
		if err != nil {
			return err
		}
		if n <= 0 {
			return fmt.Errorf("recvSerial: Read returned %d", n)
		}
		o += n
	}
	return nil
}

func (c *Conn) sendVPUT(reqs []snes.WriteRequest) (int, error) {
	// VGET:
	sb := make([]byte, 64)
	sb[0] = byte('U')
	sb[1] = byte('S')
	sb[2] = byte('B')
	sb[3] = byte('A')
	sb[4] = byte(OpVPUT)
	sb[5] = byte(SpaceSNES)
	sb[6] = byte(FlagDATA64B | FlagNORESP)

	expected := 0
	for i := 0; i < len(reqs); i++ {
		// 4-byte struct: 1 byte size, 3 byte address
		sb[32+(i*4)] = reqs[i].Size
		sb[33+(i*4)] = byte((reqs[i].Address >> 16) & 0xFF)
		sb[34+(i*4)] = byte((reqs[i].Address >> 8) & 0xFF)
		sb[35+(i*4)] = byte((reqs[i].Address >> 0) & 0xFF)
		expected += int(reqs[i].Size)
	}

	return expected, sendSerial(c.f, sb)
}

func (c *Conn) sendVPUTBatch(batch []snes.WriteRequest) error {
	total, err := c.sendVPUT(batch)
	if err != nil {
		return err
	}

	// determine total size of accompanying data:
	packets := total / 64
	remainder := total & 63
	// round up to accommodate any remainder:
	if remainder > 0 {
		packets++
	}

	// concatenate all accompanying data together in one large slice:
	whole := make([]byte, packets*64)
	o := 0
	for i := 0; i < len(batch); i++ {
		copy(whole[o:], batch[i].Data)
		o += len(batch[i].Data)
	}

	// send whole slice over serial port:
	err = sendSerial(c.f, whole)
	if err != nil {
		return err
	}

	// reply to each write request:
	for i := 0; i < len(batch); i++ {
		cb := batch[i].Completed
		if cb != nil {
			cb(snes.ReadOrWriteResponse{
				IsWrite: true,
				Address: batch[i].Address,
				Size:    batch[i].Size,
				Data:    batch[i].Data,
			})
		}
	}

	return nil
}
