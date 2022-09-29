package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d = 10
	d += 5
	fmt.Println(d)

	var e = 10
	e -= 5
	fmt.Println(e)

	var f = 10
	f *= 5
	fmt.Println(f)

	var g = 10
	g /= 5
	fmt.Println(g)

	var h = 10
	h %= 5
	fmt.Println(h)

	var i = 10
	i++
	fmt.Println(i)

	var j = 10
	j--
	fmt.Println(j)

	var k = 10
	k = -k
	fmt.Println(k)

	var l = -10
	l = +l
	fmt.Println(l)

	var m = true
	m = !m
	fmt.Println(m)

	n := !false
	fmt.Println(n)
}