package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

const (
	nameCorpID     = "corp_id"
	nameCorpSecret = "corp_secret"
)

var (
	corpID     string
	corpSecret string
)

// weiXinWorkAccessTokenCmd 获取企业微信接口调用凭证
var weiXinWorkAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get weixin work access token",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinWorkGetAccessToken(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkAccessTokenCmd)

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpID, nameCorpID, "i", "", "weixin corp id (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(nameCorpID)

	weiXinWorkAccessTokenCmd.Flags().StringVarP(&corpSecret, nameCorpSecret, "s", "", "weixin corp secret (required)")
	weiXinWorkAccessTokenCmd.MarkFlagRequired(nameCorpSecret)
}

// WeiXinWorkGetAccessToken 获取企业微信接口调用凭证
func WeiXinWorkGetAccessToken() error {

	if accessToken != "" {
		return fmt.Errorf("flags %q not required", nameAccessToken)
	}

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	accessTokenResp, err := token.GetAccessToken(corpID, corpSecret)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, accessTokenResp))

	return nil
}
