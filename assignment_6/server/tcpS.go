package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	var response string
Loop:
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		channels := make(map[string]string)
		input := strings.Split(netData, " ")

		command := string(input[0])
		fmt.Println("command " + command)

		user := input[1]
		fmt.Println("user " + user)

		channel := input[2]
		fmt.Println("channel " + channel)

		switch command {
		case "Reg":
			if user[0] != '@' {
				// fmt.Println("Invalid user name")
				response = "Invalid user name\n"
				fmt.Print("-> ", string(response))
				fmt.Println()
				c.Write([]byte(response))
				continue Loop
			}
			if channel[0] != '#' {
				// fmt.Println("Invalid channel name")
				response = "Invalid channel name\n"
				fmt.Print("-> ", string(response))
				fmt.Println()
				c.Write([]byte(response))
				continue Loop
			}
			channels[user] = channel
			// fmt.Println("User " + user + " registered to channel " + channel)
			response = "User " + user + " registered to channel " + channel + "\n"
			fmt.Print("-> ", string(response))
			fmt.Println()
			c.Write([]byte(response))
			continue Loop
		case "Msg":
			if _, ok := channels[user]; !ok {
				// fmt.Println("User " + user + " is not registered")
				response = "User " + user + " is not registered\n"
				fmt.Print("-> ", string(response))
				fmt.Println()
				c.Write([]byte(response))
				continue Loop
			}
			response = "User " + user + " sent message to channel " + channels[user] + "\n"
			// fmt.Println("User " + user + " sent message to channel " + channels[user]+"\n")
			fmt.Print("-> ", string(response))
			fmt.Println()
			c.Write([]byte(response))
			continue Loop

		case "Quit":
			if _, ok := channels[user]; !ok {
				// fmt.Println("User " + user + " is not registered")
				response = "User " + user + " is not registered\n"
				fmt.Print("-> ", string(response))
				fmt.Println()
				c.Write([]byte(response))
				continue Loop
			}
			// fmt.Println("User " + user + " quit from channel " + channels[user] + "\n")
			response = "User " + user + " quit from channel " + channels[user] + "\n"
			fmt.Print("-> ", string(response))
			fmt.Println()
			c.Write([]byte(response))
			delete(channels, user)
			continue Loop
		case "STOP":
			fmt.Println("TCP server exiting...")
			return
		default:
			fmt.Println("Invalid command")
			return
		}
	}
}
