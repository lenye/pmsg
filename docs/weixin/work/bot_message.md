### 推送企业微信群机器人消息

命令参数说明

```text
$ pmsg weixin work bot -h

-a, --user_agent string     http user agent

-k, --key string            企业微信群机器人key (必填)
-m, --msg_type string       消息类型 (必填)，text(文本消息)、markdown(markdown消息)、
                                           image(图片消息)、news(图文消息)、file(文件消息)
-o, --to_user string        文本消息时，提醒群中的指定成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人。
                            如果开发者获取不到userid，可以使用to_mobile
-b, --to_mobile string      文本消息时，提醒手机号对应的群成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人

args                        参数：消息内容
```

消息内容

1. 文本消息 --msg_type text
    ```text
    "HelloWorld"
    ```

1. markdown消息 --msg_type markdown
    ```text
    "markdown"
    ```

1. 图片消息 --msg_type image
    ```text
    文件名称含路径
    ```

1. 图文消息 --msg_type news
   ```json
   {
      "articles": [
         {
            "title": "中秋节礼品领取",
            "description": "今年中秋节公司有豪礼相送",
            "url": "www.qq.com",
            "picurl": "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
         }
      ]
   }
   ```

1. 文件消息 --msg_type file
    ```text
    "MEDIA_ID"
    ```

样例

linux

```shell
$ pmsg weixin work bot -k key -m text 'HelloWorld'

ok

$ pmsg weixin work bot -k key -m image /img/app.png

ok
```

[推送企业微信群机器人消息开发文档](https://developer.work.weixin.qq.com/document/path/91770)