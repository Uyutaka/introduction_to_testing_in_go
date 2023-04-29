package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTest := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"0", 0, false, "0 is not prime, by definition!"},
		{"negative number", -1, false, "Negative numbers are not prime, by definition!"},
		{"not prime", 4, false, "4 is not a prime number because it is divisible by 2"},
	}

	for _, e := range primeTest {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if msg != e.msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// read os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from out read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "> " {
		t.Errorf("expected > but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// read os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from out read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a number to see if it is prime, or 'q' to quit.") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a valid number!"},
		{name: "prime number", input: "7", expected: "7 is a prime number!"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
