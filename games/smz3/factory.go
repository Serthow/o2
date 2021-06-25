package smz3

import (
  "log"
	"o2/games"
	"o2/snes"
)

var gameName = "SMZ3"

type Factory struct{}

var factory *Factory

func FactoryInstance() *Factory { return factory }

func (f *Factory) IsROMSupported(rom *snes.ROM) bool {
  log.Printf("ROM ID Values", rom.Header.HeaderVersion(), rom.Header.MapMode, rom.Header.ROMSize, rom.Header.OldMakerCode)
	if rom.Header.HeaderVersion() != 1 {
    log.Printf("HeaderVersion Failed", rom.Header.HeaderVersion())
		return false
	}
	if rom.Header.MapMode != 0x35{
    log.Printf("MapMode Failed")
		return false
	}
	if rom.Header.ROMSize < 0x0D {
    log.Printf("ROMSize Failed")
		return false
	}
	if rom.Header.OldMakerCode != 0x01 {
    log.Printf("OldMakerCode Failed")
		return false
	}
	return true
}

func (f *Factory) CanPlay(rom *snes.ROM) (ok bool, whyNot string) {
	// TODO: read header of ROM to determine what variants are supported or not
	return true, ""
}

func (f *Factory) Patcher(rom *snes.ROM) games.Patcher {
	return &Patcher{rom: rom}
}

func init() {
	factory = &Factory{}
	games.Register(gameName, factory)
}
