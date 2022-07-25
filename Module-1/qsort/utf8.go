package main

import (
	"fmt"
)

func main() {
	rs := []rune{655360, 192}
	fmt.Println(encode(rs))
	fmt.Println([]byte(string(rs)))
	fmt.Println(decode(encode(rs)))
	fmt.Println(rs)
}

func encode(utf32 []rune) []byte {
	arrayByte := make([]byte, 0)
	for _, value := range utf32 {
		num := int(value)
		if num < 128 {
			arrayByte = append(arrayByte,
				array2ToByte(intToArray2(num), 0))
		} else if num < 2048 {
			arrayByte = append(arrayByte,
				array2ToByte(intToArray2(num)[6:], 192),
				array2ToByte(intToArray2(num)[0:6], 128))
		} else if num < 65536 {
			arrayByte = append(arrayByte,
				array2ToByte(intToArray2(num)[12:], 224),
				array2ToByte(intToArray2(num)[6:12], 128),
				array2ToByte(intToArray2(num)[0:6], 128))
		} else {
			if len(intToArray2(num)) > 17 {
				arrayByte = append(arrayByte,
					array2ToByte(intToArray2(num)[18:], 240),
					array2ToByte(intToArray2(num)[12:18], 128),
					array2ToByte(intToArray2(num)[6:12], 128),
					array2ToByte(intToArray2(num)[0:6], 128))
			} else {
				var temp = []int{0}
				arrayByte = append(arrayByte,
					array2ToByte(temp, 240),
					array2ToByte(intToArray2(num)[12:18], 128),
					array2ToByte(intToArray2(num)[6:12], 128),
					array2ToByte(intToArray2(num)[0:6], 128))
			}
		}
	}
	return arrayByte
}

func intToArray2(value int) []int {
	array1 := make([]int, 0)
	for num := value; num != 1; num /= 2 {
		array1 = append(array1, num%2)
	}
	array1 = append(array1, 1)
	return array1
}

func array2ToByte(value []int, res byte) byte {
	for index, value := range value {
		res += byte(value << index)
	}
	return res
}

func decode(utf8 []byte) []rune {
	arrayRune := make([]rune, 0)
	flag := 1
	tempArrayByte := make([]byte, 0)
	for _, value := range utf8 {
		num := int(value)
		if num < 128 {
			arrayRune = append(arrayRune, rune(num))
			continue
		}
		if flag == 1 {
			if num < 224 {
				flag = 2
				tempArrayByte = append(tempArrayByte, value)
			} else if num < 240 {
				flag = 3
				tempArrayByte = append(tempArrayByte, value)
			} else {
				flag = 4
				tempArrayByte = append(tempArrayByte, value)
			}
		} else {
			tempArrayByte = append(tempArrayByte, value)
			flag--
			if flag == 1 {
				lenTempArrayByte := len(tempArrayByte)
				tempArrayInt := make([]int, 0)
				if lenTempArrayByte < 3 {
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[1]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[0]))[0:5]...)
					arrayRune = append(arrayRune, array2ToRune(tempArrayInt))
				} else if lenTempArrayByte < 4 {
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[2]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[1]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[0]))[0:4]...)
					arrayRune = append(arrayRune, array2ToRune(tempArrayInt))
				} else {
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[3]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[2]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[1]))[0:6]...)
					tempArrayInt = append(tempArrayInt,
						intToArray2(int(tempArrayByte[0]))[0:3]...)
					arrayRune = append(arrayRune, array2ToRune(tempArrayInt))
				}
				tempArrayByte = make([]byte, 0)
			}

		}
	}
	return arrayRune
}

func array2ToRune(value []int) rune {
	res := 0
	for index, value := range value {
		res += int(value << index)
	}
	return rune(res)
}
