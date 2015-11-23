package main

import "fmt"
import "net"
import "bufio" 
import "os"
import "time"

func main() {
	p := fmt.Println
	writeBuffer := make(chan string,5)
	p("welcome")
	conn, err := net.Dial("tcp","localhost:6667")
	if err != nil{
		p("error:",err)
	}
	p("Enter your name:")
	r := bufio.NewReader(os.Stdin)
	nick, _ := r.ReadString('\n')
	fmt.Fprintf(conn, "NICK " + nick + "\r\n")
	fmt.Fprintf(conn, "USER gotest localhost irc.mytest.net Matthew \r\n")
	fmt.Fprintf(conn, "JOIN #test \r\n")

	go read(conn)
	go ping(writeBuffer)
	go send(conn,writeBuffer)
	write(writeBuffer)

}

func read(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
		if message == "PING" {
			fmt.Println("recieved ping")
			fmt.Fprintf(conn, "PRIVMSG #test : PONG \r\n")
		}
	}
}

func send(conn net.Conn,ch chan string){
	for {
		fmt.Fprintf(conn, <-ch)
	}
}

func write(ch chan string){
	for {
		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')
		ch <- fmt.Sprint("PRIVMSG #test :",text,"\r\n")
	}
}

func ping(ch chan string){
	for {
		time.Sleep(15 * time.Second)
		ch <- "PONG #test\r\n"
	}
}