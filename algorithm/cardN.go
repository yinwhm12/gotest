package algorithm

import "fmt"

/*
根据 鬼(王)牌 数量进行 划分
 */
func ZeroSpecialCards(a []int)  {
	var similarMap map[int]int
	//var countOrderCards int //统计 顺子数目
	length := len(a)
	CountSimilarCards(a,similarMap)//统计 对数
	if len(similarMap) == 1{//一种情况 对子只有一个 //说明可以胡牌
	//	判断是否 是3n
		for i := 0; i < length;{
			//j := i
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
		//fmt.Println("similarCards=",a[index])
		similarMap[a[index]]++
		//if similarMap[a[index]] == 2 {//两个数量
		//	similarMap[2] ++
		//}else if similarMap[a[index]] == 3{//三个 数量
		//	similarMap[3]++
		//	similarMap[2]--
		//}
	}
	//return 0
}

//
func IsHu(a []int,similarMap map[int]int)  {
	length := len(a)
	if length < 3{
		return
	}
	for i := 0 ; i < len(a) ; i++{
		//similarMap[]
	//	三个
		if v, ok := similarMap[a[i]]; ok{
			if v == 3{//是三个

			}else if v == 2 {

			}else {//一个
				//if v1,ok1 := similarMap[a[i]+1]; ok1{
				//	if v1 >0{
				//		similarMap[a[i+1]]--
				//	}else {
				//	//	缺少
				//	}
				//}

			}
		}
	}
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

func TestHu1(a []int,similarMap map[int]int) int {
	count := 0
	for _,v := range a{
		if similarMap[v] == 3{
			if v % 100 == 4{//东南西北
				//count++
				//similarMap[3] --
			}else {
				if similarMap[v+1] > 0 && similarMap[v+2] > 0{
					if similarMap[v+1] == 1&& similarMap[v+2] >= 1{
						if similarMap[v+2] == 1 &&similarMap[v+3] > 0{ //555 6 7 8 d d  后结合
							similarMap[v+1]--
							similarMap[v+2]--
							similarMap[v+3]--
							count++
						}else {//解决 111 2 333 44 55 / 555 6 7 888 前结合
							if similarMap[v+2] == 3{//原来是三个的
								//similarMap[3]--
								//similarMap[2]++
							}
							similarMap[v]--
							similarMap[v+1]--
							similarMap[v+2]--
							count++
							//similarMap[3]--
							//similarMap[2]++
						}
					}else if similarMap[v+1] == 2 && similarMap[v+2] == 2&& similarMap[v+3] == 1{ //555 66 77 8
						similarMap[v]--
						similarMap[v+1]--
						similarMap[v+2]--
						count++
						//similarMap[3]--
						//3个变成2个，所以这里+1 2对的都没了，所以这里-2 所以下面直接-1
						//similarMap[2]--
					}else{//其他情况 不能与其他牌进行有效结合
						//count++
						//similarMap[3] -= 3
					}
				}
			}
		}else if similarMap[v] == 1{
			if similarMap[v+1]  > 0 && similarMap [v+2] > 0{//6 7 8

				//if similarMap[v+1] == 2{//2变1
				//	similarMap[2]--
				//}else if similarMap[v+1] == 3{//3变2
				//	similarMap[3]--
				//	similarMap[2]++
				//}else if similarMap[v+1] == 4{//4变3
				//	similarMap[4]--
				//	similarMap[3]++
				//}
				//
				//if similarMap[v+2] == 2{
				//	similarMap[2]--
				//}else if similarMap[v+2] == 3{
				//	similarMap[3]--
				//	similarMap[2]++
				//}else if similarMap[v+2] == 4{
				//	similarMap[4]--
				//	similarMap[3]++
				//}
				similarMap[v] --
				similarMap[v+1] --
				similarMap[v+2] --
				count++

			}
		}else if similarMap[v] == 2 {
			if similarMap[v+1] >= 2 && similarMap[v+2] >= 2{
				//if similarMap[v+1] == 3{
				//	similarMap[3]--
				//}else if similarMap[v+1] == 4{
				//	similarMap[4]--
				//	similarMap[2]++
				//}
				//if similarMap[v+2] == 3{
				//	similarMap[3]--
				//}else if similarMap[v+2] == 4{
				//	similarMap[4]--
				//	similarMap[2]++
				//}
				similarMap[v] = similarMap[v] - 2
				similarMap[v+1] = similarMap[v+1] - 2
				similarMap[v+2] = similarMap[v+2] - 2
				//similarMap[2] -= 3
				count += 2
			}
		}else if similarMap[v] == 4{
			if similarMap[v+1]>=1 && similarMap[v+2]>=1{
				similarMap[v]--
				similarMap[v+1]--
				similarMap[v+2]--
				count++
			}
		}else {//异常

		}
	}
	return count
}

func EndSum(similarMap map[int]int) (a [4]int) {
	for i, v := range similarMap{
		if i != 1000{//去掉王牌的统计 排除 m[1000]会添加到a[]的情况里
			if v == 0 {
				a[0]++
			}else if v == 1  {
				a[1]++
			}else if v == 2{
				a[2]++
			}else if v == 3 {
				a[3]++
			}
		}
	}
	return
}

func TestHu(similarMap map[int]int)  int{
	count := 0
	for i,v := range similarMap{
		//fmt.Println("testhu---------")
		if v == 3{
			if i % 100 == 4{//东南西北
				//count++
				//similarMap[3] --
			}else {
				if similarMap[i+1] > 0 && similarMap[i+2] > 0{
					if similarMap[i+1] == 1&& similarMap[i+2] >= 1{
						if similarMap[i+2] == 1 &&similarMap[i+3] > 0{ //555 6 7 8 d d  后结合
							similarMap[i+1]--
							similarMap[i+2]--
							similarMap[i+3]--
							count++
						}else {//解决 111 2 333 44 55 / 555 6 7 888 前结合
							if similarMap[i+2] == 3{//原来是三个的
								similarMap[3]--
								similarMap[2]++
							}
							similarMap[i]--
							similarMap[i+1]--
							similarMap[i+2]--
							count++
							similarMap[3]--
							similarMap[2]++
						}
					}else if similarMap[i+1] == 2 && similarMap[i+2] == 2&& similarMap[i+3] == 1{ //555 66 77 8
						similarMap[i]--
						similarMap[i+1]--
						similarMap[i+2]--
						count++
						similarMap[3]--
						//3个变成2个，所以这里+1 2对的都没了，所以这里-2 所以下面直接-1
						similarMap[2]--
					}else{//其他情况 不能与其他牌进行有效结合
						//count++
						//similarMap[3] -= 3
					}
				}
			}
		}else if v == 1{
			if similarMap[i+1]  > 0 && similarMap [i+2] > 0{//6 7 8
				//if similarMap[i+1] ==1 && similarMap[i+2] ==1{
				//	count++
				//	delete(similarMap,i)
				//	delete(similarMap,i+1)
				//	delete(similarMap,i+2)
				//}else {
				//
				//}
				similarMap[i] --
				similarMap[i+1] --
				similarMap[i+2] --
				count++

			}
		}else if v == 2 {
			if similarMap[i+1] >= 2 && similarMap[i+2] >= 2{
				similarMap[i] = similarMap[i] - 2
				similarMap[i+1] = similarMap[i+1] - 2
				similarMap[i+2] = similarMap[i+2] - 2
				similarMap[2] -= 3
				count += 2
			}
		}
	}
	return count
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

//统计牌 包括王牌 新生成一个数组 去掉王牌的数组 maxCard 为王牌
func CountCards(a []int,similarMap map[int]int, maxCard int)(b []int)  {
	length := len(a)
	countMax := 0
	for index := 0; index < length; index++{
		//if _, ok := similarMap[a[index]]; ok{
		//	similarMap[a[index]]++
		//}else {
		//
		//}
		//fmt.Println("similarCards=",a[index])
		if a[index] == maxCard{//是 王牌
			countMax ++
		}else {
			b = append(b,a[index])//重新生成的数组
			similarMap[a[index]]++
		}
		//if similarMap[a[index]] == 2 {//两个数量
		//	similarMap[2] ++
		//}else if similarMap[a[index]] == 3{//三个 数量
		//	similarMap[3]++
		//	similarMap[2]--
		//}
	}
	similarMap[1000] = countMax//用1000来存放王牌
	return
	//return 0
}

//m=1时 王牌一张时
func OneSpecialCards(similarMap map[int]int,b []int,length int)  {
	count := TestHu1(b,similarMap)
	//fmt.Println("count==",count)
	c := EndSum(similarMap)
	//fmt.Println("c=",c)
	m := similarMap[1000]
	left := count *3 + c[2] *2 + c[3] * 3
	if c[2] == 0 && c[1] == 1 {//没有对子
		if length - left == 1 + m{
			fmt.Println("hu 一对 单一 王牌一")
		}else {
			fmt.Println("no hu")
		}
	}else if c[2] == 1 && c[1] == 2{//有一个对子 且 两个单的
		//两个单 与一个王牌 且 两个单为有序的
		if length - left == 2 + m && WacthBlackOrder(b,similarMap) == m{
			fmt.Println("hu 暗顺1 王牌一")
		}else {
			fmt.Println("no hu")
		}
	}else if c[2] == 2 && c[1] == 0{//两个对子 且没有单的
		if length - left == m{
			fmt.Println("hu 两对")
		}else{
			fmt.Println("no hu")
		}
	}else {
		fmt.Println("no hu" )
	}

}

//m =2 是 王牌数量为2 时
func TwoSpecialCards(similarMap map[int]int,b []int,length int)  {
	count := TestHu1(b,similarMap)
	c := EndSum(similarMap)
	//fmt.Println("2--c=",c," count=",count)
	m := similarMap[1000]
	left := count *3 + c[2] *2 + c[3] * 3
	if c[2] == 0 && c[1] == 0{//缺少对子 俩个王牌当
		if length - left == m {
			fmt.Println("hu 缺一对 双王")
		}else {
			fmt.Println("no hu")
		}
	}else if  c[2] == 0 && c[1] == 3 && WacthBlackOrder(b,similarMap) == 1{//3个单的 但是其中两个为有序的
		if length - left == c[1] + m{
			fmt.Println("hu 3单 其中一暗")
		}else {
			fmt.Println("no hu")
		}
	}else if c[2] == 1 && c[1] == 4 && WacthBlackOrder(b,similarMap) == m{//两个暗子 分别 有顺序
		if length - left == c[1] + m{
			fmt.Println("hu 4单 其中有两队 暗顺序")
		}else {
			fmt.Println("no hu")
		}
	}else if c[2] == 2 && c[1] ==2 && WacthBlackOrder(b,similarMap) == 1{//1个暗子 两个对子
		if length - left == c[1] + m{
			fmt.Println("hu 两个单 1个暗子 两个对子")
		}else{
			fmt.Println("no hu" )
		}
	}else if c[2] == 3 && c[1] == 0 {
		if length - left == m{
			fmt.Println("hu 3个对子")
		}else{
			fmt.Println("no hu")
		}
	}else if  c[2] == 1 && c[1] == 1{//剩下一张单
		if length - left == c[1]+m{
			fmt.Println("hu 剩下一张单牌")
		}else {
			fmt.Println("no hu")
		}
	}else{
		fmt.Println("no hu")
	}
}

//王牌 m=3
func ThreeSpecialCards(similarMap map[int]int,b []int,length int)  {
	count := TestHu1(b,similarMap)
	c := EndSum(similarMap)
	m := similarMap[1000]
	left := count *3 + c[2] *2 + c[3]*3
	if c[2] == 1{//一个对子存在的情况
		if c[1] == 0&&length - left == m{//全齐全
			fmt.Println("hu 三张王牌 可以成3个或者顺子")
		}else if c[1] == 6 && WacthBlackOrder(b,similarMap) == m && length -left == c[1]+m{
		//	三个暗子
			fmt.Println("hu 三个暗子 补 3个顺子")
		}else{
			fmt.Println("no hu")
		}
	}else if c[2] == 2 {//两个对子存在的情况
		if c[1] == 4&& length - left == c[1] + m && WacthBlackOrder(b,similarMap) == 2{
		//	存在2个暗子
			fmt.Println("hu 两个对子 2个暗子")
		}else if c[1] ==1 && length - left == c[1] + m{
		//	存在单一个
			fmt.Println("hu 两个对子 1个单")
		} else{
			fmt.Println("no hu")
		}
	}else if c[2] == 3 {
		if c[1] == 2&& length - left == c[1] +m &&WacthBlackOrder(b,similarMap) == 1{
		//	存在一个暗子
			fmt.Println("hu 三个对子 且 一个暗子")
		}else{
			fmt.Println("no hu")
		}
	}else if c[2] == 4{
		if c[1] == 0 && length -left == m{
			fmt.Println("hu 4个对子")
		}else{
			fmt.Println("no hu")
		}
	}else if c[2] == 0{
		if c[1] == 2 && length - left == c[1]+m && WacthBlackOrder(b,similarMap) == 1{
		//	存在 一个暗子 无对子
			fmt.Println("hu 一个暗子 无对子")
		}else if c[1] == 5 && length - left == c[1] +m && WacthBlackOrder(b,similarMap) == 2 {
		//	存在 两个暗子 一个单
			fmt.Println("hu 两个暗子 一个单")
		}else {
			fmt.Println("no hu" )
		}
	}else{
		fmt.Println("no hu")
	}
}

//判断 暗顺 参数hasCount 为有多少个暗顺 标记过的 单 将其k-v值 增 10 以区别
func WacthBlackOrder(b []int,similarMap map[int]int) (hasCount int){
	for _, v := range b{
		if similarMap[v] == 1 && similarMap[v+1] == 1{
			similarMap[v] = 1 + 10
			similarMap[v+1] = 1 + 10
			hasCount++
		}else if similarMap[v] == 1&&similarMap[v+2] == 1 {
			similarMap[v] = 1 + 10
			similarMap[v+2] = 1 + 10
			hasCount++
		}
	}
	return hasCount
}