package card
import (
	"sort"
	"github.com/name5566/leaf/log"
	"math/rand"
)

/**
 * WuMing
 *2017/5/24 下午7:01
 *
 */

var (
	g_NeedHunCount int = 4
	mahjongArr         = []int{
		101, 102, 103, 104, 105, 106, 107, 108, 109, //#万
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		101, 102, 103, 104, 105, 106, 107, 108, 109,
		201, 202, 203, 204, 205, 206, 207, 208, 209, //#饼
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		201, 202, 203, 204, 205, 206, 207, 208, 209,
		301, 302, 303, 304, 305, 306, 307, 308, 309, //#条
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		301, 302, 303, 304, 305, 306, 307, 308, 309,
		401, 402, 403, 404, 405, 406, 407, //# 东 西 南 北 中 发 白
		401, 402, 403, 404, 405, 406, 407,
		401, 402, 403, 404, 405, 406, 407,
		401, 402, 403, 404, 405, 406, 407,
	}
	callTime int = 0
)

/**
 吴名
 2017/5/22 下午2:43
 洗牌
 */
func ShuffleCards() []int {
	return Shuffle1(mahjongArr)
}
func Shuffle1(items []int) []int {
	i := len(items) - 1
	for i >= 1 {
		j := rand.Int63n(int64(i + 1))
		items[i], items[j] = items[j], items[i]
		i--
	}

	return items
}


/**
 * 吴名
 *2017/5/23 下午8:04
 *根据类型分组(万,条,筒,风,真宝)  ++++++++瑞金麻将中还有 花牌++++++++++
 */
func seperateArr(mjArr []int, hunMj int) [][]int {
	reArr := make([][]int, 5)
	ht := hunMj / 100 //确定类型等
	hv := hunMj % 10  //确定几万
	for _, mj := range mjArr {
		t := mj / 100
		v := mj % 10
		if ht == t && hv == v {
			t = 0 //真宝index:0数组
		}
		reArr[t] = append(reArr[t], mj)
		//append之后重新排序
		//升序排列
		sort.Sort(sort.IntSlice(reArr[t]))
	}
	return reArr
}

/**
 吴名
 2017/5/24 下午8:53
 判断顺子
 */
func test3Combine(mj1, mj2, mj3 int) bool {
	t1, t2, t3 := mj1/100, mj2/100, mj3/100
	//牌型不同不能组合
	if t1 != t2 || t1 != t3 {
		return false
	}
	v1, v2, v3 := mj1%10, mj2%10, mj3%10
	// 重牌
	if v1 == v2 && v1 == v3 {
		return true
	}
	if t3 == 4 {
		return false
	}
	if (v1+1) == v2 && (v1+2) == v3 {

		return true
	}
	return false
}

/**
 * 吴名
 *2017/5/23 下午6:00
 *判断3n+2模型中,取模后,组成一对将或一个扑,至少需要的癞子数量
 */
func getModNeedNum(arrLen int, isJiang bool) int {
	if arrLen <= 0 {
		return 0
	}

	modNum := arrLen % 3
	needNumArr := [3]int{0, 2, 1}
	if isJiang {
		needNumArr = [3]int{2, 1, 0}
	}
	return needNumArr[modNum]
}

/**
 * Zhou Zihao
 *2017/5/23 11:30
 *一组牌中需要多少个赖子
 */
