package httpUtil

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SendFile(name string) error {
	url := "http://133.167.127.211:8080/upload?path=" + name

	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)

	f, err := os.Open("./iw/tmp/" + name)
	if err != nil {
		return err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return err
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
