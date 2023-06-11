package Solution_1

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  nil,
		out: nil,
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.move()
	}

	if len(this.out) > 0 {
		value := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return value
	}

	panic("empty queue")
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.move()
	}

	if len(this.out) > 0 {
		return this.out[len(this.out)-1]
	}

	panic("empty queue")
}

func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

func (this *MyQueue) move() {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}

	this.in = this.in[:0]
}
