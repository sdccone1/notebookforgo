package demo
/**
@author:David Ma
@Content:test for fmt package
@Date:2020-11-18 19:33
*/
func Hello(feeling string) string {
	prefix := ""
	const name = "Lisa"
	//注意go中的switch的case不用在写break强制中断匹配了，当匹配成功后会自动终止，不再继续匹配下去了！
	switch feeling {
	case "good":
		prefix = "Love you"
	case "bad":
		prefix = "Fucking you"
	default:
		prefix = "Hi "
	}
	return prefix + name
}

