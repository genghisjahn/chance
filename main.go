package main

import (
	"fmt"
	"math/rand"
	"time"
)

const winchance = .6

type stake struct {
	Funds      float32
	MaxFunds   float32
	Wins       int
	Losses     int
	WinStreak  int
	LossStreak int
	lastbet    bool
}

func (s stake) CalcBet() float32 {

	//return s.Funds * .6
	return 100.0
}

func (s *stake) Bet(amount float32) error {
	if amount > s.Funds {
		return fmt.Errorf("%v is more than you have %v\n", amount, s.Funds)
	}
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	if r.Float32() <= winchance {
		s.Funds += amount
		s.Wins++

	} else {
		s.Funds -= amount
		s.Losses++
		s.lastbet = false
	}
	if s.Funds > s.MaxFunds {
		s.MaxFunds = s.Funds
	}
	return nil
}

func main() {
	var f float32 = 1000.0
	s := stake{Funds: f, MaxFunds: f}
	for i := 0; i < 100; i++ {
		err := s.Bet(s.CalcBet())
		if err != nil || s.Funds < 1.0 {
			break
		}
		if s.Funds == 0 {
			break
		}
	}
	fmt.Println("MaxFunds:", s.MaxFunds)
	fmt.Println("Funds:", s.Funds)
	fmt.Println("Wins/Losses: ", s.Wins, s.Losses)

}
