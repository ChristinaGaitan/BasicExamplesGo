// Chat server

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

// Mensajes entrantes
var (
	entrantes = make(chan client)
	salientes = make(chan client)
	mensajes = make(chan string)
)
// Componente que emite los mensajes salientes
func broadcaster() {

	// Función broadcaster (emisor)
	// Todos los clientes conectados
	clients := make(map[client]bool)

	for {
		select {

		case msg := <- mensajes:
			// Emitir el mensaje entrante a todos los clientes
			// que están conectados por medio de channels
			for cli := range clients {
				cli <- msg
			}

		case cli := <- entrantes:
			clients[cli] = true

		case cli := <- salientes: 
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn){
	// Mensajes salientes
	ch := make(chan string)
	go clientWriter(conn, ch)

	quien := conn.RemoteAddr().String()
	ch <- "Tú eres " + quien

	mensajes <- quien + " se ha conectado a la sala"
	entrantes <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		mensajes <- quien + ": " + input.Text()
	}

	salientes <- ch
	mensajes <- quien + " se ha ido"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string){
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}