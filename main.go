package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func wordCount(counts map[string]int, str string) map[string]int {
	wordList := strings.Fields(str)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {

	counts := make(map[string]int)
	strLine := os.Args[1]
	if strLine != "GoLang_Test.txt" {
		counts = wordCount(counts, strLine)
	} else {
		file, err := os.Open(strLine)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			counts = wordCount(counts, scanner.Text())
			// fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	p := make(PairList, len(counts))
	i := 0
	for k, v := range counts {
		p[i] = Pair{k, v}
		i++
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Value > p[j].Value
	})

	resultJson, err := json.Marshal(p[:10])
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(resultJson))
}
