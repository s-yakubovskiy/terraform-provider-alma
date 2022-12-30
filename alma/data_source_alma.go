package alma

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var HostURL string = "http://localhost:8000"

// dataSourceServices is the services data source which will pull information on all services served by services-catalog.
func dataSourceAlmaServices() *schema.Resource {
	return &schema.Resource{
		//But watch the Schema, here KEYs are 'Computed: true' not 'Required: true'
		//because we don't want to provide these values while read.
		ReadContext: resourceAlmaReadServices,
		Schema: map[string]*schema.Schema{
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of service",
						},
						"domain": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of domain",
						},
					},
				},
			},
		},
	}
}

// func getDetailsDataSourceRead(d *schema.ResourceData, meta interface{}) (err error) {
// dataSourceService is the service data source which will pull information on selected service served by services-catalog.
func dataSourceAlmaService() *schema.Resource {
	return &schema.Resource{
		//But watch the Schema, here KEYs are 'Computed: true' not 'Required: true'
		//because we don't want to provide these values while read.
		ReadContext: resourceAlmaReadService,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Computed:    false,
				Required:    true,
				Description: "name of service",
				Type:        schema.TypeString,
				Elem:        schema.TypeString,
			},
			"service": &schema.Schema{
				Computed:    true,
				Description: "meta of service",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
