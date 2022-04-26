package practice_21

func Run() {
	client := &client{}
	mac := &macComputer{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsPC := &windowsComputer{}
	windowsPCAdapter := &windowsAdapter{
		windowsPC: windowsPC,
	}

	client.insertLightningConnectorIntoComputer(windowsPCAdapter)
}
