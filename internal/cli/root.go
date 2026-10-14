package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	env "github.com/tomato2718/docker-registry-manager/internal/env"
)

var rootCmd = &cobra.Command{
	Use:   "docker-registry-manager",
	Short: "Docker register manager",
	Run: func(cmd *cobra.Command, args []string) {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
	},
}

func Execute() {
	if env.RegistryUrl == "" {
		println("Env unset")
		os.Exit(1)
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
