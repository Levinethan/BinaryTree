package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

//不平衡二叉树  解决增删查改综合效率
//需要掌握技巧
//data 递归左 递归右
//递归左 data 递归右
//递归左 递归右 data
//非递归

type Node struct {
	Data int
	Left *Node  //指向左边节点
	Right *Node  //指向右边节点

}

type BinaryTree struct {
	Root  *Node  //任何一个二叉树 都有一个根节点
	Size  int		//数据的大小
	
}
//当我们围绕一个 数据结构 实现一系列东西的时候  New
func NewBinaryTree() *BinaryTree  { //*BinaryTree 返回一个指针地址
	bst := &BinaryTree{
		Root: nil,
		Size: 0,
	}
	return bst
}

//获取二叉树的大小
func (bst *BinaryTree)GetSize() int  {
	return bst.Size
}

//判断是否为空
func ( bst *BinaryTree)IsEmpty() bool  {
	return bst.Size==0
}

//根节点插入
func (bst *BinaryTree)Add(data int)  {
	bst.Root = bst.add(bst.Root,data)
}

//不断得用到递归 实现插入
//插入节点
func (bst *BinaryTree)add (n *Node,data int) *Node {
	if n==nil{
		bst.Size ++
		return &Node{
			Data:  data,
			Left:  nil,
			Right: nil,
		}
	}else {
		if data < n.Data{
			n.Left = bst.add(n.Left,data) //递归调用 比我小去左边
		}else  if data > n.Data{
			n.Right = bst.add(n.Right,data)
		}
		return n
	}
}
//判断数据是否存在
func (bst *BinaryTree)Isin (data int)bool  {
	return bst.isin(bst.Root,data)  //从根节点开始查找
	//二叉树里的二分查找  无论是红黑树 AVL树也好  也是一样的方法
}


// 但是我们判断是否存在的时候 需要增加一个节点  有时候需要跑到左节点 或者是右节点去寻找
func (bst *BinaryTree)isin (n *Node , data int)bool  {
	if n == nil{
		return false  //如果树是空树 ，找不到
	}
	if data == n.Data{
		return true
	}else if data < n.Data{
		return bst.isin(n.Left,data)  //实现一个递归调用  一层树枝没找到 继续调用递归去下一层
	}else {
		return bst.isin(n.Right,data)
	}
}

//找二叉树最大最小值
func (bst *BinaryTree)FindMax()int  {
	if bst.Size == 0{
		panic("二叉树节点里为空")
	}
	return bst.findMax(bst.Root).Data //取得最大  只要你右边节点 不等于nil  那么我就能找到最大
	//同理 最小值也是
}

func (bst *BinaryTree)findMax (n*Node) *Node  {
	if n.Right == nil{   //寻找最大  right
		return n
	}else {
		return bst.findMax(n.Right)
	}

}

func (bst *BinaryTree)FindMin() int  {
	if bst.Size == 0{
		panic("二叉树节点里为空")
	}

	return bst.findMin(bst.Root).Data
}

func (bst *BinaryTree)findMin (n*Node) *Node {
	if n.Left == nil{
		return n
	}else {
		return bst.findMin(n.Left)
	}
}

//那么就完成了  二叉平衡树的极值

//前序遍历
func (bst *BinaryTree)PreOrder()  {
	bst.preorder(bst.Root)
}

func (bst *BinaryTree)PreOrderNoRecursion() []int  {
	mybst := bst.Root //备份二叉树
	mystack := list.New()  //生成一个栈
	res := make([]int,0)  //生成数组，容纳中序数据
	for mybst!= nil || mystack.Len()!=0{
		for mybst !=nil{
			res = append(res,mybst.Data)
			mystack.PushBack(mybst)
			mybst = mybst.Left  //代表一个节点
		}
		if mystack.Len()!=0{
			//如果栈的长度不等于0  不断从栈中取出数据
			v := mystack.Back()			//挨个取出节点
			mybst = v.Value.(*Node)   //实例化
			mybst = mybst.Right
			mystack.Remove(v) //删除
		}
	}
	return res  //非递归实现遍历
}

