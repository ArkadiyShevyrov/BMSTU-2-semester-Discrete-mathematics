package main

import (
	"fmt"
	"math"
)

type Fraction struct {
	x int
	y int
}

func main() {
	arrayN, n, err := myScan()
	if err != nil {
		fmt.Println("Error: myScan")
		return
	}

	SequentialExclusion(arrayN, n)
	/*	for _, fractions := range arrayN {
		fmt.Println(fractions)
	}*/
	res := make([]Fraction, n)
	if n > 3 {
		res = ReverseConvolutionFive(arrayN, n)
	} else {
		Transform(arrayN, n)
		res = ReverseConvolution(arrayN, n)
	}

	erorr := 0
	for i := 0; i < n; i++ {
		if res[i].y == 0 {
			erorr++
		}
	}
	if erorr == 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d/%d\n", res[i].x, res[i].y)
		}
	} else {
		fmt.Println("No solution")
	}
}

func myScan() ([][]Fraction, int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error: scan N")
		return make([][]Fraction, 0), 0, err
	}
	arrayN := make([][]Fraction, n)
	for index, _ := range arrayN {
		arrayN[index] = make([]Fraction, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			var num int
			_, err := fmt.Scan(&num)
			if err != nil {
				fmt.Println("Error: scan numbers")
				return make([][]Fraction, 0), 0, err
			}
			arrayN[i][j] = Fraction{x: num, y: 1}
		}
	}
	return arrayN, n, err
}

func ReverseConvolutionFive(arrayN [][]Fraction, n int) []Fraction {
	res := make([]Fraction, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arrayN[i][j].x != 0 {
				return res
			}
		}
	}
	for i := 0; i < n; i++ {
		if arrayN[i][i].x == 0 {
			return res
		}
	}
	for i := 1; i <= n; i++ {
		res[n-i] = arrayN[n-i][n]
		for j := 1; j < i; j++ {
			res[n-i] = difference(res[n-i], multiplication(arrayN[n-i][n-j], res[n-j]))
			//fmt.Println(res[n-i], arrayN[n-i][n-j], res[n-j])
		}
	}
	return res
}

func SequentialExclusion(arrayN [][]Fraction, n int) [][]Fraction {
	for k := 0; k < n; k++ {
		for i := k; i < n; i++ {
			divider := arrayN[i][k]
			if divider.x != 0 {
				for j := k; j < n+1; j++ {
					arrayN[i][j] = division(arrayN[i][j], divider)
				}
			}
		}
		//обмен

		if arrayN[k][k].x == 0 {
			for i := k + 1; i < n; i++ {
				if arrayN[i][k].x != 0 {
					for j := 0; j < n+1; j++ {
						temp := arrayN[k][j]
						arrayN[k][j] = arrayN[i][j]
						arrayN[i][j] = temp
					}
					break
				}
			}
		}
		for i := k + 1; i < n; i++ {
			if arrayN[i][k].x != 0 {
				for j := k; j < n+1; j++ {
					arrayN[i][j] = difference(arrayN[i][j], arrayN[k][j])
				}
			}
		}
	}
	return arrayN
}

func difference(a, b Fraction) Fraction {
	var c Fraction
	c.x = a.x*b.y - b.x*a.y
	c.y = a.y * b.y
	if c.x == 0 {
		c.y = 1
	}
	if c.x < 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x > 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x == 0 {
		return c
	}
	minNumber := GCDEuclidean(int(math.Abs(float64(c.x))), int(math.Abs(float64(c.y))))
	for value := minNumber; value >= 2; value-- {
		if c.x%value == 0 && c.y%value == 0 {
			c.x /= value
			c.y /= value
		}
	}
	return c
}

func subtraction(a, b Fraction) Fraction {
	var c Fraction
	c.x = a.x*b.y + b.x*a.y
	c.y = a.y * b.y
	if c.x == 0 {
		c.y = 1
	}
	if c.x < 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x > 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x == 0 {
		return c
	}
	minNumber := GCDEuclidean(int(math.Abs(float64(c.x))), int(math.Abs(float64(c.y))))
	for value := minNumber; value >= 2; value-- {
		if c.x%value == 0 && c.y%value == 0 {
			c.x /= value
			c.y /= value
		}
	}
	return c
}

func division(a, b Fraction) Fraction {
	var c Fraction
	c.x = a.x * b.y
	c.y = a.y * b.x
	if c.x == 0 {
		c.y = 1
	}
	if c.x == c.y {
		c.x = 1
		c.y = 1
	}
	if c.x < 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x > 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}

	if c.x == 0 {
		return c
	}
	minNumber := GCDEuclidean(int(math.Abs(float64(c.x))), int(math.Abs(float64(c.y))))
	for value := minNumber; value >= 2; value-- {
		if c.x%value == 0 && c.y%value == 0 {
			c.x /= value
			c.y /= value
		}
	}
	return c
}

func multiplication(a, b Fraction) Fraction {
	var c Fraction
	c.x = a.x * b.x
	c.y = a.y * b.y
	if c.x == c.y {
		c.x = 1
		c.y = 1
	}
	if c.x < 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x > 0 && c.y < 0 {
		c.x *= -1
		c.y *= -1
	}
	if c.x == 0 {
		return c
	}
	minNumber := GCDEuclidean(int(math.Abs(float64(c.x))), int(math.Abs(float64(c.y))))
	for value := minNumber; value >= 2; value-- {
		if c.x%value == 0 && c.y%value == 0 {
			c.x /= value
			c.y /= value
		}
	}
	return c
}
func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func Transform(arrayN [][]Fraction, n int) [][]Fraction {
	for i := 0; i < n; i++ {
		max := 1
		for j := 0; j < n; j++ {
			if int(math.Abs(float64(arrayN[i][j].y))) > max {
				max = arrayN[i][j].y
			}
		}
		tempMulti := Fraction{max, 1}
		for j := 0; j < n+1; j++ {
			arrayN[i][j] = multiplication(arrayN[i][j], tempMulti)
		}
	}
	return arrayN
}

func ReverseConvolution(arrayN [][]Fraction, n int) []Fraction {
	res := make([]Fraction, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arrayN[i][j].x != 0 {
				return res
			}
		}
	}
	for i := 0; i < n; i++ {
		if arrayN[i][i].x == 0 {
			return res
		}
	}
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			for arrayN[i-1-j][i].x != 0 {
				if arrayN[i-1-j][i].x*arrayN[i-1-j][i].y > 0 {
					for k := 0; k < n+1; k++ {
						arrayN[i-1-j][k] =
							difference(arrayN[i-1-j][k], arrayN[i][k])
					}
				} else {
					for k := 0; k < n+1; k++ {
						arrayN[i-1-j][k] =
							subtraction(arrayN[i-1-j][k], arrayN[i][k])
					}
				}
			}
		}
		divider := arrayN[i-1][i-1]
		for j := 0; j < n+1; j++ {
			arrayN[i-1][j] = division(arrayN[i-1][j], divider)
		}
	}
	for i := 0; i < n; i++ {
		res[i] = arrayN[i][n]
	}
	return res
}

/*
3
2 4 1 36
5 2 1 47
2 3 4 37

*/

/*
3
-4  -1   8   2
 7  -7   7   3
 5  -1  -4   7

*/
