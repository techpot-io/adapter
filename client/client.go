package client

import (
	"fmt"

	"github.com/praiakov/adapter/computer"
)

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com computer.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
