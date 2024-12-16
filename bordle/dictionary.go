package bordle

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Dictionary struct {
	words []string
}

// Function to create a new dictionary using
// list of words provided.
func NewDictionary(words []string) *Dictionary {
	d := &Dictionary{
		words: words,
	}
	return d
}

// Function to load dictionary
func LoadDictionary(filePath string) ([]string, error) {
	wordMap := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		wordMap = append(wordMap, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordMap, nil
}

// Function to lookup a guess (word from user) against the dictionary
// a positive hit will return the word.
func (d *Dictionary) Lookup(needle string) (string, error) {
	if d == nil {
		fmt.Println("Dic is nil, cannot access fields")
		return "", errors.New("Dic is nil, cannot access fields")
	}

	if len(d.words) == 0 {
		fmt.Println("Dic.words is empty")
		return "", errors.New("Dic.words is empty")
	}

	found := indexOf(d.words, needle)

	if found != -1 {
		return d.words[found], nil
	}
	return "", errors.New("word does not exist")
}

// Function to return the index of an element if
// found in a slice
func indexOf(slice []string, element string) int {
	for i, v := range slice {
		if strings.ToLower(v) == strings.ToLower(element) {
			return i
		}
	}
	return -1 // Return -1 if the element is not found
}

// Getter function to return words in the dictionary
func (d *Dictionary) GetWords() []string {
	return d.words
}
