package main

type Player struct {
	ID int
}

type Thing struct{
	frame int
	state *[]byte
	inputs map[*Player]int
}

func Simulate() {

}

func Revert() {

}