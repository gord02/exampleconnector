// Code generated by mdatagen. DO NOT EDIT.

package metadata

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for resourcedetectionprocessor/eks resource attributes.
type ResourceAttributesConfig struct {
	CloudPlatform ResourceAttributeConfig `mapstructure:"cloud.platform"`
	CloudProvider ResourceAttributeConfig `mapstructure:"cloud.provider"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		CloudPlatform: ResourceAttributeConfig{
			Enabled: true,
		},
		CloudProvider: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}
