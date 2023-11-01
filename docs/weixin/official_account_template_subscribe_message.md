### 发送微信公众号一次性订阅消息

命令参数说明

```text
$ pmsg weixin offiaccount template subscribe -h

Aliases:
  subscribe, sub

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-o, --to_user string         接收人的open_id (必填)
-p, --template_id string     模版id (必填)
    --mini stringToString    跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar
    --url string             用户点击后跳转的url
    --scene string           订阅场景值 (必填)
    --title string           消息标题，15字以内 (必填)
    --raw                    模板数据是原始字符串字面值，没有任何转义处理

args                         参数：模板数据，默认是解释字符串，支持“\”转义
```

样例

windows

```shell
pmsg.exe weixin offiaccount template subscribe -i app_id -s app_secret -p template_id -o open_id --scene scene --title title "{\"first\":{\"value\":\"测试\"}}"

ok
```

linux

```shell
$ pmsg weixin offiaccount template subscribe -i app_id -s app_secret -p template_id -o open_id --scene scene --title title '{"first":{"value":"测试"}}'

使用命令别名
$ pmsg wx mp tpl sub -i app_id -s app_secret -p template_id -o open_id --scene scene --title title '{"first":{"value":"测试"}}'

ok
```

官方开发文档 [微信公众号一次性订阅消息](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html)
