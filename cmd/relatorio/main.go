package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Vulnerável de propósito: entrada do usuário vai para o shell.
	out, err := exec.Command("sh", "-c", "wc -l "+os.Args[1]).Output()
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Print(string(out))
}
