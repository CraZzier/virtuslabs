package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//Help serves for handling help command
func Help() {
	fmt.Print(`
	help - showing available commands
	version - current version of an application
	run --file <filename> serves HTML file to a port 3000 on 127.0.0.1
	`)
}

//Version serves for handling version command
func Version() {
	fmt.Println("Its the 1.0.0 version")
}

//Run serves for handling run command
func Run(words []string) {
	if len(words) == 1 {
		fmt.Println("No flag added")
		return
	}
	switch words[1] {
	case "--file":
		if len(words) == 2 {
			fmt.Println("No file specified")
			return
		}
		if words[2] != "" {
			_, err := os.OpenFile(filepath.Join("./", words[2]), os.O_RDWR, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			go func() {
				http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
					http.ServeFile(res, req, filepath.Join("./", words[2]))
				})
				fmt.Println("Server has been started on 127.0.0.1:3000 - " + words[2])
				http.ListenAndServe(":3000", nil)
			}()
		}
	default:
		fmt.Println("Unknown flag")
	}
}
func main() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		//Spliting by spaces
		commandArg := strings.Fields(input)
		if len(commandArg) > 3 {
			fmt.Println("Unknown command structure")
			continue
		}
		switch commandArg[0] {
		case "help":
			Help()
		case "version":
			Version()
		case "run":
			Run(commandArg)
		default:
			fmt.Println("Unknown command")
		}
	}

}
