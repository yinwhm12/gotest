package algorithm

import (
	"fmt"
	"sort"
	"mime/multipart"
)

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

//统计 能形成顺子的数量(原始牌进行统计)
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

//统计牌的数量 单张的多少 对的多少 三个的多少
func EndSum(similarMap map[int]int) (a [4]int) {
	for i, v := range similarMap{
		if i != 1000{//去掉王牌的统计 排除 m[1000]会添加到a[]的情况里
			if v == 0 {
				a[0]++
			}else if v == 1  {//单的
				a[1]++
			}else if v == 2{//对的
				a[2]++
			}else if v == 3 {//三个的
				a[3]++
			}
		}
	}
	return
}

//统计原始的每种牌的数量
func EndOriginalCountCards(similarMap map[int]int) (a [5]int) {
	for i, v := range similarMap{
		if i != 1000{//去掉王牌的统计 排除 m[1000]会添加到a[]的情况里
			if v == 0 {
				a[0]++
			}else if v == 1  {//单的
				a[1]++
			}else if v == 2{//对的
				a[2]++
			}else if v == 3 {//三个的
				a[3]++
			}else if v == 4 {
				a[4]++
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
			if a[index] == 405{//红中记录
				similarMap[1001]++// 1001放置红中的数量
			}
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
func OneSpecialCards(similarMap map[int]int,b []int,c [4]int,length,count int)  {
	//count := TestHu1(b,similarMap)
	//fmt.Println("count==",count)
	//c := EndSum(similarMap)
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
func TwoSpecialCards(similarMap map[int]int,b []int,c [4]int,length,count int)  {
	//count := TestHu1(b,similarMap)
	//c := EndSum(similarMap)
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
//参数
func ThreeSpecialCards(similarMap map[int]int,b []int,c [4]int,length,count int,)  {
	//count := TestHu1(b,similarMap)
	//c := EndSum(similarMap)
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
			fmt.Println("hu 两个暗子 5个单")
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



//七对的情况
//c[] 为统计牌数
//3张王牌
func HuBySevenThreeW(length int,c [5]int)  {
	//length := len(a)
	//c
	//left := length - c[2] * 2
	if length == 14{
		switch c[2] {
			case 4:
				if c[3]==1{
				//	一个杠的七对 龙巧对
					fmt.Println("hu 已有4个对 有3个1 可以形成一个杠 龙巧对")
				}else if c[1] == 3{
					//七巧对
					fmt.Println("hu 已有4个对,3个单 王钓七巧对")
				}else {

				}
			case 3:
				if c[3] == 1&& c[1] == 2{//有一个杠 七对 王钓龙巧
					fmt.Println("hu 已有3个 且3个的有一个c[3]==1 单的2个 王钓七巧对")
				}else if c[4] ==1&&c[1] ==1&& c[3]==0{
					fmt.Println("hu 4个的1个 单的一个 王钓龙巧对")
				}else {

				}
			case 2:
				if c[4] == 1{//一个暗杠 龙巧对
					fmt.Println("hu 已有2个 且4个的有一个c[4]==1 王钓龙巧对")
				}else if c[3] == 2 && c[1] == 1{//有两个杠的 11 22 333 444 5 王钓龙巧
					fmt.Println("hu 已有2个 且3个的有2个 1个单 王钓七巧对")
				}else {
					
				}
			case 0:
				if c[4] == 2{
					//有杠的七对
					fmt.Println("hu 没有对的，且有4个的两个 王钓龙巧对")
				}else if c[4] == 1&& c[3] ==2 && c[1] ==1 {
					fmt.Println("hu 没对 4个的1个 3个的2个 单一个 王钓龙巧对")
				}else {

				}
			case 1:
				if c[3] == 3{ //111 111 111 11
					//有杠的 七对 王钓龙巧
					fmt.Println("hu 有一对 且3个3的 王钓七巧对")
				}else if c[4] == 2 && c[1] == 1 {
					fmt.Println("hu 有一对 4个的2个 单个1个 王钓龙巧对")
				}else if c[4] ==1 && c[3] ==1 && c[1] ==2 {
					fmt.Println("hu 有一对 4个的1个 3个的1个 单个的2个 王钓龙巧对")
				}else {

				}
			case 5:
				if c[1] == 1{//七巧对
					fmt.Println("hu 已有5对 + 1个单 王钓七巧对")
				}
			default:
				fmt.Println("no hu")
		}
	}else {//不能胡

	}
}

//王牌 为2时 能胡的七对
func HuBySevenTwoW(length int,c [5]int)  {
	if length == 14{
		switch c[2] {
		case 6:
			if c[1] == 0{
				fmt.Println("hu 双王当 6对 王钓七巧对")
			}else {
			}
		case 5:
			if c[1] == 2 {
				fmt.Println("hu  5对 单个的2 双王 王钓七巧对")
			}else {

			}
		case 4:
			if c[3] == 1&& c[1] ==1{//
				fmt.Println("hu 4对 三个的一个 王钓七巧对")
			}else if c[4] == 1 && c[1] == 0{
				fmt.Println("hu 4对 4个的一个 王钓龙巧对")
			}else {

			}
		case 3:
			if c[3] == 2 && c[1] == 0{
				fmt.Println("hu 3对 3个的2个 王钓七巧对")
			}else if c[4] == 1 && c[1] == 2{
				fmt.Println("hu 3对 4个的一个 2个单 王钓龙巧对")
			}else {

			}
		case 2:
			if c[4] == 2 && c[3] == 0 && c[1] == 0{
				fmt.Println("hu 2对 4个的2个 单的2个 王钓龙巧对")
			}else if c[4] == 1 && c[3] == 1&& c[1] == 1{
				fmt.Println("hu 2对 4个的1个 3个的1个 1个单的 王钓龙巧对")
			}else {

			}
		case 1:
			if c[4] == 2 && c[3] == 0 && c[1] == 2{
				fmt.Println("hu 1对 4个的2个 3个的0 单的2个 王钓龙巧对")
			}else if c[4] == 1&& c[3] == 2&& c[1] == 0{
				fmt.Println("hu 1对 4个的1个 3个的2个 单的0个 王钓龙巧对")
			}else {

			}
		default:
			if c[4] == 3&& c[3] ==0 &&c[1] ==0{
				fmt.Println("hu 4个的3个 双王 王钓龙巧对")
			}else if c[4] ==2 && c[3] ==1 && c[1] ==1{
				fmt.Println("hu 4个的2个 3个的1个 单一个 王钓龙巧对")
			}else {
				fmt.Println("no hu")
			}
		}
	}else {
		fmt.Println("no hu")
	}
}

//王牌为1时
func HuBySevenOneW(length int, c [5]int)  {
	if length == 14{
		switch c[2] {
		case 6:
			if c[4] == 0&& c[3] == 0&& c[1] == 1{
				fmt.Println("hu 6对 单一个 王牌1 6对 王钓七巧对")
			} else {

			}
		case 5:
			if c[4] == 0&& c[3] ==1 && c[1] ==0 {
				fmt.Println("hu 5对 3个一个 王牌1 5对 王钓七巧对")
			}else {

			}
		case 4:
			if c[4] == 1&& c[3] ==0 && c[1] ==1 {
				fmt.Println("hu 4对 4个的一个 单的一个 王牌1 4对 王钓龙巧对")
			}else {

			}
		case 3:
			if c[4] == 1 && c[3] == 1&& c[1] ==0{
				fmt.Println("hu 4个的1 3个的1 王牌1 3对 王钓龙巧对")
			}else {

			}
		case 2:
			if c[4] == 2 && c[3] == 0 && c[1] ==1{
				fmt.Println("hu 4个的2个 单的1个 王牌1 2对 王钓龙巧对")
			}else {
			}
		case 1:
			if c[4] ==2 && c[3] ==1 && c[1] ==0{
				fmt.Println("hu 4个的2个 3个的1个 王牌1 1对 王钓龙巧对")
			}else {

			}
		default:
			if c[4] ==3 && c[3] ==0&& c[1] ==1{
				fmt.Println("hu 4个的3个 单个的1 王牌1 王钓龙巧对")
			}else {
				fmt.Println("no hu")
			}
		}
	}else {

	}
}

// 胡牌 没有王牌
func HuBySevenZeroW(length int, c [5]int)  {
	if length == 14{
		switch c[2] {
		case 7:
			if c[4] == 0&& c[3]== 0 && c[1] ==0{
				fmt.Println("hu 7对 没有王牌 硬巧对")
			}else{

			}
		case 5:
			if c[4] == 1 && c[3] == 0&& c[1]== 0{
				fmt.Println("hu 5对 4个的1个 没有王牌 硬巧对+龙巧对")
			}else{

			}
		case 3:
			if c[4] == 2 && c[3] == 0 && c[1] ==0{
				fmt.Println("hu 3对 4个的2个 没有王牌  硬巧对+龙巧对")
			}else{

			}
		case 1:
			if c[4] == 3 && c[3] == 0&& c[1]==0{
				fmt.Println("hu 1对 4个额3个 没有王牌 硬巧对+龙巧对")
			}else{
				fmt.Println("no hu")
			}
		}
	}else {
		fmt.Println("no hu")
	}
}


//十三烂 情况
func HuByBadCards(a []int,similarMap map[int]int,count int,c [4]int)  {
	if count == 0{
		if c[2] == 0 && c[3] == 0&& c[1] == len(a){
			blackCount := WacthBlackOrder(a,similarMap)
			if blackCount == 0{//没有暗顺子
				sevenCount := 0
				for {//统计 东南西北中发白
					if similarMap[401+sevenCount] == 1{
						sevenCount++
					}else {
						break
					}
				}
				if sevenCount == 7{
					fmt.Println("七风到位⼗三烂")
				}else{
					fmt.Println("普通的十三烂")
				}
			}else {
				fmt.Println("no hu 13")
			}
		}
	}
}

//没有王牌的胡
//参数 length 长度 count 顺子数
func HuByOrdinaryCards(length,count,maxCard int, c [4]int,similarMap map[int]int,a []int)  {
	left := count *3 + c[2] *2 + c[3] *3
	if length - left == 0 && c[2] == 1 && c[1] == 0{
		if similarMap[1001] == 0{
			fmt.Println("hu l 大道")
		}else {
			fmt.Println(" hu 普通")
		}
	}else if similarMap[1001] > 0{//有红中的情况
		//b := a//复制a
		b := make([]int,len(a))
		copy(b,a)
		if maxCard == 405{//王牌是 中 不要变化

		}else {
			for i,v := range b  {
				if v == 405{
					b[i] = maxCard
				}
			}
			sort.Ints(b)
			m := make(map[int]int)
			CountCards(b,m,maxCard)
			newCount := TestHu1(b,m)
			cc := EndSum(m)
			if length !=2 {//相当于 新的 牌型 但是 始终是 没有王牌的
				newLeft := newCount * 3 + cc[2] *2 + c[3]*3
				if length - newLeft == 0&& cc[2] == 1&& cc[1] == 0{
					fmt.Println("hu 小道")
				}

			////	七对 不能 替换的
			//	go func() {
			//		HuByBadCards(a,m,newCount,cc)//十三烂
			//	}()
			//	以上两个协程 可以同时 胡的情况
			}else{//一张手牌
				if cc[2] == 1{
					fmt.Println("hu 全求人 + 小道 没有王牌")
				}
			}
		}
		fmt.Println("no hu ordinary")
	}else {
		fmt.Println("no hu ordinary")
	}
}

//统计 牌 色(类型)
//判断 是否是清一色 先去掉 王牌 在统计 
//返回 bool    b[] 去掉 王牌的数组
func IsOneColorCards(b []int) bool  {
	length := len(b)
	count := 0
	for i, v := range b{
		//similarMap[v]
		if i + 1 < length{
			if v / 100 == b[i+1] / 100{
				count++
			}else {
				break
			}
		}
	}
	return length - count == 0
}

// 有一套 或者 2套以上
//胡牌后才触发
// 一套 至少 2个对子     ×
// 两套 至少 4个对子	×
//func HowManyTao(a []int) int  {
//	if len(a) < 8{//少于 8张手牌 无法成立 套型
//
//	}else{
//		m := make(map[int]int)
//		b := CountCards(a,m,0)//不排除王牌 连王牌也统计上
//		for _, v := range a{
//			if m[v] == 2{
//				if m[v-1] == 1{
//					if m[v+1] >= 2{
//						m[v-1]--
//						m[v] -= 2
//						m[v+1] -= 2
//					}
//				}else if m[v+1] >= 2 {
//					if m[v+2]== 1{
//						m[v] -= 2
//						m[v] -= 2
//					}
//					m[v] -= 2
//					m[v + 1] -= 2
//					m[v + 2] -= 2
//				}
//			}
//		}
//
//	}
//}

//参数 a[]原始数组 similarMap 经过取顺子后的 flag 是否是3n*2胡的类型
func Tao(a []int, similarMap map[int]int,flag bool) int {
	count := 0//记数
	//flag := true
	if len(a) < 8{//手牌 少于8 不能形成 套型
		return count
	}
	m := make(map[int]int)
	CountCards(a,m,405)
	fmt.Println("map =m =",m)
	c := EndOriginalCountCards(m)
	fmt.Println("c =",c)
	if c[2] + c[3] + c[4] < 2{ //手牌上的对子少于 2对 不能形成 套型
		return count
	}
	if len(a) == 14 && !flag{
		similarMap = m
	}
	for _, v := range a{
		fmt.Println("v=",v)
		if m[v] >=2 && m[v+1] >= 2{
			fmt.Println("in-first v=",v)
					if flag{
						//胡 不是七对胡 3n+2 >= 8
						if vv, ok := similarMap[v]; ok {//
							fmt.Println("vv,",vv)
							if vvv, ok1 := similarMap[v+1]; ok1 {
								fmt.Println("vvv",vvv)
								//if m[v] - vv >= 2 && m[v+1] - vvv >=2{
								if m[v-1] == 1|| m[v+2] >0{
									m[v] -= 2
									m[v+1] -= 2
									count ++
								}
								//}
							}
						}
					}else{
					//	是 七对胡
						count++
						m[v] -=2
						m[v+1] -=2
					}
		}else if m[v] >= 2&& m[v+2] >=2 {
			if !flag{//七对胡
				if m[v + 1] >0{
					m[v] -= 2
					m[v+1] -= 2
					count++
				}
			}else {
				if vv, ok := similarMap[v]; ok {
					if vvv, ok1 := similarMap[v+2]; ok1 {
						if m[v] - vv >= 2 && m[v+2] - vvv >=2{
							count ++
							m[v] -= 2
							m[v+1] -= 2
						}
					}
				}
			}
		}
	}
	//for _, v := range m {
	//	fmt.Println("v--v=",v)
	//	if m[v] >=2 && m[v+1] >= 2{
	//		if vv, ok := similarMap[v]; ok {
	//			if vvv, ok1 := similarMap[v+1]; ok1 {
	//				if m[v] - vv >= 2 && m[v+1] - vvv >=2{
	//					count ++
	//					m[v] -= 2
	//					m[v+1] -= 2
	//				}
	//			}
	//		}
	//	}
	//}
	return count
}


//套类型
//手牌 8张
//参数count 顺子数，maxCard 王牌值	a[] 原始数据(手牌)
func TaoByEightCards(count,maxCard int, a []int)  {
	m := make(map[int]int)
	CountCards(a,m,maxCard)
	fmt.Println("map =m =",m)
	c := EndOriginalCountCards(m)
	for _,v := range a{
		if (m[v] >= 2 &&m[v+1] >=2) ||(m[v]>=2 && m[v+2] >=2){//对子的才有 22 33 / 22 44 类型的
			if count == 2{//两个顺子
				fmt.Println("有一套")
			}else if m[1000] == 3{
				fmt.Println("有一套")
			} else if m[1000] == 2 {
				if c[2] == 3{//三个对子
					fmt.Println("有一套")
				}else if c[2] == 1&&c[4] == 1 {//2222 33 _ _
					fmt.Println("有一套")
				} else{
					if GetOrderCounts(v,m){//
						fmt.Println("有一套")
					}else {
						fmt.Println("没有一套")
					}
				}
			}else if m[1000] == 1{
			//	只要 有一个暗顺 即可
				if GetOrderCounts(v,m){//1 22 33 77 _
					fmt.Println("有一套")
				}else {
					fmt.Println("没哟一套")
				}
				
			}else {
				fmt.Println("没有一套")
			}
		}
	}
}

//11张牌
//  count 顺子的数量(胡后)
func TaoByElevenCards(count,maxCard int,a []int,similarMap map[int]int)   {
	m := make(map[int]int)
	CountCards(a,m,maxCard)
	fmt.Println("map =m =",m)
	c := EndOriginalCountCards(m)
	for _, v := range a{
		if (m[v] >= 2 &&m[v+1] >=2) ||(m[v]>=2 && m[v+2] >=2){
			if count == 3{
				fmt.Println("有一套")
			}else if m[1000] == 3 {
				if count == 2|| c[3] ==2{
					fmt.Println("有一套")
				}else if GetOrderCounts(v,m){//有一个顺子即可 要对子参与
					fmt.Println("有一套")
				}else {
					if count == 1 || c[3] == 1 || c[2] == 4{
						fmt.Println("有一套")
					}else{

					}
				}
			}else if m[1000] == 2{
				if GetOrderCounts(v,m){//参与暗顺
					fmt.Println("有一套")
					break
				}else {//没有参与
					if c[2] == 3{
						if count == 1 || c[3] == 1{
							fmt.Println("有一套")
						}else {

						}
					}else {
						fmt.Println("没有一套")
					}
				}
			}else if m[1000]== 1{
				if GetOrderCounts(v,m){
					fmt.Println("有一套")
				}
			}else{
				fmt.Println("没有一套")
			}
		}
	}
}

//14张牌
//
func TaoByFourteenCards(count,maxCard int,a []int,similarMap map[int]int)  {
	m := make(map[int]int)
	CountCards(a,m,maxCard)
	fmt.Println("map =m =",m)
	c := EndOriginalCountCards(m)
	for _, v := range m{
		if (m[v] >= 2 &&m[v+1]>=2)||(m[v]>=2 && m[v+2] >=2){
			if c[2] >= 2{//有套 基本条件
				if c[2] >= 4{//有两套 基本条件

				}
			}
		}
	}
}

//判断 是否是连续的 对子 基本有顺
func AboveDoubleArray(m map[int]int) (bMap map[int]int){
	//falg := false
	for i,v := range m{
		if v >= 2{
			if m[i+1] >= 2 || m[i+2] >= 2{
					bMap[i] = v
			}
		}
	}
}

// 统计 2

//顺子 计算 不管是对或者三个的 只要符合 都算上
//为了 统计 1套 对子是否参与
func GetOrderCounts(v int,m map[int]int) bool {
	//flag := false
		if m[v+1] >=2 {//顺子的暗顺
			if m[v-1] > 0 || m[v+2] >0 {
			//	有一套 即有一个顺子
				return true
			}else{
				return false
			}
		}else if m[v+2] >= 2{//间隔 的暗顺
			if m[v+1] > 0{
				//fmt.Println("有一套")
				return true
			}else {
				return false
			}
		}else {
			return  false
		}
}