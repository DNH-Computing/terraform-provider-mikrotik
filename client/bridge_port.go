package client

import (
	"fmt"
	"log"
)

type BridgePort struct {
	Id          string `mikrotik:".id"`
	Bridge      string `mikrotik:"bridge"`
	Interface   string `mikrotik:"interface"`
	Pvid        int    `mikrotik:"pvid"`
	Edge        string `mikrotik:"edge"`
	TagStacking bool   `mikrotik:"tag-stacking"`
}

func (context Mikrotik) AddBridgePort(port *BridgePort) (*BridgePort, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/port/add", port)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Bridge Port creation response: `%v`", response)

	if err != nil {
		return nil, err
	}

	id := response.Done.Map["ret"]

	return context.FindBridgePort(id)
}

func (context Mikrotik) FindBridgePort(id string) (*BridgePort, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/bridge/port/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Bridge Port find response: %v", response)

	if err != nil {
		return nil, err
	}

	port := BridgePort{}
	err = Unmarshal(*response, &port)

	if err != nil {
		return nil, err
	}
	if port.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Bridge Port `%s` not found", id))
	}

	return &port, nil
}

func (context Mikrotik) UpdateBridgePort(port *BridgePort) (*BridgePort, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/port/set", port)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return context.FindBridgePort(port.Id)
}

func (context Mikrotik) DeleteBridgePort(id string) error {
	client, err := context.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/interface/bridge/port/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)
	return err
}
