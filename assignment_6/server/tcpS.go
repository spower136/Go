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
	var channels = make(map[string]string)
	// channels["#general"] = "#general"

	users := make(map[string]string)
	// users["user1"] = "user1"

	msgs := make(map[string]string)
	// msgs["#general"] = "#general"

Loop:
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input := strings.Split(netData, " ")
		command := input[0]

		switch command {
		case "REG":
			if len(input) == 1 {
				response = "REGISTER: Please provide a username"
			} else {
				username := input[1]
				if username[0] != '@' {
					response = "REGISTER: Username must start with '@'"
				} else {
					if _, ok := users[username]; ok {
						response = "REGISTER: Username already exists"
					} else {
						users[username] = username
						response = "REGISTER: Successfully registered"
					}
				}
			}
		case "LEAVE":
			if len(input) == 1 {
				response = "LEAVE: Please provide a channel"
			} else {
				channel := input[1]
				if _, ok := channels[channel]; ok {
					delete(channels, channel)
					response = "LEAVE: Successfully left"
				} else {
					response = "LEAVE: Channel does not exist"
				}
			}
		case "MSG":
			if len(input) == 1 {
				response = "MSG: Please provide a message"
			} else {
				message := input[1]
				msgs[message] = message
				response = "MSG: " + message
			}
		case "LIST":
			// if input[1] == "LIST" {
			if len(channels) == 0 {
				response = "CHNS: There are no channels"
			} else {
				response = "CHNS: "
				for _, value := range channels {
					response += value //+ " "
				}
			}
			// }
		case "JOIN":
			if len(input) == 1 {
				response = "JOIN: Please provide a channel"
			} else {
				channel := input[1]
				if channel[0] != '#' {
					response = "JOIN: Channel must start with '#'"
				} else {
					if _, ok := channels[channel]; ok {
						response = "JOIN: Channel already exists"
					} else {
						channels[channel] = channel
						response = "JOIN: Successfully joined"
					}
				}
			}
		case "STOP":
			response = "CHNS: "
			for _, value := range channels {
				response += value + " "
			}
			response += "\nMSGS: "
			for _, value := range msgs {
				response += value + " "
			}
			response += "\nUSERS: "
			for _, value := range users {
				response += value + " "
			}
			break Loop
		default:
			response = "UNKNOWN: Command not recognized"
		}
		c.Write([]byte(response + "\n"))

	}

}
