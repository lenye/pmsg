### 发送微信小程序客服消息

命令参数说明

```text
$ pmsg weixin miniprogram customer -h

Aliases:
  customer, kf

-a, --user_agent string     http user agent

-t, --access_token string   微信接口调用凭证
-i, --app_id string         微信app_id
-s, --app_secret string     微信app_secret

如果没有提供 access_token，需要提供微信 app_id 和 app_secret 获取 access_token

-o, --to_user string        接收人的open_id (必填)
-m, --msg_type string       消息类型 (必填)，text(文本消息)、image(图片消息)、link(图文链接)、miniprogrampage(小程序卡片)
    --raw                   消息内容是原始字符串字面值，没有任何转义处理

args                        参数：消息内容，默认是解释字符串，支持“\”转义
```

消息内容

1. 文本消息 --msg_type text
    ```text
    HelloWorld
    ```

1. 图片消息 --msg_type image
    ```text
    MEDIA_ID
    ```

1. 图文链接 --msg_type link
    ```json
    {
      "title": "Happy Day",
      "description": "Is Really A Happy Day",
      "url": "URL",
      "thumb_url": "THUMB_URL"
    }
    ```

1. 小程序卡片 --msg_type miniprogrampage
    ```json
    {
      "title": "title",
      "pagepath": "pagepath",
      "thumb_media_id": "thumb_media_id"
    }
    ```

样例

windows

发送文本消息

```shell
pmsg.exe weixin miniprogram customer -i app_id -s app_secret -o open_id -m text "HelloWorld"

使用命令别名
pmsg.exe wx mini kf -i app_id -s app_secret -o open_id -m text "HelloWorld"

ok
```

发送小程序卡片

```shell
pmsg.exe weixin miniprogram customer -i app_id -s app_secret -o open_id -m miniprogrampage "{\"title\":\"title\",\"pagepath\":\"pagepath\",\"thumb_media_id\":\"thumb_media_id\"}"

ok
```

linux

发送文本消息

```shell
$ pmsg weixin miniprogram customer -i app_id -s app_secret -o open_id -m text 'HelloWorld'

使用命令别名
$ pmsg wx mini kf -i app_id -s app_secret -o open_id -m text 'HelloWorld'

ok
```

发送小程序卡片

```shell
pmsg weixin miniprogram customer -i app_id -s app_secret -o open_id -m miniprogrampage '{"title":"title","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}'

ok
```

官方开发文档 [微信小程序客服消息](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html)