func getNeedHunInSub(subArr []int, hNum int) {
	callTime += 1

	if g_NeedHunCount == 0 {
		return
	}

	lArr := len(subArr)

	if hNum+getModNeedNum(lArr, false) >= g_NeedHunCount {
		return
	}

	if lArr == 0 {
		g_NeedHunCount = Min(hNum, g_NeedHunCount) //获取最小数
		return
	} else if lArr == 1 {
		g_NeedHunCount = Min(hNum+2, g_NeedHunCount)
		return
	} else if lArr == 2 {
		t := subArr[0] / 100 //获取牌类型(万条筒等)
		v0 := subArr[0] % 10 //确定为
		v1 := subArr[1] % 10
		//牌为东南西北中发白,花牌
		if t == 4 { // 东南西北中发白（无顺）
			if v0 == v1 {
				g_NeedHunCount = Min(hNum+1, g_NeedHunCount)
				return
			}
		} else if (v1 - v0) < 3 {
			g_NeedHunCount = Min(hNum+1, g_NeedHunCount)
		}
		return
	} else if lArr >= 3 { //大于三张牌
		t := subArr[0] / 100
		v0 := subArr[0] % 10
		v2 := subArr[2] % 10

		//第一个和另外两个一铺
		arrLen := len(subArr)
		for i := 1; i < arrLen; i++ {
			if hNum+getModNeedNum(lArr-3, false) >= g_NeedHunCount {
				break
			}
			v1 := subArr[i] % 10
			//13444   134不可能连一起
			if v1-v0 > 1 {
				break
			}
			if (i + 2) < arrLen {
				if (subArr[i+2] % 10) == v1 {
					continue
				}
			}
			if i+1 < arrLen {
				tmp1, tmp2, tmp3 := subArr[0], subArr[i], subArr[i+1]
				if test3Combine(tmp1, tmp2, tmp3) {
					subArr = RemoveArray(subArr, 0)
					subArr = RemoveArray(subArr, i)
					subArr = RemoveArray(subArr, i+1)
					subLen := len(subArr)
					log.Debug("getNeedHunInSub_subLen:", subLen)
					getNeedHunInSub(subArr, hNum)
					subArr = append(subArr, tmp1)
					subArr = append(subArr, tmp2)
					subArr = append(subArr, tmp3)
					sort.Sort(sort.IntSlice(subArr))
				}

			}
			v1 = subArr[1] % 10
			if hNum+getModNeedNum(lArr-2, false)+1 < g_NeedHunCount {
				if t == 4 { //东南西北中发白（无顺）
					if v0 == v1 {
						tmp1 := subArr[0]
						tmp2 := subArr[1]
						subArr = RemoveArray(subArr, tmp1)
						subArr = RemoveArray(subArr, tmp2)
						getNeedHunInSub(subArr, hNum+1)
						subArr = append(subArr, tmp1)
						subArr = append(subArr, tmp2)
						sort.Sort(sort.IntSlice(subArr))
					}
				} else {
					arrLen = len(subArr)
					for j := 1; j < arrLen; j++ {
						if hNum+getModNeedNum(lArr-2, false)+1 >= g_NeedHunCount {
							break
						}
						v1 := subArr[j] % 10
						//#如果当前的value不等于下一个value则和下一个结合避免重复
						if (j + 1) != arrLen {
							v2 = subArr[j+1] % 10
							if v1 == v2 {
								continue
							}
						}
						mius := v1 - v0
						if mius < 3 {
							tmp1 := subArr[0]
							tmp2 := subArr[j]
							subArr = RemoveArray(subArr, tmp1)
							subArr = RemoveArray(subArr, tmp2)
							getNeedHunInSub(subArr, hNum+1)
							subArr = append(subArr, tmp1)
							subArr = append(subArr, tmp2)
							sort.Sort(sort.IntSlice(subArr))
							if mius >= 1 {
								break
							}
						} else {
							break
						}

					}
				}
			}
			//# 第一个自己一铺
			if hNum+getModNeedNum(lArr-1, false)+2 < g_NeedHunCount {
				tmp := subArr[0]
				subArr = RemoveArray(subArr, tmp)
				getNeedHunInSub(subArr, hNum+2)
				subArr = append(subArr, tmp)
				sort.Sort(sort.IntSlice(subArr))
			}

		}

	} else {
		return
	}
}

/**
 * 吴名
 *2017/5/23 下午6:00
 *判断两个是不是对子
 */
func test2Combine(mj1, mj2 int) bool {
	t1, t2 := mj1/100, mj2/100 //确定万,筒,条
	v1, v2 := mj1%10, mj2%10   //确定几万
	if t1 == t2 && v1 == v2 {
		return true
	}
	return false
}

/**
 * 吴名
 *2017/5/23 下午6:00
 *
 */
