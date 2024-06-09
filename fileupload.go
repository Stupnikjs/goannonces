package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) ProcessMultipartReq(r *http.Request) error {

	for _, headers := range r.MultipartForm.File {

		for _, h := range headers {
			file, err := h.Open()
			fmt.Println(h.Size, h.Header)
			if err != nil {
				return err
			}

			defer file.Close()

			buf, err := io.ReadAll(file)

			err = app.LoadToBucket(h.Filename, buf)
			if err != nil {
				return err
			}

			fmt.Printf("File uploaded successfully: %v", h.Filename)

		}

	}
	return nil

}
