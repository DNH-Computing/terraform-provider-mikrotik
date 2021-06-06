package client

import (
	"fmt"
	"log"
)

type Bonding struct {
	Id     string `mikrotik:".id"`
	Mode   string `mikrotik:"mode"`
	Mtu    int    `mikrotik:"mtu"`
	Name   string `mikrotik:"name"`
	Slaves string `mikrotik:"slaves"`
}

func (client Mikrotik) AddBonding(p *Bonding) (*Bonding, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := Marshal("/interface/bonding/add", p)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] Bonding creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return client.FindBonding(id)
}

func (client Mikrotik) ListBondings() ([]Bonding, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/interface/bonding/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Found bondings: %v", r)

	bondings := []Bonding{}

	err = Unmarshal(*r, &bondings)

	if err != nil {
		return nil, err
	}

	return bondings, nil
}

func (client Mikrotik) FindBonding(id string) (*Bonding, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/interface/bonding/print", "?.id=" + id}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] Bonding response: %v", r)

	if err != nil {
		return nil, err
	}

	bonding := Bonding{}
	err = Unmarshal(*r, &bonding)

	if err != nil {
		return nil, err
	}
	if bonding.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("bonding `%s` not found", id))
	}

	return &bonding, nil
}

func (client Mikrotik) FindBondingByName(name string) (*Bonding, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/interface/bonding/print", "?name=" + name}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] Bonding response: %v", r)

	if err != nil {
		return nil, err
	}

	bonding := Bonding{}
	err = Unmarshal(*r, &bonding)

	if err != nil {
		return nil, err
	}

	if bonding.Name == "" {
		return nil, NewNotFound(fmt.Sprintf("bonding `%s` not found", name))
	}

	return &bonding, nil
}

func (client Mikrotik) UpdateBonding(p *Bonding) (*Bonding, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bonding/set", p)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindBonding(p.Id)
}

func (client Mikrotik) DeleteBonding(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/interface/bonding/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
