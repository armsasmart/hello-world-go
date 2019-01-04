package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) getLastName() string {
	return p.lastName
}

func (p *Person) getFirstName() string {
	return p.firstName
}

func main() {
	a := 2
	b := 5
	if a == b {
		fmt.Println("a == b is true")
	} else {
		fmt.Println("a == b is false")
	}

	var arrayA [2]string
	arrayA[0] = "one"
	arrayA[1] = "two"
	fmt.Println(arrayA)

	colors := [3]string{"red", "green", "blue"}
	var s []string = colors[0:2]
	fmt.Println(s)

	var numbers = make([]string, 0)
	numbers = append(numbers, "ONE")
	numbers = append(numbers, "TWO")
	numbers = append(numbers, "THREE")
	numbers = append(numbers, "FOUR")
	numbers = append(numbers, "FIVE")
	numbers = append(numbers, "SIX")
	printSlice("numbers", numbers)

	fmt.Println(sum(a, b))

	p1 := Person{
		"Sirawit",
		"Arayatrakullikit",
		24,
	}
	fmt.Println(p1.getFirstName())
	fmt.Println(p1.getAge())
	fmt.Println(p1.getLastName())

	animals := map[string]string{
		"D1": "Dog",
		"D2": "Cat",
		"D3": "Chopper",
	}
	fmt.Println(animals)
	fmt.Println(animals["D1"])

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	x := 1
	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	default:
		fmt.Println("x unknow")
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) int {
	return a + b
}
