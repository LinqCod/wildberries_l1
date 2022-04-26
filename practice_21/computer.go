package practice_21

import "fmt"

type computer interface {
	insertIntoLightningPort()
}

type macComputer struct{}

func (m *macComputer) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac successfully!")
}

type windowsComputer struct{}

func (w *windowsComputer) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows computer successfully!")
}
