package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	errNotImplemented = errors.New("not implemented yet")
	errNoWords        = errors.New("no words in file")
)

func main() {
	var in *string
	var out *string

	in = flag.String("in", "input.txt", "the filename of the text to scramble")
	out = flag.String("out", "output.txt", "the filename of the scrambled output file")
	flag.Parse()

	fullText, err := readFile(*in)
	if err != nil {
		fmt.Printf("error in readFile: %v", err)
	}

	scrambledText, err := scrambleText(fullText)
	if err != nil {
		fmt.Printf("error in scrambleText: %v", err)
	}

	err = writeFile(*out, scrambledText)
	if err != nil {
		fmt.Printf("error in writeFile: %v", err)
	}
}

func readFile(filename string) (string, error) {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func scrambleText(text string) (string, error) {
	words := strings.Split(text, " ")
	if len(words) == 0 {
		return "", errNoWords
	}

	scrambled := ""
	for _, word := range words {
		scrambled = scrambled + scrambleWord(word) + " "
	}

	return scrambled, nil
}

func writeFile(filename, scrambled string) error {
	fmt.Printf("\nOUTPUT:\n\n %s\n\n", scrambled)
	err := os.WriteFile(filename, []byte(scrambled), 0644)
	if err != nil {
		fmt.Printf("error in writeFile: %v", err)
		return err
	}

	return nil
}

func scrambleWord(word string) string {
	firstLetter := word[0:1]
	lastLetter := ""
	if len(word) > 1 {
		lastLetter = word[len(word)-1:]
	}
	restWord := ""
	if len(word)-1 >= 1 {
		restWord = word[1 : len(word)-1]
	}

	newWord := firstLetter + scrambleMiddle(restWord) + lastLetter
	return newWord
}

func scrambleMiddle(middle string) string {
	// no possibility to switch
	if len(middle) < 2 {
		return middle
	}

	// length two, always switch
	if len(middle) == 2 {
		return string(middle[1]) +
			string(middle[0])
	}

	// from length three, truly scramble
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	slice := []byte(middle)
	r2.Shuffle(len(middle), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	return string(slice)
}
