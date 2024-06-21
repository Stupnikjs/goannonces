package util

import (
	"bufio"
	"io"
	"log"
)

func ByteFromMegaFile(file io.Reader) ([]byte, error) {

	reader := bufio.NewReader(file)

	finalByteArr := make([]byte, 0, 2048*1000)

	for {
		soloByte, err := reader.ReadByte()
		if err != nil {
			log.Println(err)
			break
		}

		finalByteArr = append(finalByteArr, soloByte)
	}

	return finalByteArr, nil

}
