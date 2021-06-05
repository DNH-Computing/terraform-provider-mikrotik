package mikrotik

import (
	"errors"

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
		},
	}
}

func resourceInterfaceBridgePortCreate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
	// vlan := prepareVlan(d)

	// c := m.(client.Mikrotik)

	// vlan, err := c.AddVlan(vlan)
	// if err != nil {
	// 	return err
	// }

	// vlanToData(vlan, d)
	// return nil
}

func resourceInterfaceBridgePortRead(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
	// c := m.(client.Mikrotik)

	// vlan, err := c.FindVlan(d.Id())

	// if err != nil {
	// 	d.SetId("")
	// 	return nil
	// }

	// if vlan == nil {
	// 	d.SetId("")
	// 	return nil
	// }

	// vlanToData(vlan, d)
	// return nil
}

func resourceInterfaceBridgePortUpdate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
	// c := m.(client.Mikrotik)

	// vlan := prepareVlan(d)
	// vlan.Id = d.Id()

	// vlan, err := c.UpdateVlan(vlan)
	// vlan.Dynamic = vlan.Dynamic

	// if err != nil {
	// 	return err
	// }

	// vlanToData(vlan, d)
	// return nil
}

func resourceInterfaceBridgePortDelete(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
	// c := m.(client.Mikrotik)

	// err := c.DeleteVlan(d.Id())

	// if err != nil {
	// 	return err
	// }

	// d.SetId("")
	// return nil
}

// func vlanToData(vlan *client.Vlan, d *schema.ResourceData) error {
// 	d.SetId(vlan.Id)
// 	d.Set("name", vlan.Name)
// 	d.Set("vlan_id", vlan.VlanId)
// 	return nil
// }

// func prepareVlan(d *schema.ResourceData) *client.Vlan {
// 	vlan := new(client.Vlan)

// 	vlan.Name = d.Get("name").(string)
// 	vlan.VlanId = d.Get("vlan_id").(int)

// 	return vlan
// }
