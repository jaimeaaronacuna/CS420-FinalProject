package main

import "fmt"

type C1 struct {
	X int
	Y int
}

func (c *C1) SetX1(v int) { c.X = v }
func (c *C1) SetY1(v int) { c.Y = v }
func (c *C1) GetX1() int  { return c.X }
func (c *C1) GetY1() int  { return c.Y }

type C2 struct {
	C1
	Y int
}

func (c *C2) SetY2(v int) { c.Y = v }
func (c *C2) GetX2() int  { return c.X }
func (c *C2) GetY2() int  { return c.Y }

func main() {

	var o2 C2

	o2.SetX1(101)
	o2.SetY1(102)

	o2.SetY2(999)

	fmt.Println(o2.GetX1())
	fmt.Println(o2.GetY1())
	fmt.Println(o2.GetX2())
	fmt.Println(o2.GetY2())
}

