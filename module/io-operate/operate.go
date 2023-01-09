package io_operate

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func init() {
	log.Println("demon modules io operate")
	//ioCopy()
	//ioCopyBuf()
	//ioCopyN()
	//ioPipe()
	//ioReadAll()
	//ioReadAtLeast()
	//ioReadFull()
	ioWriteString()
}

func ioCopy() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func ioCopyBuf() {
	d := strings.NewReader("dst io.Reader\n")
	s := strings.NewReader("src io.Reader\n")
	b := make([]byte, 8)

	if _, err := io.CopyBuffer(os.Stdout, d, b); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, s, b); err != nil {
		log.Fatal(err)
	}
}

func ioCopyN() {
	r := strings.NewReader("some io.Reader \n")

	if _, err := io.CopyN(os.Stdout, r, 8); err != nil {
		log.Fatal(err)
	}
}

func ioPipe() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read \n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func ioReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", b)
}

func ioReadAtLeast() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 18)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}
}

func ioReadFull() {
	r := strings.NewReader("some io.Reader to be read\n")

	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}
}

func ioWriteString() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
