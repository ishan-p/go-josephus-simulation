package josephus

type SuicideCircle struct {
	survivors []Soldier
	killed    []Soldier
	hops      int
	josephus  Soldier
}

func (circle *SuicideCircle) Init(soldiers []Soldier, hops int) {
	circle.survivors = soldiers
	circle.killed = make([]Soldier, 0, len(soldiers))
	circle.hops = hops
}
