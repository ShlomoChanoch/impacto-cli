package cmd

import (
	"fmt"
	"strings"

	"github.com/ShlomoChanoch/impacto-cli/internal/generator"
	"github.com/spf13/cobra"
)

var mrCmd = &cobra.Command{
	Use:   "mr <provider> <resource>",
	Short: "Generate a Crossplane Managed Resource template",
	Long: fmt.Sprintf(`Outputs a production-ready Managed Resource manifest for the given provider and resource to stdout.
The manifest is ready for customization with clear CHANGE comments.

Available providers:
  %s

Examples:
  impacto mr oci compartment
  impacto mr oci bucket > my-bucket.yaml
  impacto mr oci cluster > oke-cluster.yaml`,
		strings.Join(generator.ValidProviders(), ", ")),
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := args[0]
		resource := args[1]
		return generator.PrintTemplate(provider, resource)
	},
}

func init() {
	rootCmd.AddCommand(mrCmd)
}