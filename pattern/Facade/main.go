package main

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light: turned on")
}

func (l *Light) Off() {
	fmt.Println("Light: turned off")
}

type Thermostat struct{}

func (t *Thermostat) SetTemperature(temp int) {
	fmt.Printf("Thermostat: Установить температуру на %d°C\n", temp)
}

type SecuritySystem struct{}

func (s *SecuritySystem) Activate() {
	fmt.Println("SecuritySystem: active")
}

func (s *SecuritySystem) Deactivate() {
	fmt.Println("SecuritySystem: not active")
}

type SmartHome struct {
	light          *Light
	thermostat     *Thermostat
	securitySystem *SecuritySystem
}

func NewSmartHome() *SmartHome {
	return &SmartHome{
		light:          &Light{},
		thermostat:     &Thermostat{},
		securitySystem: &SecuritySystem{},
	}
}

func (sh *SmartHome) LeaveHome() {
	fmt.Println("SmartHome: preparing for exit")
	sh.light.Off()
	sh.thermostat.SetTemperature(18)
	sh.securitySystem.Activate()
}

func (sh *SmartHome) ReturnHome() {
	fmt.Println("SmartHome: preparing for return")
	sh.light.On()
	sh.thermostat.SetTemperature(22)
	sh.securitySystem.Deactivate()
}

func main() {
	home := NewSmartHome()

	home.LeaveHome()

	home.ReturnHome()
}
