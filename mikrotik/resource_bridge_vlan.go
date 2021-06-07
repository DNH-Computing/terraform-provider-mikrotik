package mikrotik

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client"

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
					Type: schema.TypeString,
				},
			},
			"tagged": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"untagged": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceInterfaceBridgeVlanCreate(d *schema.ResourceData, m interface{}) error {
	vlan := prepareBridgeVlan(d)

	c := m.(client.Mikrotik)

	vlan, err := c.AddBridgeVlan(vlan)
	if err != nil {
		return err
	}

	return bridgeVlanToData(vlan, d)
}

func resourceInterfaceBridgeVlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	vlan, err := c.FindBridgeVlan(d.Id())

	if err != nil {
		return err
	}

	return bridgeVlanToData(vlan, d)
}

func resourceInterfaceBridgeVlanUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	vlan := prepareBridgeVlan(d)
	vlan.Id = d.Id()

	vlan, err := c.UpdateBridgeVlan(vlan)

	if err != nil {
		return err
	}

	return bridgeVlanToData(vlan, d)
}

func resourceInterfaceBridgeVlanDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	return c.DeleteBridgeVlan(d.Id())
}

func bridgeVlanToData(vlan *client.BridgeVlan, d *schema.ResourceData) error {
	d.SetId(vlan.Id)
	if err := d.Set("bridge", vlan.Bridge); err != nil {
		return err
	}
	if err := d.Set("vlan_ids", commaSeparatedStringToSlice(vlan.VlanIds)); err != nil {
		return err
	}
	if err := d.Set("tagged", commaSeparatedStringToSlice(vlan.Tagged)); err != nil {
		return err
	}
	if err := d.Set("untagged", commaSeparatedStringToSlice(vlan.Untagged)); err != nil {
		return err
	}
	return nil
}

func prepareBridgeVlan(d *schema.ResourceData) *client.BridgeVlan {
	vlan := new(client.BridgeVlan)

	vlan.Bridge = d.Get("bridge").(string)

	vlan.VlanIds = setToCommaSeparatedString(d.Get("vlan_ids").(*schema.Set))

	vlan.Tagged = setToCommaSeparatedString(d.Get("tagged").(*schema.Set))

	vlan.Untagged = setToCommaSeparatedString(d.Get("untagged").(*schema.Set))

	return vlan
}
