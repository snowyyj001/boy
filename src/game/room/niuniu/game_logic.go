package niuniu

import (
	"github.com/snowyyj001/loumiao/util"
)

var CardsResp = [POKER_NUMBER]int{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d,
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d,
	//0x41, 0x42
}

func l_getColor(card int) int {
	c := card & 0xf0
	return c >> 8
}

func l_getValue(card int) int {
	v := card & 0x0f
	return v
}

func l_getLogicValue(card int) int {
	v := card & 0x0f
	if v > 10 {
		return 10
	}
	return v
}

func l_shuffle(cards []int) {
	sz := len(cards)
	for i := 0; i < sz; i++ {
		ri := util.Random(sz)
		tmp := cards[i]
		cards[i] = cards[ri]
		cards[ri] = tmp
	}
}

func l_sort(cards []int, up bool) {
	sz := len(cards)
	for i := 0; i < sz; i++ {
		for j := i + 1; j < sz; j++ {
			if up {
				if l_getValue(cards[i]) > l_getValue(cards[j]) {
					tmp := cards[i]
					cards[i] = cards[j]
					cards[j] = tmp
				}
			} else {
				if l_getValue(cards[i]) < l_getValue(cards[j]) {
					tmp := cards[i]
					cards[i] = cards[j]
					cards[j] = tmp
				}
			}
		}
	}
}
