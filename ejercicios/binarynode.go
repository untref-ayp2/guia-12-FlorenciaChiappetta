package ejercicios

import (
	"fmt"
	"math"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func NewBinaryNode(data int, left *BinaryNode, right *BinaryNode) *BinaryNode {
	return &BinaryNode{left: left, right: right, data: data}
}

func (n *BinaryNode) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

func (n *BinaryNode) Height() int {
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
func (n *BinaryNode) SumarNodos() int {
	if n == nil {
		return 0
	}
	suma := n.data
	if n.left != nil {
		suma += n.left.SumarNodos()
	}
	if n.left != nil {
		suma += n.right.SumarNodos()
	}
	return suma
}

func (n *BinaryNode) SumarNodos2() int {
	if n == nil {
		return 0
	}

	return n.data + n.left.SumarNodos() + n.right.SumarNodos()
}

func (n *BinaryNode) SumarHojasByNode() int {

	if n == nil {
		return 0
	}
	sumaizq := n.left.SumarHojasByNode()
	sumader := n.right.SumarHojasByNode()
	if n.left == nil && n.right == nil {
		return n.data
	}

	return sumaizq + sumader
}

func (n *BinaryNode) SumaValoresPares() int {

	if n == nil {
		return 0
	}
	sumaI := n.left.SumaValoresPares()
	sumaD := n.right.SumaValoresPares()
	if n.data%2 == 0 {
		return n.data + sumaI + sumaD
	}
	return sumaI + sumaD
}

func (n *BinaryNode) SumarNodosMayoresA6ByNode() int {
	if n == nil {
		return 0
	}
	sumaIzq := n.left.SumarNodosMayoresA6ByNode()
	sumaDer := n.right.SumarNodosMayoresA6ByNode()
	if n.data > 6 {
		return n.data + sumaDer + sumaIzq
	}
	return sumaDer + sumaIzq
}
func (n *BinaryNode) MayorAlturaByNode() int {
	// Si el árbol está vacío, devolver -1
	if n == nil {
		return -1
	}
	alturaIzquierda := n.left.MayorAlturaByNode() + 1
	alturaDerecha := n.right.MayorAlturaByNode() + 1
	// Devolver la mayor altura
	if alturaIzquierda > alturaDerecha {
		return alturaIzquierda
	} else {
		return alturaDerecha
	}
}
func (n *BinaryNode) SumaHojasIzquierdas(esHijoIzquierdo bool) int {
	if n == nil {
		return 0
	}
	// Si el nodo es una hoja hija izquierda, sumar su valor
	if n.left == nil && n.right == nil && esHijoIzquierdo {
		return n.data
	}
	// Calcular la suma de las hojas hijas izquierdas del sub-árbol izquierdo
	sumaIzquierda := n.left.SumaHojasIzquierdas(true)
	// Calcular la suma de las hojas hijas izquierdas del sub-árbol derecho
	sumaDerecha := n.right.SumaHojasIzquierdas(false)
	// Devolver la suma total
	return sumaIzquierda + sumaDerecha
}

func reemplazarConSumaHijos(raiz *Nodo) {
	// Si el árbol está vacío, no hacer nada
	if raiz == nil {
		return
	}
	// Reemplazar los valores de los nodos del sub-árbol izquierdo
	reemplazarConSumaHijos(raiz.izquierda)
	// Reemplazar los valores de los nodos del sub-árbol derecho
	reemplazarConSumaHijos(raiz.derecha)
	// Calcular la suma de los valores de los hijos
	sumaHijos := 0
	if raiz.izquierda != nil {
		sumaHijos += raiz.izquierda.valor
	}
	if raiz.derecha != nil {
		sumaHijos += raiz.derecha.valor
	}
	// Reemplazar el valor del nodo actual con la suma de los valores de los hijos
	raiz.valor = sumaHijos
}

func altura(raiz *Nodo) int {
	// Si el árbol está vacío, devolver -1
	if raiz == nil {
		return -1
	}
	// Calcular la altura del sub-árbol izquierdo
	alturaIzquierda := altura(raiz.izquierda)
	// Calcular la altura del sub-árbol derecho
	alturaDerecha := altura(raiz.derecha)
	// Devolver la mayor altura más uno
	if alturaIzquierda > alturaDerecha {
		return alturaIzquierda + 1
	} else {
		return alturaDerecha + 1
	}
}

func sumaNodosDerechosImpares(raiz *Nodo, esHijoDerecho bool) int {
	// Si el árbol está vacío, devolver 0
	if raiz == nil {
		return 0
	}
	// Si el nodo es un hijo derecho cuyo valor es impar, sumar su valor
	if esHijoDerecho && raiz.valor%2 == 1 {
		return raiz.valor + sumaNodosDerechosImpares(raiz.izquierda, false) + sumaNodosDerechosImpares(raiz.derecha, true)
	} else {
		return sumaNodosDerechosImpares(raiz.izquierda, false) + sumaNodosDerechosImpares(raiz.derecha, true)
	}
}
func sumaHojasIzquierdas(raiz *Nodo, esHijoIzquierdo bool) int {
	// Si el árbol está vacío, devolver 0
	if raiz == nil {
		return 0
	}
	// Si el nodo es una hoja hija izquierda, sumar su valor
	if raiz.izquierda == nil && raiz.derecha == nil && esHijoIzquierdo {
		return raiz.valor
	}
	// Calcular la suma de las hojas hijas izquierdas del sub-árbol izquierdo
	sumaIzquierda := sumaHojasIzquierdas(raiz.izquierda, true)
	// Calcular la suma de las hojas hijas izquierdas del sub-árbol derecho
	sumaDerecha := sumaHojasIzquierdas(raiz.derecha, false)
	// Devolver la suma total
	return sumaIzquierda + sumaDerecha
}
func mayorAltura(raiz *Nodo) int {
	// Si el árbol está vacío, devolver -1
	if raiz == nil {
		return -1
	}
	// Calcular la altura del sub-árbol izquierdo
	alturaIzquierda := mayorAltura(raiz.izquierda) + 1
	// Calcular la altura del sub-árbol derecho
	alturaDerecha := mayorAltura(raiz.derecha) + 1
	// Devolver la mayor altura
	if alturaIzquierda > alturaDerecha {
		return alturaIzquierda
	} else {
		return alturaDerecha
	}
}
func sumaValoresPares(raiz *Nodo) int {
	// Si el árbol está vacío, devolver 0
	if raiz == nil {
		return 0
	}
	// Calcular la suma de los valores pares del subárbol izquierdo
	sumaIzquierda := sumaValoresPares(raiz.izquierda)
	// Calcular la suma de los valores pares del subárbol derecho
	sumaDerecha := sumaValoresPares(raiz.derecha)
	// Si el valor del nodo actual es par, sumarlo
	if raiz.valor%2 == 0 {
		return raiz.valor + sumaIzquierda + sumaDerecha
	} else {
		return sumaIzquierda + sumaDerecha
	}
}
