// Пользователь принимает сообщения от сервера и исправляет регистр:

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	message := make([]byte, 1024) // создадим буфер

	for i := 0; i < 3; i++ {
		n, err := conn.Read(message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println(err)
			return
		}
		fmt.Println(strings.ToUpper(string(message[:n])))
	}
}

// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:4545")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer conn.Close()
// 	for {
// 		var source string
// 		fmt.Print("Введите слово: ")
// 		_, err := fmt.Scan(&source)
// 		if err != nil {
// 			fmt.Println("Некорректный ввод", err)
// 			continue
// 		}
// 		// отправляем сообщение
// 		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		// получим ответ
// 		fmt.Println("Ответ:")
// 		conn.SetReadDeadline(time.Now().Add(time.Second * 5))

// 		// Получение данных вынесено в отдельный цикл,
// 		// поэтому даже если сервер получит больше
// 		// много данных, он их обработает
// 		for {
// 			buff := make([]byte, 1024)
// 			n, err := conn.Read(buff)
// 			if err != nil {
// 				break
// 			}
// 			fmt.Print(string(buff[0:n]))
// 			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
// 		}
// 	}
// }
