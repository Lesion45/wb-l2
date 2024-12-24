package main

import "fmt"

type State interface {
	Start()
	Stop()
}

type Machine struct {
	state State
}

func (m *Machine) SetState(state State) {
	m.state = state
}

func (m *Machine) Start() {
	m.state.Start()
	if _, isOffState := m.state.(*OffState); isOffState {
		m.SetState(&OnState{})
	}
}

func (m *Machine) Stop() {
	m.state.Stop()
	if _, isOnState := m.state.(*OnState); isOnState {
		m.SetState(&OffState{})
	}
}

type OffState struct{}

func (o *OffState) Start() {
	fmt.Println("Machine is turned on.")
}

func (o *OffState) Stop() {
	fmt.Println("Machine is already off.")
}

type OnState struct{}

func (o *OnState) Start() {
	fmt.Println("Machine is already on.")
}

func (o *OnState) Stop() {
	fmt.Println("Machine is turned off.")
}

type WorkingState struct{}

func (w *WorkingState) Start() {
	fmt.Println("Machine is already working.")
}

func (w *WorkingState) Stop() {
	fmt.Println("Machine is stopped.")
}

func main() {
	machine := &Machine{}
	machine.SetState(&OnState{})

	machine.Start()
	machine.Stop()
	machine.Start()
	machine.Stop()
}
