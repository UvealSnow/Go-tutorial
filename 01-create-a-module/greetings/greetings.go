package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map that associates each of the names with its message.
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func ParseNames(names string) ([]string, error) {
	names = strings.TrimSpace(names)
	if names == "" {
		return nil, errors.New("no names given")
	}

	var cleanNames []string
	for _, name := range strings.Split(names, ",") {
		if cleaned := strings.TrimSpace(name); cleaned != "" {
			cleanNames = append(cleanNames, cleaned)
		}
	}

	if len(cleanNames) == 0 {
		return nil, errors.New("no valid names given")
	}

	return cleanNames, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Nice to see you, %v!",
		"Hullo, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