func canHu(hunNum int, arr []int) bool {
	tmpArr := []int{}
	tmpArr = append(tmpArr, arr...)
	arrLen := len(tmpArr)
	if arrLen <= 0 {
		if hunNum >= 2 {
			return true
		}
		return false
	}

	if hunNum < getModNeedNum(arrLen, true) {
		return false
	}

	for index, v := range tmpArr {
		if index == arrLen-1 { //如果是最后一张牌
			if hunNum > 0 {
				hunNum = hunNum - 1
				tmpArr = RemoveArray(tmpArr, index)
				g_NeedHunCount = 4
				getNeedHunInSub(tmpArr, 0)
				if g_NeedHunCount <= hunNum {
					// print 'type:',tmp/100, 'value', tmp%10, 1
					return true
				}
				hunNum = hunNum + 1
				tmpArr = append(tmpArr, v)
				//append之后重新排序
				//升序排列
				sort.Sort(sort.IntSlice(tmpArr))
			}
		} else {
			//如果是倒数第二个元素,或??????
			if index == arrLen-2 || (tmpArr[index]%10) != (tmpArr[index+2]%10) {
				if test2Combine(tmpArr[index], tmpArr[index+1]) {
					tmp1 := tmpArr[index]
					tmp2 := tmpArr[index+1]
					tmpArr = append(tmpArr[:index], tmpArr[index+2:]...)
					g_NeedHunCount = 4
					getNeedHunInSub(tmpArr, 0)
					if g_NeedHunCount <= hunNum {
						//# print 'type:',tmp1/100, 'value', tmp1%10, 2
						return true
					}
					tmpArr = append(tmpArr, tmp1, tmp2)
					//append之后重新排序
					//升序排列
					sort.Sort(sort.IntSlice(tmpArr))
				}
			}
			if hunNum > 0 && v%10 != tmpArr[index+1]%10 {
				hunNum = hunNum - 1
				//删除元素
				tmpArr = append(tmpArr[:index], tmpArr[index+1:]...)
				g_NeedHunCount = 4
				getNeedHunInSub(tmpArr, 0)
				if g_NeedHunCount <= hunNum {
					//# print 'type:',tmp/100, 'value', tmp%10, 3
					return true
				}
				hunNum = hunNum + 1
				tmpArr = append(tmpArr, v)
				//append之后重新排序
				//升序排列
				sort.Sort(sort.IntSlice(tmpArr))
			}
		}
	}
	return false
}

/**
 * 吴名
 *2017/5/23 下午7:28
 *判断胡牌
 */
func TestHu(mj int, mjArr []int, hunMj int) bool {
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...)
	if mj != 0 {
		tmpArr = append(tmpArr, mj) //插入一个麻将
	}
	sptArr := seperateArr(tmpArr, hunMj)
	curHunNum := len(sptArr[0])
	//有4个宝,直接胡牌
	if curHunNum > 3 {
		return true
	}

	ndHunArr := []int{} //每个分类需要混的数组
	for i := 1; i < 5; i++ {
		g_NeedHunCount = 4
		getNeedHunInSub(sptArr[i], 0)
		ndHunArr = append(ndHunArr, g_NeedHunCount)
	}
	isHu := false
	// 将在万中
	// 如果需要的混小于等于当前的则计算将在万中需要的混的个数
	ndHunAll := ndHunArr[1] + ndHunArr[2] + ndHunArr[3]
	if ndHunAll <= curHunNum {
		hasNum := curHunNum - ndHunAll
		isHu = canHu(hasNum, sptArr[1])
		if isHu {
			return true
		}
	}

	//将在饼中
	ndHunAll = ndHunArr[0] + ndHunArr[2] + ndHunArr[3]
	if ndHunAll <= curHunNum {
		hasNum := curHunNum - ndHunAll
		isHu = canHu(hasNum, sptArr[2])
		if isHu {
			return true
		}
	}
	//将在条中
	ndHunAll = ndHunArr[0] + ndHunArr[1] + ndHunArr[3]
	if ndHunAll <= curHunNum {
		hasNum := curHunNum - ndHunAll
		isHu = canHu(hasNum, sptArr[3])
		if isHu {
			return true
		}
	}

	//将在风中
	ndHunAll = ndHunArr[0] + ndHunArr[1] + ndHunArr[2]
	if ndHunAll <= curHunNum {
		hasNum := curHunNum - ndHunAll
		isHu = canHu(hasNum, sptArr[4])
		if isHu {
			return true
		}
	}
	return false
}

/**
 吴名
 2017/5/25 下午2:15
 mj:上家出牌
 */
func testGang(mj int, mjArr []int, hunMj int) bool {
	t := mj / 100
	v := mj % 10
	c := 0
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...)
	sptArr := seperateArr(tmpArr, hunMj)
	if len(sptArr[t]) < 2 {
		//该类型牌小于2张,加上mj也不能杠
		return false
	} else {
		//遍历这一类型麻将数组,计算mj这张牌的数量
		for _, tmj := range sptArr[t] {
			if tmj%10 == v {
				c = c + 1
			}
		}
		//明杠(别人出一张,自己三张) 暗杠(已有三张,自摸一张)
		if c == 3 {
			return true
		}
	}
	return false
}

