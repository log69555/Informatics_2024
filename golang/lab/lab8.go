package lab

import (
	"fmt"
	"os"
	"path/filepath"

	"isuct.ru/informatics2022/lab4"
)

func Show_lab8() {
	halh_way, err := filepath.Abs(filepath.Dir(filepath.Join(".", "lab8.go")))
	if err != nil {
		fmt.Println(err)
	}

	full_way := filepath.Join(halh_way, "input.txt")

	var constants []byte
	file, err := os.Create(full_way)

	input := input_value()

	constants = []byte(input)

	if err != nil {
		fmt.Println("Не удалось создать файл, ошибка:", err)
		return
	}

	defer file.Close()
	file.Write(constants)
	fmt.Println("Файл input.txt успешно создан")

	var A []float64
	var B []float64

	list := get_value("input.txt")

	a := list[0]
	b := list[1]

	A = lab4.TaskA(list[2], list[3], list[4])
	B = list[5:]

	fmt.Println("Задача A:")
	lab4.Enter(a, b, A)

	fmt.Println("Задача B:")
	lab4.Enter(a, b, B)
}
