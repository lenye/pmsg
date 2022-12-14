### 获取微信接口调用凭证（公众号、小程序）

命令参数说明

```text
$ pmsg weixin token -h

-a, --user_agent string     http user agent

-i, --app_id string         微信app_id (必填)
-s, --app_secret string     微信app_secret (必填)
```

样例

```shell
$ pmsg weixin token -i app_id -s app_secret

使用命令别名
$ pmsg wx token -i app_id -s app_secret

ok; access_token: "access_token", expires_in: 7200, expire_at: "2022-09-20T15:00:20+08:00"
```

官方开发文档 [获取微信接口调用凭证](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html)