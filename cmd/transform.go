package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/QodeSrl/infrar-engine/pkg/engine"
	"github.com/QodeSrl/infrar-engine/pkg/types"
	"github.com/spf13/cobra"
)

var (
	provider   string
	pluginDir  string
	capability string
	inputFile  string
	outputFile string
)

var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform Infrar SDK code to provider-specific code",
	Long: `Transform converts your Infrar SDK code to native cloud provider code.

The transformation happens at build/deploy time, resulting in zero runtime overhead.
Your deployed code contains only native provider SDKs (boto3, google-cloud-storage, etc.).

Examples:
  # Transform from file to AWS
  infrar transform --provider aws --input app.py --output app_aws.py

  # Transform from stdin to GCP
  cat app.py | infrar transform --provider gcp

  # Transform with custom plugin directory
  infrar transform --provider aws --plugins ./my-plugins --input app.py`,
	RunE: runTransform,
}

func init() {
	rootCmd.AddCommand(transformCmd)

	transformCmd.Flags().StringVarP(&provider, "provider", "p", "aws", "Target cloud provider (aws, gcp, azure)")
	transformCmd.Flags().StringVar(&pluginDir, "plugins", "../infrar-plugins/packages", "Path to plugins directory")
	transformCmd.Flags().StringVarP(&capability, "capability", "c", "storage", "Capability to transform (storage, database, etc.)")
	transformCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file to transform (or use stdin)")
	transformCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (or use stdout)")
}

func runTransform(cmd *cobra.Command, args []string) error {
	// Parse provider
	var targetProvider types.Provider
	switch provider {
	case "aws":
		targetProvider = types.ProviderAWS
	case "gcp":
		targetProvider = types.ProviderGCP
	case "azure":
		targetProvider = types.ProviderAzure
	default:
		return fmt.Errorf("unknown provider '%s' (use: aws, gcp, azure)", provider)
	}

	// Create engine
	eng, err := engine.New()
	if err != nil {
		return fmt.Errorf("error creating engine: %w", err)
	}

	// Load rules
	if err := eng.LoadRules(pluginDir, targetProvider, capability); err != nil {
		return fmt.Errorf("error loading rules: %w", err)
	}

	// Read input
	var sourceCode string
	if inputFile != "" {
		content, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("error reading input file: %w", err)
		}
		sourceCode = string(content)
	} else {
		// Read from stdin
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading stdin: %w", err)
		}
		sourceCode = string(content)
	}

	// Transform
	result, err := eng.Transform(sourceCode, targetProvider)
	if err != nil {
		return fmt.Errorf("transformation error: %w", err)
	}

	// Write output
	if outputFile != "" {
		if err := os.WriteFile(outputFile, []byte(result.TransformedCode), 0644); err != nil {
			return fmt.Errorf("error writing output file: %w", err)
		}
		fmt.Fprintf(os.Stderr, "✓ Transformed code written to %s\n", outputFile)
		showMetadata(result)
	} else {
		// Write ONLY code to stdout (for piping)
		fmt.Print(result.TransformedCode)
		// Metadata goes to stderr
		showMetadata(result)
	}

	return nil
}

func showMetadata(result *types.TransformationResult) {
	// Show warnings
	if len(result.Warnings) > 0 {
		fmt.Fprintln(os.Stderr, "\nWarnings:")
		for _, w := range result.Warnings {
			fmt.Fprintf(os.Stderr, "  - %s\n", w.Message)
		}
	}

	// Show metadata
	if len(result.Imports) > 0 {
		fmt.Fprintln(os.Stderr, "\nImports added:")
		for _, imp := range result.Imports {
			fmt.Fprintf(os.Stderr, "  - %s\n", imp)
		}
	}

	if len(result.Requirements) > 0 {
		fmt.Fprintln(os.Stderr, "\nDependencies required:")
		for _, req := range result.Requirements {
			fmt.Fprintf(os.Stderr, "  - %s %s\n", req.Package, req.Version)
		}
	}
}
