package main

import (
	"gotest/algorithm"
	"fmt"
	"github.com/name5566/leaf/go"
)

func main()  {
	//a := []int{105,105,105,106,107,108,201,202,203,401,401,401,403,403}
	//a := []int{105,105,105,106,107}
	//a := []int{106,107,108,108,108}
	//a := []int{106,107,108,108,108,201,202,203}
	//a := []int{106,106,106,107,107,107,108,108,108,109,109}
	//a := []int{105,105,105,106,107,108,401,401}
	//a := []int{101,101,101,102,103,103,103,104,104,105,105}
	//a := []int{101,102,103,104,105,106,107,108,109,109,109}
	//a := []int{101,101,101,101,102,103,104,104}//1111 2 3 44
	//a := []int{101,101,101,101,102,102,102,103}//1111 222 3
	//a := []int{101,101,101,101,102,102,103,103}//1111 22 33
	//a := []int{101,101,101,101,102,103,103,103} //1111 2 333
	//a := []int{101,101,101,101,102,102,103,103,104,105,105}//1111 22 33 4 55
	//a := []int{101,102,103,204,205,206,303,304,305,306,306}//
	//a := []int{101,102,103,204,205,206,209,303,304,305,306}//1 2 3 四五六 345条 一个王牌 一对 单一 王牌一
	//a := []int{101,102,103,204,205,206,206,207,209,303,304,305,306,306}//暗顺1 王牌一
	//a := []int{101,102,103,204,205,206,206,206,209,303,304,305,306,306}//两对

	//a := []int{101,102,103,204,205,206,209,209,303,304,305}//缺一对 双王
	//a := []int{101,102,103,103,104,106,108,204,205,206,209,209,303,303}//4单 其中2暗
	//a := []int{101,102,103,103,104,106,106,204,205,206,209,209,303,303}//两个单 1个暗子 两个对子
	//a := []int{101,102,103,103,103,106,106,204,205,206,209,209,303,303}//3个对子
	//a := []int{101,102,103,103,204,205,206,209,209,303,303}//剩下一张单牌

	//a := []int{101,102,103,204,205,206,209,209,209,303,303}//三张王牌 可以成3个或者顺子
	//a := []int{103,104,203,204,205,206,208,209,209,209,301,301,401,401}// 两个对子 2个暗子
	//a := []int{103,203,204,205,209,209,209,301,301,401,401}// 两个对子 1个单
	//a := []int{103,103,203,204,205,206,207,209,209,209,301,301,401,401}// 三个对子 且 一个暗子
	//a := []int{103,103,203,204,205,206,206,209,209,209,301,301,401,401}// 4个对子
	//a := []int{203,204,205,206,207,209,209,209}//  一个暗子 无对子
	//a := []int{203,204,205,206,207,209,209,209,302,304,401}// 两个暗子 一个单

	//a := []int{102,103,201,201,201,203,204,205,305,305,402,402,402,405}// 23万 3个一同，345同，一对5条，3个南 一个王牌(405)
	//a := []int{106,107,108,207,208,209,301,301,302,303,304,309,309,309}// 678w 789t 11d 234d 999d
	//a := []int{103,104,105,206,207,208,301,302,303,305,306,307,401,405}// 345w 678t 123d 567d 东中(王)
	//a := []int{101,101,202,203,204,302,302,302,309,309,309,401,401,405}// 11w 234t 222d 999d 东东 中(王)
	//a := []int{109,109,109,204,204,301,301,301,307,307,307,309,309,309}// 999w 44t 111d 777d 999d
	//a := []int{101,101,101,101,207,207,207,207,208,208,209,301,405,405}

	//a := []int{102,103,103,104,104,207,207,208,208,208,301,301,405,405}
	a := []int{103,103,104,104,207,207,208,208,208,405,405}
	m := make(map[int]int)
	w := 405 //王牌
	//algorithm.CountSimilarCards(a,m)
	//fmt.Println()
	//fmt.Printf("m = %v\n",m)
	b := algorithm.CountCards(a,m,w)
	o := algorithm.EndOriginalCountCards(m)//统计原始的每种牌的数量
	fmt.Println("b=",b)
	fmt.Println("oo=",o)
	fmt.Println("f-m=",m)
	length := len(a)
	count := algorithm.TestHu1(a,m)
	c := algorithm.EndSum(m)//统计变化后的每种牌 的数量
	d := g.New(10)
	//chanC := d.NewLinearContext()
	switch m[1000] {
		case 1:
			if length == 2{//一张手牌
				fmt.Println("hu 全求人")
			}else {//
				d.Go(func() {
					algorithm.OneSpecialCards(m,b,c,length,count)
				}, nil)

				d.Go(func() {
					algorithm.HuBySevenOneW(length,o)
				},nil)
				//algorithm.OneSpecialCards(m,b,c,length,count)
				//algorithm.HuBySevenOneW(length,o)
				d.Cb(<-d.ChanCb)
				d.Cb(<-d.ChanCb)

			}
			//algorithm.OneSpecialCards(m,b,c,length,count)
		case 2:
			//d.Go(func() {
			//	algorithm.TwoSpecialCards(m,b,c,length,count)
			//},nil)
			//d.Go(func() {
			//	algorithm.HuBySevenTwoW(length,o)
			//},nil)
			//d.Cb(<-d.ChanCb)
			//d.Cb(<-d.ChanCb)
			algorithm.TwoSpecialCards(m,b,c,length,count)
			algorithm.HuBySevenTwoW(length,o)


		case 3:
			//algorithm.ThreeSpecialCards(m,b,c,length,count)
			////同时进行看是否能有两种胡牌
			//algorithm.HuBySevenThreeW(length,o)
			//d.Go(func() {
			//	algorithm.ThreeSpecialCards(m,b,c,length,count)
			//},nil)
			//d.Go(func() {
			//	algorithm.HuBySevenThreeW(length,o)
			//},nil)
			algorithm.ThreeSpecialCards(m,b,c,length,count)
			algorithm.HuBySevenThreeW(length,o)

	default:
		//if c[1] != 0{
		//	fmt.Println("no hu")
		//}else {
			//if count == 0{//没有顺子
			//	//下面始 十三烂情况
			//	if c[2] == 0 && c[3] == 0&& c[1] == length{//没有对子 三个的
			//		//algorithm.HuByBadCards(a,m)//十三烂
			//		//blackCount := algorithm.WacthBlackOrder(a,m)
			//		//if blackCount == 0{//没有暗顺子
			//		//	sevenCount := 0
			//		//	for {//统计 东南西北中发白
			//		//		if m[401+sevenCount] == 1{
			//		//			sevenCount++
			//		//		}else {
			//		//			break
			//		//		}
			//		//	}
			//		//	if sevenCount == 7{
			//		//		fmt.Println("七风到位⼗三烂")
			//		//	}else{
			//		//		fmt.Println("普通的十三烂")
			//		//	}
			//		//}else {
			//		//	fmt.Println("no hu")
			//		//}
			//	}else {
			//		//没有顺子 但是有这种情况 111 111 111 11
			//		left := count *3 + c[2] *2 + c[3] *3//count == 0
			//		if length - left == 0 && c[2] == 1{
			//			fmt.Println("hu l")
			//		}else {
			//			fmt.Println("no hu")
			//		}
			//		//fmt.Println("no hu")
			//	}
			//}else {//有顺子
			//	fmt.Println("没有王牌 ",c)
			//	left := count *3 + c[2] *2 + c[3] *3
			//	if length - left == 0 && c[2] == 1{
			//		fmt.Println("hu l")
			//	}else {
			//		fmt.Println("no hu")
			//	}
			//}
		if length !=2{
			algorithm.HuByOrdinaryCards(length,count,w,c,m,a)//普通胡牌
			algorithm.HuBySevenZeroW(length,o)//七对 胡牌
			algorithm.HuByBadCards(a,m,count,c)//十三烂 胡牌
		}else {//一张手牌
			if c[2] == 1{
				fmt.Println("hu 全求人 没有王牌")
			}
		}
	//	}
	}
	//if m[1000] > 0{//有王牌
	//
	//}else {//没有王牌
	//	count := algorithm.TestHu1(a,m)
	//	c := algorithm.EndSum(m)
	//	if c[1] != 0{
	//		fmt.Println("no hu")
	//	}else {
	//		left := count *3 + c[2] *2 + c[3] *3
	//		if length - left == 0{
	//			fmt.Println("hu l")
	//		}else {
	//			fmt.Println("no hu")
	//		}
	//	}
	//}
}
