# go-quick-env
### Quickly read variables from environment files

The best way to import environment variables to your code, is by using `.env` files. 

This library offers a quick way to do exactly that!

#### Example

```go
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

```

Enjoy ;)