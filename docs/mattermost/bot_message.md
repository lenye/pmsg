### 推送Mattermost机器人消息

命令参数说明

```text
$ pmsg mattermost bot -h

-a, --user_agent string     http user agent

    --url string            mattermost webhook url
    --raw                   消息内容是原始字符串字面值，没有任何转义处理

args                        参数：消息内容，默认是解释字符串，支持“\”转义
```

样例

linux

```shell
$ pmsg mattermost bot --url webhook_url 'Hello, World!'

ok
```

官方开发文档 [推送mattermost机器人消息](https://developers.mattermost.com/integrate/webhooks/incoming/)