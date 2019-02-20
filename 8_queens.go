package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	// QUEEN 皇后的个数
	QUEEN = 16
)

// wg 计数器
var wg sync.WaitGroup

func calcChessWhenFirstRowIs(firstRow int, result chan int64) {

	// QueenPos 用来记录列里面哪些已经放了queen，不用再放
	var QueenPos uint64 = 0x00000000

	// QueenLeftPos 用来记录左斜里面放了，n行，m列，对应 n+m 位
	var QueenLeftPos uint64 = 0x0000000000000000

	// QueenRightPos 用来记录右斜里面放了，n行，m列，对应 Q - m + n - 1 位
	var QueenRightPos uint64 = 0x0000000000000000

	var oneBit uint64 = 0x0000000000000001

	defer wg.Done()

	// 整个棋盘
	var chess [QUEEN]int

	// 初始化
	chess[0] = firstRow
	QueenPos = oneBit << uint(firstRow)
	QueenLeftPos = oneBit << uint(firstRow)
	QueenRightPos = oneBit << uint(QUEEN-1-firstRow)

	var ii int
	for ii = 1; ii < QUEEN; ii++ {
		chess[ii] = -1
	}

	var row int = 1
	var s int64

	for {

		if row == 0 {
			break
		}

		var origin = chess[row]
		var column = origin
		var pos, leftPos, rightPos uint64

		var originLeft = row + origin
		var originRight = QUEEN - 1 + row - origin
		var ok bool
		var result bool = true

		for ok = true; ok; ok = (QueenPos&pos) > 0 || (QueenLeftPos&leftPos) > 0 || (QueenRightPos&rightPos) > 0 {

			column = column + 1

			if column >= QUEEN {
				// 已经是最后一个位置，超过了
				chess[row] = -1

				if origin != -1 { // 清除原来的位置
					QueenPos = QueenPos & ^(oneBit << uint(origin))
					QueenLeftPos = QueenLeftPos & ^(oneBit << uint(originLeft))
					QueenRightPos = QueenRightPos & ^(oneBit << uint(originRight))
				}

				result = false
				break
			}
			pos = oneBit << uint(column)
			leftPos = oneBit << uint(row+column)
			rightPos = oneBit << uint(QUEEN-1+row-column)
		}

		if result == true {
			chess[row] = column

			QueenPos = QueenPos | (oneBit << uint(column))
			QueenLeftPos = QueenLeftPos | (oneBit << uint(column+row))
			QueenRightPos = QueenRightPos | (oneBit << uint(QUEEN-1+row-column))

			if origin != -1 { // 清除原来的位置
				QueenPos = QueenPos & ^(oneBit << uint(origin))
				QueenLeftPos = QueenLeftPos & ^(oneBit << uint(originLeft))
				QueenRightPos = QueenRightPos & ^(oneBit << uint(originRight))
			}

		} else {
			row--
			continue
		}

		if row == QUEEN-1 {
			s++
			var pos = chess[row]
			QueenPos = QueenPos & ^(oneBit << uint(pos))
			QueenLeftPos = QueenLeftPos & ^(oneBit << uint(QUEEN-1+pos))
			QueenRightPos = QueenRightPos & ^(oneBit << uint(QUEEN-1+QUEEN-1-pos))
			chess[row] = -1
			row--

		} else {
			row++
		}
	}

	result <- s
}

func main() {

	fmt.Printf("Calculating %d queens...\n", QUEEN)

	var now = time.Now().UnixNano()

	var sizes chan int64 = make(chan int64)
	var i int

	for i = 0; i < QUEEN; i++ {
		wg.Add(1)
		go calcChessWhenFirstRowIs(i, sizes)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var size, total int64
	for size = range sizes {
		total += size
	}

	var afterCaculation = time.Now().UnixNano()

	fmt.Printf("There is total %d solutions \n", total)
	fmt.Println((afterCaculation-now)/1000, "us")
}
