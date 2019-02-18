package main

import "fmt"

const (
	// QUEEN 皇后的个数
	QUEEN = 17
)

var chess [QUEEN]int

// QueenPos 用来记录列里面哪些已经放了queen，不用再放
var QueenPos uint64 = 0x00000000

// QueenLeftPos 用来记录左斜里面放了，n行，m列，对应 n+m 位
var QueenLeftPos uint64 = 0x0000000000000000

// QueenRightPos 用来记录右斜里面放了，n行，m列，对应 Q - m + n - 1 位
var QueenRightPos uint64 = 0x0000000000000000

var oneBit uint64 = 0x0000000000000001

func initChess() {
	var i int
	for i = 0; i < QUEEN; i++ {
		chess[i] = -1
	}

	fmt.Println("calculating...")
}

func placeQueen(row int) bool {

	var origin = chess[row]
	var i = origin
	var pos uint64
	var leftPos, rightPos uint64

	var originLeft = row + origin
	var originRight = QUEEN - 1 + row - origin
	var ok bool

	for ok = true; ok; ok = (QueenPos&pos) > 0 || (QueenLeftPos&leftPos) > 0 || (QueenRightPos&rightPos) > 0 {

		i = i + 1

		if i >= QUEEN {
			// 已经是最后一个位置，超过了
			chess[row] = -1

			if origin != -1 { // 清除原来的位置
				QueenPos = QueenPos & ^(oneBit << uint(origin))
				QueenLeftPos = QueenLeftPos & ^(oneBit << uint(originLeft))
				QueenRightPos = QueenRightPos & ^(oneBit << uint(originRight))
			}

			return false
		}
		pos = oneBit << uint(i)
		leftPos = oneBit << uint(row+i)
		rightPos = oneBit << uint(QUEEN-1+row-i)
	}

	chess[row] = i

	QueenPos = QueenPos | (oneBit << uint(i))
	QueenLeftPos = QueenLeftPos | (oneBit << uint(i+row))
	QueenRightPos = QueenRightPos | (oneBit << uint(QUEEN-1+row-i))

	if origin != -1 { // 清除原来的位置
		QueenPos = QueenPos & ^(oneBit << uint(origin))
		QueenLeftPos = QueenLeftPos & ^(oneBit << uint(originLeft))
		QueenRightPos = QueenRightPos & ^(oneBit << uint(originRight))
	}
	return true
}

func main() {
	initChess()

	var i, s int
	for {
		if i == -1 {
			break
		}

		var result = placeQueen(i)

		if result == false {
			i--
			continue
		}

		if i == QUEEN-1 {
			s++
			var pos = chess[i]
			QueenPos = QueenPos & ^(oneBit << uint(pos))
			QueenLeftPos = QueenLeftPos & ^(oneBit << uint(QUEEN-1+pos))
			QueenRightPos = QueenRightPos & ^(oneBit << uint(QUEEN-1+QUEEN-1-pos))
			chess[i] = -1
			i--
		} else {
			i++
		}
	}

	fmt.Printf("There is total %d solutions\n", s)
}
