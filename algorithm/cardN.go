package algorithm




/*
根据 鬼(王)牌 数量进行 划分
 */
func ZeroSpecialCards(a []int)  {
	var similarMap map[int]int
	var countOrderCards int //统计 顺子数目
	length := len(a)
	CountSimilarCards(a,similarMap)//统计 对数
	if len(similarMap) == 1{//一种情况 对子只有一个 //说明可以胡牌
	//	判断是否 是3n
		for i := 0; i < length;{
			j := i
			i = FindTypeCards(a,i)
		}

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
		//if _, ok := similarMap[a[index]]; ok{
		//	similarMap[a[index]]++
		//}else {
		//
		//}
		similarMap[a[index]]++
		if similarMap[a[index]] {

		}
	}
	//return 0
}

/*
统计同类型下的 顺子(三个)数目 3n
 */
func CountOrderCards(a []int)(count int)  {
	length := len(a)
	if length < 3{
		count = 0

	}else {
		for i := 0 ; i < length; i = i + 2{
			if a[i + 2] - a[i] == a[i + 1]{//顺子
				count++

			}else if a[i] == a[i + 2]{//三个的
				count++
			}else {//都不是

			}
		}
	}
	return  count
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