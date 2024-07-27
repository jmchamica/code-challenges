/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Codec struct {
    
 }
 
 const (
	 sep = ";"
 )
 
 func Constructor() Codec {
	 return Codec{}
 }
 
 
 func atoi(s string) int {
	 res, err := strconv.Atoi(s)
	 if err != nil {
		 panic(err)
	 }
 
	 return res
 }
 
 // Serializes a tree to a single string.
 func (this *Codec) serialize(root *TreeNode) string {
	 if root == nil {
		 return ""
	 }
 
	 res := fmt.Sprintf("%d", root.Val)
	 stack := make([]*TreeNode, 0)
	 stack = append(stack, root)
 
	 for len(stack) > 0 {
		 iter := stack[0]
		 stack = append(stack[:0], stack[1:]...)
 
		 l := ""
		 if iter.Left != nil {
			 l = fmt.Sprintf("%d", iter.Left.Val)
		 }
 
		 r := ""
		 if iter.Right != nil {
			 r = fmt.Sprintf("%d", iter.Right.Val)
		 }
		 res += sep + l + sep + r
 
		 if iter.Left != nil {
			 stack = append(stack, iter.Left)
		 }
 
		 if iter.Right != nil {
			 stack = append(stack, iter.Right)
		 }
	 }
 
	 return res
 }
 
 type Node struct {
	 val    *TreeNode
	 parent *TreeNode
	 left   bool
 }
 
 // Deserializes your encoded data to tree.
 func (this *Codec) deserialize(data string) *TreeNode {
	 nodes := strings.Split(data, sep)
 
	 if len(nodes) <= 1 {
		 return nil
	 }
 
	 rootNode := Node{
		 val:    nil,
		 parent: nil,
		 left:   false,
	 }
 
	 stack := make([]*Node, 0)
	 stack = append(stack, &rootNode)
 
	 for _, v := range nodes {
		 if len(stack) <= 0 {
			 panic("Stack size is insufficient")
		 }
 
		 // pop
		 iter := stack[0]
		 stack = append(stack[:0], stack[1:]...)
 
		 if len(v) <= 0 {
			 continue
		 }
 
		 // alloc node
		 curr := TreeNode{
			 Val: atoi(v),
		 }
		 iter.val = &curr
 
		 // push children
		 left := Node{
			 val:    nil,
			 parent: iter.val,
			 left:   true,
		 }
		 right := Node{
			 val:    nil,
			 parent: iter.val,
			 left:   false,
		 }
		 stack = append(stack, &left)
		 stack = append(stack, &right)
 
		 if iter.parent == nil {
			 continue
		 }
 
		 if iter.left {
			 iter.parent.Left = iter.val
			 continue
		 }
 
		 iter.parent.Right = iter.val
	 }
 
	 return rootNode.val
 }