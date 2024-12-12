package main

import "fmt"

type (
	// inheritance	HunmanClass struct{}

	HunmanClass struct{}

	Wanpire struct{ HunmanClass }  

	Devel struct{ HunmanClass }
)

func (h HunmanClass) Walk() {
	fmt.Println("human walk")
}

func (h HunmanClass) Eat() {
	fmt.Println("human eat")
}

func (h HunmanClass) Breath() {
	fmt.Println("human Breath")
}

func (w Wanpire) KinBlood() {
	fmt.Println("kinblod human")
}

func (w Wanpire) Ply() {
	fmt.Println("ply..........")
}

func (w Devel) Killed() {
	fmt.Println("killed ")
}
