package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
	passLength   = flag.Int("n", 8, "specify generate password length")
	format       = flag.String("f", "LUDS", "generated password using you can select characters")
	uuidFlag     = flag.Bool("uuid", false, "generated uuid v4")
	helpFlag     = flag.Bool("-help", false, "help")
)

func shuffle(val []rune) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(val) > 0 {
		n := len(val)
		randIndex := r.Intn(n)
		val[n-1], val[randIndex] = val[randIndex], val[n-1]
		val = val[:n-1]
	}
}

func isLowerLetters() bool {
	return strings.Contains(*format, "L")
}

func isUpperLetters() bool {
	return strings.Contains(*format, "U")
}

func isDigits() bool {
	return strings.Contains(*format, "D")
}

func isSymbols() bool {
	return strings.Contains(*format, "S")
}

func characters() []rune {
	source := ""
	if isLowerLetters() {
		source += lowerLetters
	}
	if isUpperLetters() {
		source += upperLetters
	}
	if isDigits() {
		source += digits
	}
	if isSymbols() {
		source += symbols
	}
	return []rune(source)
}

func copyToClipboard(password string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = in.Write([]byte(password))
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func uuidV4() string {
	return uuid.New().String()
}

func main() {
	flag.Parse()

	generated := ""
	if *uuidFlag {
		generated = uuidV4()
	} else if *helpFlag {
		fmt.Printf("help")
	} else {
		passSource := characters()
		shuffle(passSource)
		for _, r := range passSource[:*passLength] {
			generated += string(r)
		}
	}

	fmt.Println(generated)
	if err := copyToClipboard(generated); err != nil {
		_, e := fmt.Fprintf(os.Stderr, "error occured: cannot copy to clipboard: %s", err.Error())
		if e != nil {
			panic(err)
		}
	}
}
