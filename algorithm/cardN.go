package algorithm




/*
根据 鬼(王)牌 数量进行 划分
 */
func ZeroSpecialCards(a []int)  {
	var similarMap map[int]int
	length := len(a)
	for i := 0; i < length;{
		j := i
		i = FindTypeCards(a,i)
		CountSimilarCards(a[j:i],similarMap)
	}
	if len(similarMap) == 1{//一种情况 对子只有一个
	//	判断是否 是3n
	}else {
	//	胡不了
	}

}

/*
统计牌的对数
 */
func CountSimilarCards(a []int,similarMap map[int]int)  {
	length := len(a)
	for index := 0; index < length; index++{
		if _, ok := similarMap[a[index]]; ok{
			similarMap[a[index]]++
		}
	}
	//return 0
}

/*
区分类型
 */
func FindTypeCards(a []int, index int) int {
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