package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (app *application) ProcessMultipartReq(r *http.Request) error {

	for _, headers := range r.MultipartForm.File {

		for _, h := range headers {
			file, err := h.Open()
			fmt.Println(h.Header["Content-Disposition"][0])
			if err != nil {
				return err
			}

			defer file.Close()

			finalByteArr, err := ByteFromMegaFile(file)

			fileSave, err := os.Create(h.Filename)
			if err != nil {
				return err
			}
			_, err = fileSave.Write(finalByteArr)
			if err != nil {
				return err
			}
			defer fileSave.Close()

		}

	}
	return nil

}

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
