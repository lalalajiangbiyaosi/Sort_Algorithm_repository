package main

import (
	// "fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close()
	}()
	MustCopy(os.Stdout, conn)
	log.Println("done")
}
func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
