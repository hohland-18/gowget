package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
        "github.com/brentley/go-pkg-optarg"
)

func main() {
	optarg.Add("o", "output-document", "output filename", "")

	var (
		filename string
	)

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "o":
			filename = opt.String()
		}
	}

	if len(optarg.Remainder) == 1 {
		url := optarg.Remainder[0]

		if len(filename) == 0 {
			_, filename = path.Split(url)
		}

		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{
			Transport: transport,
		}

		res, err := client.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		n, err := io.Copy(file, res.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(n, "bytes downloaded.")
	} else {
		optarg.Usage()
	}
}
