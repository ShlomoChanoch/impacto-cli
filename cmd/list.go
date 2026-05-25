package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/ShlomoChanoch/impacto-cli/internal/generator"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available providers and resource types",
	Long: `Display a comprehensive catalog of providers and resource blueprints 
supported by the Impacto CLI.

Use 'impacto mr <provider> <resource>' to generate a Managed Resource manifest.
Each resource template includes clear CHANGE comments for customization.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "PROVIDER\tRESOURCE\tDESCRIPTION")
		fmt.Fprintln(w, "--------\t--------\t-----------")

		providers := generator.ValidProviders()
		for _, provider := range providers {
			resources, err := generator.ValidResources(provider)
			if err != nil {
				continue
			}

			// Sort resources for consistent output
			sort.Strings(resources)

			for i, resource := range resources {
				desc := generator.ResourceDescription(provider, resource)
				if i == 0 {
					fmt.Fprintf(w, "%s\t%s\t%s\n", provider, resource, desc)
				} else {
					fmt.Fprintf(w, "%s\t%s\t%s\n", "", resource, desc)
				}
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
