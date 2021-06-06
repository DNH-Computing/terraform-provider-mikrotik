package client

import (
	"fmt"
	"log"
)

type Vlan struct {
	Id     string `mikrotik:".id"`
	Name   string `mikrotik:"name"`
	VlanId int    `mikrotik:"vlan-id"`
}

func (client Mikrotik) AddVlan(vlan *Vlan) (*Vlan, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/vlan", vlan)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] vlan creation response: `%v`", response)

	if err != nil {
		return nil, err
	}

	id := response.Done.Map["ret"]
	return client.FindVlan(id)
}

func (client Mikrotik) FindVlan(id string) (*Vlan, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/vlan", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] Vlan find response: %v", response)

	if err != nil {
		return nil, err
	}

	port := Vlan{}
	err = Unmarshal(*response, port)

	if err != nil {
		return nil, err
	}
	if port.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Vlan `%s` not found", id))
	}

	return &port, nil
}

func (client Mikrotik) UpdateVlan(vlan *Vlan) (*Vlan, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/vlan/set", vlan)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindVlan(vlan.Id)
}

func (context Mikrotik) DeleteVlan(id string) error {
	client, err := context.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/interface/vlan/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)
	return err
}
