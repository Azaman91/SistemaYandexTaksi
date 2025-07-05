package main

import (
	"fmt"
)

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name    string
	species string
	age     int
	sound   string
}

func (a animal) MakeSound() string {
	return a.sound
}

func (a animal) GetName() string {
	return a.name
}

func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d\n", a.name, a.species, a.age)
}

func NewAnimal(name, species string, age int, sound string) animal {
	return animal{name: name, species: species, age: age, sound: sound}
}

func ZooShow(animals []Animal) {
	for _, num := range animals {
		fmt.Print(num.GetInfo())
		fmt.Println(num.MakeSound())
	}
}

type ZooKeeper struct {
	animal
}

func (z ZooKeeper) Feed(a Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!.", a.GetName(), a.MakeSound())
}

func main() {
	tiger := NewAnimal("Барсик", "Тигр", 5, "Ррр")
	penguin := NewAnimal("Пиня", "Пингвин", 2, "Кря")
	ZooShow([]Animal{tiger, penguin})

	keeper := ZooKeeper{}
	keeper.Feed(tiger)
}
