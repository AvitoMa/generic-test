package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const serverAddr = "localhost:9999"

type Result struct {
	Number    string
	Factorial string
}

func main() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Silly favicon request")
	})

	http.HandleFunc("/", calculateFactorial)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func calculateFactorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	//TODO : Explain why we need this line
	reply = bytes.Trim(reply, "\u0000")

	result := &Result{
		Number:    param,
		Factorial: string(reply),
	}

	println("reply from server=", string(reply))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
