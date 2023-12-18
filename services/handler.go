package services

import (
	"net/http"

	"github.com/urfave/cli/v2"
)

func Handler(c *cli.Context, url string) error {
	println("url: ", url)
	req, err := http.Get(url)
	println("status: ", req.Status)
	if err != nil {
		return err
	}

	println("content-length: ", req.ContentLength/1024/1024, "MB")
	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	file := NewFile(url)
	defer file.Close()
	file.Open()
	file.Write(&body)
	return nil
}

func HandlerRoot(c *cli.Context) error {
	if c.Args().First() == "" {
		return nil
	}
	println("Root handler")
	return Handler(c, c.Args().First())
}
