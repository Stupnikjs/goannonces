package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) ProcessMultipartReq(r *http.Request) error {
	for i := range 2 {
		headers := r.MultipartForm.File[fmt.Sprintf("file%d", i)]

		for _, h := range headers {
			file, err := h.Open()
			if err != nil {
				return err
			}

			defer file.Close()

			buf, err := io.ReadAll(file)

			fmt.Printf("File uploaded successfully: %v", h.Filename)

			err = app.LoadToBucket(h.Filename, buf)
			if err != nil {
				return err
			}

		}
		/*

		 */
	}
	return nil
}
