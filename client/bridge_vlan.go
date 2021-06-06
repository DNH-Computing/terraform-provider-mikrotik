package client

import (
	"fmt"
	"log"
)

type BridgeVlan struct {
	Id       string `mikrotik:".id"`
	Bridge   string `mikrotik:"bridge"`
	VlanIds  string `mikrotik:"vlan-ids"`
	Tagged   string `mikrotik:"tagged"`
	Untagged string `mikrotik:"untagged"`
}

func (context Mikrotik) AddBridgeVlan(vlan *BridgeVlan) (*BridgeVlan, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/vlan/add", vlan)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Bridge Port creation response: `%v`", response)

	if err != nil {
		return nil, err
	}

	id := response.Done.Map["ret"]

	return context.FindBridgeVlan(id)
}

func (context Mikrotik) FindBridgeVlan(id string) (*BridgeVlan, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/bridge/vlan/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	response, err := client.RunArgs(cmd)

	log.Printf("[DEBUG] Bridge Port find response: %v", response)

	if err != nil {
		return nil, err
	}

	vlan := BridgeVlan{}
	err = Unmarshal(*response, &vlan)

	if err != nil {
		return nil, err
	}
	if vlan.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("Bridge Port `%s` not found", id))
	}

	return &vlan, nil
}

func (context Mikrotik) UpdateBridgeVlan(vlan *BridgeVlan) (*BridgeVlan, error) {
	client, err := context.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/vlan/set", vlan)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return context.FindBridgeVlan(vlan.Id)
}

func (context Mikrotik) DeleteBridgeVlan(id string) error {
	client, err := context.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/interface/bridge/vlan/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = client.RunArgs(cmd)
	return err
}
