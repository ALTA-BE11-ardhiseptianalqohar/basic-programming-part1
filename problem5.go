
package main

import "fmt"

func main()  {
	var r float32 
	var T float32 
	var lp float32
	var pi float32 = 3.14
	fmt.Println("Masukan r: ")
	fmt.Scanln(&r)
	fmt.Println("Masukan T: ")
	fmt.Scanln(&T)

	lp = 2 * pi * r * (r + T)
	fmt.Println("Hasil penjualan", lp)
}