package lab

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

func Calculate(a, b, x float64) float64 {
	y := math.Cbrt(a*x+b) / (math.Log(x) * math.Log(x))

	return y
}

func Enter(a, b float64, arr []float64) {
	for index, value := range arr {
		y := Calculate(a, b, value)
		fmt.Printf("X%d = %g\n", index, y)
	}
}

func TaskA(x_f, x_end, step float64) []float64 {
	list := []float64{}

	for i := x_f; i < x_end+step; i += step {
		list = append(list, i)
	}
	return list
}

func get_value(f string) []float64 {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, 64)
	constant := ""

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		constant += string(data[:n])
	}

	example := `\d\.\d+`
	re := regexp.MustCompile(example)
	values := re.FindAllString(constant, -1)
	list := make([]float64, 0, len(values))

	for _, value := range values {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, f)
	}

	defer file.Close()
	return list
}

func Show_lab8() {
	constants := []byte("a = 1.35\nb = 0.98")
	file, err := os.Create("input.txt")

	if err != nil {
		fmt.Println("Не удалось создать файл, ошибка:", err)
		return
	}

	defer file.Close()
	file.Write(constants)
	fmt.Println("Файл input.txt успешно создан")

	A := TaskA(1.14, 4.24, 0.62)
	var B []float64 = []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	list := get_value("input.txt")

	a := list[0]
	b := list[1]

	fmt.Println("Задача A:")
	Enter(a, b, A)

	fmt.Println("Задача B:")
	Enter(a, b, B)
}
