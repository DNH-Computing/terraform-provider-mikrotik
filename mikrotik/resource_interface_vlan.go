package mikrotik

import (
	"errors"

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
	return errors.New("Not yet implemented")
}

func resourceInterfaceVlanRead(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
}
func resourceInterfaceVlanUpdate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
}
func resourceInterfaceVlanDelete(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not yet implemented")
}
