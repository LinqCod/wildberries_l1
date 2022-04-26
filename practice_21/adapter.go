package practice_21

import "fmt"

type windowsAdapter struct {
	windowsPC *windowsComputer
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal into USB")
	w.windowsPC.insertIntoUSBPort()
}
