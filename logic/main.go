package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	anagramValue = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func Anagram() [][]string {
	list := make(map[string][]string)

	for _, word := range anagramValue {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	var resp [][]string
	for _, val := range list {
		var item []string
		for idx := range val {
			item = append(item, val[idx])
		}
		resp = append(resp, item)
	}

	return resp
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	fmt.Println(Anagram())
}
