package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Eres más feo que Truman Capote"
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("https://www.koko.eco")
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
