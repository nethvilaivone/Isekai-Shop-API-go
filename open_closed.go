package main

import "fmt"

type Animals interface {
	Speak()
}

type Dog struct{}
type MockDog struct{}

func NewDog() Animals {
	return Dog{}
}
func (d Dog) Speak() {

	fmt.Println("Speak")
}

func NewMockDog() Animals {
	return MockDog{}
}

func (md MockDog) Speak() {
	fmt.Println("Speak")
}
func SpeakAnymals(an Animals){
  an.Speak()
}