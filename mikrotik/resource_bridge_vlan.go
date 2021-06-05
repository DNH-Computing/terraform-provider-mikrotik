package mikrotik

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceBridgeVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgeVlanCreate,
		Read:   resourceInterfaceBridgeVlanRead,
		Update: resourceInterfaceBridgeVlanUpdate,
		Delete: resourceInterfaceBridgeVlanDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bridge": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tagged": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"untagged": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceInterfaceBridgeVlanCreate(d *schema.ResourceData, m interface{}) error {
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

func resourceInterfaceBridgeVlanRead(d *schema.ResourceData, m interface{}) error {
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

func resourceInterfaceBridgeVlanUpdate(d *schema.ResourceData, m interface{}) error {
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

func resourceInterfaceBridgeVlanDelete(d *schema.ResourceData, m interface{}) error {
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
