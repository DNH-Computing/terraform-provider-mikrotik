package mikrotik

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceVlanCreate,
		Read:   resourceInterfaceVlanRead,
		Update: resourceInterfaceVlanUpdate,
		Delete: resourceInterfaceVlanDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceInterfaceVlanCreate(d *schema.ResourceData, m interface{}) error {
	vlan := prepareVlan(d)

	client := m.(client.Mikrotik)

	vlan, err := client.AddVlan(vlan)
	if err != nil {
		return err
	}

	vlanToData(vlan, d)
	return nil
}

func resourceInterfaceVlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	vlan, err := c.FindVlan(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if vlan == nil {
		d.SetId("")
		return nil
	}

	vlanToData(vlan, d)
	return nil
}
func resourceInterfaceVlanUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	vlan := prepareVlan(d)
	vlan.Id = d.Id()

	vlan, err := c.UpdateVlan(vlan)

	if err != nil {
		return err
	}

	return vlanToData(vlan, d)
}
func resourceInterfaceVlanDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	err := c.DeleteVlan(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func vlanToData(vlan *client.Vlan, d *schema.ResourceData) error {
	d.SetId(vlan.Id)
	if err := d.Set("name", vlan.Name); err != nil {
		return err
	}
	if err := d.Set("vlan_id", vlan.VlanId); err != nil {
		return err
	}
	return nil
}

func prepareVlan(d *schema.ResourceData) *client.Vlan {
	vlan := new(client.Vlan)

	vlan.Name = d.Get("name").(string)
	vlan.VlanId = d.Get("vlan_id").(int)

	return vlan
}
