package main

import "fmt"

type Human struct {
	age    int8
	height float32
	weight float32
}

func (h *Human) GetAge() int8 {
	return h.age
}

func (h *Human) GetHeight() float32 {
	return h.height
}

func (h *Human) GetWeight() float32 {
	return h.weight
}

func (h *Human) Eat() {
	fmt.Println("I eat!")
}

func (h *Human) Sleep() {
	fmt.Println("I sleep Zzz...")
}

func (h *Human) Code() {
	fmt.Println("I code B-)")
}

type Action struct {
	Human    // embedding(встраивание)
	Duration int64
}

func (a *Action) Sleep() {
	fmt.Printf("No sleep for %d\n", a.Duration)
}

func main() {
	h := Human{
		age:    18,
		height: 172.0,
		weight: 65.2,
	}

	fmt.Println("Human:")
	fmt.Println("Age:", h.GetAge(), "height:", h.GetHeight(), "weight", h.GetWeight())
	h.Eat()
	h.Sleep()
	h.Code()

	// Встроенное поле иницилизируется явным образом
	a := Action{
		Human: Human{
			age:    100,
			height: 0.0,
			weight: 1000.0,
		},
		Duration: 1000,
	}
	fmt.Println("\nAction:")
	// Встраивание позволяет обращаться к методам и полям встроенного поля напрямую при помощи оператора "точка"
	fmt.Println("Age:", a.GetAge(), "height:", a.GetHeight(), "weight", a.GetWeight())
	// Но при этом возможность обращаться к встроенному полю по названию типа также остается (a.GetAge() == a.Human.GetAge())
	_ = a.Human.GetAge()
	a.Eat()
	// Встраивание не запрещает создавать метод в новой структуре, который по названию и сигнатуре совпадает с методом встраиваемого поля
	a.Sleep()
	// При этом не отнимая возможности обращаться к методу встраиваемой структуры
	a.Human.Sleep()
	a.Code()
}
