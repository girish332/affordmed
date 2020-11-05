package main

import "fmt"

type icart struct {
	id   string
	name string
}

var cartSlice = make([]icart, 0)

func main() {

	c1 := icart{
		id:   "123",
		name: "abc",
	}

	cartSlice = append(cartSlice, c1)

	c2 := icart{
		id:   "345",
		name: "bcd",
	}

	c3 := icart{
		id:   "3",
		name: "Girish",
	}

	get(c1)
	add(c2)

	update(c2, c3)
	getAll()
	remove(c1)
	getAll()

}

func get(c icart) {

	for _, v := range cartSlice {
		if v == c {
			fmt.Println(v)
			return
		}
	}
	fmt.Println("Object not present in the slice")

}

func getAll() {

	for _, v := range cartSlice {
		fmt.Println(v)
	}
}

func add(c icart) {
	cartSlice = append(cartSlice, c)

}

func update(c1 icart, c2 icart) {

	for _, v := range cartSlice {
		if v == c1 {
			v.id = c2.id
			v.name = c2.name
			return
		}
	}
	fmt.Println("Object not present in the slice")

}

func remove(c icart) {

	for i, v := range cartSlice {

		if v == c {
			cartSlice = cartSlice[:i+copy(cartSlice[i:], cartSlice[i+1:])]
			fmt.Println("object removed is ", c)
			return

		}
	}

	fmt.Println("Object not present")

}
