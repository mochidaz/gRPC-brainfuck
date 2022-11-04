package tests

import (
	"testing"
	"mochidaz/simple-service/utils"
)

func exec(source string) string {
	ex, err := utils.Execute(source)

	if err != nil {
		panic(err)
	}

	return ex
}

func TestInterpret(t *testing.T) {

	case_1 := "+++++++++[>++++++++>+++++++++++>++++>++++++++++<<<<-]>.>++.+++++++..+++.>----.>---.<<.+++.------.--------.>+."
	case_2 := "+++++++++[>+++++++++>++++++++++++>++++>++++++++<<<<-]>+++.>----.+.++++++++++.>----.<----------.++++++++++.>.>-.<<<--.--.>>>----.<+."
	case_3 := "+++++++++[>++++++++++>+++++++++++>++++>+++++++++++++>++++++++<<<<<-]>---.>++.>----.<----.>>---.<<++++.>.>++.<<.>>-.+.<<++++.+++++.-------.>.>.<<+.---.>.>>-.<<<<-----.--.>>>>----.<<+."

	if !(exec(case_1) == "Hello World!") {
		panic("Test Failed!")
	}

	if !(exec(case_2) == "This is GRPC!") {
		panic("Test Failed!")
	}

	if !(exec(case_3) == "We are testing the GRPC!") {
		panic("Test Failed!")
	}

}