// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharedAppleDeviceUser undocumented
type SharedAppleDeviceUser struct {
	// Object is the base model of SharedAppleDeviceUser
	Object
	// UserPrincipalName User name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DataToSync Data to sync
	DataToSync *bool `json:"dataToSync,omitempty"`
	// DataQuota Data quota
	DataQuota *int `json:"dataQuota,omitempty"`
	// DataUsed Data quota
	DataUsed *int `json:"dataUsed,omitempty"`
}