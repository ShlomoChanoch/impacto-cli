package generator

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ShlomoChanoch/impacto-cli/templates"
)

// ProviderResources maps resource names to their descriptions for each provider.
var ProviderResources = map[string]map[string]string{
	"oci": {
		"compartment": "OCI Compartment (pré-requisito para todos os recursos)",
		"vcn":         "Virtual Cloud Network (pré-requisito de rede)",
		"subnet":      "Subnet dentro de uma VCN",
		"bucket":      "Object Storage Bucket",
	},
}

// ValidProviders returns a sorted list of available provider names.
func ValidProviders() []string {
	providers := make([]string, 0, len(ProviderResources))
	for p := range ProviderResources {
		providers = append(providers, p)
	}
	sort.Strings(providers)
	return providers
}

// ValidResources returns a sorted list of available resource names for a given provider.
func ValidResources(provider string) ([]string, error) {
	resources, ok := ProviderResources[provider]
	if !ok {
		return nil, fmt.Errorf("unknown provider '%s'. Available providers: %s",
			provider, strings.Join(ValidProviders(), ", "))
	}

	names := make([]string, 0, len(resources))
	for r := range resources {
		names = append(names, r)
	}
	sort.Strings(names)
	return names, nil
}

// ResourceDescription returns the description for a given provider and resource.
func ResourceDescription(provider, resource string) string {
	if resources, ok := ProviderResources[provider]; ok {
		if desc, ok := resources[resource]; ok {
			return desc
		}
	}
	return ""
}

// PrintTemplate reads a Managed Resource template from the embedded filesystem
// and writes it to stdout.
// provider must be a valid provider (e.g., "oci")
// resource must be a valid resource for that provider (e.g., "bucket", "instance")
func PrintTemplate(provider, resource string) error {
	// Validate provider
	if _, ok := ProviderResources[provider]; !ok {
		return fmt.Errorf("unknown provider '%s'. Available providers: %s\nRun 'impacto list' for details.",
			provider, strings.Join(ValidProviders(), ", "))
	}

	// Validate resource
	if _, ok := ProviderResources[provider][resource]; !ok {
		validRes, _ := ValidResources(provider)
		return fmt.Errorf("unknown resource '%s' for provider '%s'. Available resources: %s\nRun 'impacto list' for details.",
			resource, provider, strings.Join(validRes, ", "))
	}

	// Build file path
	filePath := fmt.Sprintf("%s/%s.yaml", provider, resource)

	// Read from embedded filesystem
	data, err := templates.FS.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("template not found: %s. This may be a bug — the resource is registered but the template file is missing.", filePath)
	}

	// Write to stdout
	_, err = os.Stdout.Write(data)
	return err
}
