package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetWifiConfigModel(d *schema.ResourceData) *models.NetWifiConfig {
	var crypto *models.NetWifiConfigNetcryptoblock // NetWifiConfigNetcryptoblock
	cryptoInterface, cryptoIsSet := d.GetOk("crypto")
	if cryptoIsSet && cryptoInterface != nil {
		cryptoMap := cryptoInterface.([]interface{})
		if len(cryptoMap) > 0 {
			crypto = NetWifiConfigNetcryptoblockModelFromMap(cryptoMap[0].(map[string]interface{}))
		}
	}
	cryptoKey, _ := d.Get("crypto_key").(string)
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := d.GetOk("encryptedSecrets")
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	identity, _ := d.Get("identity").(string)
	var keyScheme *models.NetworkWiFiKeyScheme // NetworkWiFiKeyScheme
	keySchemeInterface, keySchemeIsSet := d.GetOk("key_scheme")
	if keySchemeIsSet {
		keySchemeModel := keySchemeInterface.(string)
		keyScheme = models.NewNetworkWiFiKeyScheme(models.NetworkWiFiKeyScheme(keySchemeModel))
	}
	priorityInt, _ := d.Get("priority").(int)
	priority := int32(priorityInt)
	var secret *models.NetWifiConfigSecrets // NetWifiConfigSecrets
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = NetWifiConfigSecretsModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	wifiSSID, _ := d.Get("wifi_s_s_id").(string)
	return &models.NetWifiConfig{
		Crypto:           crypto,
		CryptoKey:        cryptoKey,
		EncryptedSecrets: encryptedSecrets,
		Identity:         identity,
		KeyScheme:        keyScheme,
		Priority:         priority,
		Secret:           secret,
		WifiSSID:         wifiSSID,
	}
}

func NetWifiConfigModelFromMap(m map[string]interface{}) *models.NetWifiConfig {
	var crypto *models.NetWifiConfigNetcryptoblock // NetWifiConfigNetcryptoblock
	cryptoInterface, cryptoIsSet := m["crypto"]
	if cryptoIsSet && cryptoInterface != nil {
		cryptoMap := cryptoInterface.([]interface{})
		if len(cryptoMap) > 0 {
			crypto = NetWifiConfigNetcryptoblockModelFromMap(cryptoMap[0].(map[string]interface{}))
		}
	}
	//
	cryptoKey := m["crypto_key"].(string)
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := m["encrypted_secrets"]
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	identity := m["identity"].(string)
	var keyScheme *models.NetworkWiFiKeyScheme // NetworkWiFiKeyScheme
	keySchemeInterface, keySchemeIsSet := m["key_scheme"]
	if keySchemeIsSet {
		keySchemeModel := keySchemeInterface.(string)
		keyScheme = models.NewNetworkWiFiKeyScheme(models.NetworkWiFiKeyScheme(keySchemeModel))
	}
	priority := int32(m["priority"].(int))  // int32
	var secret *models.NetWifiConfigSecrets // NetWifiConfigSecrets
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = NetWifiConfigSecretsModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	wifiSSID := m["wifi_s_s_id"].(string)
	return &models.NetWifiConfig{
		Crypto:           crypto,
		CryptoKey:        cryptoKey,
		EncryptedSecrets: encryptedSecrets,
		Identity:         identity,
		KeyScheme:        keyScheme,
		Priority:         priority,
		Secret:           secret,
		WifiSSID:         wifiSSID,
	}
}

// Update the underlying NetWifiConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetWifiConfigResourceData(d *schema.ResourceData, m *models.NetWifiConfig) {
	d.Set("crypto", SetNetWifiConfigNetcryptoblockSubResourceData([]*models.NetWifiConfigNetcryptoblock{m.Crypto}))
	d.Set("crypto_key", m.CryptoKey)
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	d.Set("identity", m.Identity)
	d.Set("key_scheme", m.KeyScheme)
	d.Set("priority", m.Priority)
	d.Set("secret", SetNetWifiConfigSecretsSubResourceData([]*models.NetWifiConfigSecrets{m.Secret}))
	d.Set("wifi_s_s_id", m.WifiSSID)
}

// Iterate through and update the NetWifiConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetWifiConfigSubResourceData(m []*models.NetWifiConfig) (d []*map[string]interface{}) {
	for _, NetWifiConfigModel := range m {
		if NetWifiConfigModel != nil {
			properties := make(map[string]interface{})
			properties["crypto"] = SetNetWifiConfigNetcryptoblockSubResourceData([]*models.NetWifiConfigNetcryptoblock{NetWifiConfigModel.Crypto})
			properties["crypto_key"] = NetWifiConfigModel.CryptoKey
			properties["encrypted_secrets"] = NetWifiConfigModel.EncryptedSecrets
			properties["identity"] = NetWifiConfigModel.Identity
			properties["key_scheme"] = NetWifiConfigModel.KeyScheme
			properties["priority"] = NetWifiConfigModel.Priority
			properties["secret"] = SetNetWifiConfigSecretsSubResourceData([]*models.NetWifiConfigSecrets{NetWifiConfigModel.Secret})
			properties["wifi_s_s_id"] = NetWifiConfigModel.WifiSSID
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetWifiConfig resource defined in the Terraform configuration
func NetWifiConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"crypto": {
			Description: ``,
			Type:        schema.TypeList, //GoType: NetWifiConfigNetcryptoblock
			Elem: &schema.Resource{
				Schema: NetWifiConfigNetcryptoblockSchema(),
			},
			Optional: true,
		},

		"crypto_key": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encrypted_secrets": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"identity": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"key_scheme": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"priority": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"secret": {
			Description: ``,
			Type:        schema.TypeList, //GoType: NetWifiConfigSecrets
			Elem: &schema.Resource{
				Schema: NetWifiConfigSecretsSchema(),
			},
			Optional: true,
		},

		"wifi_s_s_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetWifiConfig resource
func GetNetWifiConfigPropertyFields() (t []string) {
	return []string{
		"crypto",
		"crypto_key",
		"encrypted_secrets",
		"identity",
		"key_scheme",
		"priority",
		"secret",
		"wifi_s_s_id",
	}
}
