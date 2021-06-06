package client

import (
	"fmt"
	"log"
)

type Ethernet struct {
	Id        string `mikrotik:".id"`
	Name      string `mikrotik:"name"`
	Mtu       int    `mikrotik:"mtu"`
	L2Mtu     int    `mikrotik:"l2mtu"`
	Advertise string `mikrotik:"advertise"`
}

func (context Mikrotik) FindEthernet(id string) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/ethernet/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Ethernet find response: %v", response)

	if err != nil {
		return nil, err
	}

	ethernet := Ethernet{}
	err = Unmarshal(*response, &ethernet)

	if err != nil {
		return nil, err
	}
	if ethernet.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Ethernet `%s` not found", id))
	}

	return &ethernet, nil
}

func (context Mikrotik) FindEthernetByName(name string) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/ethernet/print", "?name=" + name}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Ethernet find response: %v", response)

	if err != nil {
		return nil, err
	}

	ethernet := Ethernet{}
	err = Unmarshal(*response, &ethernet)

	if err != nil {
		return nil, err
	}
	if ethernet.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Ethernet `%s` not found", name))
	}

	return &ethernet, nil
}

func (context Mikrotik) UpdateEthernet(ethernet *Ethernet) (*Ethernet, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/ethernet/set", ethernet)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return context.FindEthernet(ethernet.Id)
}
