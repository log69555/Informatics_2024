package lab

import (
	"fmt"
	"os"
)

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
}
