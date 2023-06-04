package tests

import (
	"guia12/ejercicios"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSumaDeHojas(t *testing.T){
	btree := ejercicios.NewBinaryTree(2)
	btree.InsertLeft(ejercicios.NewBinaryTree(3))
	btree.InsertRight(ejercicios.NewBinaryTree(5))

	sumaHojas := btree.SumaDeHojas()

	assert.Equal(t, 8, sumaHojas)
}

func TestSumaDeNodosInternos(t *testing.T){
	btree1 := ejercicios.NewBinaryTree(3)
	btree1.InsertLeft(ejercicios.NewBinaryTree(6))
	btree2 := ejercicios.NewBinaryTree(2)
	btree2.InsertRight(ejercicios.NewBinaryTree(4))
	// btree3 := ejercicios.NewBinaryTree(5)

}


// Ej 03





// Ej 04
func TestSumaNodosPares(t *testing.T){
	tree1:=ejercicios.NewBinaryTree(2)
	tree1.InsertRight(ejercicios.NewBinaryTree(3))
	tree1.InsertLeft(ejercicios.NewBinaryTree(4))
	sumaNodosPares:= tree1.SumaNodosPares()
	assert.Equal(t, 6, sumaNodosPares)
}

// Ej 05
func TestSumaNodosMayoresA6(t *testing.T){
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

	sumaMayorSeis:= tree.SumaNodosMayoresA6()

	assert.Equal(t, 33, sumaMayorSeis)
}

// Ej 06
func TestComparacionSubArboles(t *testing.T){
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

	alturaSubarboles:= tree.CompararacionAlturasEntreSubarboles()

	assert.Equal(t, 3, alturaSubarboles)
}




// Ej 07
func TestSumaDeHojasIzquierdas(t *testing.T){
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

	sumaHojasIzq:= tree.SumaDeHojasIzquierdas()

	assert.Equal(t, 44, sumaHojasIzq)
}


// Ej 8
func TestSumaNodosDerechasImpares(t *testing.T){
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

	sumaNodosDerImpares:= tree.SumaNodosDerechosImpares()
	assert.Equal(t, 55, sumaNodosDerImpares)
}