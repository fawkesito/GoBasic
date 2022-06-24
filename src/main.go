package main

import (
	"fmt"
)

// func main() {
// 	// Declaración de constantes
// 	const pi float64 = 3.141592653589793
// 	const pi2 = 3.141592653589793
// 	fmt.Println("pi:", pi)
// 	fmt.Println("pi2:", pi2)

// 	// Declaración de variables enteras

// 	base:=12
// 	var altura int = 10
// 	var area int = base * altura
// 	fmt.Println("area:", area)

// 	//Zero values

// 	var a int
// 	var b float64
// 	var c bool
// 	var d string

// 	fmt.Println(a, b, c, d)
// }

// func main() {
// 	//Matemáticas

// 	//suma
// 	x:=10
// 	y:=5

// 	result:=x+y
// 	fmt.Println("suma:", result)

// 	//resta
// 	result=x-y
// 	fmt.Println("resta:", result)

// 	//multiplicación
// 	result=x*y
// 	fmt.Println("multiplicación:", result)

// 	//división
// 	result=x/y
// 	fmt.Println("división:", result)

// 	//Módulo
// 	result=x%y
// 	fmt.Println("modulo:", result)

// 	//Incremento
// 	x++
// 	fmt.Println("incremento:", x)

// 	//Decremento
// 	x--
// 	fmt.Println("decremento:", x)

// }

// func main() {
// 	// Trapecio
// 	basemenor := 4
// 	basemayor := 8
// 	altura := 6

// 	area := ((basemenor + basemayor) * altura) / 2

// 	fmt.Println("area:", area)

// 	//Circulo
// 	const pi = 3.141592653589793
// 	const radio = 4
// 	areaCirculo := pi * (radio*radio)
// 	fmt.Println("Área círculo:", areaCirculo)
// }

// func main() {
// 	// Declaración de variables
// 	helloMessage := "Hello"
// 	worldMessage := "World"

// 	//Println
// 	fmt.Println(helloMessage, worldMessage)
// 	fmt.Println(helloMessage, worldMessage)

// 	//Printf
// 	nombre := "Juan"
// 	edad := 30
// 	fmt.Printf("%s tiene %d años\n", nombre, edad) // %s es para string, %d es para int y %f es para float buena practica
// 	fmt.Printf("%v tiene %v años\n", nombre, edad)

// 	//Sprintf
// 	message := fmt.Sprintf("%s tiene %d años", nombre, edad)
// 	fmt.Println(message)

// 	//Tipo de dato
// 	fmt.Printf("Hellomessage es de tipo %T\n", helloMessage)
// }

// func normalFunction(message string) {
// 	fmt.Println(message)
// }

// func tripleArgument(a, b, c int) {
// 	fmt.Println(a + b + c)
// }

// func returnValue(a int) int {
// 	return a + 1	// return es para regresar un valor
// }

// func doubleReturn(a int) (c, d int) {
// 	return a, a*2
// }

// func main() {
// 	normalFunction("Hello World")
// 	tripleArgument(1, 2, 3)
// 	value := returnValue(2)
// 	fmt.Println("Value:", value)
// 	value1, value2 := doubleReturn(2)
// 	fmt.Println("Value1:", value1)
// 	fmt.Println("Value2:", value2)
// }

// func rectangleArea(base, height int) int {
// 	return base * height
// }
// func circleArea(radius int) int {
// 	return 3 * radius * radius
// }

// func main() {
// 	fmt.Println("Hello World")
// 	fmt.Println("Area Rectangle:", rectangleArea(10, 5))
// 	fmt.Println("Area Circle:", circleArea(5))
// }

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(1, 2))
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func main() {

// 	// //for condicional
// 	// for i := 0; i <= 10; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// //for while
// 	// i := 0
// 	// for i <= 10 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	// //for forever
// 	// iForever := 0
// 	// for {
// 	// 	fmt.Println(iForever)
// 	// 	iForever++
// 	// }

// 	//convertir texto a número
// 	value, err := strconv.Atoi("53")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Value is:", value)
// 	userPassword("fawkes", "ola")
// }
// func userPassword(userName, password string) bool {
// 	return userName == "fawkes" && password == "ola"
// }

// func main() {

// 	switch modulo := 4 % 2; modulo {
// 	case 0:
// 		fmt.Println("Es par")
// 	default:
// 		fmt.Println("Es impar")
// 	}

// 	//Sin condición
// 	value := 50
// 	switch {
// 	case value > 100:
// 		fmt.Println("Es mayor a 100")
// 	case value < 0:
// 		fmt.Println("Es menor a 0")
// 	default:
// 		fmt.Println("No condición")
// 	}
// }

// func main() {
// 	//Defer
// 	defer fmt.Println("Hola") // Ejecuta todo el codigo debajo y la deja para el final
// 	fmt.Println("Mundo")

// 	//Continue y break

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)

// 		//Continue
// 		if i == 2 {
// 			fmt.Println("ES 2")
// 			continue
// 		}

// 		//Break
// 		if i == 8 {
// 			fmt.Println("Break")
// 			break
// 		}
// 	}
// }

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

func main() {
	fmt.Println(split(50))
}

// Outside a function, every statement begins with a keyword, so := is not available outside a function