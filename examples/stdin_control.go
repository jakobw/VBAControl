package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vba_control"
)

func getInputFromStdin(vba *vba_control.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		vba.SendInput(strings.TrimSpace(line))
	}
}

func main() {
	rom := os.Args[1]
	vba := vba_control.New(rom)

	go vba.Start()
	fmt.Println("Starting VBA with", rom)

	getInputFromStdin(vba)
}
