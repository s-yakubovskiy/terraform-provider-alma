package alma

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v1 "gitlab.com/fahecom/platform/release-eng/service-catalog-app/api/service-catalog/v1"
	sc "gitlab.com/fahecom/platform/release-eng/service-catalog-app/pkg/model"
)

func resourceAlmaReadServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceAlmaReadServices", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	res, err := c.almaclient.GetAllServices()
	if err != nil {
		return diag.FromErr(err)
	}
	if res != nil {
		//As the return item is a []Services, lets Unmarshal it into "services"
		resItems := flattenServices(res)
		if err := d.Set("services", resItems); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	log.Printf("[DEBUG] %s: resourceAlmaReadServices finished successfully", d.Id())
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenServices(servicesList []*sc.Service) []interface{} {
	if servicesList != nil {
		services := make([]interface{}, len(servicesList))
		for i, service := range servicesList {
			sl := make(map[string]interface{})

			sl["name"] = service.Name
			sl["domain"] = service.Domain
			services[i] = sl
		}
		return services
	}
	return make([]interface{}, 0)
}

func resourceAlmaReadService(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceAlmaReadService", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	resourceName := d.Get("name").(string)

	res, err := c.almaclient.GetService(resourceName)
	if err != nil {
		return diag.FromErr(err)
	}
	if res != nil {
		//As the return item is a []Services, lets Unmarshal it into "services"
		resItems := flattenService(res)
		if err := d.Set("service", resItems); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	log.Printf("[DEBUG] %s: resourceAlmaReadServics finished successfully", d.Id())
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenService(servicesList *v1.GetServiceReply) interface{} {
	if servicesList != nil {
		sl := make(map[string]interface{})

		sl["name"] = servicesList.Service.Name
		sl["domain"] = servicesList.Service.Domain
		return sl
	}
	return nil
}
