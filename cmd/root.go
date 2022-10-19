package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/version"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pmsg",
	Short: "publish message",
	Long: `publish message:
 weixin offiaccount template message,
 weixin offiaccount template subscribe message (onetime),
 weixin offiaccount subscribe message,
 weixin offiaccount customer message,
 weixin miniprogram subscribe message,
 weixin miniprogram customer message,
 work weixin app message,
 work weixin appchat message,
 work weixin linkedcorp message,
 work weixin externalcontact message,
 work weixin customer message`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}
`)
	rootCmd.Version = version.Print()
}
