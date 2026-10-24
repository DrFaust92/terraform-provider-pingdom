package pingdom

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePingdomProbes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingdomProbesRead,

		Schema: map[string]*schema.Schema{
			"only_active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Only return active probes.",
			},
			"region": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NA", "EU", "APAC", "LATAM"}, false),
				Description:  "Only return probes in this region.",
			},
			"ipv4_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ipv6_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"probes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id":          {Type: schema.TypeInt, Computed: true},
						"name":        {Type: schema.TypeString, Computed: true},
						"active":      {Type: schema.TypeBool, Computed: true},
						"hostname":    {Type: schema.TypeString, Computed: true},
						"ip":          {Type: schema.TypeString, Computed: true},
						"ipv6":        {Type: schema.TypeString, Computed: true},
						"city":        {Type: schema.TypeString, Computed: true},
						"country":     {Type: schema.TypeString, Computed: true},
						"country_iso": {Type: schema.TypeString, Computed: true},
						"region":      {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourcePingdomProbesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Clients).Pingdom

	params := map[string]string{}
	if d.Get("only_active").(bool) {
		params["onlyactive"] = "true"
	}

	probes, err := client.Probes.List(params)
	if err != nil {
		return diag.Errorf("Error retrieving probes: %s", err)
	}

	region := d.Get("region").(string)
	ipv4 := make([]string, 0, len(probes))
	ipv6 := make([]string, 0, len(probes))
	result := make([]map[string]interface{}, 0, len(probes))
	for _, p := range probes {
		if region != "" && !strings.EqualFold(p.Region, region) {
			continue
		}

		if p.IP != "" {
			ipv4 = append(ipv4, p.IP)
		}
		if p.IPv6 != "" {
			ipv6 = append(ipv6, p.IPv6)
		}
		result = append(result, map[string]interface{}{
			"id":          p.ID,
			"name":        p.Name,
			"active":      p.Active,
			"hostname":    p.Hostname,
			"ip":          p.IP,
			"ipv6":        p.IPv6,
			"city":        p.City,
			"country":     p.Country,
			"country_iso": p.CountryISO,
			"region":      p.Region,
		})
	}

	d.SetId("probes" + strings.ToLower(region))
	if err := d.Set("ipv4_addresses", ipv4); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ipv6_addresses", ipv6); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("probes", result); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
