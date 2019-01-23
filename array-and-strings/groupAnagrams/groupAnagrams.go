package main

import (
	"fmt"
	"sort"
	"strings"
)

// 将字符串按字母从小到大排序
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// 方法一：现将每个单词按字母排序，存入一个数组，数组中元素唯一
// 再将单词一一与这个数组的元素进行比较，如果是已有元素，则填入对应的列中，
// 如果这个元素第一次出现，则新增加一个列
// 扫描完所有单词后，返回结果
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	var sortedStrs []string

outerLoop:
	for _, str := range strs {
		sortedStr := sortString(str)

		for i, item := range sortedStrs {
			if sortedStr == item {
				res[i] = append(res[i], str)
				continue outerLoop
			}
		}
		sortedStrs = append(sortedStrs, sortedStr)
		res = append(res, []string{str})
	}
	return res
}

// 方法二：改进，使用map[string][]string存储，而不是[][]string存储
func groupAnagrams2(strs []string) [][]string {
	anagramMap := make(map[string][]string, 0)

	for _, str := range strs {
		sortedStr := sortString(str)

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	res := make([][]string, 0)
	for _, value := range anagramMap {
		res = append(res, value)
	}
	return res
}

// 方法三：先统计一个单词中每个词的出现次数，作为字母频率数组
// 再利用map结构，以数组作为key，这个词的序号作为value，达到聚合的效果
// 最后将这个map结构转换为对应的二维数组结构并返回结果
// Source: https://leetcode.com/problems/group-anagrams/discuss/212228
func groupAnagrams3(strs []string) [][]string {
	length := len(strs)
	charFrequencyArray := make([][26]int, length)
	for i := 0; i < length; i++ {
		for j := 0; j < len(strs[i]); j++ { //统计每个单词中的字母出现情况
			charFrequencyArray[i][strs[i][j]-'a']++
		}
	}
	indexesMap := make(map[[26]int][]int)
	for i := 0; i < length; i++ {
		indexesMap[charFrequencyArray[i]] = append(indexesMap[charFrequencyArray[i]], i)
	}
	var results [][]string
	for _, indexes := range indexesMap {
		var tmp []string
		for j := 0; j < len(indexes); j++ {
			tmp = append(tmp, strs[indexes[j]])
		}
		results = append(results, tmp)
	}
	return results
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams2(strs)
	fmt.Println(res)
}
