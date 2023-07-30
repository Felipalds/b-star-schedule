/***********************************************************
 * RBTree.go: Implementação de uma árvore-B (B-Tree) em Go *
 ***********************************************************/
package main

import (
	"fmt"
	"strings"
)

const t = 2 // Grau (ou ordem) da árvore B
// No. mínimo de chaves/nó = t - 1
// No. máximo de chaves/nó = 2*t - 1

type DataType Index

/******************************
 * Declaração da classe BTree *
 ******************************/

/*******************
 * Estrutura do nó *
 *******************/
type BTreeNode struct {
	leaf     bool         // identifica se um nó é folha (true) ou não (false)
	keys     []DataType   // vetor de chaves em cada nó
	children []*BTreeNode // vetor de ponteiros em cada nó
}

var INDEX_NIL Index

/*********************************************************
 * InitNode(leaf): Criação de um novo nó da árvore B *
 * leaf indica se o novo nó será uma folha ou não        *
 *********************************************************/
func InitNode(leaf bool) *BTreeNode {
	return &BTreeNode{
		leaf:     leaf,
		keys:     []DataType{},
		children: []*BTreeNode{},
	}
}

/**************************************
 * Definição da estrutura da Árvore B *
 **************************************/
type BTree struct {
	root *BTreeNode
}

/*************************************
 * Init(): Inicialização da árvore-B *
 *************************************/
func Init() *BTree {
	return &BTree{
		root: InitNode(true),
	}
}

/*********************************************************
 * Impressão da árvore B em forma de árvore de diretório *
 *********************************************************/
func (node *BTreeNode) Print(indent string, last bool) {
	fmt.Print(indent)
	if last {
		fmt.Print("└─ ")
		indent += "    "
	} else {
		fmt.Print("├─ ")
		indent += "|   "
	}
	keys := make([]string, len(node.keys))
	fmt.Print("[")
	for i, key := range node.keys {
		keys[i] = fmt.Sprintf("%v", key)
	}
	fmt.Println(strings.Join(keys, "|"), "]")

	childCount := len(node.children)
	for i, child := range node.children {
		child.Print(indent, i == childCount-1)
	}
}

/*********************************************************
 * Impressão da árvore B em forma de árvore de diretório *
 *********************************************************/
func (node *BTreeNode) PrintContacts() {
	for i, _ := range node.keys {
		index := Index(node.keys[i])
		getAndPrintContact(&index)
	}
	for _, child := range node.children {
		child.VisitInOrder()
	}
}

/*********************************************************
 * B Tree in order *
 *********************************************************/
func (node *BTreeNode) VisitInOrder() {

	for i, _ := range node.keys {
		index := Index(node.keys[i])
		insertIndexInFile(&index)
	}
	for _, child := range node.children {
		child.VisitInOrder()
	}
}

/*********************************************************
 * splitChild(i): implementa a Divisão de um filho cheio *
 * i é o ponto onde o nó será dividido                   *
 *********************************************************/
func (node *BTreeNode) splitChild(i int16) {
	child := node.children[i]
	newChild := InitNode(child.leaf)

	// Move as chaves e os filhos para o novo filho
	newChild.keys = append(newChild.keys, child.keys[t:]...)
	child.keys = child.keys[:t]
	if !child.leaf { // divide o nó em dois
		newChild.children = append(newChild.children, child.children[t:]...)
		child.children = child.children[:t]
	}

	// Insere o novo filho no nó
	node.children = append(node.children, nil)
	copy(node.children[i+2:], node.children[i+1:])
	node.children[i+1] = newChild

	// Move a chave correspondente para cima
	node.keys = append(node.keys, DataType(INDEX_NIL))
	copy(node.keys[i+1:], node.keys[i:])
	node.keys[i] = child.keys[t-1]
	child.keys = child.keys[:t-1]
}

/***********************************************************
 * Insert(key): Inserção de uma chave em um nó da árvore B *
 * key é a chave que será inserida                        *
 ***********************************************************/
func (node *BTreeNode) Insert(key DataType) {
	if !node.leaf {
		// Encontra o filho apropriado para inserir a chave
		i := len(node.keys) - 1
		for i >= 0 && key.key < node.keys[i].key {
			i--
		}

		// Insere a chave no filho apropriado
		if len(node.children[i+1].keys) == 2*t-1 {
			node.splitChild(int16(i) + 1)
			if key.key > node.keys[i+1].key {
				i++
			}
		}
		node.children[i+1].Insert(key)
	} else {
		// Insere a chave no nó folha
		i := len(node.keys) - 1
		node.keys = append(node.keys, DataType(INDEX_NIL))
		for i >= 0 && key.key < node.keys[i].key {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	}
}

/******************************************************
 * Insert(key): Inserção de uma chave na árvore B     *
 * Esta é a função que deve ser chamada para realizar *
 * a inserção. key é a chave a ser inserida.          *
 ******************************************************/
func (tree *BTree) Insert(key DataType) {
	root := tree.root
	if len(root.keys) == 2*t-1 {
		newRoot := InitNode(false)
		newRoot.children = append(newRoot.children, root)
		newRoot.splitChild(0)
		tree.root = newRoot
	}
	tree.root.Insert(key)
}

// Busca de uma chave na árvore B
func (node *BTreeNode) Search(key string) *DataType {
	i := 0
	for i < len(node.keys) && key > node.keys[i].key {
		i++
	}

	if i < len(node.keys) && key == node.keys[i].key {
		return &node.keys[i]
	} else if node.leaf {
		return nil
	} else {
		return node.children[i].Search(key)
	}
}

// Busca de uma chave na árvore B
func (tree *BTree) Search(key string) *DataType {
	return tree.root.Search(key)
}
