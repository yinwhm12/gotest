package card

import (
	"math/rand"
	"time"
	"github.com/name5566/leaf/log"
)

func init() {
	//根据时间设置随机数种子,这样确保随机的数据不会出现重复
	rand.Seed(time.Now().UnixNano())

}

/**
给到非自己的人发广播
*/
//func BroadcastMsg(msg interface{}, table map[int]*User, pos int) {
//	for k, user := range table {
//		if k != pos && user.state != UserLogout {
//			user.WriteMsg(msg)
//		}
//	}
//}

/**
获取一个数组中的最小数
*/
func Min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}

/**
 * Zhou Zihao
 *2017/5/24 14:41
 *移除数组元素
 */
func RemoveArray(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

/*
-- To shuffle an array a of n elements (indices 0..n-1):
for i from n−1 downto 1 do
j ← random integer such that 0 ≤ j ≤ i
exchange a[j] and a[i]
*/
func Shuffle(items []int) []int {
	i := len(items) - 1
	for i >= 1 {
		j := rand.Int63n(int64(i + 1))
		items[i], items[j] = items[j], items[i]
		i--
	}

	return items
}

/**
 吴名
 2017/5/26 上午10:57
 判断int切片中是否包含值为v的元素
 */
func IntArrContain(arr []int, v int) bool {
	for _, value := range arr {
		if value == v {
			return true
		}
	}
	return false
}

/**
 吴名
 2017/6/7 下午8:06
 获取出了牌之后剩下的牌
 */
func getRemain(arr []int, chu int) []int {
	remain := []int{}
	for index, v := range arr {
		if v == chu {
			remain = append(remain, arr[index+1:]...)
			log.Debug("未出的牌:%v,要出掉牌:%v,剩下的牌:%v", arr, chu, remain)
			break
		} else {
			remain = append(remain, v)
		}
	}
	return remain
}

/**
 吴名
 2017/6/23 下午5:36
 比较两个int切片内容是否相等(循环遍历方式,还有reflect.deepEqual方式,效率低)
 */
func intSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	//BCE特性:bounds check能够明确保证v != b[i]中的b[i]不会出现越界错误，从而避免了b[i]中的越界检查从而提高效率
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/**
 吴名
 2017/6/23 下午5:50
 判断一维数组是否存在于二维数组中
 */
func twoDSliContain(twoDArr [][]int, arr []int) bool {
	if arr == nil || twoDArr == nil {
		return false
	}

	for _, v := range twoDArr {
		if intSliceEqualBCE(v, arr) {
			return true
		}
	}
	return false
}

