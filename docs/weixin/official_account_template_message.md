### 发送微信公众号模板消息

命令参数说明

```text
$ pmsg weixin offiaccount template -h

Aliases:
  template, tpl

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-o, --to_user string         接收人的open_id (必填)
-p, --template_id string     模版id (必填)
-c, --client_msg_id string   防重入id
    --color string           模板内容字体颜色，不填默认为黑色
    --mini stringToString    跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar
    --url string             用户点击后跳转的url
    --raw                    模板数据是原始字符串字面值，没有任何转义处理

args                         参数：模板数据，默认是解释字符串，支持“\”转义    
```

样例

windows

```shell
pmsg.exe weixin offiaccount template -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok; msgid: 1234567890
```

linux

```shell
$ pmsg weixin offiaccount template -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

使用命令别名
$ pmsg wx mp tpl -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok; msgid: 1234567890
```

官方开发文档 [微信公众号模板消息](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#%E5%8F%91%E9%80%81%E6%A8%A1%E6%9D%BF%E6%B6%88%E6%81%AF)
