package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ShlomoChanoch/impacto-cli/internal/generator"
	"github.com/spf13/cobra"
)

// These variables are populated at build time using:
// go build -ldflags "-X github.com/ShlomoChanoch/impacto-cli/cmd.version=v0.2.0 ..."
var (
	version = "dev"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version and supported providers",
	Long: `Displays the release version of the Impacto CLI binary,
supported cloud providers, and Crossplane compatibility.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "Impacto CLI Version: %s\n", version)
		fmt.Fprintf(os.Stdout, "Crossplane Compatibility: 2.2+\n")
		fmt.Fprintf(os.Stdout, "Supported Providers: %s\n", strings.Join(generator.ValidProviders(), ", "))

		// List available resources per provider
		for _, provider := range generator.ValidProviders() {
			resources, err := generator.ValidResources(provider)
			if err != nil {
				continue
			}
			fmt.Fprintf(os.Stdout, "  %s: %s\n", provider, strings.Join(resources, ", "))
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Bind to the root command's version flag structure natively supported by Cobra
	rootCmd.Version = version
}
