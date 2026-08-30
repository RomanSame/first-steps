package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Print("watsup!")                                                                       // bez perenosa
	fmt.Println("eto main nahuy")                                                              // s perenosom
	fmt.Println("my dick is " + fmt.Sprint(rand.Intn(30)) + " cantimeters long")               //combo no sprint poka hz che eto
	fmt.Printf("nickname: %s from zero to ten - %v, ghoul: %v - 7\n", "samerism", 9.832, 1000) // formatirovanni

	//%s - stroka %d - celoe chislo %f - drobnoe %v - universal absolute typeshi %T - pokazat

	sdelal := 1
	fmt.Println(sdelal)
	fmt.Println(add(1337, 1448))
	a, b := swap("swop", "swip")
	fmt.Println(a, b)
}
