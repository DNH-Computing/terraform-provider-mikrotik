package mikrotik

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceBonding() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBondingCreate,
		Read:   resourceInterfaceBondingRead,
		Update: resourceInterfaceBondingUpdate,
		Delete: resourceInterfaceBondingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transmit_hash_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slaves": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceInterfaceBondingCreate(d *schema.ResourceData, m interface{}) error {
	bonding := prepareBonding(d)

	c := m.(client.Mikrotik)

	bonding, err := c.AddBonding(bonding)
	if err != nil {
		return err
	}

	return bondingToData(bonding, d)
}

func resourceInterfaceBondingRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	bonding, err := c.FindBonding(d.Id())

	if err != nil {
		return err
	}

	return bondingToData(bonding, d)
}

func resourceInterfaceBondingUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	bonding := prepareBonding(d)
	bonding.Id = d.Id()

	bonding, err := c.UpdateBonding(bonding)

	if err != nil {
		return err
	}

	return bondingToData(bonding, d)
}

func resourceInterfaceBondingDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	return c.DeleteBonding(d.Id())
}

func bondingToData(bonding *client.Bonding, d *schema.ResourceData) error {
	d.SetId(bonding.Id)
	if err := d.Set("mode", bonding.Mode); err != nil {
		return err
	}
	if err := d.Set("mtu", bonding.Mtu); err != nil {
		return err
	}
	if err := d.Set("name", bonding.Name); err != nil {
		return err
	}
	if err := d.Set("transmit_hash_policy", bonding.TransmitHashPolicy); err != nil {
		return err
	}
	if err := d.Set("slaves", commaSeparatedStringToSlice(bonding.Slaves)); err != nil {
		return err
	}
	return nil
}

func prepareBonding(d *schema.ResourceData) *client.Bonding {
	bonding := new(client.Bonding)

	bonding.Mode = d.Get("mode").(string)
	bonding.Mtu = d.Get("mtu").(int)
	bonding.Name = d.Get("name").(string)
	bonding.TransmitHashPolicy = d.Get("transit_hash_policy").(string)

	bonding.Slaves = setToCommaSeparatedString(d.Get("slaves").(*schema.Set))

	return bonding
}
