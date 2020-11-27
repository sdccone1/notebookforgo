package demo

/**
@author:David Ma
@Content:定义一个Equals()用于比较两个map是否相等
@Date:2020-11-27 15:33
*/
func Equals(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k1, v1 := range map1 {
		v2, isExists := map2[k1]
		if !isExists {
			return false
		}
		if v2 != v1 {
			return false
		}
	}
	return true
}
