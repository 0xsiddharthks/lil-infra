package redis

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type parsedCommand struct {
	keyword string
	args    []string
}

const (
	PING = "ping"
	ECHO = "echo"
	SET  = "set"
	GET  = "get"
	DEL  = "del"
)

func Main() {
	fmt.Println("starting server.. on port 6380")

	l, err := net.Listen("tcp", "127.0.0.1:6380")
	if err != nil {
		fmt.Println("Failed to bind to port 6380")
		os.Exit(1)
	}
	var conn net.Conn

	for {
		conn, err = l.Accept()
		fmt.Println("Accepted new connection request from: ", conn.RemoteAddr().String())
		if err != nil {

			fmt.Println("Failed to accept connection")
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("connected to new client...")
	response := make([]byte, 1024)
	for {
		n, err := conn.Read(response)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client closed connection")
				break
			}
			conn.Write([]byte("-ERR Error reading: " + err.Error() + "\r\n"))
			break
		}
		fmt.Println("Received:", string(response[:n]))
		handleCommand(conn, response[:n])
	}
	conn.Close()
}

func handleCommand(conn net.Conn, command []byte) {
	parsedCommand, err := parseCommand(string(command))
	if err != nil {
		handleError(conn, err)
		return
	}
	fmt.Println("parsed command")
	switch parsedCommand.keyword {
	case PING:
		handlePing(conn, parsedCommand.args)
	case ECHO:
		handleEcho(conn, parsedCommand.args)
	case SET:
		handleSet(conn, parsedCommand.args)
	case GET:
		handleGet(conn, parsedCommand.args)
	default:
		handleError(conn, errors.New("ERR: command \""+parsedCommand.keyword+"\" not implemented"))
	}
}

func handleError(conn net.Conn, err error) {
	conn.Write([]byte("-" + err.Error() + "\r\n"))
}

func handlePing(conn net.Conn, args []string) {
	if (len(args)) > 1 {
		conn.Write([]byte("-ERR: invalid arguments for command: 'ping'\r\n"))
		return
	}
	if len(args) == 0 {
		conn.Write([]byte("+PONG"))
		conn.Write([]byte("\r\n"))
		return
	}
	fmt.Println("Response: ", encodeResponse(args[0:1]))
	conn.Write([]byte(encodeResponse(args[0:1])))
}

func handleEcho(conn net.Conn, args []string) {
	if (len(args)) != 1 {
		conn.Write([]byte("-ERR: invalid arguments for command: 'echo'\r\n"))
		return
	}
	conn.Write([]byte(encodeResponse(args[0:1])))
}

func handleSet(conn net.Conn, args []string) {
	if (len(args)) < 2 || (len(args)) > 3 {
		conn.Write([]byte("-ERR: invalid arguments for command: 'set'\r\n"))
		return
	}
	if (len(args)) == 2 {
		Instance().set(args[0], args[1], -1)
	} else {
		ttl, err := strconv.Atoi(args[2])
		if err != nil {
			conn.Write([]byte("-ERR: invalid ttl argument for command: 'set'\r\n"))
			return
		}
		expiry := time.Now().Unix()*1000 + int64(ttl)
		Instance().set(args[0], args[1], expiry)
	}
	conn.Write([]byte("+OK\r\n"))
}

func handleGet(conn net.Conn, args []string) {
	if (len(args)) != 1 {
		conn.Write([]byte("-ERR: invalid arguments for command: 'get'\r\n"))
		return
	}
	val, err := Instance().get(args[0])
	if err != nil {
		conn.Write([]byte(nil))
	}
	conn.Write([]byte(encodeResponse([]string{val})))
}

func parseCommand(command string) (parsedCommand, error) {
	commandLiterals := strings.Split(command, "\r\n")
	commandLiterals = removeEvenElementsFromSlice(commandLiterals)
	if len(commandLiterals) < 2 {
		return parsedCommand{}, errors.New("ERR: too few keywords")
	}
	commandName := strings.ToLower(commandLiterals[1])
	return parsedCommand{commandName, commandLiterals[2:]}, nil
}

// TODO: replace with general purpose decoder
func removeEvenElementsFromSlice(slice []string) []string {
	var result []string
	for i, v := range slice {
		if i%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func encodeString(response string) string {
	return "$" + strconv.Itoa(len(response)) + "\r\n" + response + "\r\n"
}

func encodeResponse(response []string) string {
	size := len(response)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return encodeString(response[0])
	}
	header := "*" + strconv.Itoa(size) + "\r\n"
	body := []string{}

	for _, v := range response {
		body = append(body, encodeString(v))
	}
	return header + strings.Join(body, "")
}

// func decodeResponse(response []string) []string {
// 	return []string{}
// }
