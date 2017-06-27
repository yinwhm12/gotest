package algorithm


// 没有鬼(王)牌的时候 a[]---手中的牌
func NoSpecialCard(a []int) bool {
	length := len(a)
	var sameCount int //对子 记数
	var blackOrderCount	int //暗顺子 记数
	var otherCount	int //东南西北牌 且不是对子 是单个存在的
	for i := 0; i < length;{
		j := i //记录上个 下标
		i = Similar(a, j)
		if i - j == 1{ //仅有一个该类型的牌
			
		}else if i - j ==  2{//该类型的牌有两张
		//	是对子? 或者 是 暗顺?
			if IsSame(a[j:i]){
				sameCount++
			}else if IsBlackOrder(a[j:i]){
				blackOrderCount++
			}else {//其他情况 东南西北 这些
				otherCount++
			}
		}else{//该类型的牌三张以上
			
		}

	}

	return false
}

//类型判断 返回下标
func Similar(a []int,index int) int {
	length := len(a)
	//if index < length{//存在 多于
	//
	//}

	for ; index < length ; index++{
		if (a[index] % 100) == (a[index + 1] % 100){//相同类型
			if (index + 1) == length{//仅有两个(剩下) 进行比较
				return length
			} else {
				Similar(a,index + 1)
			}
		}else {//不同类型
			return index + 1
		}
	}
	return length

}

//是否是对子
func IsSame(b []int) bool {
	if len(b) == 2{
		return b[1] - b[0] == 0
	}else {
		return false
	}
}

//是否是暗顺(如1万 2万，1万 3万)
func IsBlackOrder(b []int) bool  {
	if len(b) == 2{
		return b[1] - b[0] == 1 || b[1] - b [0] == 2
	}else {
		return false
	}
}

//三张牌以上 对子数 三个相同的数 顺子数 或者其他
//map 是一个 统计 对子 顺子 暗顺的
//func DealMoreCards(b []int)(mapCards map[int]int){
//
//}

// 顺子 判断


