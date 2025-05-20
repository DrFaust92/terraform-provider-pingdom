package pingdom

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PINGDOM_API_TOKEN", nil),
			},
			"solarwinds_user": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLARWINDS_USER", nil),
			},
			"solarwinds_passwd": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLARWINDS_PASSWD", nil),
			},
			"solarwinds_org_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLARWINDS_ORG_ID", nil),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PINGDOM_MAX_RETRIES", 3),
				Description: "Maximum number of retries for API calls.",
			},
			"retry_wait_min": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PINGDOM_RETRY_WAIT_MIN", 1),
				Description: "Minimum wait time in seconds between retries.",
			},
			"retry_wait_max": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PINGDOM_RETRY_WAIT_MAX", 30),
				Description: "Maximum wait time in seconds between retries.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pingdom_check":       resourcePingdomCheck(),
			"pingdom_team":        resourcePingdomTeam(),
			"pingdom_contact":     resourcePingdomContact(),
			"pingdom_integration": resourcePingdomIntegration(),
			"pingdom_maintenance": resourcePingdomMaintenance(),
			"pingdom_occurrence":  resourcePingdomOccurrences(),
			"pingdom_tms_check":   resourcePingdomTmsCheck(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"pingdom_contact":      dataSourcePingdomContact(),
			"pingdom_contacts":     dataSourcePingdomContacts(),
			"pingdom_team":         dataSourcePingdomTeam(),
			"pingdom_teams":        dataSourcePingdomTeams(),
			"pingdom_integration":  dataSourcePingdomIntegration(),
			"pingdom_integrations": dataSourcePingdomIntegrations(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing Pingdom client")
	return config.Client()
}