func (bst *BinaryTree)preorder(node *Node)  {
	//如果是前序遍历  递归如何实现？  从判断开始
	if node == nil{
		return
	}
	fmt.Println(node.Data)
	//递归
	bst.preorder(node.Left)
	bst.preorder(node.Right)
}
//中序遍历
func (bst *BinaryTree)InOrder()  {
	bst.inorder(bst.Root)
}

func (bst *BinaryTree)InOrderNoRecursion() []int  {
	mybst := bst.Root //备份二叉树
	mystack := list.New()  //生成一个栈
	res := make([]int,0)  //生成数组，容纳中序数据
	for mybst!= nil || mystack.Len()!=0{
		for mybst !=nil{
			mystack.PushBack(mybst)
			mybst = mybst.Left  //代表一个节点
		}
		if mystack.Len()!=0{
			//如果栈的长度不等于0  不断从栈中取出数据
			v := mystack.Back()			//挨个取出节点
			mybst = v.Value.(*Node)   //实例化
			res = append(res,mybst.Data) //压入数据
			mybst = mybst.Right
			mystack.Remove(v) //删除
		}
	}
	return res  //非递归实现遍历
}

func (bst *BinaryTree)inorder(node *Node)  {
	if node == nil{
		return
	}

	//递归
	bst.inorder(node.Left)
	fmt.Println(node.Data)
	bst.inorder(node.Right)
}
//后续遍历
func (bst *BinaryTree)PostOrder()  {
	bst.postorder(bst.Root)
}
func (bst *BinaryTree)PostOrderNoRecursion() []int  {
	mybst := bst.Root //备份二叉树
	mystack := list.New()  //生成一个栈
	res := make([]int,0)  //生成数组，容纳中序数据

	var PreVisited *Node  //提前访问的节点

	//后续遍历
	for mybst!= nil || mystack.Len()!=0{
		for mybst !=nil{
			mystack.PushBack(mybst)
			mybst = mybst.Left  //代表一个节点  左边循环
		}

		v := mystack.Back() //取出栈中节点
		top := v.Value.(*Node) //实例化
		if top.Left == nil&& top.Right==nil || (top.Right==nil)&&PreVisited==top.Left || (PreVisited==top.Right){
			//如果左右==nil 或者 右边=nil 上一个节点=left 或者 Pre=right 都代表循环到尽头
			res = append(res,top.Data)
			PreVisited = top
			mystack.Remove(v)
		}else {
			mybst = top.Right    //右边循环
		}
	}
	return res  //非递归实现遍历
}

func (bst *BinaryTree)postorder(node *Node)  {
	if node == nil{
		return
	}
	//递归  深度遍历
	bst.postorder(node.Left)
	bst.postorder(node.Right)
	fmt.Println(node.Data)
}

func (bst *BinaryTree)String () string  {
	var buffer bytes.Buffer   //保存字符串
	bst.GenerateBSTstring(bst.Root,0,&buffer) //调用函数实现遍历
	return buffer.String()
}

func (bst *BinaryTree)GenerateBSTstring(node *Node,depth int,buffer *bytes.Buffer){
	if node == nil{
		buffer.WriteString(bst.GenerateDepthstring(depth)+"nil\n") //空节点
		return
	}

	bst.GenerateBSTstring(node.Left,depth+1,buffer)

	buffer.WriteString(bst.GenerateDepthstring(depth)+strconv.Itoa(node.Data)+"\n")

	bst.GenerateBSTstring(node.Right,depth+1,buffer)
}

func (bst *BinaryTree)GenerateDepthstring(depth int) string  {
	var buffer bytes.Buffer

	for i := 0 ; i<depth  ;i ++{
		buffer.WriteString("--")  //深度为0  1-- 2--
	}
	return buffer.String()
}

