package quickenv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadDotEnvFile() error {
	return LoadFile(".env")
}

func LoadFile(filePath string) error {
	lines, err := readLines(filePath)
	if err != nil {
		return fmt.Errorf("cnnot load file: %s", err.Error())
	}

	for i := range lines {
		parts := strings.SplitN(lines[i], "=", 2)
		if len(parts) < 2 {
			return fmt.Errorf("bad syntax at line %d", i+1)
		}
		if err := os.Setenv(parts[0], parts[1]); err != nil {
			return fmt.Errorf("cannot set environment variable: %s", err)
		}
	}

	return nil
}

func GetEnvAsString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		value = strings.Trim(value, `'`)
		value = strings.Trim(value, `"`)
		return value
	}

	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnvAsString(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func GetEnvAsFloat(name string, defaultVal float64) float64 {
	valueStr := GetEnvAsString(name, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}

	return defaultVal
}

func GetEnvAsBool(name string, defaultVal bool) bool {
	valueStr := GetEnvAsString(name, "")
	fmt.Println(valueStr)
	if val, err := strconv.ParseBool(valueStr); err == nil {
		return val
	}

	return defaultVal
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
