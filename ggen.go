package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
	passLength   = 8
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

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		l, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("error: cannno generate password")
		}
		passLength = l
	}
	passSource := []rune(lowerLetters + upperLetters + digits + symbols)
	shuffle(passSource)

	password := ""
	for _, r := range passSource[:passLength] {
		password += string(r)
	}
	fmt.Println(password)
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = in.Write([]byte(password))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = in.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}
