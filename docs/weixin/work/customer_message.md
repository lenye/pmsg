### 发送企业微信客服消息

命令参数说明

```text
$ pmsg workweixin customer -h

-a, --user_agent string     http user agent

-t, --access_token string   企业微信接口调用凭证
-i, --corp_id string        企业微信corp_id
-s, --corp_secret string    企业微信corp_secret

如果没有提供 access_token，需要提供企业微信 corp_id 和 corp_secret 获取 access_token

-c, --msg_id string         指定消息ID
-k, --open_kf_id string     指定发送消息的客服帐号ID (必填)
-o, --to_user string        指定接收消息的客户UserID (必填)
-m, --msg_type string       消息类型 (必填) text(文本消息)、image(图片消息)、
                                           voice(语音消息)、video(视频消息)、file(文件消息)、
                                           link(图文链接消息)、miniprogram(小程序消息)、
                                           msgmenu(菜单消息)、location(地理位置消息)
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
    ```text
    MEDIA_ID
    ```

1. 文件消息 --msg_type file
    ```text
    MEDIA_ID
    ```

1. 图文链接消息 --msg_type link
   ```json
   {
      "title": "企业如何增长？企业微信给出3个答案",
      "desc": "今年中秋节公司有豪礼相送",
      "url": "URL",
      "thumb_media_id": "MEDIA_ID"
   }
   ```

1. 小程序消息 --msg_type miniprogram
   ```json
   {
      "appid": "APPID",
      "title": "欢迎报名夏令营",
      "thumb_media_id": "MEDIA_ID",
      "pagepath": "PAGE_PATH"
   }
   ```

1. 菜单消息 --msg_type msgmenu
   ```json
   {
      "head_content": "您对本次服务是否满意呢? ",
      "list": [
         {
            "type": "click",
            "click": {
               "id": "101",
               "content": "满意"
            }
         },
         {
            "type": "click",
            "click": {
               "id": "102",
               "content": "不满意"
            }
         },
         {
            "type": "view",
            "view": {
               "url": "https://work.weixin.qq.com",
               "content": "点击跳转到自助查询页面"
            }
         },
         {
            "type": "miniprogram",
            "miniprogram": {
               "appid": "wx123123123123123",
               "pagepath": "pages/index.html?userid=zhangsan&amp;orderid=123123123",
               "content": "点击打开小程序查询更多"
            }
         },
         {
            "type": "text",
            "text": {
               "content": "纯文本，支持\n换行"
            }
         }
      ],
      "tail_content": "欢迎再次光临"
   }
   ```

1. 地理位置消息 --msg_type location
   ```json
   {
      "name": "测试小区",
      "address": "实例小区，不真实存在，经纬度无意义",
      "latitude": 0,
      "longitude": 0
   }
   ```

样例

linux

发送文本消息

```shell
$ pmsg workweixin customer -i corp_id -s corp_secret -o user_id -k kf_id -m text 'HelloWorld'

ok; msgid: "msgid"
```

官方开发文档 [企业微信客服消息](https://developer.work.weixin.qq.com/document/path/94677)