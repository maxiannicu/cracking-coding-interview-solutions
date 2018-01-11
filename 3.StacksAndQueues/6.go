package main

import (
	"container/list"
	"fmt"
)

const cat = 1
const dog = 2

type Shelter struct {
	list *list.List
}

type Animal struct {
	animalType int
	name       string
}

func NewShelter() *Shelter {
	return &Shelter{
		list: list.New(),
	}
}

func (this *Shelter) Enqueue(animalType int, name string) {
	this.list.PushBack(Animal{
		animalType: animalType,
		name:       name,
	})
}

func (this *Shelter) DequeueAny() (int, string) {
	first := this.list.Front()
	this.list.Remove(first)

	return first.Value.(Animal).animalType, first.Value.(Animal).name
}

func (this *Shelter) DequeueByType(animalType int) (bool, string) {
	first := this.list.Front()
	for first != nil && first.Value.(Animal).animalType != animalType {
		first = first.Next()
	}
	if first == nil {
		return false, ""
	}
	this.list.Remove(first)

	return true, first.Value.(Animal).name
}

func (this *Shelter) DequeueDog() (bool, string) {
	return this.DequeueByType(dog)
}

func (this *Shelter) DequeueCat() (bool, string) {
	return this.DequeueByType(cat)
}

func main() {
	shelter := NewShelter()
	shelter.Enqueue(dog, "dog a")
	shelter.Enqueue(cat, "cat a")
	shelter.Enqueue(dog, "dog b")
	shelter.Enqueue(cat, "cat b")
	shelter.Enqueue(cat, "cat c")

	fmt.Println(shelter.DequeueCat())
	fmt.Println(shelter.DequeueDog())
	fmt.Println(shelter.DequeueAny())
}
