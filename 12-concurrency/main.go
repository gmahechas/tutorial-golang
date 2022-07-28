package main

import (
	"fmt"
	"net/http"
)

func checkServer(server string, channel chan string) {
	_, error := http.Get(server)
	if error != nil {
		// fmt.Println(server, "is down")
		channel <- server + " is down"
	} else {
		// fmt.Println(server, "is up")
		channel <- server + " is up"
	}
}

func main() {
	channel := make(chan string)
	servers := []string{"https://www.google.com/", "https://www.facebook.com/", "https://go.dev/"}
	for _, server := range servers {
		go checkServer(server, channel)
	}
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}
}