/**
 吴名
 2017/5/25 下午2:38
 判断牌型是否能碰
 */
func testPeng(mj int, mjArr []int, hunMj int) bool {
	t := mj / 100
	v := mj % 10
	c := 0
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...)
	sptArr := seperateArr(tmpArr, hunMj)
	if len(sptArr[t]) < 2 {
		//该类型牌小于2张,加上mj也不能碰
		return false
	} else {
		//遍历这一类型麻将数组,计算mj这张牌的数量
		for _, tmj := range sptArr[t] {
			if tmj%10 == v {
				c = c + 1
			}
		}
		//别人出一张mj,自己手上有2张或3张,可以碰
		if c == 2 || c == 3 {
			return true
		}
	}
	return false
}

/**
 吴名
 2017/5/25 下午2:49
 */
func analyzeAnGang(mjArr []int, hunMj int) []int {
	result := []int{}
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...)
	sptArr := seperateArr(tmpArr, hunMj)
	for _, subArr := range sptArr {
		subLen := len(subArr)
		if subLen < 4 {
			continue
		} else {
			for j, subMj := range subArr {
				if subLen-1-j < 3 {
					break
				}
				if (subMj%10 == subArr[j+1]%10) && (subArr[j+1]%10 == subArr[j+2]%10) && (subArr[j+2]%10 == subArr[j+3]%10) {
					result = append(result, subMj)
				}
			}
		}
	}
	return result
}

/**
 吴名
 2017/5/25 下午5:25
 */
func getJiangNeedHum(arr []int) int {
	minNeedNum := 4
	tmpArr := []int{}
	tmpArr = append(tmpArr, arr...)
	arrLen := len(tmpArr)
	if arrLen <= 0 {
		return 2
	}
	for i, v := range tmpArr {
		if i == arrLen-1 { //如果是最后一张牌
			tmpArr = RemoveArray(tmpArr, i)
			g_NeedHunCount = 4
			getNeedHunInSub(tmpArr, 0)
			minNeedNum = Min(minNeedNum, g_NeedHunCount+1)

			tmpArr = append(tmpArr, v)
			//排序.恢复原状
			sort.Sort(sort.IntSlice(tmpArr))
		} else {
			if (i+2) == arrLen || (v%10) != (tmpArr[i+2]%10) {
				if test2Combine(v, tmpArr[i+1]) {
					tmp2 := tmpArr[i+1]
					tmpArr = append(tmpArr[:i], tmpArr[i+2:]...)
					g_NeedHunCount = 4
					getNeedHunInSub(tmpArr, 0)

					minNeedNum = Min(minNeedNum, g_NeedHunCount)
					tmpArr = append(tmpArr, v)
					tmpArr = append(tmpArr, tmp2)
					//排序.恢复原状
					sort.Sort(sort.IntSlice(tmpArr))
				}
			}
			if v%10 != (tmpArr[i+1] % 10) {
				tmpArr = append(tmpArr[:i], tmpArr[i+1:]...)
				g_NeedHunCount = 4
				getNeedHunInSub(tmpArr, 0)
				minNeedNum = Min(minNeedNum, g_NeedHunCount+1)

				tmpArr = append(tmpArr, v)
				//排序
				sort.Sort(sort.IntSlice(tmpArr))
			}
		}
	}
	return minNeedNum
}

