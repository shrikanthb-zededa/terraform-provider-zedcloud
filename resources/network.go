package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiclient "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_network_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

/*
EdgeNetworkConfiguration edge network configuration API
*/

func NetworkResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateNetwork,
		DeleteContext: DeleteNetwork,
		ReadContext:   ReadNetwork,
		UpdateContext: UpdateNetwork,
		Schema:        zschema.Network(),
	}
}

func NetworkDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadNetwork,
		Schema:      zschema.Network(),
	}
}

func CreateNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ToModel(d)
	params := config.CreateNetworkParams()
	params.SetBody(model)

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.Network.CreateNetwork(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := ReadNetwork(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	network, diags := readNetwork(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	zschema.SetNetworkResourceData(d, network)
	d.SetId(network.ID)

	return diags
}

func readNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Network, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := config.GetNetworkByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return nil, diags
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.Network.ReadNetwork(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	network := resp.GetPayload()
	zschema.SetNetworkResourceData(d, network)

	return network, diags
}

func UpdateNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewEdgeNetworkConfigurationUpdateEdgeNetworkParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ToModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.Network.UpdateNetwork(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := ReadNetwork(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewEdgeNetworkConfigurationDeleteEdgeNetworkParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.Network.DeleteNetwork(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
