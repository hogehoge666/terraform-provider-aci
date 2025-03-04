package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciL3ExtSubnet() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAciL3ExtSubnetRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"external_network_instance_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"aggregate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"scope": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		}),
	}
}

func dataSourceAciL3ExtSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)

	ip := d.Get("ip").(string)

	rn := fmt.Sprintf("extsubnet-[%s]", ip)
	ExternalNetworkInstanceProfileDn := d.Get("external_network_instance_profile_dn").(string)

	dn := fmt.Sprintf("%s/%s", ExternalNetworkInstanceProfileDn, rn)

	l3extSubnet, err := getRemoteL3ExtSubnet(aciClient, dn)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(dn)
	_, err = setL3ExtSubnetAttributes(l3extSubnet, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
