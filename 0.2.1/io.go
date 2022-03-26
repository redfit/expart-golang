package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

type neverEnding byte

func (b neverEnding) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}

func TestIsPNG(t *testing.T) {
	n, want := int64(10), false
	r := io.LimitReader(neverEnding('x'), n)
	got, err := IsPNG1(r)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Error(want, "!=", got)
	}
}

/*
func TestUpperCount(t *testing.T) {
	str, want := "AbcD", 2
	var buf bytes.Buffer
	r := io.TeeReader(strings.NewReader(str), &buf)
	got, err := UpperCount(r)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Error(want, "!=", got)
	}
	if str != buf.String() {
		t.Error("読み込んだ文字列が一致しない")
	}
}
*/

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	w := io.MultiWriter(f, h)
	_, err = io.WriteString(w, "hello")
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))

}

func NewPNG(r io.Reader) (io.Reader, error) {
	magicNum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicNum))
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	if !bytes.Equal(magicNum, buf) {
		return nil, errors.New("not a png")
	}
	pngImg := io.MultiReader(bytes.NewReader(magicNum), r)
	return pngImg, nil

}

func IsPNG1(r io.Reader) (bool, error) {
	magicNum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicNum))
	//_, err := io.ReadAtLeast(r, buf, len(magicNum))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return false, err
	}
	return bytes.Equal(buf, magicNum), nil
}

func IsPNG2(r io.ReadSeeker) (bool, error) {
	magicNum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicNum))
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}
	return bytes.Equal(buf, magicNum), nil
}

/*
func Post(m *Message) (rerr error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		enc := json.NewEncoder(pw)
		err := enc.Encode(m)
		if err != nil {
			rerr = err
		}
	}()
	const url = "http://example.com"
	const contentType = "application/json"
	_, err := http.Post(url, contentType, pr)
	if err != nil {
		return err
	}
	return nil
}
*/
