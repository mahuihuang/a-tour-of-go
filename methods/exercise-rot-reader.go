package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		v := p[i]
		switch {
		case v >= 'a' && v <= 'z':
			p[i] = 'a' + (v-'a'+13)%26
		case v >= 'A' && v <= 'Z':
			p[i] = 'A' + (v-'A'+13)%26
		default:
		}
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
