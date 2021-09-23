package onecore

import (
	"fmt"
	"time"
)

// Gopher
type Gopher struct {
	Name       string
	Id         int
	CoffeeName string
}

var coffeeMachine = CoffeeMachine{}

func (this *Gopher) MakeCoffee(CoffeeName string) {
	if coffeeMachine.CoffeeName == "" {
		coffeeMachine.CoffeeName = CoffeeName
		coffeeMachine.Gopher = *this
		fmt.Println("Gopher", this.Id, "Make Coffee", coffeeMachine.CoffeeName)
		time.Sleep(10 * time.Second)
	}
	this.TakeCoffee()
	this.DrinkCoffee()
}

func (this *Gopher) TakeCoffee() {
	if coffeeMachine.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Take Coffee", coffeeMachine.CoffeeName)
		this.CoffeeName = coffeeMachine.CoffeeName
		time.Sleep(5 * time.Second)

		coffeeMachine.CoffeeName = ""
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Take")
		this.CoffeeName = ""
	}
}
func (this *Gopher) DrinkCoffee() {
	if this.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Drink")
	}
}
