package fungi_test

import (
	"errors"
	"fmt"

	"github.com/bionichound/fungi"
)

func ExampleMap() {
	ints := []int{1, 2, 4}
	strs := []string{"hello", "friends"}
	strcts := []struct{ Name string }{
		{
			Name: "Bionic",
		},
		{
			Name: "Hound",
		},
	}

	intsPlusOne := fungi.Map(ints, func(i int) int { return i + 1 })
	lenOfStrings := fungi.Map(strs, func(i string) int { return len(i) })
	strctsNames := fungi.Map(strcts, func(i struct{ Name string }) string { return i.Name })

	fmt.Printf("Original ints: %v\n", ints)
	fmt.Printf("Ints plus one: %v\n", intsPlusOne)
	fmt.Printf("Original strings: %v\n", strs)
	fmt.Printf("Length of strings: %v\n", lenOfStrings)
	fmt.Printf("Original structs: %v\n", strcts)
	fmt.Printf("Names of each struct: %v\n", strctsNames)
	// Output:
	// Original ints: [1 2 4]
	// Ints plus one: [2 3 5]
	// Original strings: [hello friends]
	// Length of strings: [5 7]
	// Original structs: [{Bionic} {Hound}]
	// Names of each struct: [Bionic Hound]
}

func ExampleFold() {
	numbers := []int{1, 2, 3, 4}
	args := []string{"go", "help", "test"}

	add := func(a, b int) int {
		return a + b
	}

	noOfChars := func(item string, curr int) int {
		return curr + len(item)
	}

	fmt.Println(fungi.Fold(add, 0, numbers))
	fmt.Println(fungi.Fold(noOfChars, 0, args))
	// Output:
	// 10
	// 10
}

func ExampleFilter() {
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{
			name: "Bob",
			age:  8,
		},
		{
			name: "James",
			age:  41,
		},
		{
			name: "Jen",
			age:  17,
		},
		{
			name: "Alex",
			age:  18,
		},
		{
			name: "Alice",
			age:  15,
		},
	}

	over18 := func(p Person) bool {
		return p.age >= 18
	}

	fmt.Println(fungi.Filter(over18, people))
	// Output:
	// [{James 41} {Alex 18}]
}

func ExampleIncludes() {
	names := []string{"Aquaman", "Batman", "WonderWoman", "The Rock"}
	fmt.Println(fungi.Includes(names, "The Rock"))
	fmt.Println(fungi.Includes(names, "Robin"))
	// Output:
	// true
	// false
}

func ExampleNumbers() {
	fromZero := fungi.Numbers{}
	fmt.Println(fromZero.Next())
	fmt.Println(fromZero.Next())
	fmt.Println(fromZero.Next())
	fromTwenty := fungi.Numbers{
		Current: 20,
	}
	fmt.Println(fromTwenty.Next())
	fmt.Println(fromTwenty.Next())
	fmt.Println(fromTwenty.Next())
	// Output:
	// 0
	// 1
	// 2
	// 20
	// 21
	// 22
}

func ExampleTake() {
	fromZero := &fungi.Numbers{}
	fmt.Println(fungi.Take[*fungi.Numbers, int](fromZero, 10))
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleDo() {
	errorEither := &fungi.Either[int]{
		Left:  errors.New("oh no something went wrong"),
		Right: 0,
	}
	numberEither := &fungi.Either[int]{
		Left:  nil,
		Right: 15,
	}

	addOne := func(a int) int {
		return a + 1
	}

	fmt.Println(fungi.Do(errorEither, addOne))
	fmt.Println(fungi.Do(numberEither, addOne))
	// Output:
	// &{oh no something went wrong 0}
	// &{<nil> 16}
}

func ExampleBind() {
	errorEither := &fungi.Either[int]{
		Left:  errors.New("oh no something went wrong"),
		Right: 0,
	}
	numberEither := &fungi.Either[int]{
		Left:  nil,
		Right: 15,
	}
	failingCase := &fungi.Either[int]{
		Left:  nil,
		Right: 13,
	}

	possiblyFail := func(a int) (int, error) {
		if a == 13 {
			return 0, errors.New("that's an unlucky number")
		}
		return a + 22, nil
	}

	fmt.Println(fungi.Bind(errorEither, possiblyFail))
	fmt.Println(fungi.Bind(numberEither, possiblyFail))
	fmt.Println(fungi.Bind(failingCase, possiblyFail))
	// Output:
	// &{oh no something went wrong 0}
	// &{<nil> 37}
	// &{that's an unlucky number 0}
}

func ExampleResult() {
	div := func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("can't divide by 0")
		}
		return a / b, nil
	}

	fmt.Println(fungi.Result(div(4, 2)))
	fmt.Println(fungi.Result(div(1, 0)))
	// Output:
	// &{<nil> 2}
	// &{can't divide by 0 0}
}
