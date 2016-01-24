package gogen

// ResourceContainer is used to set and get
// resources among the generators.
type ResourceContainer struct {
	Container map[string]interface{}
}

// Set adds resource to the resource container
func (rc *ResourceContainer) Set(name string, v interface{}) {
	rc.Container[name] = v
}

// Get retrieves resource from the resource container
func (rc *ResourceContainer) Get(name string) interface{} {
	return rc.Container[name]
}
