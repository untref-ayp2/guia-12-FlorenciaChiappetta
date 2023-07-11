package ejercicios

type BinaryTree struct {
	root *BinaryNode
}

func NewBinaryTree(data int) *BinaryTree {
	node := NewBinaryNode(data, nil, nil)
	return &BinaryTree{root: node}

}

func (t *BinaryTree) InsertLeft(bt *BinaryTree) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}
func (t *BinaryTree) InsertRight(bt *BinaryTree) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

func (t *BinaryTree) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

func (t *BinaryTree) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

func (t *BinaryTree) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

func (t *BinaryTree) Empty() {
	t.root = nil
}

func (t *BinaryTree) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree) Size() int {
	if t.root != nil {
		return t.root.Size()
	} else {
		return 0
	}
}

func (t *BinaryTree) Height() int {
	if t.root != nil {
		return t.root.Height()
	} else {
		return -1
	}
}
func (t *BinaryTree) SumarNodos() int {

	if t.root != nil {
		return t.root.SumarNodos()
	}
	return 0
}

func (t *BinaryTree) SumarNodos2() int {

	if t.root != nil {
		return t.root.SumarNodos2()
	}
	return 0
}
func (t *BinaryTree) SumarHojas() int {

	if t.root == nil {
		return 0
	}

	return t.root.SumarHojasByNode()
}

func (t *BinaryTree) SumarHojasPares() int {

	if t.root == nil {
		return 0
	}

	return t.root.SumaValoresPares()
}

func (t *BinaryTree) SumarNodosMayoresA6() int {
	if t.root == nil {
		return 0
	}
	return t.root.SumarNodosMayoresA6ByNode()
}

func (t *BinaryTree) MayorAltura() int {
	if t.root == nil {
		return -1
	}
	return t.root.MayorAlturaByNode()
}

func (t *BinaryTree) SumaHojasIzquierdas() int {
	if t.root == nil {
		return 0
	}
	return t.root.SumaHojasIzquierdas(false)
}
