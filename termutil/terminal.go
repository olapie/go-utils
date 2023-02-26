package termutil

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
)

func ReadOne() (byte, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}

	var b [1]byte
	_, err = os.Stdin.Read(b[:])
	term.Restore(int(os.Stdin.Fd()), oldState)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func ReadPassword(msg ...any) string {
	var pass []byte
	var err error
	for len(pass) == 0 {
		log.Print(msg...)
		log.Print(": ")
		pass, err = terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(err)
		}
		log.Println()
	}
	return string(pass)
}

func ReadConfirmedPassword(prompt1, prompt2 string) *string {
	for i := 0; i < 3; i++ {
		pass1 := ReadPassword(prompt1)
		pass2 := ReadPassword(prompt2)
		if pass1 == pass2 {
			return &pass1
		}
	}
	return nil
}

func ConfirmInput(answer string) bool {
	answer = strings.TrimSpace(answer)
	if answer == "" {
		panic("answer cannot be empty")
	}
	log.Printf("Enter '%s' to confirm: ", answer)
	var actual string
	fmt.Scanln(&actual)
	return actual == answer
}
