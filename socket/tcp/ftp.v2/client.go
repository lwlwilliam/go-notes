// FTPClient
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

// strings used by the user interface
var commands = map[string]string{
	"cd":    "cd",
	"clear": "clear",
	"help":  "help",
	"ls":    "ls",
	"pwd":   "pwd",
	"quit":  "quit",
}

// strings used across the network
const (
	LS  = "LS"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	var host string
	if len(os.Args) != 2 {
		host = "localhost"
	} else {
		host = os.Args[1]
	}

	conn, err := net.Dial("tcp", host+":1202")
	checkError(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")

		line, err := reader.ReadString('\n')
		// lose trailing whitespace
		line = strings.TrimRight(line, " \t\r\n")
		if err != nil {
			break
		}

		// split into command + arg
		strs := strings.SplitN(line, " ", 2)
		// decode user request
		switch strs[0] {
		case commands["ls"]:
			dirRequest(conn)
		case commands["cd"]:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case commands["pwd"]:
			pwdRequest(conn)
		case commands["quit"]:
			conn.Close()
			os.Exit(0)
		case commands["clear"]:
			clear()
		case commands["help"]:
			help()
		default:
			fmt.Println("Unknown command")
		}
	}
}

func help() {
	var sl []string
	for k, _ := range commands {
		sl = append(sl, k)
	}

	sort.Strings(sl)

	for _, v := range sl {
		fmt.Printf("%-15s", commands[v])
	}

	fmt.Printf("\n")
}

// clear the terminal
func clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Println("Unknown clear command on", runtime.GOOS)
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(LS + " "))

	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		// read till we hit a blank line
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir \"" + s + "\"")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
