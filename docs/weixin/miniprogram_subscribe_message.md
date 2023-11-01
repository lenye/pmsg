### 发送微信小程序订阅消息

命令参数说明

```text
$ pmsg weixin miniprogram subscribe -h

Aliases:
  subscribe, sub

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-o, --to_user string             接收人的open_id (必填)
-p, --template_id string         模版id (必填)
-g, --miniprogram_state string   跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
    --lang string                进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
    --page string                点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
    --raw                        模板数据是原始字符串字面值，没有任何转义处理

args                             参数：模板数据，默认是解释字符串，支持“\”转义
```

样例

windows

```shell
pmsg.exe weixin miniprogram subscribe -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok


使用命令别名
pmsg.exe wx mini sub -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok

```


linux

```shell
$ pmsg weixin miniprogram subscribe -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok


使用命令别名
$ pmsg wx mini sub -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok
```

官方开发文档 [微信小程序订阅消息](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)