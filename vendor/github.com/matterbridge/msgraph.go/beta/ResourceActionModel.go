// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ResourceAction undocumented
type ResourceAction struct {
	// Object is the base model of ResourceAction
	Object
	// AllowedResourceActions Allowed Actions
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// NotAllowedResourceActions Not Allowed Actions.
	NotAllowedResourceActions []string `json:"notAllowedResourceActions,omitempty"`
}