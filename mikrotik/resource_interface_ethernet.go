package mikrotik

import (
	"log"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetCreate,
		Read:   resourceInterfaceEthernetRead,
		Update: resourceInterfaceEthernetUpdate,
		Delete: resourceInterfaceEthernetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l2mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceEthernetCreate(d *schema.ResourceData, m interface{}) error {
	port := prepareEthernet(d)
	// vlan := prepareVlan(d)

	c := m.(client.Mikrotik)

	existingPort, err := c.FindEthernetByName(port.Name)

	if err != nil {
		return err
	}

	d.SetId(existingPort.Id)

	return resourceInterfaceEthernetUpdate(d, m)
}

func resourceInterfaceEthernetRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	port, err := c.FindEthernet(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if port == nil {
		d.SetId("")
		return nil
	}

	return ethernetToData(port, d)
}

func resourceInterfaceEthernetUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	port := prepareEthernet(d)
	port.Id = d.Id()

	port, err := c.UpdateEthernet(port)

	if err != nil {
		return err
	}

	return ethernetToData(port, d)
}

func resourceInterfaceEthernetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[WARN] Not deleting ethernet interface `%s` because you can't", d.Get("name"))
	return nil
}

func ethernetToData(port *client.Ethernet, d *schema.ResourceData) error {
	d.SetId(port.Id)
	d.Set("name", port.Name)
	d.Set("mtu", port.Mtu)
	d.Set("l2mtu", port.L2Mtu)
	return nil
}

func prepareEthernet(d *schema.ResourceData) *client.Ethernet {
	port := new(client.Ethernet)

	port.Name = d.Get("name").(string)
	port.Mtu = d.Get("mtu").(int)
	port.L2Mtu = d.Get("l2mtu").(int)

	return port
}
