package main

import "testing"

func TestPhilo_creation(t *testing.T) {
	cs := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		cs[i] = new(ChopS)
	}

	philos2 := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos2[i] = &Philo{
			cs[i],
			cs[(i+1)%5],
			rune(i),
			3,
			nil.shiftChannel,
		}
	}
	philos := philos2

	if len(philos) != 5 {
		t.Fatal("error creating philos ")
	}

	for _, philo := range philos {
		if philo.leftCS == nil || philo.rightCS == nil {
			t.Fatal("No chopsticks")
		}
	}
}
