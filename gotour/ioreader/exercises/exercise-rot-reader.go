package exercises

import (
	"io"
	"os"
	"strings"
)

// Rot13Reader sample
type Rot13Reader struct {
	r io.Reader
}

func (rt *Rot13Reader) Read(b []byte) (int, error) {
	n, err := rt.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'a' && b[i] <= 'm' || b[i] >= 'A' && b[i] <= 'M':
			b[i] += 13
		case b[i] >= 'n' && b[i] <= 'z' || b[i] >= 'N' && b[i] <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

// ExecuteRotReader just execute de function sample
func ExecuteRotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
