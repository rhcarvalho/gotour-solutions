package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (z *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = z.r.Read(p)
    lkp := map[byte]byte{
        'A': 'N', 'B': 'O', 'C': 'P', 'D': 'Q',
        'E': 'R', 'F': 'S', 'G': 'T', 'H': 'U',
        'I': 'V', 'J': 'W', 'K': 'X', 'L': 'Y',
        'M': 'Z', 'N': 'A', 'O': 'B', 'P': 'C',
        'Q': 'D', 'R': 'E', 'S': 'F', 'T': 'G',
        'U': 'H', 'V': 'I', 'W': 'J', 'X': 'K',
        'Y': 'L', 'Z': 'M',
        'a': 'n', 'b': 'o', 'c': 'p', 'd': 'q',
        'e': 'r', 'f': 's', 'g': 't', 'h': 'u',
        'i': 'v', 'j': 'w', 'k': 'x', 'l': 'y',
        'm': 'z', 'n': 'a', 'o': 'b', 'p': 'c',
        'q': 'd', 'r': 'e', 's': 'f', 't': 'g',
        'u': 'h', 'v': 'i', 'w': 'j', 'x': 'k',
        'y': 'l', 'z': 'm'}
    for i := 0; i < n; i++ {
        if v, ok := lkp[p[i]]; ok {
            p[i] = v
        }
    }
    return
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
