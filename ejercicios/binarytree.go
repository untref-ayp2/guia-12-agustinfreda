package ejercicios

type BinaryTree[data int] struct {
	root *BinaryNode[int]
}

func NewBinaryTree(data int) *BinaryTree[int] {
	node := NewBinaryNode(data, nil, nil)
	return &BinaryTree[int]{root: node}

}

func (t *BinaryTree[int]) InsertLeft(bt *BinaryTree[int]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}
func (t *BinaryTree[int]) InsertRight(bt *BinaryTree[int]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

func (t *BinaryTree[int]) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

func (t *BinaryTree[int]) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

func (t *BinaryTree[int]) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

func (t *BinaryTree[int]) Empty() {
	t.root = nil
}

func (t *BinaryTree[int]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree[int]) Size() int {
	if t.root != nil {
		return int(t.root.Size())
	} else {
		return -1
	}
}

func (t *BinaryTree[int]) Height() int {
	if t.root != nil {
		return int(t.root.Height())
	} else {
		return 0
	}
}

// Ejercicios


// Ejercicio 1
func (t *BinaryTree[int]) SumaDeNodos()int{
	if t.root == nil{
		return 0
	}
	return int(t.root.SumaDeNodos())
}


// Ejercicio 2
func (t *BinaryTree[int]) SumaDeHojas() int {
	if t.root == nil {
		return 0
	}

	return int(t.SumaDeHojas())

}


// Ejercicio 3
func (t *BinaryTree[int]) SumaNodosInternos()int{
	if t.root == nil{
		return 0
	}
	return int(t.root.SumaDeNodos() - t.root.SumaDeHojas() - t.root.data)
}

// Ejercicio 4
func (t *BinaryTree[int]) SumaNodosPares()int{
	if t.root == nil{
		return 0
	}

	return int(t.root.SumaNodosPares())
}


// Ejercicio 5
func (t *BinaryTree[int]) SumaNodosMayoresA6()int{
	if t.root == nil{
		return 0
	}

	return int(t.root.SumaNodosMayoresA6())
}

// Ejercicio 6
func (t *BinaryTree[int]) CompararacionAlturasEntreSubarboles() int{
	if t.root == nil{
		return -1
	}

	return int(t.root.CompararacionAlturasEntreSubarboles())
}


// Ejercicio 7
func (t *BinaryTree[int]) SumaDeHojasIzquierdas()int{
	
	if t.root == nil{
		return 0
	}

	return int(t.root.SumaDeHojasIzquierdas())
}


// Ejercicio 8 
func (t *BinaryTree[data]) SumaNodosDerechosImpares()int{
	if t.root == nil{
		return 0
	}

	return int(t.root.SumaNodosDerechosImpares())
}