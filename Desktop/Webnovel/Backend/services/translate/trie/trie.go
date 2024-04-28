package trie

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)

type Node struct {
	Children  map[rune]*Node
	EndOfNode bool
	Value     string
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: newNode(),
	}
}
func newNode() *Node {
	return &Node{
		Children:  make(map[rune]*Node),
		Value:     "",
		EndOfNode: false,
	}
}

func (trie *Trie) Insert(key, Value string) {
	node := trie.Root
	for _, char := range key {
		if node.Children[char] == nil {
			node.Children[char] = newNode()
		}
		node = node.Children[char]
	}
	node.EndOfNode = true
	node.Value = Value

}

// func (trie *Trie) Search(key string) string {
// 	node := trie.Root
// 	for _, char := range key {
// 		if node.Children[char] == nil {
// 			return ""
// 		}
// 		node = node.Children[char]
// 	}
// 	if node.EndOfNode {
// 		return node.Value
// 	}
// 	return ""
// }

func LoadDictionaryToTrie(filepath string, trie *Trie) error {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(filepath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "=")
		if len(tokens) == 2 {
			key := tokens[0]
			tk := strings.Split(tokens[1], "/")
			Value := tk[0]
			trie.Insert(key, Value)
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil

}

// func LoadNameToMap(filepath string) map[string]string {
// 	Map := make(map[string]string)
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		fmt.Println(filepath, err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		kv := strings.Split(line, "=")
// 		if len(kv) == 2 {
// 			key := kv[0]
// 			Value := kv[1]
// 			Map[key] = Value
// 		}
// 	}
// 	return Map
// }

// func (t *Trie) Translate(input []rune) string {
// 	var result strings.Builder
// 	for len(input) != 0 {
// 		Value, remaining := t.LongestMatching(input)
// 		result.WriteString(Value)
// 		input = remaining
// 	}
// 	return result.String()
// }

// func (t *Trie) LongestMatching(input []rune) (string, []rune) {
// 	node := t.Root
// 	var lastValue string
// 	lastPos := 0
// 	for i, r := range input {
// 		if nextNode, ok := node.Children[r]; ok {
// 			if nextNode.Value != "" {
// 				lastValue = nextNode.Value
// 				lastPos = i
// 			}
// 			node = nextNode
// 		} else {
// 			break
// 		}

// 	}
// 	if lastValue == "" {
// 		return string(input[0]), input[1:]
// 	}
// 	//fmt.Println(lastValue)
// 	return lastValue + " ", input[lastPos+1:]
// }
