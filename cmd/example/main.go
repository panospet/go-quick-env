package main

import (
	"fmt"
	goquickenv "github.com/panospet/go-quick-env"
)

type Example struct {
	MyBool   bool
	MyInt    int
	MyString string
	MyFloat  float64
}

func main() {
	if err := goquickenv.LoadDotEnvFile(); err != nil {
		panic(err)
	}

	var example Example

	example.MyBool = goquickenv.GetEnvAsBool("TEST_BOOL_VAR", true)
	example.MyInt = goquickenv.GetEnvAsInt("TEST_INT_VAR", 90)
	example.MyString = goquickenv.GetEnvAsString("TEST_VAR", "not")
	example.MyFloat = goquickenv.GetEnvAsFloat("TEST_FLOAT_VAR", 1.11)

	fmt.Println(example)
}
