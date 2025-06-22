package util

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"sync"
)

/*
敏感词过滤
*/

//go:embed sensitive_words.txt
var sensitiveWords []byte // 嵌入敏感词文件

// TrieNode 前缀树节点
type TrieNode struct {
	IsEnd    bool               // 关键词结束标识
	Children map[rune]*TrieNode // 子节点（可能有多个,key是下级字符，value是下级节点）
}

// SensitiveFilter 敏感词过滤器
type SensitiveFilter struct {
	Root *TrieNode
}

// GetSensitiveFilter 获取敏感词过滤器，使用sync.Once保证单例
var GetSensitiveFilter = sync.OnceValue(func() *SensitiveFilter {
	filter := &SensitiveFilter{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
	// 从嵌入文件逐行读取敏感词，并添加到前缀树中
	scanner := bufio.NewScanner(strings.NewReader(string(sensitiveWords)))
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) > 0 {
			filter.AddWord(word)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("解析敏感词失败: %v", err))
	}
	return filter
})

// AddWord 添加敏感词到前缀树
func (f *SensitiveFilter) AddWord(word string) {
	current := f.Root
	for _, char := range word { // 遍历单词中的每个字符
		// 判断当前字符是否已被添加到前缀树中,如果没有则需要新建
		if _, ok := current.Children[char]; !ok {
			current.Children[char] = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
		}
		// 指针指向子节点，进入下一轮循环
		current = current.Children[char]
	}
	// 设置敏感词结束标识
	current.IsEnd = true
}

// Filter 过滤敏感词（时间复杂度，O(n*m)，n是文本长度，m是敏感词平均长度）
func (f *SensitiveFilter) Filter(text string) string {
	if len(text) == 0 {
		return ""
	}
	runes := []rune(text)
	// 保存过滤后的文本结果
	var result strings.Builder
	currentNode := f.Root
	begin := 0 // 滑动窗口起始点
	end := 0   // 当前窗口终止点
	for end < len(runes) {
		// 检查下级节点
		currentNode = currentNode.Children[runes[end]]
		if currentNode == nil {
			// 以begin开头的字符串不是敏感词
			result.WriteRune(runes[begin])
			// 滑动窗口进入下一个位置
			begin++
			end = begin
			// 重新指向根节点
			currentNode = f.Root
		} else if currentNode.IsEnd {
			// 发现敏感词,将begin~end字符串替换掉
			result.WriteString("***")
			// 进入下一个位置
			end++
			begin = end
			// 重新指向根节点
			currentNode = f.Root
		} else {
			// 检查下一个字符
			end++
		}
	}
	// 将最后一批字符计入结果
	result.WriteString(string(runes[begin:]))
	return result.String()
}
