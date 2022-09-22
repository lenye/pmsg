### 发送微信公众号模板消息

命令参数说明

```text
$ pmsg weixin mp template -h

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

    --client_msg_id string   防重入id
-c, --color string           模板内容字体颜色，不填默认为黑色
-m, --mini stringToString    跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar
-o, --open_id string         接收人的open_id (必填)
-p, --template_id string     模版id (必填)
-u, --url string             用户点击后跳转的url
```

样例

windows

```shell
pmsg.exe weixin mp template -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok; msgid: 1234567890
```

linux

```shell
$ pmsg weixin mp template -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok; msgid: 1234567890
```

[微信公众号模板消息开发文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#%E5%8F%91%E9%80%81%E6%A8%A1%E6%9D%BF%E6%B6%88%E6%81%AF)
