package main

import(
	"fmt"
)

func powerSeries(a int) (int, int, int, int) {
	return a*a, a*a*a, a*a*a*a, a*a*a*a*a
}

func main () {

	var square int
	var cube int
	var quartic int
        var quintic int

	square,cube,quartic,quintic = powerSeries(5) 
	fmt.Println("square:",square, "cube:",cube, "quartic:", quartic, "quintic:",quintic)
}
