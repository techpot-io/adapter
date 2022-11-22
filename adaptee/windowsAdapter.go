package adaptee

import (
	"fmt"

	"github.com/praiakov/adapter/adapter"
)

type WindowsAdapter struct {
	WindowMachine *adapter.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lighting signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
