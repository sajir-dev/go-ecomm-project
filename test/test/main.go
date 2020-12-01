package main

type animal struct {
	id     uint8
	name   string
	age    uint8
	veg    bool
	origin string
}

var t *animal = animal{32, "cow", 5, true, "India"}

func main() {
	a, e := update(3, "tiger")
	fmt.Printl(a, e)
}
