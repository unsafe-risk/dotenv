package dotenv

import (
	"bufio"
	"bytes"
	"os"
)

// Read reads the .env file and returns the values as a map.
func Read(path string) (map[string]string, error) {
	result := make(map[string]string)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	for {
		key, err := reader.ReadBytes('=')
		if err != nil {
			break
		}
		value, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		result[string(key[:len(key)-1])] = string(value[:len(value)-1])
	}
	return result, nil
}

// Apply reads the .env file and sets the values in the environment.
func Apply(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	for {
		key, err := reader.ReadBytes('=')
		if err != nil {
			break
		}
		value, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		os.Setenv(string(key[:len(key)-1]), string(value[:len(value)-1]))
	}
	return nil
}
