package main

import "fmt"

func main()  {
	bst := NewBinaryTree()
	node1 := &Node{Data:  4, Left:  nil, Right: nil,}
	node2 := &Node{Data:  2, Left:  nil, Right: nil,}
	node3 := &Node{Data:  6, Left:  nil, Right: nil,}
	node4 := &Node{Data:  1, Left:  nil, Right: nil,}
	node5 := &Node{Data:  3, Left:  nil, Right: nil,}
	node6 := &Node{Data:  5, Left:  nil, Right: nil,}
	node7 := &Node{Data:  7, Left:  nil, Right: nil,}
	bst.Root=node1
	node1.Left=node2
	node1.Right=node3
	node2.Left=node4
	node2.Right=node5
	node3.Left=node6
	node3.Right=node7
	bst.Size = 7
	//手动创建平衡树
	fmt.Println("-------------中序遍历------------")
	bst.InOrder()
	fmt.Println(bst.InOrderNoRecursion())
	fmt.Println("-------------中序遍历------------")

	fmt.Println("-------------前序遍历------------")
	bst.PreOrder()
	fmt.Println(bst.PreOrderNoRecursion())
	fmt.Println("-------------前序遍历------------")

	fmt.Println("-------------后序遍历------------")
	bst.PostOrder()
	fmt.Println(bst.PostOrderNoRecursion())
	fmt.Println("-------------后序遍历------------")

	fmt.Println("last----------------------------")
	fmt.Println(bst.String())


	fmt.Println("面试题")
	fmt.Println("--------------队列广度遍历level--------------")
	bst.LevelS()
	fmt.Println("--------------队列广度遍历level--------------")

	fmt.Println("--------------栈深度遍历level--------------")
	bst.Levelstack(bst.Root)
	fmt.Println("--------------栈深度遍历level--------------")


	fmt.Println("--------------二叉树最小公共祖先--------------")
	nodelast := bst.FindlowerstAncestor(bst.Root,node3,node4)
	fmt.Println(nodelast)
	fmt.Println("--------------二叉树最小公共祖先--------------")

	fmt.Println("------------递归二叉树深度----------------")

	fmt.Println(bst.GetDepth(bst.Root))
	fmt.Println("--------------递归二叉树深度--------------")

}
