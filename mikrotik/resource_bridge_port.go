package mikrotik

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceBridgePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgePortCreate,
		Read:   resourceInterfaceBridgePortRead,
		Update: resourceInterfaceBridgePortUpdate,
		Delete: resourceInterfaceBridgePortDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bridge": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pvid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"edge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_stacking": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceInterfaceBridgePortCreate(d *schema.ResourceData, m interface{}) error {
	port := prepareBridgePort(d)

	client := m.(client.Mikrotik)

	port, err := client.AddBridgePort(port)
	if err != nil {
		return err
	}

	bridgePortToData(port, d)
	return nil
}

func resourceInterfaceBridgePortRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	port, err := c.FindBridgePort(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if port == nil {
		d.SetId("")
		return nil
	}

	bridgePortToData(port, d)
	return nil
}

func resourceInterfaceBridgePortUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	port := prepareBridgePort(d)
	port.Id = d.Id()

	port, err := c.UpdateBridgePort(port)

	if err != nil {
		return err
	}

	bridgePortToData(port, d)
	return nil
}

func resourceInterfaceBridgePortDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	err := c.DeleteBridgePort(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func bridgePortToData(port *client.BridgePort, d *schema.ResourceData) error {
	d.SetId(port.Id)
	if err := d.Set("bridge", port.Bridge); err != nil {
		return err
	}
	if err := d.Set("interface", port.Interface); err != nil {
		return err
	}
	if err := d.Set("pvid", port.Pvid); err != nil {
		return err
	}
	if err := d.Set("edge", port.Edge); err != nil {
		return err
	}
	if err := d.Set("tag_stacking", port.TagStacking); err != nil {
		return err
	}
	return nil
}

func prepareBridgePort(d *schema.ResourceData) *client.BridgePort {
	port := new(client.BridgePort)

	port.Bridge = d.Get("bridge").(string)
	port.Interface = d.Get("interface").(string)
	port.Pvid = d.Get("pvid").(int)
	port.Edge = d.Get("edge").(string)
	port.TagStacking = d.Get("tag_stacking").(bool)

	return port
}
