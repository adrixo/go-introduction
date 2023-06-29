package main

import (
	"fmt"
	"sync"
	"time"
)

const DEBUG = false

type ChopS struct {
	sync.Mutex
	id int
}

type Philo struct {
	leftCS, rightCS   ChopS
	id                int
	numberOfMeals     int
	shiftChannel      chan int
	permissionChannel chan int
	seats             []seat
}

type Host struct {
	shiftChannel       chan int
	permissionChannels []chan int
	seats              []seat
}

type seat struct {
	sync.Mutex
	id int
}

func (h *Host) routine() {
	go h.dispatch(h.seats[0])
	go h.dispatch(h.seats[1])
}

func (h *Host) dispatch(seat seat) {
	for {
		if DEBUG {
			fmt.Printf("[host] starting to wait for seat %d\n", seat.id)
		}
		philoID := <-h.shiftChannel
		if DEBUG {
			fmt.Printf("[host] s-%d allow to eat to philo %d\n", seat.id, philoID)
		}
		h.permissionChannels[philoID] <- seat.id
		<-h.permissionChannels[philoID]

	}
}

func (p *Philo) routine() {
	for i := 0; i < p.numberOfMeals; i++ {
		if DEBUG {
			fmt.Printf("[p-%d] asking for a seat\n", p.id)
		}
		p.shiftChannel <- p.id
		assignedSeatID := <-p.permissionChannel
		if DEBUG {
			fmt.Printf("[p-%d] assigned to seat %d\n", p.id, assignedSeatID)
		}
		seat := p.seats[assignedSeatID]
		seat.Lock()

		if DEBUG {
			fmt.Printf("[p-%d] seated in %d\n", p.id, seat.id)
		}
		p.eat(seat)
		seat.Unlock()
		if DEBUG {
			fmt.Printf("[p-%d] away from seat %d\n", p.id, seat.id)
		}
		p.permissionChannel <- assignedSeatID

	}
}

func (p *Philo) eat(seat seat) {
	if DEBUG {
		fmt.Printf("[p-%d] taking chop l-%d\n", p.id, p.leftCS.id)
	}
	p.leftCS.Lock()
	if DEBUG {
		fmt.Printf("[p-%d] taking chop r-%d\n", p.id, p.rightCS.id)
	}
	p.rightCS.Lock()
	fmt.Printf("[p-%d] starting to eat <%d> seat: %d\n", p.id, p.id, seat.id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("[p-%d] finish eating <%d>\n", p.id, p.id)
	p.leftCS.Unlock()
	if DEBUG {
		fmt.Printf("[p-%d] left chop l-%d\n", p.id, p.leftCS.id)
	}
	p.rightCS.Unlock()
	if DEBUG {
		fmt.Printf("[p-%d] left chop r-%d\n", p.id, p.rightCS.id)
	}
}

func main() {
	numberOfPhilos := 5
	numberOfSeats := 2
	seats := make([]seat, numberOfSeats)
	for i, _ := range seats {
		seats[i].id = i
	}

	permissionChannels := make([]chan int, numberOfPhilos)
	for i := 0; i < numberOfPhilos; i++ {
		permissionChannels[i] = make(chan int)
	}

	host := &Host{
		make(chan int),
		permissionChannels,
		seats,
	}

	cs := make([]ChopS, numberOfPhilos)
	for i := 0; i < numberOfPhilos; i++ {
		cs[i].id = i
	}

	philos := make([]*Philo, numberOfPhilos)
	for i := 0; i < numberOfPhilos; i++ {
		philos[i] = &Philo{
			cs[i],
			cs[(i+1)%numberOfPhilos],
			i,
			3,
			host.shiftChannel,
			permissionChannels[i],
			seats,
		}
	}
	go host.routine()

	for i := 0; i < numberOfPhilos; i++ {
		go philos[i].routine()
	}
	time.Sleep(3 * time.Second)
}
