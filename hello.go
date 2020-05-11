package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Printf("hello, world\n")

	// var i int

	// for i = 1; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println("son değer : ", i)

	//arrays
	// var m [2]int
	// m[0] = 25

	// fmt.Printf("%#v\n", m)

	// n := [3]int{23, 24}

	// fmt.Printf("%#v\n", n)

	// z := [...]string{"foo", "bar"}
	// fmt.Printf("%#v\n", z)

	// fmt.Println("-----")

	// var a [2][3]string

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		a[i][j] = fmt.Sprintf("%d-%d", i+1, j+1)
	// 	}
	// }

	// fmt.Println(a)

	//slices
	// fmt.Println("...slices...")
	// nn := []int{1, 2, 3, 4}

	// fmt.Printf("%#v\n", nn)

	// chars := [6]string{"a", "b", "c", "d", "e", "f"}

	// var slice []string = chars[1:4]

	// fmt.Printf("%#v\n", slice)

	// chars := [6]string{"a", "b", "c", "d", "e", "f"}

	// var slice []string = chars[1:4]

	// chars[2] = "x"

	// fmt.Printf("%#v", slice)

	// s := make([]int, 2)

	// s[0] = 10
	// s[1] = 20

	// fmt.Printf("%#v:\n", s)

	// s = append(s, 30, 40, 50)

	// fmt.Printf("%#v:\n", s)

	// list := []string{"cat", "dog", "rabbit"}

	// // for key, value := range list {
	// // 	fmt.Println(key, value)
	// // }

	// for _, value := range list {
	// 	fmt.Println(value)
	// }

	// sonuc := toplam(5, 10, 15, 20, 40, 50, 60, 70)

	// fmt.Println(sonuc)

	// o1, o2 := (islem(5, 6))
	// fmt.Printf("o1 : %v, o2 : %v", o1, o2)

	// sonuc, err := bol(5, 0)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("sonuc : %v", sonuc)

	// r := Rect{
	// 	width:  5.0,
	// 	height: 5.0,
	// }

	// r.setWidth(2.0)
	// r.setHeight(3.0)

	// fmt.Printf("%+v", r)

	// var newRect Shaper = Rect{
	// 	width:  5.0,
	// 	height: 4.0,
	// }

	// fmt.Printf("\ntype of newRect is %T\n", newRect)
	// fmt.Printf("value of newRect is %v\n", newRect)
	// fmt.Printf("value of newRect area is %2.f\n\n", newRect.Area())

	newRect := Rect{width: 5.0, height: 4.0}
	newCircle := Circle{radius: 3.0}

	fmt.Printf("rect area is : %2.f\n\n", Calculator(newRect))
	fmt.Printf("circle area is : %2.f\n\n", Calculator(newCircle))

}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func Calculator(s Shaper) float64 {
	return s.Area()
}

type Shaper interface {
	Area() float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) setWidth(w float64) {
	r.width = w
}

func (r *Rect) setHeight(h float64) {
	r.height = h
}

func toplam(n ...int) int {
	t := 0

	for _, v := range n {
		t += v
	}

	return t
}

func islem(a int, b int) (int, float64) {
	x := a * b
	y := float64(a) * 0.25

	return x, y
}

func bol(sayi, bolen float64) (float64, error) {
	if bolen == 0 {
		return 0, errors.New("bölen 0 olamaz")
	}

	sonuc := sayi / bolen

	return sonuc, nil
}
