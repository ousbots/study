package implement_queue_using_stacks

type MyQueue struct {
	forward []int
	reverse []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	for len(this.forward) > 0 {
		this.reverse = append(this.reverse, this.forward[len(this.forward)-1])
		this.forward = this.forward[:len(this.forward)-1]
	}

	this.reverse = append(this.reverse, x)
}

func (this *MyQueue) Pop() int {
	for len(this.reverse) > 0 {
		this.forward = append(this.forward, this.reverse[len(this.reverse)-1])
		this.reverse = this.reverse[:len(this.reverse)-1]
	}

	val := this.forward[len(this.forward)-1]
	this.forward = this.forward[:len(this.forward)-1]

	return val
}

func (this *MyQueue) Peek() int {
	for len(this.reverse) > 0 {
		this.forward = append(this.forward, this.reverse[len(this.reverse)-1])
		this.reverse = this.reverse[:len(this.reverse)-1]
	}

	return this.forward[len(this.forward)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.forward) == 0 && len(this.reverse) == 0
}
