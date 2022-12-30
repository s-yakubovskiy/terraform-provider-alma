package alma

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	almaclient "github.com/s-yakubovskiy/alma-client"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

// NewApiClient will return a new instance of ApiClient using which we can communicate with services-backend
func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	client, err := c.NewAlmaClient()
	if err != nil {
		return c, diag.FromErr(err)
	}
	c.almaclient = client
	return c, nil
}

type ApiClient struct {
	data       *schema.ResourceData
	almaclient *almaclient.Client
}

func (a *ApiClient) NewAlmaClient() (*almaclient.Client, error) {
	host := a.data.Get("host").(string)
	c, err := almaclient.NewClient(&host, nil, nil)
	if err != nil {
		return c, err
	}
	return c, nil
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"alma_services": dataSourceAlmaServices(),
			"alma_service":  dataSourceAlmaService(),
		},
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALMA_BACKEND_HOST_URL", "http://service-catalog.platform.k8s.dev.cnm.team"),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}
