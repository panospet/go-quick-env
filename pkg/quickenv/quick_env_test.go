package quickenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	err := LoadDotEnvFile()
	assert.Nil(t, err)

	A := GetEnvAsString("STRING_VAR_A", "aaa")
	assert.Equal(t, "hello", A)

	B := GetEnvAsString("STRING_VAR_B", "bbb")
	assert.Equal(t, "my name is", B)

	C := GetEnvAsString("STRING_VAR_C", "ccc")
	assert.Equal(t, "john,doe", C)

	C2 := GetEnvAsString("STRING_VAR_NOT_EXISTING", "default")
	assert.Equal(t, "default", C2)

	D := GetEnvAsInt("INT_VAR", 666)
	assert.Equal(t, 123, D)

	D2 := GetEnvAsInt("INT_VAR_NOT_EXISTING", 666)
	assert.Equal(t, 666, D2)

	E := GetEnvAsFloat("FLOAT_VAR", 66.666)
	assert.Equal(t, 123.456, E)

	E2 := GetEnvAsFloat("FLOAT_VAR_NOT_EXISTING", 66.666)
	assert.Equal(t, 66.666, E2)

	F := GetEnvAsBool("BOOL_VAR_FALSE", true)
	assert.False(t, F)

	G := GetEnvAsBool("BOOL_VAR_TRUE", false)
	assert.True(t, G)
}
