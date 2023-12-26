package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func tcpClient(hostAndPort string) {
	fmt.Println("Starting TCP client")
	c, err := net.Dial("tcp", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func handleTcpConnection(c net.Conn) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		if temp == "ping" {
			_, err := c.Write([]byte("pong "))
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		_, err = c.Write([]byte(string(counter)))
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Close()
}

func tcpServer(hostAndPort string) {
	fmt.Println("Starting TCP server")
	l, err := net.Listen("tcp4", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleTcpConnection(c)
		count++
	}
}

func udpClient(hostAndPort string) {
	fmt.Println("Starting UDP client")
	s, err := net.ResolveUDPAddr("udp4", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func udpServer(hostAndPort string) {
	fmt.Println("Starting UDP server")
	s, err := net.ResolveUDPAddr("udp4", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("-> ", string(buffer[0:n-1]))

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {

	hostArg := flag.String("host", "127.0.0.1", "Host or ip address")
	portArg := flag.String("port", "3000", "Port")
	clientArg := flag.Bool("client", false, "Start as client (defaults as server)")
	udpArg := flag.Bool("udp", false, "Use UDP (defaults to TCP)")

	flag.Parse()
	//fmt.Println("args:", flag.Args())

	if *hostArg == "" {
		fmt.Println("Warning: Host not provided")
		return
	}

	hostAndPort := *hostArg + ":" + *portArg
	if *clientArg {
		if *udpArg {
			udpClient(hostAndPort)
		} else {
			tcpClient(hostAndPort)
		}
	} else {
		if *udpArg {
			udpServer(hostAndPort)
		} else {
			tcpServer(hostAndPort)
		}
	}

}
