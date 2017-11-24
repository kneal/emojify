package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp/syntax"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type GithubEmojis struct {
	Emojis map[string][]string
}

func (emojis *GithubEmojis) getGithubEmojis() *GithubEmojis {

	filename, _ := filepath.Abs("./resources/github_emojis.yml")

	yamlFile, err := ioutil.ReadFile(filename)
	check(err)

	err = yaml.Unmarshal(yamlFile, &emojis)
	check(err)

	return emojis
}

func checkIfEmojiExists(word string, emojis map[string][]string) bool {

	for _, key := range emojis {
		fmt.Printf("Checking key: %s \n", key)
		for _, emoji := range key {
			if word == emoji {
				fmt.Printf("%s == %s \n", word, emoji)
				return true
			} else {
				fmt.Printf("%s != %s \n", word, emoji)
			}
		}
	}
	return false
}

func main() {
	var githubEmojis GithubEmojis
	emojis := githubEmojis.getGithubEmojis().Emojis

	file, err := os.Open("./resources/test.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			if syntax.IsWordChar([]rune(word)[0]) {
				checkIfEmojiExists(word, emojis)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
