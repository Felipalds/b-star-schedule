package main

import (
)

const t = 2

type DataType uint16

type Node struct {
    isLeaf bool
    keys []DataType
    children []*Node
}

type BTree struct {
    root *Node
}

func Init() *BTree {
    return &BTree {
        root: InitNode(true),
    }
}

func InitNode(isLeaf bool) *Node {
    return &Node {
        isLeaf: isLeaf,
        keys: []DataType{},
        children: []*Node{},
    }
}

func (node *Node) Insert (key DataType) {

    i := len(node.keys) - 1
    if !node.isLeaf {
        for i >= 0 && key < node.keys[i] {
            i--
        }

        if len(node.children[i+1].keys) == 2*t - 1 {
            node.splitChild(int16(i) + 1)
            if key > node.keys[i + 1] {
                i++
            }
        }
        node.children[i+1].Insert(key)
    } else {
        node.keys = append(node.keys, 0)
        for i >= 0 && key < node.keys[i] {
            node.keys[i + 1] = node.keys[i]
            i--
        }
        node.keys[i + 1] = key
    }
}

func (tree *BTree) Insert (key DataType) {
    root := tree.root
    if len( root.keys ) == 2*t - 1 {
        newRoot := InitNode(false)
        newRoot.children = append(newRoot.children, root)
        newRoot.splitChild(0)
        tree.root = newRoot
    }
    tree.root.Insert(key)
}

func (node *Node) splitChild (i int16) {
    child := node.children[i]
    newChild := InitNode(child.isLeaf)

    newChild.keys = append(newChild.keys, child.keys[t:]...)
    child.keys = child.keys[:t]

    if !child.isLeaf {
        newChild.children = append(newChild.children, child.children[t:]...)
        child.children = child.children[:t]
    }

    node.children = append(node.children, nil)
    copy(node.children[i+2:], node.children[i+1:])
    node.children[i+1] = newChild

    node.keys = append(node.keys, 0)
    copy(node.keys[i+1:], node.keys[i:])
    node.keys[i] = child.keys[t-1]
    child.keys = child.keys[:t-1]
}
