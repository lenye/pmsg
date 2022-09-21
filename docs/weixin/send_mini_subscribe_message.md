### 发送微信小程序订阅消息

命令参数说明

```text
$ pmsg weixin mini subscribe -h

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

    --lang string                进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
    --page string                点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。    
-g, --miniprogram_state string   跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
-o, --open_id string             接收人的open_id (必填)
-p, --template_id string         模版id (必填)
```

样例

windows

```shell
pmsg.exe weixin mini subscribe -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok
```

linux

```shell
$ pmsg weixin mini subscribe -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok
```

[微信小程序订阅消息开发文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)