package main

import "fmt"

type USWallSocket struct{}

func (s *USWallSocket) ProvideUSPower() {
	fmt.Println("Питание через американскую розетку 110V")
}
type EuropeanPlug interface {
	ChargeWithEuropeanPlug()
}
type Traveler struct{}

func (t *Traveler) ChargeLaptop(plug EuropeanPlug) {
	fmt.Println("турист пытается зарядить ноутбук")
	plug.ChargeWithEuropeanPlug()
	fmt.Println("Ноутбук заряжается! ")
}

type USToEuropeanAdapter struct {

	usSocket *USWallSocket
}


func (a *USToEuropeanAdapter) ChargeWithEuropeanPlug() {
	fmt.Println("Адаптер преобразует вызов")
	a.usSocket.ProvideUSPower()
}

func main() {
	traveler := &Traveler{}
	usSocket := &USWallSocket{}

	adapter := &USToEuropeanAdapter{
		usSocket: usSocket,
	}
	traveler.ChargeLaptop(adapter)
}