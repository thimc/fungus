package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const randomDirection = "><^V"

type Fungus struct {
	width, height int
	Matrix        [][]byte
	x, y          int
	dx, dy        int
	stack         []int
	str           bool
}

func NewFungus(w, h, sz int) *Fungus {
	f := &Fungus{
		width:  w,
		height: h,
		Matrix: make([][]byte, h),
		stack:  make([]int, sz),
	}
	for i := range f.Matrix {
		f.Matrix[i] = make([]byte, w)
		for j := range f.Matrix[i] {
			f.Matrix[i][j] = ' '
		}
	}
	f.dx = 1
	return f
}

func (f *Fungus) pop() int {
	if len(f.stack) <= 0 {
		panic("stack underflow")
	}
	ln := len(f.stack) - 1
	v := f.stack[ln]
	f.stack = f.stack[:ln]
	return v
}

func (f *Fungus) push(v int) {
	f.stack = append(f.stack, v)
}

func (f *Fungus) move() {
	f.x = (f.x + f.dx + f.width) % f.width
	f.y = (f.y + f.dy + f.height) % f.height
}

func (f *Fungus) Run() {
	ch := f.Matrix[f.y][f.x]

	// String mode
	if ch == '"' {
		f.str = !f.str
	} else if f.str {
		f.push(int(ch))
		f.move()
		f.Run()
		return
	}

	switch ch {
	case '+': // Addition
		x := f.pop()
		y := f.pop()
		f.push(x + y)
	case '-': // Subtraction
		x := f.pop()
		y := f.pop()
		f.push(y - x)
	case '*': // Multiplication
		x := f.pop()
		y := f.pop()
		f.push(x * y)
	case '/': // Division
		x := f.pop()
		y := f.pop()
		f.push(y / x)
	case '%': // Modulo
		x := f.pop()
		y := f.pop()
		f.push(y % x)
	case '!': // Logical not
		x := f.pop()
		if x == 0 {
			f.push(1)
		} else {
			f.push(0)
		}
	case '`': // Greater than
		x := f.pop()
		y := f.pop()
		if x > y {
			f.push(1)
		} else {
			f.push(0)
		}
	case '>': // Right
		f.dy = 0
		f.dx = 1
	case '<': // Left
		f.dy = 0
		f.dx = -1
	case '^': // Up
		f.dy = -1
		f.dx = 0
	case 'v': // Down
		f.dy = 1
		f.dx = 0
	case '_': // Horizontal if
		x := f.pop()
		f.dy = 0
		if x == 0 {
			f.dx = 1
		} else {
			f.dx = -1
		}
	case '|': // Vertical If
		x := f.pop()
		f.dx = 0
		if x == 0 {
			f.dy = 1
		} else {
			f.dy = -1
		}
	case ':': // Duplicate push
		x := f.pop()
		f.push(x)
		f.push(x)
	case '\\': // Swap stack
		x := f.pop()
		y := f.pop()
		f.push(x)
		f.push(y)
	case '$': // Discard stack top
		f.pop()
	case '.': // Output as integer
		fmt.Printf("%d", f.pop())
	case ',': // Output as ASCII char
		fmt.Printf("%c", f.pop())
	case '#': // Bridge jump command
		f.move()
	case 'g': // Get call
		y := f.pop()
		x := f.pop()
		f.push(int(f.Matrix[y][x]))
	case 'p': // Put call
		y := f.pop()
		x := f.pop()
		z := f.pop()
		f.Matrix[y][x] = byte(int(z))
	case '&': // Get integer input
		x := 0
		fmt.Scanf("%d", &x)
		f.push(int(x))
	case '~': // Get character input
		r := bufio.NewReader(os.Stdin)
		ch, _ := r.ReadByte()
		f.push(int(ch))
	case '?': // Random direction
		ch = randomDirection[rand.Intn(len(randomDirection))]
		f.move()
		f.Run()
		return
	case '@': // Exit
		return
	default: // Push literal to stack
		if ch >= '0' && ch < '9' {
			f.push(int(ch - '0'))
		}
	}

	f.move()
	f.Run()
}

func main() {
	fungus := NewFungus(80, 24, 1024)

	file := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		file = f
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	y := 0
	for s.Scan() {
		line := s.Text()
		for i := 0; i < len(line); i++ {
			ch := byte(line[i])
			if ch != '\n' {
				fungus.Matrix[y][i] = ch
			}
		}
		y++
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fungus.Run()
}
