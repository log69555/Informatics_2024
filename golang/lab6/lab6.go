package lab6

import "fmt"

type Character struct {
	name  string
	class string
	level int
}

func (c *Character) change_class(new_class string) {
	(*c).class = new_class
}

func (c *Character) change_level(new_level int) {
	(*c).level = new_level
}

func (c *Character) new_characters(new_name, new_class string, new_level int) {
	(*c).name = new_name
	(*c).class = new_class
	(*c).level = new_level
}

func (c Character) Name() string {
	return c.name
}

func (c Character) Class() string {
	return c.class
}

func (c Character) Level() int {
	return c.level
}

func Show_lab6() {
	Hero := Character{
		name:  "Игорь",
		class: "Рыцарь",
		level: 4,
	}

	fmt.Println("Характеристики персонажа: имя - ", Hero.Name(), " класс - ", Hero.Class(), " уровень - ", Hero.Level())

	Hero.change_class("Волшебник")
	Hero.change_level(6)

	fmt.Println("================================================================================================")
	fmt.Println("Изменены класс и уровень с помощью методов change_class и change_level:")
	fmt.Println("Характеристики персонажа: имя - ", Hero.Name(), " класс - ", Hero.Class(), " уровень - ", Hero.Level())

	Hero.new_characters("Паша", "Целитель", 10)

	fmt.Println("================================================================================================")
	fmt.Println("Полностью изменем персонаж с помощью метода new_character:")
	fmt.Println("Характеристики персонажа: имя - ", Hero.Name(), " класс - ", Hero.Class(), " уровень - ", Hero.Level())
}
