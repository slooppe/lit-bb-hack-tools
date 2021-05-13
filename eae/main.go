package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

//main
func main() {
	extractExtensions(ScanTargets())
}

//ScanInput return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return result
}

//extractExtensions
func extractExtensions(input []string) {
	set := make(map[string]int)
	for _, elem := range input {
		u, err := url.Parse(AddProtocol(elem))
		if err == nil {
			elem = u.Path
			fmt.Println(elem)
			i := strings.LastIndex(elem, ".")
			if i >= 0 {
				extension := elem[i:]
				_, exists := set[extension]
				if exists {
					set[extension] += 1
				} else {
					set[extension] = 1
				}
			}
		}
	}
	//sort reverse
	n := map[int][]string{}
	var a []int
	for k, v := range set {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("[ %d ] %s\n", k, s)
		}
	}
}

//AddProtocol
func AddProtocol(input string) string {
	if input[:7] != "http://" {
		return "http://" + input
	}
	return input
}