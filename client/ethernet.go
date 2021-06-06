package client

import (
	"fmt"
	"log"
)

type Ethernet struct {
	Id   string `mikrotik:".Id"`
	Name string `mikrotik:"name"`
	Mtu  int    `mikrotik:"mtu"`
}

func (context Mikrotik) FindEthernet(id string) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/ethernet", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Ethernet find response: %v", response)

	if err != nil {
		return nil, err
	}

	port := Ethernet{}
	err = Unmarshal(*response, port)

	if err != nil {
		return nil, err
	}
	if port.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Ethernet `%s` not found", id))
	}

	return &port, nil
}

func (context Mikrotik) FindEthernetByName(name string) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/ethernet", "?name=" + name}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Ethernet find response: %v", response)

	if err != nil {
		return nil, err
	}

	port := Ethernet{}
	err = Unmarshal(*response, port)

	if err != nil {
		return nil, err
	}
	if port.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Ethernet `%s` not found", name))
	}

	return &port, nil
}

func (context Mikrotik) UpdateEthernet(port *Ethernet) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/ethernet/set", port)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return context.FindEthernet(port.Id)
}
