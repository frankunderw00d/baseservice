# baseservice


```
package main

import "log"

const (
	EncryptionKey = "frgok"

	// 默认公约数，在最大公因数为1和二者本身的时候为此值
	DefaultCommonDivisor = 5
)

var OriginValue = "This is 0123456789abcdABCD!@#$%^&*(){}[]你好吗？，。"

/*
f : 补齐位📔数
r : 数值加数
a : 左移位数
n : 是否与 255 非运算
k : 交换位数和间隔
*/

func main() {
	keys := []uint8(EncryptionKey)
	v := []uint8(OriginValue)

	log.Printf("EncryptionKey : %+v", keys)
	log.Println("============================")

	log.Println(v)
	log.Println(len(v))

	v = makeUp(keys[0], v)

	log.Println(v)
	log.Println(len(v))

	v = blurValue(keys[1], v)

	log.Println(v)
	log.Println(len(v))

	v = leftMove(keys[2], v)

	log.Println(v)
	log.Println(len(v))

	v = bitOperation(keys[3], v)

	log.Println(v)
	log.Println(len(v))

	v = spaceExchange(keys[4], v)

	log.Println(v)
	log.Println(len(v))

	//============================ reverse
	//log.Println("============================================= reverse")
	//
	//v = revBitOperation(keys[3], v)
	//
	//log.Println(v)
	//log.Println(len(v))
	//log.Println(string(v))
	//
	//v = revLeftMove(keys[2], v)
	//
	//log.Println(v)
	//log.Println(len(v))
	//log.Println(string(v))
	//
	//v = revBlurValue(keys[1], v)
	//
	//log.Println(v)
	//log.Println(len(v))
	//log.Println(string(v))
	//
	//v = revMakeUp(v)
	//
	//log.Println(v)
	//log.Println(len(v))
	//log.Println(string(v))

}

// 补齐位📔数，补 0
func makeUp(n uint8, v []uint8) []uint8 {
	remainder := int(n) - (len(v) % int(n))

	fix := make([]uint8, remainder)

	v = append(v, fix...)

	return v
}

// 去掉补齐位📔数，去掉所有的 0
func revMakeUp(v []uint8) []uint8 {
	a := v
	for i := len(v) - 1; i > 0; i-- {
		if a[i] == 0 {
			a = a[:i]
		}
	}

	return a
}

// 数值模糊
func blurValue(n uint8, v []uint8) []uint8 {
	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, uint8((int(v[i])+int(n))%256))
	}
	return a
}

// 清晰数值
func revBlurValue(n uint8, v []uint8) []uint8 {
	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, uint8((int(v[i])+256-int(n))%256))
	}
	return a
}

// 左移数值
func leftMove(n uint8, v []uint8) []uint8 {
	// 最多左移7位，8位就是原来的值
	n = n % 8

	f := func(i uint8, v uint8) uint8 {
		for i > 0 {
			a := v
			v = v << 1
			if a >= 128 {
				v += 1
			}
			i--
		}
		return v
	}

	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, f(n, v[i]))
	}

	return a
}

// 逆转左移数值
func revLeftMove(n uint8, v []uint8) []uint8 {
	// 最多左移7位，8位就是原来的值
	n = n % 8

	f := func(i uint8, v uint8) uint8 {
		for i > 0 {
			a := v
			v = v >> 1
			if a%2 != 0 {
				v += 128
			}
			i--
		}
		return v
	}

	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, f(n, v[i]))
	}

	return a
}

// 位运算 是否与 255 非运算
func bitOperation(n uint8, v []uint8) []uint8 {
	if n%2 == 0 {
		return v
	}

	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, v[i]^255)
	}

	return a
}

// 逆转位运算 是否与 255 非运算
func revBitOperation(n uint8, v []uint8) []uint8 {
	if n%2 == 0 {
		return v
	}

	a := make([]uint8, 0)
	for i := 0; i < len(v); i++ {
		a = append(a, v[i]^255)
	}

	return a
}

// 空间交换
func spaceExchange(n uint8, v []uint8) []uint8 {
	divisor := maxCommonDivisor(int(n), len(v))
	if divisor == 1 {
		divisor = DefaultCommonDivisor
	}

	gap := len(v) % divisor

	log.Printf("Divisor : %d", divisor)
	log.Printf("Gap : %d", gap)

	a := make([]uint8, 0)
	b := make([]uint8, 0)
	for i := 0; i < len(v)-gap; i++ {
		if i%(divisor*2) >= divisor {
			a = append(a, v[i])
		} else {
			b = append(b, v[i])
		}
	}

	return append(a, b...)
}

// 逆转空间交换
func revSpaceExchange(n uint8, v []uint8) []uint8 {
	return []uint8{}
}

// 最大公约数
func maxCommonDivisor(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	gap := a % b
	smaller := a
	if b < a {
		smaller = b
	}

	return maxCommonDivisor(smaller, gap)
}

```