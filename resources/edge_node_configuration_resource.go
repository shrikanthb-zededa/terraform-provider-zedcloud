// // Code generated by go-swagger; DO NOT EDIT.
package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiclient "github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

/*
EdgeNodeConfiguration edge node configuration API

Note, an edge-node and a device-config are the same thing. Due is inconcistency in the API definition both terms are used interchangeably in the resulting generated code,.
*/

func EdgeNodeConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext:   GetEdgeNode,
		CreateContext: CreateEdgeNode,
		UpdateContext: UpdateEdgeNode,
		DeleteContext: DeleteEdgeNode,
		Schema:        zschema.DeviceConfigSchema(),
	}
}

func DataResourceEdgeNodeConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetEdgeNode,
		Schema:      zschema.DeviceConfigSchema(),
	}
}

func CreateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DeviceConfigModel(d)
	params := edge_node_configuration.NewEdgeNodeConfigurationCreateEdgeNodeParams()
	params.SetBody(model)

	fmt.Println("------CREATE---------------------------------------")
	if err := os.WriteFile("/tmp/req", []byte("==========REQ=============\n"+spew.Sdump(params)), 0644); err != nil {
		fmt.Println(err)
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationCreateEdgeNode(params, nil)
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
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSucess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	// due to api design, we need to fetch the newly created edge-node/device-config
	// and update the the base os in a subsequent request
	deviceConfig, diags := getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	// publish the api response to local state and d
	zschema.SetDeviceConfigResourceData(d, deviceConfig)
	d.SetId(deviceConfig.ID)

	// to update the base-os-image, the api requires several requests.
	// 1. set the image config with a create or udpate request
	// 2. publish the image config
	// 3. apply the image config
	// image in tf config file
	// var localBaseImageVersion string
	// localBaseImages := params.Body.BaseImage
	// if len(localBaseImages) == 1 && localBaseImages[0] != nil && localBaseImages[0].Version != nil {
	// 	localBaseImageVersion = *(localBaseImages[0].Version)
	// }
	// // image in api response
	// var remoteBaseImageVersion string
	// var remoteBaseImageIsActive bool
	// remoteBaseImages := deviceConfig.BaseImage
	// if len(remoteBaseImages) == 1 && remoteBaseImages[0] != nil {
	// 	if remoteBaseImages[0].Version != nil {
	// 		remoteBaseImageVersion = *(remoteBaseImages[0].Version)
	// 	}
	// 	if remoteBaseImages[0].Activate != nil {
	// 		remoteBaseImageIsActive = *(remoteBaseImages[0].Activate)
	// 	}
	// }

	// // do not update if local config euqals api config
	// updateBaseImage := localBaseImageVersion != remoteBaseImageVersion && remoteBaseImageIsActive
	// if updateBaseImage {
	// 	if diags := PublishBaseOS2(ctx, d, m); len(diags) != 0 {
	// 		return diags
	// 	}
	// 	if diags := ApplyBaseOS(ctx, d, m); len(diags) != 0 {
	// 		return diags
	// 	}
	// }

	// to update admin-state the api requires separate requests
	if diags := setAdminState(ctx, d, m, params.Body.AdminState, deviceConfig.AdminState); len(diags) > 0 {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetEdgeNode(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	fmt.Println("------END CREATE---------------------------------------")
	return diags
}

func setAdminState(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
	localAdminState, remoteAdminState *models.AdminState,
) diag.Diagnostics {

	// no config
	if localAdminState == nil {
		return nil
	}

	// config same as api state
	if remoteAdminState != nil {
		if *localAdminState == *remoteAdminState {
			return nil
		}
	}

	// states differ
	fmt.Println("====================================")
	fmt.Println(*localAdminState)
	fmt.Println(*remoteAdminState)
	fmt.Println("====================================")

	if *localAdminState == models.ADMINSTATE_ACTIVE {
		// do not activate if already registered
		if *remoteAdminState == models.ADMINSTATE_REGISTERED {
			return nil
		}
		if diags := ActivateEdgeNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}
	if *localAdminState == models.ADMINSTATE_INACTIVE {
		if diags := DeactivateEdgeNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}

	return nil
}

func getEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.DeviceConfig, diag.Diagnostics) {
	var diags diag.Diagnostics
	fmt.Println("------get---------------------------------------")

	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: name")...)
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNodeByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	deviceConfig := resp.GetPayload()
	if err := os.WriteFile("/tmp/req", []byte("==========RESP=============\n"+spew.Sdump(deviceConfig)), 0644); err != nil {
		fmt.Println(err)
	}
	fmt.Println("------END get---------------------------------------")

	return deviceConfig, diags
}

func GetEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("------GET---------------------------------------")
	deviceConfig, diags := getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	zschema.SetDeviceConfigResourceData(d, deviceConfig)
	d.SetId(deviceConfig.ID)
	fmt.Println("------END GET---------------------------------------")

	return diags
}

func UpdateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// d.Partial(true)

	// params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeParams()

	// xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	// if xRequestIdIsSet {
	// 	params.XRequestID = xRequestIdVal.(*string)
	// }

	// params.SetBody(zschema.DeviceConfigModel(d))

	// idVal, idIsSet := d.GetOk("id")
	// if idIsSet {
	// 	id, _ := idVal.(string)
	// 	params.ID = id
	// } else {
	// 	diags = append(diags, diag.Errorf("missing client parameter: id")...)
	// 	return diags
	// }

	// // loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	// // props := zschema.GetEdgeNodeConfigurationPropertyFields()
	// // for _, v := range props {
	// // 	if d.HasChange(v) {
	// // 	} else {
	// // 		props = utils.Remove(props, v)
	// // 	}
	// // }

	// // makes a bulk update for all properties that were changed
	// client := m.(*apiclient.Zedcloudapi)
	// resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNode(params, nil)
	// log.Printf("[TRACE] response: %v", resp)
	// if err != nil {
	// 	return append(diags, diag.Errorf("unexpected: %s", err)...)
	// }

	// responseData := resp.GetPayload()
	// if responseData != nil && len(responseData.Error) > 0 {
	// 	for _, err := range responseData.Error {
	// 		diags = append(diags, diag.FromErr(errors.New(err.Details))...)
	// 	}
	// 	return diags
	// }

	// // the zedcloud API does not return the partially updated object but a custom response.
	// // thus, we need to fetch the object and populate the state.
	// if errs := EdgeNodeConfiguration_GetEdgeNodeByName(ctx, d, m); err != nil {
	// 	return append(diags, errs...)
	// }

	d.Partial(false)

	return diags
}

// func EdgeNodeConfiguration_GetDeviceInterfaceTags(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetDeviceInterfaceTagsParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	filterObjIdVal, filterObjIdIsSet := d.GetOk("filter_obj_id")
// 	if filterObjIdIsSet {
// 		params.FilterObjID = filterObjIdVal.(*string)
// 	}

// 	filterObjNameVal, filterObjNameIsSet := d.GetOk("filter_obj_name")
// 	if filterObjNameIsSet {
// 		params.FilterObjName = filterObjNameVal.(*string)
// 	}

// 	nextOrderByVal, nextOrderByIsSet := d.GetOk("next_order_by")
// 	if nextOrderByIsSet {
// 		params.NextOrderBy = nextOrderByVal.([]string)
// 	}

// 	nextPageNumVal, nextPageNumIsSet := d.GetOk("next_page_num")
// 	if nextPageNumIsSet {
// 		params.NextPageNum = nextPageNumVal.(*int64)
// 	}

// 	nextPageSizeVal, nextPageSizeIsSet := d.GetOk("next_page_size")
// 	if nextPageSizeIsSet {
// 		params.NextPageSize = nextPageSizeVal.(*int64)
// 	}

// 	nextPageTokenVal, nextPageTokenIsSet := d.GetOk("next_page_token")
// 	if nextPageTokenIsSet {
// 		params.NextPageToken = nextPageTokenVal.(*string)
// 	}

// 	nextTotalPagesVal, nextTotalPagesIsSet := d.GetOk("next_total_pages")
// 	if nextTotalPagesIsSet {
// 		params.NextTotalPages = nextTotalPagesVal.(*int64)
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetDeviceInterfaceTags(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_GetEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_GetEdgeNodeAttestation(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeAttestationParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNodeAttestation(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_GetEdgeNodeBySerial(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeBySerialParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	serialnoVal, serialnoIsSet := d.GetOk("serialno")
// 	if serialnoIsSet {
// 		params.Serialno = serialnoVal.(string)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: serialno")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNodeBySerial(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_GetEdgeNodeOnboarding(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeOnboardingParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNodeOnboarding(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_GetEdgeviewClientScript(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeviewClientScriptParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeviewClientScript(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

// func EdgeNodeConfiguration_QueryEdgeNodes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := edge_node_configuration.NewEdgeNodeConfigurationQueryEdgeNodesParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	fieldsVal, fieldsIsSet := d.GetOk("fields")
// 	if fieldsIsSet {
// 		params.Fields = fieldsVal.([]string)
// 	}

// 	summaryVal, summaryIsSet := d.GetOk("summary")
// 	if summaryIsSet {
// 		params.Summary = summaryVal.(*bool)
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationQueryEdgeNodes(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	respModel := resp.GetPayload()
// 	zschema.SetDeviceConfigResourceData(d, respModel)

// 	return diags
// }

func ActivateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationActivateEdgeNodeParams()

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
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationActivateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSucess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	return diags
}

func DeactivateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationDeActivateEdgeNodeParams()

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
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationDeActivateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSucess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	return diags
}

// func EdgeNodeConfiguration_BaseOsUpgradeRetryEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationBaseOsUpgradeRetryEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationBaseOsUpgradeRetryEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_Offboard(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationOffboardParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationOffboard(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_PreparePowerOff(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationPreparePowerOffParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationPreparePowerOff(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_Reboot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationRebootParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationReboot(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_StartDebugEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationStartDebugEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.EdgeNodeConfigurationModel(d))
// 	// EdgeNodeConfigurationStartDebugEdgeNodeBody

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationStartDebugEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_StartEdgeviewEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationStartEdgeviewEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.EdgeNodeConfigurationModel(d))
// 	// EdgeNodeConfigurationStartEdgeviewEdgeNodeBody

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationStartEdgeviewEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_StopDebugEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationStopDebugEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationStopDebugEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeNodeConfiguration_StopEdgeviewEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationStopEdgeviewEdgeNodeParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationStopEdgeviewEdgeNode(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

func ApplyBaseOS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParams()

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

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNodeBaseOS(params, nil)
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

	return diags
}

func PublishBaseOS2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params()

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

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNodeBaseOS2(params, nil)
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

	return diags
}

// func EdgeNodeConfiguration_UpdateEdgeNodeBaseOS3(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Params()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
// 	props := zschema.GetEdgeNodeConfigurationPropertyFields()
// 	for _, v := range props {
// 		if d.HasChange(v) {
// 		} else {
// 			props = utils.Remove(props, v)
// 		}
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNodeBaseOS3(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

func DeleteEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationDeleteEdgeNodeParams()

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

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationDeleteEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
