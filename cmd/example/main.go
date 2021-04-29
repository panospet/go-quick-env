package main

import (
	"fmt"
	"go-quick-env/internal/quickenv"
)

type Example struct {
	MyBool   bool
	MyInt    int
	MyString string
	MyFloat  float64
}

func main() {
	if err := quickenv.LoadDotEnvFile(); err != nil {
		panic(err)
	}

	var example Example

	example.MyBool = quickenv.GetEnvAsBool("TEST_BOOL_VAR", true)
	example.MyInt = quickenv.GetEnvAsInt("TEST_INT_VAR", 90)
	example.MyString = quickenv.GetEnvAsString("TEST_VAR", "not")
	example.MyFloat = quickenv.GetEnvAsFloat("TEST_FLOAT_VAR", 1.11)

	fmt.Println(example)
}
