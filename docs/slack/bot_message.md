### 推送slack机器人消息

命令参数说明

```text
$ pmsg slack bot -h

-a, --user_agent string     http user agent

    --url string   slack webhook url

args                        参数：消息内容
```

样例

linux

```shell
$ pmsg slack bot --url webhook_url '{"text": "Hello, World!"}'

ok
```

官方开发文档 [推送slack机器人消息](https://api.slack.com/messaging/webhooks)