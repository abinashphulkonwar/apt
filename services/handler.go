package services

import (
	"io"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Handler(c *cli.Context, url string) error {
	println("url: ", url)
	req, err := http.Get(url)

	if err != nil {
		return err
	}
	println("status: ", req.Status)

	println("content-length: ", req.ContentLength/1024/1024, "MB", req.ContentLength, "bytes")
	file := NewFile(url)
	file.Open()

	defer file.Close()
	err = ReadBody(req, file)
	if err != nil {
		return err
	}
	return nil
}

func HandlerRoot(c *cli.Context) error {
	if c.Args().First() == "" {
		return nil
	}
	println("Root handler")
	return Handler(c, c.Args().First())
}

func ReadBody(req *http.Response, file *File) error {
	body := make([]byte, 4096*1000)
	i, err := req.Body.Read(body)
	if err != nil {
		if err == io.EOF {
			println("end: ", i)
			chuck := body[0:i]
			file.Write(&chuck)
			return nil
		}
		return err
	}
	println("chunks get: ", i/1024, "KB")
	chuck := body[0:i]
	file.Write(&chuck)
	return ReadBody(req, file)
}
