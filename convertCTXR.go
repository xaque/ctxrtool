package main

import (
	"fmt"
	"os"
	"encoding/binary"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2{
		fmt.Println("Please provide image(s) to convert.")
		os.Exit(0)
	}
	for i := 1; i < len(os.Args); i++ {
		convertCTXRFileToPAM(os.Args[i])
	}
}

func convertCTXRFileToPAM(filename string) {
	// Open CTXR file as read-only
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	check(err)
	defer f.Close()
	fmt.Printf("Converting %s...\n", filename)

	// TODO currently ignoring the rest of the header! Assuming it's a particular kind of CTXR
	// Hardcoded values from CTXR header
	widthIndex :=  int64(0x2c)
	heightIndex := int64(0x2e)
	headerSize := int64(0x80)
	depth := uint16(4) // Assuming input is ARGB for a depth of 4

	// Read image values from CTXR header
	widthB := make([]byte, 2)
	heightB := make([]byte, 2)
	f.ReadAt(widthB, widthIndex)
	f.ReadAt(heightB, heightIndex)
	width := binary.BigEndian.Uint16(widthB)
	height := binary.BigEndian.Uint16(heightB)

	//size := width * height * depth
	fmt.Printf("Dimensions: %dx%d\n", width, height)

	// Read raster image data into buffer
	fInfo, err := f.Stat()
	buf := make([]byte, fInfo.Size() - headerSize)
	f.ReadAt(buf, headerSize)

	// Convert raster from ARGB to RGBA
	for i := int64(0); i < int64(len(buf) - 4); i+=4 {
		alpha := buf[i]
		red := buf[i+1]
		green := buf[i+2]
		blue := buf[i+3]
		//debug
		// if i == 0 {
		// 	fmt.Printf("pixel: %x\n", buf[i:4])
		// }
		buf[i] = red
		buf[i+1] = green
		buf[i+2] = blue
		buf[i+3] = alpha
		// if i == 0 {
		// 	fmt.Printf("pixel: %x\n", buf[i:4])
		// }
	}

	// Create new file for converted image
	nf, err := os.Create(filename + ".pam")
	check(err)
	defer nf.Close()

	// Generate PAM header
	header := "P7\n" + 
			  fmt.Sprintf("WIDTH %d\n", width) +
			  fmt.Sprintf("HEIGHT %d\n", height) +
			  fmt.Sprintf("DEPTH %d\n", depth) +
			  "MAXVAL 255\n" + 
			  "TUPLTYPE RGB_ALPHA\n" + 
			  "ENDHDR\n"

	// Write PAM to file
	nf.WriteString(header)
	nf.Write(buf)
	fmt.Printf("Wrote PAM file to %s.pam\n", filename)
}