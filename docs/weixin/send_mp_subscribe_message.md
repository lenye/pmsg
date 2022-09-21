### 发送微信公众号订阅通知

命令参数说明

```text
$ pmsg weixin mp subscribe -h

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

-m, --mini stringToString    跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar
-o, --open_id string         接收人的open_id (必填)
-p, --template_id string     模版id (必填)
    --page string            跳转网页时填写

```

样例

windows

```shell
pmsg.exe weixin mp subscribe -i app_id -s app_secret -p template_id -o open_id "{\"first\":{\"value\":\"测试\"}}"

ok
```

linux

```shell
$ pmsg weixin mp subscribe -i app_id -s app_secret -p template_id -o open_id '{"first":{"value":"测试"}}'

ok
```

[微信公众号订阅通知开发文档](https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#send%E5%8F%91%E9%80%81%E8%AE%A2%E9%98%85%E9%80%9A%E7%9F%A5)