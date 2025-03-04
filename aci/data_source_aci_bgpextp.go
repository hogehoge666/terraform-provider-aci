package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciL3outBgpExternalPolicy() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAciL3outBgpExternalPolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"l3_outside_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func dataSourceAciL3outBgpExternalPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)

	rn := fmt.Sprintf("bgpExtP")
	L3OutsideDn := d.Get("l3_outside_dn").(string)

	dn := fmt.Sprintf("%s/%s", L3OutsideDn, rn)

	bgpExtP, err := getRemoteL3outBgpExternalPolicy(aciClient, dn)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(dn)
	_, err = setL3outBgpExternalPolicyAttributes(bgpExtP, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
