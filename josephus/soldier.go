package josephus

type Soldier struct {
	position int
}

func (soldier *Soldier) Init(position int) {
	soldier.position = position
}
