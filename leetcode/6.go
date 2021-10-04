package leetcode

/*
 * 6. Z 字形变换
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 * 		P A H N
 * 		APLSIIG
 * 		Y I R
 *
 *	之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 * 	请你实现这个将字符串进行指定行数变换的函数：
 *
 *		string convert(string s, int numRows);
 *
 *
 * 	示例 1：
 * 		输入：s = "PAYPALISHIRING", numRows = 3
 * 		输出："PAHNAPLSIIGYIR"
 *
 * 	示例 2：
 * 		输入：s = "PAYPALISHIRING", numRows = 4
 * 		输出："PINALSIGYAHRPI"
 *
 *
 *	解题思路:
 *		找规律，通过一次遍历字符串 确认每个字符的位置 直接放到对应结果位置
 *		结果放入一个切片，将一个二维切片结构的底层指向该切片, 此时需要计算出每行的实际字符处理，通过实际字符数量确定切片大小
 *		对字符串进行逻辑处理，处理后底层切片的结果即为需要返回的数据
 *		时间复杂度O(n), 空间复杂度O(n)
 *
 *		!!! 上面方法中事先算出来每行实际字符的数量，这个属于进一步优化
 *			也可以不事先出每行实际长度，但是后续需要进行字符串拼接，复杂度比上面要高
 *
 *
 */

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	//结果存储空间 —— 真实存储的地方
	result := make([]byte, len(s))
	//基于result的逻辑存储
	resultArr := make([][]byte, numRows)

	//组内元素数量
	groupSize := numRows + numRows - 2

	groupNum := len(s) / groupSize
	remains := len(s) % groupSize

	//计算每行实际字符数量 并给起分配底层空间
	pos := 0
	for i := 0; i < numRows; i++ {
		oldPos := pos
		if i == 0 || i == numRows-1 {
			if remains >= i+1 {
				pos += groupNum + 1
			} else {
				pos += groupNum
			}
		} else {
			pos += groupNum * 2
			if remains >= i+1 {
				pos += 1
			}

			if remains >= numRows+(numRows-i-1) {
				pos += 1
			}
		}

		resultArr[i] = result[oldPos:oldPos]
	}

	//数据写入
	idx, step := 0, -1
	for _, v := range s {
		if idx == 0 || idx == numRows-1 {
			step *= -1
		}

		resultArr[idx] = append(resultArr[idx], byte(v))
		idx += step
	}
	return string(result)
}
