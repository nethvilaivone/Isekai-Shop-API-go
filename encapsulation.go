package main

import "fmt"


type (
	PlayerClass interface {
		Actack()
	}

	Worior struct {
		wepen string
	}

	Mage struct {
		spell string
	}
)

//  encapsulation 
func PlayerClassAtack(plyerclass PlayerClass) {
	plyerclass.Actack()
}

// abstraction 
func (m Worior) Actack() {
	fmt.Println(m.wepen)

}

func (m Mage) Actack() {
	fmt.Println(m.spell)
}
