package main

import (
	"fmt"
	"strings"
)

func contains(arr []string, el string) bool {
	for _, e := range arr {
		if e == el {
			return true
		}
	}
	return false
}

func dfs(graph map[string][]string, node string, path []string) int {
	if node == "end" {
		return 1
	}

	sum := 0
	for _, neighbour := range graph[node] {
		if strings.ToUpper(neighbour) == neighbour || !contains(path, neighbour) {
			sum += dfs(graph, neighbour, append(path, node))
		}
	}
	return sum
}

func main() {
	graph := make(map[string][]string)

	for {
		var input string
		if _, err := fmt.Scanf("%s", &input); err != nil {
			break
		}
		splitInput := strings.Split(input, "-")
		a := splitInput[0]
		b := splitInput[1]

		if neighbours, ok := graph[a]; ok {
			graph[a] = append(neighbours, b)
		} else {
			graph[a] = []string{b}
		}

		if neighbours, ok := graph[b]; ok {
			graph[b] = append(neighbours, a)
		} else {
			graph[b] = []string{a}
		}
	}

	fmt.Println(dfs(graph, "start", []string{}))
}
