package main

import "fmt"

type ShiftRegister struct {
	hi, lo                                          uint64
	length, feedforward, feedbackward, latestOutput int
	and                                             [2]int
}

func (r *ShiftRegister) ShfitRightByOne() {
	r.lo >>= 1
	bit := r.hi & 1
	r.lo |= bit << 63
	r.hi >>= 1
}

func (r *ShiftRegister) SetFeedBackBit(val int) {
	r.hi |= uint64(val) << (r.length - 64 - 1)
}

func (r *ShiftRegister) BitAt(pos int) int {
	if pos < 64 {
		if (1<<pos)&r.lo > 0 {
			return 1
		} else {
			return 0
		}
	}
	if (1<<(pos-64))&r.hi > 0 {
		return 1
	} else {
		return 0
	}
}

func (r *ShiftRegister) NextOutput() int {
	output := r.BitAt(0) ^ r.BitAt(r.length-r.feedforward) ^ (r.BitAt(r.length-r.and[0]) & r.BitAt(r.length-r.and[1]))
	r.latestOutput = output
	return output
}

type Trivium struct {
	A, B, C *ShiftRegister
}

func NewTrivium() *Trivium {
	return &Trivium{
		&ShiftRegister{
			// leftmost 80 bits are set by IV
			hi:           0,
			lo:           0,
			length:       93,
			feedbackward: 69,
			feedforward:  66,
			and:          [2]int{91, 92},
			latestOutput: 0,
		},
		&ShiftRegister{
			// leftmost 80 bits are set by key
			hi:           0,
			lo:           0,
			length:       84,
			feedbackward: 78,
			feedforward:  69,
			and:          [2]int{82, 83},
			latestOutput: 0,
		},
		&ShiftRegister{
			hi:           0,
			lo:           7, // three rightmost bits are set to 1
			length:       111,
			feedbackward: 87,
			feedforward:  66,
			and:          [2]int{109, 110},
			latestOutput: 0,
		},
	}
}

func (t *Trivium) NextOutput() int {
	a := t.A.NextOutput()
	t.A.ShfitRightByOne()

	b := t.B.NextOutput()
	t.B.ShfitRightByOne()

	c := t.C.NextOutput()
	t.C.ShfitRightByOne()

	t.A.SetFeedBackBit(c ^ t.A.BitAt(t.A.length-t.A.feedbackward))
	t.B.SetFeedBackBit(a ^ t.B.BitAt(t.B.length-t.B.feedbackward))
	t.C.SetFeedBackBit(b ^ t.C.BitAt(t.C.length-t.C.feedbackward))

	return a ^ b ^ c
}

func (t *Trivium) WarmUp(limit int) []int {
	if limit <= 0 {
		return make([]int, 0)
	}

	result := make([]int, limit)

	for i := 0; i < limit; i++ {
		result[i] = t.NextOutput()
	}

	return result
}

func main() {
	trivium := NewTrivium()
	fmt.Println(trivium.WarmUp(70))
}

