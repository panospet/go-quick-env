# go-quick-env
### Quickly read variables from environment files

The best way to import environment variables to your code, is by using `.env` files. 

This library offers a quick way to do exactly that!

#### Example

Create a .env file like this:
```
TEST_BOOL_VAR=false
TEST_INT_VAR=666
TEST_VAR="hello"
TEST_FLOAT_VAR=89.98
```

And your go code should be like: 
```go
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

	fmt.Println(example) // should print {false 666 hello 89.98}
}


```

Enjoy ;)