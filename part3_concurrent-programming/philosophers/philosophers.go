package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NumMeals        = 3
	NumPhilosophers = 5
	maxEating       = 2
)

type Philo struct {
	idx     int
	meals   int
	leftCS  *ChopS
	rightCS *ChopS
}

type ChopS struct {
	sync.Mutex
}

func (p Philo) eat(c chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < NumMeals; i++ {
		c <- &p
		if p.meals < NumMeals {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Println("starting to eat ", p.idx)
			p.meals++ // eating
			fmt.Println("finished eating", p.idx)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			wg.Done()
		}
	}
}

func placeChopsticks() []*ChopS {
	cs := make([]*ChopS, NumPhilosophers)
	for i := 0; i < NumPhilosophers; i++ {
		cs[i] = new(ChopS)
	}
	return cs
}

func placePhilosophers(cs []*ChopS) []*Philo {
	p := make([]*Philo, NumPhilosophers)
	for i := 0; i < NumPhilosophers; i++ {
		p[i] = &Philo{i + 1, 0, cs[i], cs[(i+1)%NumPhilosophers]}
	}
	return p
}

func host(c chan *Philo) {
	for {
		if len(c) == maxEating {
			<-c
			<-c
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	CSticks := placeChopsticks()
	philos := placePhilosophers(CSticks)

	var wg sync.WaitGroup
	wg.Add(NumPhilosophers * NumMeals)

	c := make(chan *Philo, maxEating)
	go host(c)

	for i := 0; i < NumPhilosophers; i++ {
		go philos[i].eat(c, &wg)
	}

	wg.Wait()
}
