package arrayandstrings

import (
	"fmt"
	"leetcode/utils"
)


func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	var results [][]string
	for i, _ := range strs{
		sortedStr := utils.SortString(strs[i])
		if _, ok := resultMap[sortedStr] ; !ok {
			resultMap[sortedStr] = []string{strs[i]}
		}else{
			resultMap[sortedStr] = append(resultMap[sortedStr], strs[i])
		}
	}

	for _, result := range resultMap{
		results = append(results, result)
	}

	return results

}

func RunGroupAnagrams() {
	fmt.Printf("%+v", groupAnagrams([]string{""}))
}
