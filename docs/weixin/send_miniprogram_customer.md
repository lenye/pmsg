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

如果没有提供access_token，需要提供微信app_id和app_secret来获取access_token

-o, --to_user string        接收人的open_id (必填)
    --msgtype string        消息类型 (必填)，text(文本消息)、image(图片消息)、link(图文链接)、miniprogrampage(小程序卡片)
```

消息内容

1. 文本消息 --msgtype text
    ```text
    "HelloWorld"
    ```

1. 图片消息 --msgtype image
    ```text
    "MEDIA_ID"
    ```

1. 图文链接 --msgtype link
    ```json
    {
      "title": "Happy Day",
      "description": "Is Really A Happy Day",
      "url": "URL",
      "thumb_url": "THUMB_URL"
    }
    ```

1. 小程序卡片 --msgtype miniprogrampage
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
pmsg.exe weixin miniprogram customer -i app_id -s app_secret -o open_id --msgtype text "HelloWorld"

使用命令别名
pmsg.exe wx mini kf -i app_id -s app_secret -o open_id --msgtype text "HelloWorld"

ok
```

发送小程序卡片

```shell
pmsg.exe weixin miniprogram customer -i app_id -s app_secret -o open_id --msgtype miniprogrampage "{\"title\":\"title\",\"pagepath\":\"pagepath\",\"thumb_media_id\":\"thumb_media_id\"}"

ok
```

linux

发送文本消息

```shell
$ pmsg weixin miniprogram customer -i app_id -s app_secret -o open_id --msgtype text 'HelloWorld'

使用命令别名
$ pmsg wx mini kf -i app_id -s app_secret -o open_id --msgtype text 'HelloWorld'

ok
```

发送小程序卡片

```shell
pmsg weixin miniprogram customer -i app_id -s app_secret -o open_id --msgtype miniprogrampage '{"title":"title","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}'

ok
```

[微信小程序客服消息开发文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html)