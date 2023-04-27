package main

import "testing"

func Test_isPrime(t *testing.T){
	primeTest:=[]struct{
		name string
		testNum int
		expected bool
		msg string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"0", 0, false, "0 is not prime, by definition!"},
		{"negative number", -1, false, "Negative numbers are not prime, by definition!"},
		{"not prime", 4, false, "4 is not a prime number because it is divisible by 2"},
	}

	for _, e := range primeTest{
		result, msg:=isPrime(e.testNum)
		if e.expected && !result{
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result{
			t.Errorf("%s: expected false but got true", e.name)
		}
		if msg != e.msg{
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
