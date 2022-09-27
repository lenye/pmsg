package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

// weiXinWorkAccessTokenCmd 获取企业微信接口调用凭证
var weiXinWorkAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get work weixin access token",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := token.CmdWorkTokenParams{
			UserAgent:  userAgent,
			CorpID:     corpID,
			CorpSecret: corpSecret,
		}
		if err := token.CmdWorkGetAccessToken(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkAccessTokenCmd)

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpID, flags.CorpID, "i", "", "work weixin corp id (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(flags.CorpID)

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpSecret, flags.CorpSecret, "s", "", "work weixin corp secret (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(flags.CorpSecret)
}
