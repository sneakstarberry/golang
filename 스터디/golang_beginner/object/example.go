package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
	val string
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(bread *Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	sandwitch.val += bread.val
	return sandwitch
}

func main() {
	// 1. 빵 한개를 꺼낸다.
	bread := &Bread{val: "bread"}

	jam := &StrawberryJam{}

	// 2. 딸기잼 뚜껑을 연다.
	OpenStrawberryJam(jam)

	// 3. 딸기잼을 한스푼 뜬다.
	spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(bread, spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(bread)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
