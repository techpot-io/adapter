package main

import (
	"github.com/praiakov/adapter/adaptee"
	"github.com/praiakov/adapter/adapter"
	"github.com/praiakov/adapter/client"
	"github.com/praiakov/adapter/service"
)

func main() {
	client := &client.Client{}
	mac := &service.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adaptee.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
