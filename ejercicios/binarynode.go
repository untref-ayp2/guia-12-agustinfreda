package ejercicios

import (
	"fmt"
	"math"
)

type BinaryNode[data int] struct {
	left  *BinaryNode[int]
	right *BinaryNode[int]
	data  int
}

func NewBinaryNode(data int, left *BinaryNode[int], right *BinaryNode[int]) *BinaryNode[int] {
	return &BinaryNode[int]{left: left, right: right, data: data}
}

func (n *BinaryNode[int]) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode[int]) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode[int]) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode[int]) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return int(size)
}

func (n *BinaryNode[int]) Height() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}

// ------------------- Ejercicios --------------------------


// Ejercicio 1
func (n *BinaryNode[int]) SumaDeNodos()int{
	suma:= n.data

	if n.left != nil{
		suma+=n.left.SumaDeNodos()
	}
	if n.right != nil{
		suma+=n.right.SumaDeNodos()
	}
	return int(suma)
}


// Ejercicio 2
func (n *BinaryNode[int]) SumaDeHojas()int{
	suma := 0	
	if n.Height() == 0{ // o left && right == nil
		suma += n.data
	}

	if n.left != nil{
		suma += n.left.SumaDeHojas()
	}
	if n.right != nil{
		suma += n.left.SumaDeHojas()
	}

	return int(suma)
}


// Ejercicio 4
func (n *BinaryNode[int]) SumaNodosPares()int{
	suma := 0

	if n.data % 2 == 0{
		suma += n.data
	}

	if n.left != nil{
		suma += n.left.SumaNodosPares()
	}
	if n.right != nil{
		suma += n.right.SumaNodosPares()
	}
	return int(suma)
}

// Ejercicio 5
func(n *BinaryNode[int]) SumaNodosMayoresA6()int{
	suma := 0

	if n.data >= 6{
		suma+= n.data
	}

	if n.left != nil{
		suma += n.left.SumaNodosMayoresA6()
	}
	if n.right != nil{
		suma += n.right.SumaNodosMayoresA6()
	}
	return int(suma)
}



// Ejercicio 6
func (n *BinaryNode[int]) CompararacionAlturasEntreSubarboles()int{

	if n.left.Height() > n.right.Height(){
		return int(n.left.Height())
	} else if n.right.Height() > n.left.Height(){
		return int(n.right.Height())
	}
		return n.Height()
}


// Ejercicio 7
func (n *BinaryNode[int]) SumaDeHojasIzquierdas()int{
	suma := n.data

	if n.left != nil && n.left.left == nil && n.left.right == nil {
		suma += n.left.data
	}

	if n.left != nil {
		suma += n.left.SumaDeHojasIzquierdas()
	}
	if n.right != nil {
		suma += n.right.SumaDeHojasIzquierdas()
	}

	return int(suma)
}


// Ejercicio 8
func (n *BinaryNode[int]) SumaNodosDerechosImpares()int{
	suma:=0

	if n.right != nil && n.right.data % 2 != 0{
		suma += n.right.data
	}
	if n.left != nil{
		suma+= n.left.SumaNodosDerechosImpares()
	}
	if n.right != nil{
		suma+= n.right.SumaNodosDerechosImpares()
	}

	return int(suma)
}