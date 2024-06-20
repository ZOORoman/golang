// Сервер передает сообщения разного регистра:

package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	conn.Write([]byte("message"))
	time.Sleep(10)
	conn.Write([]byte("MesSaGe"))
	time.Sleep(10)
	conn.Write([]byte("MESSAGE"))
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net"
// 	"os"
// )

// func main() {
// 	httpRequest := "GET / HTTP/1.1\n" +
// 		"Host: golang.org\n\n"
// 	// conn, err := net.Dial("tcp", "golang.org:80")
// 	conn, err := net.Dial("tcp", "172.217.16.145:80")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer conn.Close()

// 	if _, err = conn.Write([]byte(httpRequest)); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	io.Copy(os.Stdout, conn)
// 	fmt.Println("Done")
// }
