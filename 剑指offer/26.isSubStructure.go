package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	//	当在A中找到B的根节点时，进入helper递归校验
	if A.Val == B.Val {
		ret = helper(A, B)
	}

	//	ret == false，说明B的根节点不在当前A树顶中，进入A的左子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Left, B)
	}

	//	当B的根节点不在当前A树顶和左子树中，进入A的右子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Right, B)
	}

	return ret

	//return helper(A,B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// helper 校验B是否与A的一个子树拥有相同的结构和节点值
func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	//	a.Val == B.Val 递归校验A B 左右子树的结构和结点是否相同
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}

func main() {

}
