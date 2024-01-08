package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot* rot13Reader) Read(b[] byte) (int, error) {
	read, err := rot.r.Read(b);
	if err != nil && err != io.EOF {
		return read, err
	}

	for i, val := range b {
		if i >= read {
			break;
		}

		if val <= 'z' && val >= 'a' {
			b[i] = (val - 'a' + 13) % 26 + 'a'
		} else if val <= 'Z' && val >= 'A' {
			b[i] = (val - 'A' + 13) % 26 + 'A'
		}
	}

	return read, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