//添加删除 本质上是指针变换
func (bst *BinaryTree)RemoveMax() int  {
	ret := bst.FindMax()
	bst.Root = bst.removemax(bst.Root)
	return ret
}

func (bst *BinaryTree)removemax (n *Node) *Node {
	if n.Right == nil{
		//删除
		LeftNode := n.Left  //备份右边的节点  因为用到了递归 需要备份一下右边节点
		bst.Size --  //删除了一个元素
		return  LeftNode
	}
	n.Right = bst.removemax(n.Right)
	return n
}

func (bst *BinaryTree)RemoveMin() int  {
	ret := bst.FindMin()
	bst.Root = bst.removemin(bst.Root)
	return ret
}
func (bst *BinaryTree)removemin(n *Node) *Node  {
	if n.Left == nil{
		//删除
		rightNode := n.Right  //备份右边的节点  因为用到了递归 需要备份一下右边节点
		bst.Size --  //删除了一个元素
		return  rightNode
	}
	n.Left = bst.removemin(n.Left)
	return n
}

func (bst *BinaryTree)Remove()  {
	
}

func (bst *BinaryTree)remove (n *Node ,data int) *Node  {
	if n == nil{
		return nil
	}
	if data < n.Data{
		n.Left= bst.remove(n.Left,data) //递归到左边
		return n
	}else if data > n.Data{
		n.Right = bst.remove(n.Right,data)  //递归到右边
		return n
	}else {
		if  n.Left == nil{
			//当左节点为空
			rightNode := n.Right
			//备份右边的节点  因为用到了递归 需要备份一下右边节点
			n.Right = nil //如果左子树为空 把右边挪给左边/把左边挪给右边
			bst.Size --  //删除了一个元素
			return  rightNode
		}
		//处理右边为空
		if n.Right ==nil{
			LeftNode := n.Left  //备份右边的节点  因为用到了递归 需要备份一下右边节点
			bst.Size --  //删除了一个元素
			n.Left = nil
			return  LeftNode
		}

		//当左右节点都不为空
		oknode := bst.findMin(n.Right)
		oknode.Right = bst.removemin(n.Right)
		oknode.Left = n.Left
		n.Left = nil
		n.Right = nil
		return oknode
	}
}

func (bst *BinaryTree) LevelS () {
	bst.levels(bst.Root)
}
func (bst *BinaryTree) levels (n *Node) {
	queue := list.New()  //新建队列  实现广度遍历
	queue.PushBack(n) //压入节点N
	for queue.Len() > 0{
		left := queue.Front() //从前面取出数据
		right := left.Value
		queue.Remove(left) //删除
		if v , ok := right.(*Node); ok && v!=nil{
			fmt.Println(v.Data) //打印数据
			queue.PushBack(v.Left)
			queue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) Levelstack (n *Node) {
	queue := list.New()  //新建队列  实现广度遍历
	queue.PushBack(n) //压入节点N
	for queue.Len() > 0{
		left := queue.Back() //从前面取出数据 .Front 队列  .Back栈
		right := left.Value
		queue.Remove(left) //删除
		if v , ok := right.(*Node); ok && v!=nil{
			fmt.Println(v.Data) //打印数据
			queue.PushBack(v.Left)
			queue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree)FindlowerstAncestor (root *Node,a *Node,b*Node) *Node  {
	if root == nil{
		return nil
	}

	if root == a || root ==b{
		return root
	}

	left := bst.FindlowerstAncestor(root.Left,a,b)
	right := bst.FindlowerstAncestor(root.Right,a,b)

	if left !=nil && right!=nil{
		return root
	}
	if left != nil{
		return left
	}else {
		return right
	}

}

func (bst *BinaryTree)GetDepth (root *Node) int  {
	if root == nil{
		return  0
	}

	if root.Right == nil && root.Left == nil{
		return  1
	}

	lengthLeft := bst.GetDepth(root.Left)
	lengthRight := bst.GetDepth(root.Right)
	if lengthLeft > lengthRight{
		return lengthLeft +1
	}else {
		return lengthRight+1
	}

}