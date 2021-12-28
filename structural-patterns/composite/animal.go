package composite

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal // <- embed!
	Swim   func()
}
