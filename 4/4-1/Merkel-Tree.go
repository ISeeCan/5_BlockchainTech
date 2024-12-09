package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	left   *Node
	right  *Node
	hash   []byte
	isLeaf bool
}

type MerkleTree struct {
	root  *Node
	leaves []*Node // 存储叶子节点，便于定位修改
}

func hashData(data string) []byte {
	hash := sha256.Sum256([]byte(data))
	return hash[:]
}

func NewMerkleTree(data []string) *MerkleTree {
	// 创建叶子节点
	var leaves []*Node
	for _, d := range data {
		leaf := &Node{
			hash:   hashData(d),
			isLeaf: true,
		}
		leaves = append(leaves, leaf)
	}

	// 生成树的非叶节点
	for len(leaves) > 1 {
		var newLevel []*Node
		for i := 0; i < len(leaves); i += 2 {
			left := leaves[i]
			var right *Node
			if i+1 < len(leaves) {
				right = leaves[i+1]
			} else {
				right = left // 奇数情况下复制节点
			}
			parentHash := sha256.Sum256(append(left.hash, right.hash...))
			parent := &Node{
				left:   left,
				right:  right,
				hash:   parentHash[:],
				isLeaf: false,
			}
			newLevel = append(newLevel, parent)
		}
		leaves = newLevel
	}

	// 根节点
	tree := &MerkleTree{
		root:   leaves[0],
		leaves: leaves,
	}
	return tree
}

func (tree *MerkleTree) UpdateLeaf(index int, newData string) {
	// 修改叶子节点的数据和哈希
	leaf := tree.leaves[index]
	leaf.hash = hashData(newData)

	// 递归更新父节点的哈希值
	updateParentHash(leaf)
}

func updateParentHash(node *Node) {
	// 如果是根节点，不需要更新
	if node == nil || node.left == nil && node.right == nil {
		return
	}

	// 更新父节点的哈希
	parentHash := sha256.Sum256(append(node.left.hash, node.right.hash...))
	node.hash = parentHash[:]

	// 递归更新上层父节点
	updateParentHash(node)
}

func compareMerkleTree(tree1 *MerkleTree, tree2 *MerkleTree) int {
	for i := 0; i < len(tree1.leaves); i++ {
		if string(tree1.leaves[i].hash) != string(tree2.leaves[i].hash) {
			return i
		}
	}
	return -1 // 如果返回-1，表示没有找到差异
}

func main() {
	// 创建16个叶子节点的初始数据
	data := []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10", "data11", "data12", "data13", "data14", "data15", "data16"}

	// 生成初始的Merkle Tree
	tree1 := NewMerkleTree(data)

	// 生成一个复制的Merkle Tree
	tree2 := NewMerkleTree(data)

	// 修改tree2中的某个叶子节点的数据
	tree2.UpdateLeaf(5, "modified_data5")

	// 比较两个Merkle Tree并找出不同的叶子节点位置
	index := compareMerkleTree(tree1, tree2)
	fmt.Printf("Modified leaf index: %d\n", index)
}
