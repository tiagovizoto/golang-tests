package main

import "fmt"

func main(){
  //variabel declarada
  var x string = "Variavel tipo string"
  fmt.Println(x)

  var y string

  y = "Variavel tipo string 2"
  fmt.Println(y)

  //Contenação de strings
  y = y + x
  fmt.Println(y)
  //Contenação de strings
  y += y
  fmt.Println(y)

  //Algums exemplos de booleanos de strings
  fmt.Println(y==x)
  fmt.Println(y!=x)
  fmt.Println(y>x)

  //Criar um variabel sem mecionar o var e o tipo
  z := "Variabel do tipo string"
  fmt.Println(z)

  //Constantes
  const souUmVelinho int64 = 301
  fmt.Println(souUmVelinho)

  var (
    aramis = "aramis"
    porthus = "porthus"
    athos = "athos"
  )
   fmt.Println(aramis + ", " + porthus + " e " + athos)



}
