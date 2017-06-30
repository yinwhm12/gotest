package main


type HuInfoMSG struct{
//

}

const (
//	胡牌编码为2000 以上的
//	其中 有王牌胡的 从100开始 即2100 有一张王牌 构成的胡牌 (2100-2200)
//	2200 为两张...
	HUZERO_DADAO	=	2001 //没有王牌的胡牌 大道	仅有一对
	HUZERO_OO	=	2002	//普通的胡牌
	HUZERO_XIAODAO	=	2003	//小道
	HUZERO_SINGAL	=	2004	//全求人 + 小道

	HUONE_12_	=	2101	//hu 一对 单一 王牌一
	HUONE_AN_	=	2102	//hu 暗顺1 王牌一
	HUONE_2TWO_	=	2103	//hu 两对

	HUTWO_LOSEDOUBLE_	=	2201	//hu 缺一对 双王当对
	HUTWO_AN_SINGAL_	=	2202//hu 3单 其中一暗 王牌 分别补顺 以及 对
	HUTWO_TWO_AN_	=	2203	//hu 4单 其中有两队 暗顺序
	HUTWO_2DOUBLE_2SINGAL_	=	2204//两个单组成1个暗子 两个对子
	HUTWO_3DOUBLE_	=	2205	//hu 3个对子
	HUTWO_LEFTSINGAL_	=	2206	//hu 剩下一张单牌 已有一对

	HUTHREE_NONEED_	=	2301	//"hu 三张王牌 可以成3个或者顺子
	HUTHREE_3AN_	=	2302	//hu 三个暗子 补 3个顺子
	HUTHREE_2DOUBLE_3AN_	=	2303	//hu 两个对子 2个暗子
	HUTHREE_2DOUBLE_SINGAL_	=	2304	//hu 两个对子 1个单
	HUTHREE_3DOUBLE_AN_	=	2305	//hu 三个对子 且 一个暗子
	HUTHREE_4DOUBLE_	=	2306	//hu 4个对子
	HUTHREE_AN_	=	2307	//hu 一个暗子 无对子
	HUTHREE_2AN_SINGAL	=	2308	//"hu 两个暗子 1个单 即无对子

//	七对标记
	HUTHREE_SEVEN_4DOUBLE_THREE_	=	2401	//hu 已有4个对 有3个的1个 龙巧对
	HUTHREE_SEVEN_4DOUBLE_3SINGAL_	=	2402	//hu 已有4个对,3个单 王钓七巧对

)
