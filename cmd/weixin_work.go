package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// weiXinWorkCmd 企业微信
var weiXinWorkCmd = &cobra.Command{
	Use:   "work",
	Short: "publish work weixin message",
	Long: `get work weixin access token,
publish work weixin message:
 app message,
 appchat message,
 linkedcorp message,
 externalcontact message,
 customer message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinWorkCmd)
}

// weiXinWorkSetAccessTokenFlags 设置企业微信access_token或者corp_id/corp_secret命令行参数
func weiXinWorkSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "work weixin access token")

	cmd.Flags().StringVarP(&corpID, flags.CorpID, "i", "", "work weixin corp id (required if corp secret is set)")
	cmd.Flags().StringVarP(&corpSecret, flags.CorpSecret, "s", "", "work weixin corp secret (required if corp id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.CorpID)
	cmd.MarkFlagsRequiredTogether(flags.CorpID, flags.CorpSecret)
}
