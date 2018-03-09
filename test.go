package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box

// 计算 box 容积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 设置 box 的颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}

//  找出所有 box 中容积最大的 box 的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// 把所以 box 的颜色涂为黑色
func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

// 根据 box 中的 color 值获取对应的颜色
func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 10, BLUE},
		Box{20, 20, 20, WHITE},
		Box{30, 30, 30, YELLOW},
	}

	fmt.Printf("there are %d boxes.\n", len(boxes))
	fmt.Println("the volume of the first one:", boxes[0].Volume(), "cm^3")
	fmt.Println("the color of the last one is:", boxes[len(boxes)-1].color.String())
	fmt.Println("the biggest one is", boxes.BiggestColor().String())

	fmt.Println("paint them all black:")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
}
