package bitcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//BitCount0   普通法
func BitCount0(x uint64) int {
	c := 0
	for x != 0 {
		c += int(x & 1)
		x >>= 1
	}
	return c
}

//BitCount1 快速清零法，其运算次数与输入n的大小无关，只与n中1的个数有关。
//如果n的二进制表示中有k个1，那么这个方法只需要循环k次即可。
//其原理是不断清除n的二进制表示中最右边的1，同时累加计数器
func BitCount1(x uint64) int {
	c := 0
	for x != 0 {
		x &= (x - 1) //清除最低位的0
		c++
	}
	return c
}

// BitCount2 查表法
func BitCount2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//BitCount3   Hacker's Delight
func BitCount3(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

/*
//BitCount4   MIT HAKMEM   只适用于32bit 的数据
//
func BitCount4(n uint64) int {
	tmp := n - ((n >> 1) & 033333333333) - ((n >> 2) & 011111111111)
	tmp = (tmp + (tmp >> 3)) & 030707070707
	return int(tmp % 63)
}
*/
