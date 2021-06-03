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

func (circle *SuicideCircle) Kill(hops int) {
	var toBeKilledSoldier Soldier
	var survived []Soldier
	survivors := len(circle.survivors)
	if survivors == 1 {
		circle.josephus = circle.survivors[0]
		return
	}
	if hops > survivors {
		survived = append(circle.survivors[survivors-1:], circle.survivors[:survivors-1]...)
		hops = hops - (survivors - 1)
		circle.survivors = survived
	} else {
		killIndex := hops - 1
		toBeKilledSoldier = circle.survivors[killIndex]
		circle.killed = append(circle.killed, toBeKilledSoldier)
		survived = append(circle.survivors[killIndex+1:], circle.survivors[:killIndex]...)
		circle.survivors = survived
		hops = circle.hops
	}
	circle.Kill(hops)
}
