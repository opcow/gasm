// go dasm
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const maxRead = 0x10000

var (
	prAddr  = flag.Bool("a", true, "print address column")
	outFile = flag.String("o", "", "output file")
	of      *os.File
)

func main() {

	var ifName string

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: gasm infile")
		os.Exit(1)
	}
	ifName = flag.Arg(0)

	finfo, err := os.Stat(ifName)
	if err != nil {
		fmt.Println("gasm: couldn't stat input file")
		os.Exit(1)
	}
	fSize := finfo.Size()
	if fSize > maxRead {
		fmt.Fprintln(os.Stderr, "gasm: input file exceeds max size (64k)")
		os.Exit(1)
	}
	fileEnd := int(fSize)

	f, err := os.Open(ifName)
	if err != nil {
		log.Fatal(err)
	}

	if *outFile != "" {
		of, err = os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer of.Close()
	} else {
		of = os.Stdout
	}
	var b bytes.Buffer

	h2 := fmt.Sprintf(";*** file %s ***\n", filepath.Base(ifName))
	h1 := fmt.Sprintf(";%s\n", strings.Repeat("*", len(h2)-1))

	b.Write([]byte(h1))
	b.Write([]byte(h2))
	b.Write([]byte(h1))
	b.Write([]byte("    PROCESSOR 6502\n"))
	b.WriteTo(of)

	{
		var header [4]byte
		defer f.Close()
		bufr := bufio.NewReader(f)
		var totalRead int
		for segments.count < maxSegs {
			br, err := io.ReadFull(bufr, header[:2])
			if err != nil {
				log.Fatal(err)
			}
			totalRead += br
			if header[0] == 0xff && header[1] == 0xff {
				continue
			}
			br, err = io.ReadFull(bufr, header[2:4])
			if err != nil {
				log.Fatal(err)
			}
			totalRead += br
			startAdd := int(header[0]) | int(header[1])<<8
			endAdd := int(header[2]) | int(header[3])<<8
			if startAdd < 0 || startAdd > 0xffff || endAdd < startAdd || endAdd > 0xffff {
				fmt.Fprintln(os.Stderr, "gasm: address error in binary")
				os.Exit(1)
			}
			br, err = io.ReadFull(bufr, memory[startAdd:endAdd+1])
			if err != nil {
				log.Fatal(err)
			}
			segments.seg[segments.count].start = startAdd
			segments.seg[segments.count].end = endAdd + 1
			segments.count++
			totalRead += br
			if totalRead >= fileEnd {
				break
			}
		}
	}
	// fmt.Print("	struct opcode {\n		int code;\n		int legth;\n		int mode;\n		std::string mnemonic;\n	};\n\n")

	// fmt.Println("opcode opcodes[] = {")
	// for i := range opcodes {
	// 	fmt.Printf("    { %#02x, %#02x, %#02x, ", opcodes[i].code, opcodes[i].length, opcodes[i].mode)
	// 	if opcodes[i].mnemonic == "" {
	// 		fmt.Println("\"XXX\" },")
	// 	} else {
	// 		fmt.Printf("\"%s\" },\n", opcodes[i].mnemonic)
	// 	}
	// }
	// fmt.Println("};")
	initPrintFuncs()
	disAsm()
}
