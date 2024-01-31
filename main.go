package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arab() {

	fmt.Println("Введите пример(a + b, a - b, a * b, a / b.)")
	fmt.Println("Используйте цифры от 1 до 10, либо от I до X ")

	var x int = 0
	var z string = ""
	var y int = 0
	fmt.Scan(&x, &z, &y)

	switch {
	case z == "+":
		fmt.Println(fmt.Sprint(x + y))

	case z == "-":
		fmt.Println(fmt.Sprint(x - y))

	case z == "*":
		fmt.Println(fmt.Sprint(x * y))

	case z == "/":
		fmt.Println(fmt.Sprint(x / y))
	}

}

var romnum = map[string]int{
	"C":      100,
	"XC":     90,
	"LXXXI":  81,
	"LXXX":   80,
	"LXXII":  72,
	"LXX":    70,
	"LXIV":   64,
	"LXIII":  63,
	"LX":     60,
	"LVI":    56,
	"LIV":    54,
	"L":      50,
	"XLVIII": 48,
	"XLV":    45,
	"XLI":    42,
	"XL":     40,
	"XXXVI":  36,
	"XXXV":   35,
	"XXXIV":  34,
	"XXXII":  32,
	"XXX":    30,
	"XXVIII": 28,
	"XXVII":  27,
	"XXV":    25,
	"XXIV":   24,
	"XXIII":  23,
	"XXII":   22,
	"XXI":    21,
	"XX":     20,
	"XIX":    19,
	"XVIII":  18,
	"XVII":   17,
	"XVI":    16,
	"XV":     15,
	"XIV":    14,
	"XIII":   13,
	"XII":    12,
	"XI":     11,
	"X":      10,
	"IX":     9,
	"VIII":   8,
	"VII":    7,
	"VI":     6,
	"V":      5,
	"IV":     4,
	"III":    3,
	"II":     2,
	"I":      1,
}

var romnumInt = [48]int{
	100,
	90,
	81,
	80,
	72,
	70,
	64,
	63,
	60,
	56,
	54,
	50,
	48,
	45,
	42,
	40,
	36,
	35,
	34,
	32,
	30,
	28,
	27,
	25,
	24,
	23,
	22,
	21,
	20,
	19,
	18,
	17,
	16,
	15,
	14,
	13,
	12,
	11,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

const (
	Panic = "Error: Ошибочка :("
)

var a, b *int
var romInt = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"*": func() int { return *a * *b },
	"/": func() int { return *a / *b },
}

var abc []string

func base(r string) {
	var rom string
	var strf int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romanInt := make([]int, 0)
	for idx := range romInt {
		for _, val := range r {
			if idx == string(val) {
				rom += idx
				abc = strings.Split(r, rom)
			}
		}
	}

	switch {
	case len(rom) > 1:
		panic(Panic)
	case len(rom) < 1:
		panic(Panic)
	}
	for _, elem := range abc {
		num, err := strconv.Atoi(elem)
		if err != nil {
			strf++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch strf {
	case 1:
		panic(Panic)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 && numbers[1] > 0 && numbers[1] < 11
		if val, ok := romInt[rom]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(Panic)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := romnum[elem]; ok && val > 0 && val < 11 {
				romanInt = append(romanInt, val)
			} else {
				panic(Panic)
			}
		}
		if val, ok := romInt[rom]; ok {
			a, b = &romanInt[0], &romanInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(Panic)
	} else if romanResult < 0 {
		panic(Panic)
	}
	for romanResult > 0 {
		for _, elem := range romnumInt {
			for i := elem; i <= romanResult; {
				for index, value := range romnum {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
func main() {
	fmt.Println("Введите пример(a + b, a - b, a * b, a / b.)")
	fmt.Println("Используйте цифры от 1 до 10, либо от I до X ")

	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		r := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(r)))
	}
}
