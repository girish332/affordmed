package main

import (
	"fmt"
)

type product struct {
	id   string
	name string
}

type icart interface {
	get(id string) (product, int)
	getAll() ([]product, int)
	add(temp product)
	update(toBeUpdated product, updated product) int
	remove(temp product) int
}

type cart struct {
	products []product
}

func initializeCart() *cart {
	return &cart{make([]product, 0)}
}

func (c *cart) get(Pid string) (product, int) {

	if len(c.products) == 0 {
		return product{}, -1
	}
	for i, v := range c.products {
		if v.id == Pid {
			return v, i
		}
	}
	return product{}, -1
}

func (c *cart) getAll() ([]product, int) {

	if len(c.products) == 0 {
		emptyProd := []product{}
		return emptyProd, -1
	}

	return c.products, 1

}

func (c *cart) add(prod product) {

	(*c).products = append((*c).products, prod)
}

func (c *cart) update(toBeUpdated product, updated product) int {

	for _, v := range c.products {

		if v == toBeUpdated {
			v.id = updated.id
			v.name = updated.name
			return 1
		}
	}

	return -1
}

func (c *cart) remove(temp product) int {

	for i, v := range c.products {

		if v == temp {
			(*c).products = (*c).products[:i+copy((*c).products[i:], (*c).products[i+1:])]
			return 1

		}
	}

	return -1
}

func main() {

	c := initializeCart()
	prod1 := product{
		id:   "123",
		name: "Girish",
	}
	c.add(prod1)
	prodSlice1, flag := c.getAll()

	if flag == 1 {
		for _, v := range prodSlice1 {
			fmt.Println(v)
		}
	}

	flagR := c.remove(prod1)
	if flagR == -1 {
		fmt.Println("Item not present in the slice")
	} else {
		fmt.Println("Item successfully removed")
	}

	c.add(prod1)
	prod2 := product{
		id:   "124",
		name: "Not girish",
	}
	flagU := c.update(prod1, prod2)

	if flagU == 1 {
		fmt.Println("Successfully updated")
	} else {
		fmt.Println("Object not present in slice")
	}

	prod3, flagG := c.get("124")

	if flagG == 1 {
		fmt.Println("Your requested product is :")
		fmt.Println(prod3)
	} else {
		fmt.Println("Product not present in slice")
	}

	prodSlice2, flagGA := c.getAll()

	if flagGA == 1 {
		for _, v := range prodSlice2 {
			fmt.Println(v)
		}
	}

}
