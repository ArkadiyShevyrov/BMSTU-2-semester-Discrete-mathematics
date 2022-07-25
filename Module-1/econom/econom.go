package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MyScanStr() string {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	return str
}

func main() {
	str := MyScanStr()
	res := econom(str)
	fmt.Println(res)
}

func econom(str string) int {
	dictionary := make([]string, 0)

	arrayIndex := arrayIndexesBrackets(str)
	for _, value := range arrayIndex {
		tempString := str[value[0] : value[1]+1]
		if !Contains(dictionary, tempString) {
			dictionary = append(dictionary, tempString)
		}
	}
	return len(dictionary)
}

func arrayIndexesBrackets(str string) [][]int {
	arrayIndexLeft := make([]int, 0)
	arrayIndexRight := make([]int, 0)
	for i, value := range str {
		if value == '(' {
			arrayIndexLeft = append(arrayIndexLeft, i)
		} else if value == ')' {
			arrayIndexRight = append(arrayIndexRight, i)
		}
	}
	lenAIB := len(arrayIndexRight)
	arrayIndexesBracket := make([][]int, lenAIB)
	for i := 0; i < lenAIB; i++ {
		arrayIndexesBracket[i] = make([]int, 2)
	}
	count := 0
	for i := lenAIB - 1; i >= 0; i-- {
		var j int
		for j = 0; j < lenAIB; j++ {
			if arrayIndexRight[j] > arrayIndexLeft[i] {
				break
			}
		}
		arrayIndexesBracket[count][0] = arrayIndexLeft[i]
		arrayIndexesBracket[count][1] = arrayIndexRight[j]
		count++
		arrayIndexRight[j] = -1
	}
	return arrayIndexesBracket
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
