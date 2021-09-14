package main

import "./entity"

func main() {
	name := "A6L"
	aoDi := entity.AoDi{}
	aoDi.Drive(name) // Drive AoDiA6L Car
	mycar := entity.MyCar{&aoDi}
	mycar.Drive(name) // Drive AoDiA6L Car
	name = "X6"
	mycar = entity.MyCar{&entity.BMW{}}
	mycar.Drive(name)        // Drive BMWX6 Car
	mycar.IDrive.Drive(name) // Drive BMWX6 Car
}
