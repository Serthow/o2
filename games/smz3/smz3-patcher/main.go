package main

import (
	"fmt"
	"io/ioutil"
	"o2/games/smz3"
	"o2/snes"
	"os"
	"path/filepath"
)

func main() {
	var err error

	args := os.Args[1:]
	if len(args) == 0 {
		panic(fmt.Errorf("missing filename argument"))
	}

	var contents []byte
	contents, err = ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	_, name := filepath.Split(args[0])

	var rom *snes.ROM
	rom, err = snes.NewROM(name, contents)
	if err != nil {
		panic(err)
	}

	// check what the alttp Factory instance thinks of this ROM:
	factoryInstance := smz3.FactoryInstance()
	isBestProvider := factoryInstance.IsROMSupported(rom)
	supported, whyNot := factoryInstance.CanPlay(rom)
	fmt.Printf("ROM is/should be supported? %v\n", isBestProvider)
	fmt.Printf("ROM can be played as SMZ3? %v\n", supported)
	if !supported {
		fmt.Printf("  Why not? %v\n", whyNot)
	}

	if !isBestProvider || !supported {
		return
	}

	// patch the ROM:
	patcher := smz3.NewPatcher(rom)
	err = patcher.Patch()
	if err != nil {
		panic(err)
	}

	// write it out to a file:
	err = ioutil.WriteFile("patched.smc", rom.Contents, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote to patched.smc")
}
