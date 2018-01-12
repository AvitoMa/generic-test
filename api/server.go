package main

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const serverAddr = "localhost:9999"

func main() {
	http.HandleFunc("/", calculateFactorial)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func calculateFactorial(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Path
	param = strings.TrimPrefix(param, "/")

	n, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(strconv.FormatInt(n, 10)))

	println("write to server = ", n)

	//TODO : Create a bug here
	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	w.Write(reply)
}
