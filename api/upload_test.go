package api

import (
	"io"
	"os"
	"testing"
)

func TestByteFromMegaFile(t *testing.T) {

	file, _ := os.Open(`C:\Users\nboud\OneDrive\Bureau\Go_Projects\zik\static\download\test.mp3`)
	defer file.Close()
	b, _ := io.ReadAll(file)

	newFile, _ := os.Open(`C:\Users\nboud\OneDrive\Bureau\Go_Projects\zik\static\download\test.mp3`)
	respByte, _ := ByteFromMegaFile(newFile)

	if len(b) != len(respByte) {

		t.Errorf("expected len b equal to respByte %v != %v", len(b), len(respByte))
	}

}
