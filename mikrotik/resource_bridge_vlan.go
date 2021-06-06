package mikrotik

import (
	"strings"

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
	if err := d.Set("vlan_ids", strings.Split(vlan.VlanIds, ",")); err != nil {
		return err
	}
	if err := d.Set("tagged", strings.Split(vlan.Tagged, ",")); err != nil {
		return err
	}
	if err := d.Set("untagged", strings.Split(vlan.Untagged, ",")); err != nil {
		return err
	}
	return nil
}

func prepareBridgeVlan(d *schema.ResourceData) *client.BridgeVlan {
	vlan := new(client.BridgeVlan)

	vlan.Bridge = d.Get("bridge").(string)

	var vlanStrings []string
	for _, vlan := range d.Get("vlan_ids").(*schema.Set).List() {
		vlanStrings = append(vlanStrings, vlan.(string))
	}
	vlan.VlanIds = strings.Join(vlanStrings, ",")

	var taggedStrings []string
	for _, tagged := range d.Get("tagged").(*schema.Set).List() {
		taggedStrings = append(taggedStrings, tagged.(string))
	}
	vlan.Tagged = strings.Join(taggedStrings, ",")

	var untaggedStrings []string
	for _, untagged := range d.Get("untagged").(*schema.Set).List() {
		untaggedStrings = append(untaggedStrings, untagged.(string))
	}
	vlan.Untagged = strings.Join(untaggedStrings, ",")

	return vlan
}
