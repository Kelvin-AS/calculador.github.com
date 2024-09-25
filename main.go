package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 fmt.Println("Introduce el primer número:")
 firstInput, _ := reader.ReadString('\n')
 firstNum, err := strconv.ParseFloat(firstInput[:len(firstInput)-1], 64)
 if err != nil {
  fmt.Println("Por favor, ingresa un número válido.")
  return
 }

 fmt.Println("Introduce el segundo número:")
 secondInput, _ := reader.ReadString('\n')
 secondNum, err := strconv.ParseFloat(secondInput[:len(secondInput)-1], 64)
 if err != nil {
  fmt.Println("Por favor, ingresa un número válido.")
  return
 }

 sum := firstNum + secondNum
 subtraction := firstNum - secondNum
 multiplication := firstNum * secondNum

 var division float64
 if secondNum != 0 {
  division = firstNum / secondNum
 } else {
  fmt.Println("No se puede dividir por cero.")
 }

 fmt.Printf("Resultados:\n")
 fmt.Printf("Suma: %.2f\n", sum)
 fmt.Printf("Resta: %.2f\n", subtraction)
 fmt.Printf("Multiplicación: %.2f\n", multiplication)
 if secondNum != 0 {
  fmt.Printf("División: %.2f\n", division)
 }
}