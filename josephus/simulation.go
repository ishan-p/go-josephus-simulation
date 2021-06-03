package josephus

import (
	"fmt"
	"time"
)

const simulationSleeptime = 250 * time.Millisecond

func PlayJosephusSimulation(noOfSoldiers int, hops int) {
	soldiers := make([]Soldier, noOfSoldiers, noOfSoldiers)
	for i := 1; i <= noOfSoldiers; i++ {
		s := Soldier{}
		s.Init(i)
		soldiers[i-1] = s
	}

	suicideCircle := SuicideCircle{}
	suicideCircle.Init(soldiers, hops)
	fmt.Println("Let the killing begin...")
	suicideCircle.Kill(hops)
	for _, val := range suicideCircle.killed {
		fmt.Print(val.position)
		fmt.Print(" -> ")
		time.Sleep(simulationSleeptime)
	}
	fmt.Printf("Josephus stood at position: %v\n", suicideCircle.josephus.position)
}
