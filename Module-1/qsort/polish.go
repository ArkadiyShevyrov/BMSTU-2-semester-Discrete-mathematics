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
	res := polish(PolishToArray(str))
	fmt.Println(res)
}

func polish(array []int) int {
	stack := make([]int, 1000)
	lenStack := 0
	for i := len(array) - 1; i >= 0; i-- {
		value := array[i]
		if value < 58 && value > 47 {
			stack[lenStack] = value - 48
			lenStack++
		} else if value == 43 {
			lenStack--
			stack[lenStack-1] = stack[lenStack-1] + stack[lenStack]
		} else if value == 45 {
			lenStack--
			stack[lenStack-1] = stack[lenStack] - stack[lenStack-1]
		} else if value == 42 {
			lenStack--
			stack[lenStack-1] = stack[lenStack-1] * stack[lenStack]
		}
	}
	return stack[0]
}

func PresenceInDictionary(value int) bool {
	Dictionary := [13]int{'0', '1', '2', '3', '4', '5',
		'6', '7', '8', '9', '+', '-', '*'}
	for _, x := range Dictionary {
		if value == x {
			return true
		}
	}
	return false
}

func PolishToArray(str string) []int {
	res := make([]int, 0)
	for _, value := range str {
		if PresenceInDictionary(int(value)) {
			res = append(res, int(value))
		}
	}
	return res
}
