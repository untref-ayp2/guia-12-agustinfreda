package main

import (
	"fmt"

	"guia12/ejercicios"
)

func arbolEjerciciosDel1Al5(){
// ARBOL EJERCICIOS DEL 1 AL 5
	// raiz
	tree:=ejercicios.NewBinaryTree(2)

	// arbol izq de raiz
	tree7:=ejercicios.NewBinaryTree(7)
	
	tree6:=ejercicios.NewBinaryTree(6)
	tree6.InsertLeft(ejercicios.NewBinaryTree(5))
	tree6.InsertRight(ejercicios.NewBinaryTree(11))
	
	tree7.InsertLeft(ejercicios.NewBinaryTree(2))
	tree7.InsertRight(tree6)

	tree.InsertLeft(tree7)
	// arbol der de raiz
	tree5:=ejercicios.NewBinaryTree(5)
	tree9:=ejercicios.NewBinaryTree(9)
	tree4:=ejercicios.NewBinaryTree(4)
	tree9.InsertLeft(tree4)
	tree5.InsertRight(tree9)

	tree.InsertRight(tree5)

	fmt.Println("-----------PrintPostOrder----------")
	tree.PrintPostOrder() // 2 5 11 6 7 4 9 5 2
}


func main() {

	// btree1 := binarytree.NewBinaryTree("+")
	// btree1.InsertLeft(binarytree.NewBinaryTree("A"))
	// btree1.InsertRight(binarytree.NewBinaryTree("B"))

	// btree2 := binarytree.NewBinaryTree("*")
	// btree2.InsertLeft(binarytree.NewBinaryTree("C"))
	// btree2.InsertRight(binarytree.NewBinaryTree("D"))

	// btree3 := binarytree.NewBinaryTree("-")
	// btree3.InsertLeft(btree1)
	// btree3.InsertRight(btree2)

	// fmt.Println("-----------PrintInOrder-----------")
	// btree3.PrintInOrder()
	// fmt.Println("-----------PrintPreOrder----------")
	// btree3.PrintPreOrder()
	// fmt.Println("-----------PrintPostOrder----------")
	// btree3.PrintPostOrder()

	



	// ARBOL EJERCICIOS 6 AL 8

	tree:= ejercicios.NewBinaryTree(20)

	tree19:=ejercicios.NewBinaryTree(19)
	tree1:=ejercicios.NewBinaryTree(1)
	tree15:=ejercicios.NewBinaryTree(15)
	tree15.InsertLeft(ejercicios.NewBinaryTree(17))
	tree15.InsertRight(ejercicios.NewBinaryTree(11))
	tree1.InsertRight(tree15)
	tree19.InsertLeft(tree1)

	tree.InsertLeft(tree19)

	tree29:=ejercicios.NewBinaryTree(29)
	tree29.InsertLeft(ejercicios.NewBinaryTree(27))
	tree29.InsertRight(ejercicios.NewBinaryTree(54))
	tree.InsertRight(tree29)

	tree.PrintPostOrder()
}
