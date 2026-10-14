package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	env "github.com/tomato2718/docker-registry-manager/internal/env"
	"github.com/tomato2718/docker-registry-manager/internal/httpclient"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all images",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := httpclient.Get(env.RegistryUrl)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, tag := range res {
			fmt.Println(tag)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
