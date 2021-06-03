package josephus

import (
	"testing"
)

func SetupTest(noOfSoldiers int, hops int) SuicideCircle {
	soldiers := make([]Soldier, noOfSoldiers, noOfSoldiers)
	for i := 1; i <= noOfSoldiers; i++ {
		s := Soldier{}
		s.Init(i)
		soldiers[i-1] = s
	}

	suicideCircle := SuicideCircle{}
	suicideCircle.Init(soldiers, hops)
	suicideCircle.Kill(hops)
	return suicideCircle
}
func TestHowManyKilled(t *testing.T) {
	noOfSoldiers := 41
	hops := 6
	suicideCircle := SetupTest(noOfSoldiers, hops)
	killed := len(suicideCircle.killed)
	want := noOfSoldiers - 1
	if killed != want {
		t.Fatalf(`Number of people killed doesn't match expectation. Soldiers - %v, Killed - %v, Want - %v`, noOfSoldiers, killed, want)
	}
}

func TestHowOnlyOneSurvivor(t *testing.T) {
	noOfSoldiers := 32
	hops := 2
	suicideCircle := SetupTest(noOfSoldiers, hops)
	survivors := len(suicideCircle.survivors)
	want := 1
	if survivors != want {
		t.Fatalf(`Number of survivors doesn't match expectation. Soldiers - %v, Survivors - %v, Want - %v`, noOfSoldiers, survivors, want)
	}
}

type testCases struct {
	noOfSoldiers         int
	hops                 int
	josephusPositionWant int
}

func TestJosephusPosition(t *testing.T) {
	testCases := []testCases{
		{47, 3, 2},
		{10, 1, 10},
		{15, 3, 5},
		{40, 2, 17},
		{200, 10, 163},
	}
	for _, testCase := range testCases {
		suicideCircle := SetupTest(testCase.noOfSoldiers, testCase.hops)
		josephusPosition := suicideCircle.josephus.position
		want := testCase.josephusPositionWant
		if josephusPosition != want {
			t.Fatalf(`Josephus position doesn't match expectation. Soldiers - %v, Hops - %v, Position - %v, Want - %v`,
				testCase.noOfSoldiers, testCase.hops, josephusPosition, want)
		}
	}
}
