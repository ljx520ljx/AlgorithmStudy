package ImplementPrefixTree

//实现Trie(前缀树)
/*
算法思路
Trie（前缀树）是用于高效存储和检索字符串的一种树形数据结构。
它特别适用于自动补全、拼写检查等场景。Trie 的核心思想是 共享前缀，即多个单词共享相同的路径。
每个节点代表一个字符。
路径从根节点到某个节点表示一个前缀。
叶节点标记 isEnd，表示该节点是一个完整单词的结尾。
主要操作：
插入单词：从根节点开始，逐个字符插入。如果字符路径已经存在，则沿着现有路径走；否则，创建新的节点。
查找单词：从根节点开始，逐个字符查找。如果路径存在且最后一个字符是一个单词的结尾，则返回 true。
前缀查找：与查找单词类似，只要路径存在，说明有单词以该前缀开始。
*/
type Trie struct {
	children map[rune]*Trie // 存储当前节点的子节点
	isEnd    bool           // 标志当前节点是否是一个完整单词的结尾
}

// Constructor 初始化 Trie 对象
func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
		isEnd:    false,
	}
}

// Insert 向前缀树中插入字符串
func (this *Trie) Insert(word string) {
	currentNode := this
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = &Trie{
				children: make(map[rune]*Trie),
				isEnd:    false,
			}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEnd = true
}

// Search 如果字符串在前缀树中存在，则返回 true
func (this *Trie) Search(word string) bool {
	currentNode := this
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEnd
}

// StartsWith 如果有任何单词以给定前缀开始，返回 true
func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this
	for _, char := range prefix {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return true
}
