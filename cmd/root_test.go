package cmd

import (
	"bytes"
	"testing"
)

// TestRootCmdNoArgs garante que o comando raiz execute sem quebrar
func TestRootCmdNoArgs(t *testing.T) {
	originalOut := rootCmd.OutOrStdout()
	originalErr := rootCmd.ErrOrStderr()

	t.Cleanup(func() {
		rootCmd.SetOut(originalOut)
		rootCmd.SetErr(originalErr)
		rootCmd.SetArgs([]string{})
	})

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{})

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Expected no error for root command, got: %v", err)
	}
}

// TestSubcommandsExist valida se a árvore do Cobra acoplou todos os comandos do MVP
func TestSubcommandsExist(t *testing.T) {
	expectedSubcommands := []string{"mr", "list", "version"}

	for _, subcmd := range expectedSubcommands {
		cmd, _, err := rootCmd.Find([]string{subcmd})
		if err != nil {
			t.Errorf("Expected subcommand '%s' to be registered, got error: %v", subcmd, err)
		}
		if cmd == nil {
			t.Errorf("Expected subcommand '%s' to exist, got nil", subcmd)
		}
	}
}

// TestMissingArgsErrors força erros de validação isolando cada subcomando
func TestMissingArgsErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"mr missing all args", []string{"mr"}},
		{"mr missing resource arg", []string{"mr", "oci"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalOut := rootCmd.OutOrStdout()
			originalErr := rootCmd.ErrOrStderr()

			t.Cleanup(func() {
				rootCmd.SetOut(originalOut)
				rootCmd.SetErr(originalErr)
				rootCmd.SetArgs([]string{})
			})

			bufOut := new(bytes.Buffer)
			bufErr := new(bytes.Buffer)

			rootCmd.SetOut(bufOut)
			rootCmd.SetErr(bufErr)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()

			if err == nil {
				t.Errorf("Expected validation error for args %v, but execution returned nil", tt.args)
			}
		})
	}
}
