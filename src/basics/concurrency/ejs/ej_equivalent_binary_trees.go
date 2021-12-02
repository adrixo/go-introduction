package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/tour/tree"
)

func printTree(t *tree.Tree, printVoidLeaves bool) {
	var prefix2 string = ""
	var prefix *string = &prefix2
	fmt.Println("Printing tree")
	recursiveBinaryTreePrint(t, true, prefix, printVoidLeaves)
}

func recursiveBinaryTreePrint(t *tree.Tree, left bool, prefix *string, printVoidLeaves bool) {

	if t != nil || (t == nil && printVoidLeaves) {
		fmt.Printf("%s", *prefix)
		if left {
			fmt.Printf("└─")
			*prefix += "   "
		} else {
			fmt.Printf("├──")
			*prefix += "|    "
		}
	}

	if t != nil {
		fmt.Printf("<" + strconv.Itoa(t.Value) + ">\n")
		tempPrefix := *prefix

		recursiveBinaryTreePrint(t.Left, false, prefix, printVoidLeaves)
		*prefix = tempPrefix
		recursiveBinaryTreePrint(t.Right, true, prefix, printVoidLeaves)
	} else if printVoidLeaves {
		fmt.Printf("<*>\n")
	}
	return

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkAndClose(t *tree.Tree, ch chan int, stepDelay time.Duration, name string) {
	Walk(t, ch, stepDelay, name)
	close(ch)
}
func Walk(t *tree.Tree, ch chan int, stepDelay time.Duration, name string) {
	time.Sleep(stepDelay)

	if t != nil {
		// fmt.Printf("%d ", t.Value)
		Walk(t.Left, ch, stepDelay, name)
		fmt.Printf("%s.%d ", name, t.Value)
		ch <- t.Value
		Walk(t.Right, ch, stepDelay, name)
	}
	return
}

// Busqueda en profunididad
// fmt.Printf("%d ", t.Value)

// if t.Left == nil && t.Right == nil {
// } else {
// 	if t.Left != nil {
// 		Walk(t.Left, ch)
// 	}
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// }
// return

// Solución
// if t.Left == nil && t.Right == nil {
// } else {
// 	if t.Left != nil {
// 		Walk(t.Left, ch)
// 	}
// fmt.Printf("%d ", t.Value)
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// }
// return

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go WalkAndClose(t1, ch1, 100*time.Millisecond, "t1")
	go WalkAndClose(t2, ch2, 300*time.Millisecond, "t2")

	for {
		fmt.Println("Reading chanels...")
		value1, isOpen1 := <-ch1
		value2, isOpen2 := <-ch2
		fmt.Println("Evaluating t1: ", value1, " - t2: ", value2)
		if value1 != value2 {
			fmt.Println("Values are not equal!")
			return false
		} else if !isOpen1 || !isOpen2 {
			break
		}
	}

	fmt.Println("Trees are equal")

	return true
}

func main() {
	tree1 := tree.New(1)
	tree2 := tree.New(1)

	// printTree(tree, true)
	// printTree(tree1, false)
	// ch := make(chan int)
	// WalkAndClose(tree1, ch, 100*time.Millisecond, "t1")
	// WalkAndClose(tree2, ch, 200*time.Millisecond, "t2")

	printTree(tree1, false)
	printTree(tree2, false)
	Same(tree1, tree2)

	tree2 = tree.New(2)
	printTree(tree1, false)
	printTree(tree2, false)
	Same(tree1, tree2)
}
