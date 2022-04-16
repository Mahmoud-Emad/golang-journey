package main
import "fmt"


// Here we have a struct named Animals.
// The Animal struct has a field named Name of type string.
// The Animal struct has a field named Age of type int.
// The Animal struct has a field named Weight of type float64.
// The Animal struct has a field named Alive of type bool.

type Animals struct{
	Name string
	Age int
	Weight float64
	Alive bool
	belly float64
}
// So we created the Animals struct.
// The CreateNewAnimal method takes a string and an int as parameters.
// The CreateNewAnimal method returns a pointer to a new Animals struct.
// The CreateNewAnimal method has a receiver of type Animals.

func (a *Animals) CreateNewAnimal(name string, age int) *Animals{
	// Now we have a method named CreateNewAnimal.
	return &Animals{
		Name: name,
		Age: age,
		Weight: 0.0,
		Alive: true,
		belly: 0.0,
	}
}
// So we created the CreateNewAnimal method.
// The SetWeight method takes a float64 as a parameter.
// The SetWeight method has a receiver of type Animals.

func (a *Animals) SetWeight(weight float64){
	// Now we have a method named SetWeight.
	a.Weight = weight
}
// So we created the SetWeight method.
// The IsAlive method takes no parameters.
// The IsAlive method returns a bool.
// The IsAlive method has a receiver of type Animals.
func (a *Animals) IsAlive() bool{
	// Now we have a method named IsAlive.
	return a.Alive
}
// So we created the IsAlive method.
// The Kill method takes no parameters.
// The Kill method returns nothing.
// The Kill method has a receiver of type Animals.
func (a *Animals) Kill() bool {
	// Now we have a method named Kill.
	a.Alive = false
	fmt.Printf("The %s is dead :(\n", a.Name)
	return a.Alive
}
// So we created the Kill method.
// Now we have to implement Live method.
func (a *Animals) Live() string {
	// Now we have a method named Live.
	if a.IsAlive() {
		fmt.Printf("The %s is alive.\n", a.Name)
		a.belly += 0.1
		a.Weight += 1.0 * a.belly
	}
	return "Animal is alive"
}
// So we created the Live method.
// Now  we need to take an instance from this struct.
func Cat(a *Animals) *Animals {
	// Now we have a method named Cat.
	cat := a.CreateNewAnimal("Cat", 1)
	for age := a.Age; age < 7; age++ {
		cat.Live()
	}
	cat.Kill()
	return cat
}

func main() {
	// Now we have a method named main.
	// Now we have to create an instance of this struct.
	a := &Animals{}
	// Now we have an instance of this struct.
	cat := Cat(a)
	// Now we have an instance of this struct.
	fmt.Println(cat)
}