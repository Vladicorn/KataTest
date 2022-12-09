package main

import (
	"fmt"
	"strconv"
)

var numbersArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
} /*
var numbersArabicCont = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
	11: "XI",
	12: "XII",
	13: "XIII",
	14: "XIV",
	15: "XV",
	16: "XVI",
	17: "XVII",
	18: "XVIII",
	19: "XIV",
	20: "XX",
}
*/
var numbersArabicCont = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func main() {

	var key0, key1, key2, key3 string
	fmt.Scanln(&key0, &key1, &key2, &key3)

	if key1 != "" {
		if key3 == "" {
			Arithmetic(key0, key1, key2)
		} else {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
	} else {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
	}

}

// Arithmetic вводная функция
func Arithmetic(key0, key1, key2 string) {

	val1, ok1 := numbersArabic[key0]
	val2, ok2 := numbersArabic[key2]

	if ok1 && ok2 {
		result, err := RomeMath(val1, val2, key1)
		if err != "" {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}

	} else if (ok1 && !ok2) || (!ok1 && ok2) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else {
		num1, _ := strconv.Atoi(key0)
		num2, _ := strconv.Atoi(key2)
		if num1 <= 10 && num2 <= 10 {
			result, _ := ArabicMath(num1, num2, key1)
			fmt.Println(result)
		} else {
			fmt.Println("Вывод ошибки, так как числа больше 10.")
		}
	}
}

// RomeMath Функция для операций с римскими числами
func RomeMath(a, b int, operation string) (string, string) {
	var result int

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	}

	if result < 1 {
		return "", "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	}

	return RomeNubmer(result), ""
}

// ArabicMath Функция для операций с арабскими числами
func ArabicMath(a, b int, operation string) (int, string) {
	var result int

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	}

	return result, ""
}

func RomeNubmer(result int) string {
	if result <= 10 {
		return numbersArabicCont[result]
	}
	if result <= 20 {
		result = result - 10
		return "X" + numbersArabicCont[result]
	}
	if result <= 30 {
		result = result - 20
		return "XX" + numbersArabicCont[result]
	}
	if result <= 40 {
		if result == 40 {
			return "XL"
		}
		result = result - 30
		return "XXX" + numbersArabicCont[result]
	}
	if result <= 50 {
		if result == 50 {
			return "L"
		}
		result = result - 40
		return "XL" + numbersArabicCont[result]
	}
	if result <= 60 {

		result = result - 50
		return "L" + numbersArabicCont[result]
	}
	if result <= 70 {
		result = result - 60
		return "LX" + numbersArabicCont[result]
	}
	if result <= 80 {
		result = result - 70
		return "LXX" + numbersArabicCont[result]
	}
	if result <= 90 {
		if result == 90 {
			return "XC"
		}
		result = result - 80
		return "LXXX" + numbersArabicCont[result]
	}
	if result <= 100 {
		if result == 100 {
			return "C"
		}
		result = result - 90
		return "XC" + numbersArabicCont[result]
	}
	return ""
}
