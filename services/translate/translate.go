package translate

import (
	"backend/repositorys"
	trie "backend/services/translate/trie"
	"fmt"
	"log"
	"strings"
	"unicode"
)

type Token struct {
	Key    string
	Value  string
	Pos    int
	Weight int
}

var Vietphrase *trie.Trie
var Name *trie.Trie

// func isWhiteSpace(s string) bool {
// 	return len(strings.TrimSpace(s)) == 0
// }

func toHalfWidth(char rune) rune {
	fullWidthPunctuation := map[rune]rune{
		'，': ',', '。': '.', '！': '!', '？': '?', '：': ':', '；': ';',
		'“': '"', '”': '"', '（': '(', '）': ')', '【': '[', '】': ']',
		'《': '<', '》': '>', '〈': '<', '〉': '>', '、': ',', '『': '"', '』': '"',
	}
	if val, ok := fullWidthPunctuation[char]; ok {
		return val
	}
	return char
}

func isPunctuation(char string) bool {
	return strings.Contains(".,■?;：<[「!“", char)
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	s = strings.Trim(s, " ")
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func InitDictionary() {
	Vietphrase = trie.NewTrie()
	if err := trie.LoadDictionaryToTrie("VietPhrase.txt", Vietphrase); err != nil {
		log.Println(err)
	}
	Name = trie.NewTrie()
	if err := trie.LoadDictionaryToTrie("Names.txt", Name); err != nil {
		fmt.Println(err)
	}
}

func getValueOfCharAtRoot(char rune) string {
	if childNode, ok := Vietphrase.Root.Children[char]; ok {
		return childNode.Value
	}
	return string(char)
}

func TranslateText(input string) string {
	lines := strings.Split(input, "\n")
	var result strings.Builder

	for _, line := range lines {
		tokenArr := tokenize(line)
		for _, token := range tokenArr {
			line = strings.ReplaceAll(line, token.Key, token.Value+" ")
		}
		result.WriteString(line + "\n")
	}
	return strings.ReplaceAll(result.String(), "的", "")
}

func processToken(root *trie.Node, pos int, weightFactor int, runes []rune) []Token {
	i := 0
	var tokenArr []Token

	if root.Children[runes[pos]] == nil {
		tokenArr = append(tokenArr, Token{
			Key:    string(runes[pos]),
			Value:  getValueOfCharAtRoot(toHalfWidth(runes[pos])),
			Pos:    pos,
			Weight: weightFactor * (i + 1),
		})
		return tokenArr
	}
	for pos+i < len(runes) && root.Children[runes[pos+i]] != nil {
		root = root.Children[runes[pos+i]]
		if root != nil && root.Value != "" {
			tokenArr = append(tokenArr, Token{
				Key:    string(runes[pos : pos+i+1]),
				Value:  root.Value,
				Pos:    pos,
				Weight: weightFactor + i,
			})
		}
		i++
	}
	return tokenArr
}

func tokenize(line string) []Token {
	runes := []rune(line)
	length := len(runes)

	bestChoices := make([]Token, length)
	bestWeights := make([]int, length+1)

	choices := make([][]Token, length)

	for i := range bestWeights {
		bestWeights[i] = -1
	}
	bestWeights[length] = 0

	for i := length - 1; i >= 0; i-- {
		tokensFromName := processToken(Name.Root, i, 5, runes)
		tokensFromVietphrase := processToken(Vietphrase.Root, i, 2, runes)
		choices[i] = append(choices[i], tokensFromName...)
		choices[i] = append(choices[i], tokensFromVietphrase...)

		for _, token := range choices[i] {
			tokenWeight := token.Weight
			nextIndex := i + len([]rune(token.Key))
			if nextIndex <= length && (tokenWeight+bestWeights[nextIndex]) > bestWeights[i] {
				bestWeights[i] = tokenWeight + bestWeights[nextIndex]
				bestChoices[i] = token
			}
		}

		if bestWeights[i] == -1 {
			bestWeights[i] = 1
			bestChoices[i] = Token{
				Key:    string(runes[i]),
				Value:  getValueOfCharAtRoot(runes[i]),
				Pos:    i,
				Weight: 1,
			}
		}
	}

	var result []Token
	i := 0
	isPunc := false
	if length != 0 {
		bestChoices[0].Value = capitalizeFirst(bestChoices[0].Value)
	}
	for i < length {
		token := bestChoices[i]
		if isPunc {
			token.Value = " " + capitalizeFirst(token.Value)
		}
		result = append(result, token)
		if isPunctuation(token.Key) {
			isPunc = true
		} else {
			isPunc = false
		}
		i += len([]rune(token.Key))
	}
	//log.Println(result)
	return result
}
func TranslateChapter(chapterId string) (string, error) {
	input, err := repositorys.GetChapterContent(chapterId)
	if err != nil {
		return "", err
	}
	lines := strings.Split(input, "\n")
	var resString strings.Builder

	for _, line := range lines {
		tokenArr := tokenize(line)
		for _, token := range tokenArr {
			if token.Key != "" {
				//line = strings.Replace(line, token.Key, token.Value, 1)
				tokenData := fmt.Sprintf("%s  %s  %d", token.Value, token.Key, token.Pos)
				resString.WriteString(tokenData + "\t\t\t")
			}
		}
		resString.WriteString("\n")
	}
	return resString.String(), nil
}
