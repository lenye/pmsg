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
	Short: "get weixin work access token",
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

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpID, flags.NameCorpID, "i", "", "weixin corp id (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(flags.NameCorpID)

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpSecret, flags.NameCorpSecret, "s", "", "weixin corp secret (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(flags.NameCorpSecret)
}
