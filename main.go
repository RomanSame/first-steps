package main

import (
	"fmt"
	"math/rand"
)

var c, python, java bool
var i, j int = 1, 2

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	BIG   = 1 << 60
	SMALL = BIG >> 99
)

func needInt(x int64) int64       { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
func main() {
	fmt.Print("watsup!")                                                                       // bez perenosa
	fmt.Println("eto main nahuy")                                                              // s perenosom
	fmt.Println("my dick is " + fmt.Sprint(rand.Intn(30)) + " cantimeters long")               //combo no sprint poka hz che eto
	fmt.Printf("nickname: %s from zero to ten - %v, ghoul: %v - 7\n", "samerism", 9.832, 1000) // formatirovanni

	//%s - stroka %d - celoe chislo %f - drobnoe %v - universal absolute typeshi %T - pokazat type

	sdelal := 1
	fmt.Println(sdelal)
	fmt.Println(add(1337, 1448))
	a, b := swap("swop", "swip")
	fmt.Println(a, b)
	fmt.Println(split(22)) // gde func split sum = to shto v skobkah
	var i int
	fmt.Println(i, c, python, java) //v go mojno nullable znacheniya delat tak kak mi ne ukazivaem znachenie peremennoi strokoi vishe
	var c, python, java = true, false, "che eshe nado"
	fmt.Println(i, j, c, python, java) // а вот это подъеб в сторону питона и джавы, мол в гошке можно указывать значения без типа и оставляя код чистым
	p := 3.14
	fmt.Printf("p door type %T\n", p)
	fmt.Println(needInt(SMALL))
	fmt.Print(needInt(BIG))
	fmt.Println(needFloat(SMALL))
	fmt.Println(needFloat(BIG))
}
