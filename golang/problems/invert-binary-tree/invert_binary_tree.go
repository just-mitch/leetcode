package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 10)

	queue = append(queue, root)
	for len(queue) > 0 {
		n, tail := queue[0], queue[1:]
		queue = tail
		if n == nil {
			continue
		}
		tmp := n.Right
		n.Right = n.Left
		n.Left = tmp

		queue = append(queue, n.Left, n.Right)
	}

	return root

}
