package practice_21

import "fmt"

type client struct{}

func (c *client) insertLightningConnectorIntoComputer(comp computer) {
	fmt.Println("Client inserts lightning connector into computer...")
	comp.insertIntoLightningPort()
}
