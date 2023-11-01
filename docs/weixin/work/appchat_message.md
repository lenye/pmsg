### 发送企业微信群聊推送消息

命令参数说明

```text
$ pmsg workweixin appchat -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供企业微信 corp_id 和 corp_secret 获取 access_token

-c, --chat_id string        群聊id (必填) 
-m, --msg_type string       消息类型 (必填)，text(文本消息)、image(图片消息)、
                                           voice(语音消息)、video(视频消息)、file(文件消息)、
                                           textcard(文本卡片消息)、news(图文消息)、mpnews(图文消息)、
                                           markdown(markdown消息)

    --safe int              表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
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

1. 语音消息 --msg_type voice
    ```text
    MEDIA_ID
    ```

1. 视频消息 --msg_type video
    ```json
    {
      "media_id": "MEDIA_ID",
      "title": "TITLE",
      "description": "DESCRIPTION"
    }
    ```

1. 文件消息 --msg_type file
    ```text
    MEDIA_ID
    ```

1. 文本卡片消息 --msg_type textcard
    ```json
    {
      "title": "title",
      "description": "description",
      "url": "URL",
      "btntxt": "更多"
    }
    ```

1. 图文消息 --msg_type news
   ```json
   {
     "articles": [
       {
         "title": "中秋节礼品领取",
         "description": "今年中秋节公司有豪礼相送",
         "url": "URL",
         "picurl": "http://res.mail.qq.com/msg.png",
         "appid": "wx123123123123123",
         "pagepath": "pages/index?userid=zhangsan&orderid=123123123"
       }
     ]
   }
   ```

1. 图文消息 --msg_type mpnews
   ```json
   {
     "articles": [
       {
         "title": "Title",
         "thumb_media_id": "MEDIA_ID",
         "author": "Author",
         "content_source_url": "URL",
         "content": "Content",
         "digest": "Digest description"
       }
     ]
   }
   ```

1. markdown消息 --msg_type markdown
    ```text
    markdown
    ```

样例

linux

发送文本消息

```shell
$ pmsg workweixin appchat -i corp_id -s corp_secret -c chat_id -m text 'HelloWorld'

ok
```

官方开发文档 [企业微信群聊推送消息](https://developer.work.weixin.qq.com/document/path/90248)