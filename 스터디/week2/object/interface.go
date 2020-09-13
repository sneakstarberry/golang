package main

import "fmt"

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct{}
type OrangeJam struct{}

type SpoonOfStrawberryJam struct{}
type SpoonOfOrangeJam struct{}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func (o *SpoonOfOrangeJam) String() string {
	return "+ orange"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (o *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	//딸기잼 빵에 발라먹기
	bread := &Bread{}
	// jam := &StrawberryJam{}
	jam := &OrangeJam{}

	bread.PutJam(jam)

	fmt.Println(bread)
}
