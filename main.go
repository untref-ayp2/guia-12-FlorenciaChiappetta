package main

import (
	"fmt"
	"guia12/ejercicios"
)

func main() {

	btree2 := ejercicios.NewBinaryTree(4)
	a := ejercicios.NewBinaryTree(5)
	b := ejercicios.NewBinaryTree(15)
	c := ejercicios.NewBinaryTree(22)
	d := ejercicios.NewBinaryTree(29)
	c.InsertLeft(d)
	b.InsertLeft(c)
	a.InsertLeft(b)
	btree2.InsertLeft(a)
	btree2.InsertRight(ejercicios.NewBinaryTree(25))
	fmt.Println("-----------PrintInOrder-----------")
	btree2.PrintInOrder()
	fmt.Println("-----------PrintPreOrder----------")
	btree2.PrintPreOrder()
	fmt.Println("-----------PrintPostOrder----------")
	btree2.PrintPostOrder()
	fmt.Println("-----------SumarNodos----------")
	fmt.Println(btree2.SumarNodos())
	fmt.Println("-----------SumarNodos2----------")
	fmt.Println(btree2.SumarNodos2())
	fmt.Println("-----------SumarHojas----------")
	fmt.Println(btree2.SumarHojas())
	fmt.Println("-----------SumarHojasPares----------")
	fmt.Println(btree2.SumarHojasPares())
	fmt.Println("-----------SumarHojasMayoresA6----------")
	fmt.Println(btree2.SumarNodosMayoresA6())
	fmt.Println("-----------MayorAltura----------")
	fmt.Println(btree2.MayorAltura())
	fmt.Println("-----------HojasIzquierdas----------")
	fmt.Println(btree2.SumaHojasIzquierdas())
}
