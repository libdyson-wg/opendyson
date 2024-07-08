package shell

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

func PromptForPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("")
	if err != nil {
		err = fmt.Errorf("error reading password: %w", err)
	}
	return string(pass), err
}

func PromptForInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("error reading input: %w", err)
	}

	in = strings.TrimSpace(in)
	return in, err
}
