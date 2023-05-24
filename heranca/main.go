package main

import "fmt"

type Casa struct {
	Geladeira string
	Endereco string
	numero int
}

type Familia struct {
	Casa 
	Mae string
	Pai string
	Filho string
}

func main(){

	Casa := Casa {"nextel" , "RUa samuel morse" , 134}

	Familia := Familia{
		Casa : Casa,
		Mae : "aurea",
	Pai : "Julio",
	Filho : "Tiago",
	}
	fmt.Println(Familia.Casa.numero)
}