func getTingArr(mjArr []int, hunMj int) []int {
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...) //创建一个麻将数组的copy
	sptArr := seperateArr(tmpArr, hunMj)

	ndHunArr := []int{} //每个分类需要混的数组
	for i := 1; i < 5; i++ {
		g_NeedHunCount = 4
		getNeedHunInSub(sptArr[i], 0)
		ndHunArr = append(ndHunArr, g_NeedHunCount)
	}

	jaNdHunArr := []int{} //每个将分类需要混的数组
	for i := 1; i < 5; i++ {
		jdNeedHunNum := getJiangNeedHum(sptArr[i])
		jaNdHunArr = append(jaNdHunArr, jdNeedHunNum)
	}

	curHunNum := len(sptArr[0])
	tingArr := []int{}
	paiArr := [][]int{{101, 110}, {201, 210}, {301, 310}, {401, 408}}
	//是否单调将
	isAllHu := false
	needNum := 0
	for _, v := range ndHunArr {
		needNum += v
	}
	if curHunNum-needNum == 1 {
		isAllHu = true
	}
	if isAllHu {
		for _, lis := range paiArr {
			for _, v := range lis {
				tingArr = append(tingArr, v)
			}
		}
		return tingArr
	}

	for i := 0; i < 4; i++ {
		// if len(sptArr[i+1]) == 0:
		//     continue;
		//听牌是将
		needNum = 0
		for j := 0; j < 4; j++ {
			if i != j {
				needNum = needNum + ndHunArr[j]
			}
		}

		if needNum <= curHunNum {
			for k := paiArr[i][0]; k < paiArr[i][1]; k++ {
				t := []int{k}
				t = append(t, sptArr[i+1]...)
				//排序
				sort.Sort(sort.IntSlice(t))
				if canHu(curHunNum-needNum, t) {
					tingArr = append(tingArr, k)
					// print callTime
				}
			}
		}

		//听牌是扑
		for j := 0; j < 4; j++ {
			if i != j {
				needNum = 0
				for k := 0; k < 4; k++ {
					if k != i {
						if k == j {
							needNum += jaNdHunArr[k]
						} else {
							needNum += ndHunArr[k]
						}
					}
				}
				if needNum <= curHunNum {
					for k := paiArr[i][0]; k < paiArr[i][1]; k++ {
						if !IntArrContain(tingArr, k) {
							t := []int{k}
							t = append(t, sptArr[i+1]...)
							g_NeedHunCount = 4
							//排序
							sort.Sort(sort.IntSlice(t))
							getNeedHunInSub(t, 0)
							if g_NeedHunCount <= curHunNum-needNum {
								tingArr = append(tingArr, k)
							}
						}
					}
				}
			}
		}
	}
	if len(tingArr) > 0 && !IntArrContain(tingArr, hunMj) {
		tingArr = append(tingArr, hunMj)
	}
	return tingArr
}

/**
 吴名
 2017/5/26 上午11:32
 */
func getTingNumArr(mjArr []int, hunMj int) []int {
	tmpArr := []int{}
	tmpArr = append(tmpArr, mjArr...) //创建一个麻将数组的copy
	sptArr := seperateArr(tmpArr, hunMj)

	ndHunArr := []int{} //每个分类需要混的数组
	for i := 1; i < 5; i++ {
		g_NeedHunCount = 4
		getNeedHunInSub(sptArr[i], 0)
		ndHunArr = append(ndHunArr, g_NeedHunCount)
	}

	jaNdHunArr := []int{} //每个将分类需要混的数组
	for i := 1; i < 5; i++ {
		jdNeedHunNum := getJiangNeedHum(sptArr[i])
		jaNdHunArr = append(jaNdHunArr, jdNeedHunNum)
	}

	//给一个混看能不能胡
	curHunNum := len(sptArr[0]) + 1
	tingArr := []int{}

	//是否单调将
	isAllHu := false
	needNum := 0
	for i := 0; i < 4; i++ {
		needNum += ndHunArr[i]
	}
	if curHunNum-needNum == 1 {
		isAllHu = true
	}
	if isAllHu {
		tingArr = append(tingArr, tmpArr...)
		return tingArr
	}

	for i := 0; i < 4; i++ {
		for index, x := range sptArr[i+1] {
			t := append(sptArr[i+1][:index], sptArr[i+1][index+1:]...)
			//将
			needNum = 0
			for j := 0; j < 4; j++ {
				if i != j {
					needNum = needNum + ndHunArr[j]
				}
			}
			if needNum <= curHunNum && !IntArrContain(tingArr, x) {
				if canHu(curHunNum-needNum, t) {
					tingArr = append(tingArr, x)
				}
			}

			//扑
			for j := 0; j < 4; j++ {
				if len(sptArr[j+1]) == 0 {
					continue
				}
				if i != j {
					needNum = 0
					for k := 0; k < 4; k++ {
						if k != i {
							if k == j {
								needNum += jaNdHunArr[k]
							} else {
								needNum += ndHunArr[k]
							}
						}
					}
					if needNum <= curHunNum && !IntArrContain(tingArr, x) {
						g_NeedHunCount = 4
						getNeedHunInSub(t, 0)
						if g_NeedHunCount <= curHunNum-needNum {
							tingArr = append(tingArr, x)
						}
					}
				}
			}
		}
	}
	return tingArr
}

