### 推送飞书自定义机器人消息

命令参数说明

```text
$ pmsg feishu bot -h

-a, --user_agent string     http user agent

-t, --access_token string   飞书自定义机器人 token (必填), 
                             token 为 webhook 地址中 xxx 部分 https://open.feishu.cn/open-apis/bot/v2/hook/xxx
-s, --secret string         签名密钥
-m, --msg_type string       消息类型 (必填)，text(文本消息)、post(富文本)、image(图片)、
                                           share_chat(分享群名片)、interactive(消息卡片)
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
   img_ecffc3b9-8f14-400f-a014-05eca1a4310g
   ```

1. 分享群名片消息 --msg_type share_chat
   ```text
   oc_f5b1a7eb27ae2c7b6adc2a74faf339ff
   ```

1. 富文本消息 --msg_type post
   ```json
   {
      "zh_cn": {
         "title": "项目更新通知",
         "content": [
            [
               {
                  "tag": "text",
                  "text": "项目有更新: "
               },
               {
                  "tag": "a",
                  "text": "请查看",
                  "href": "http://www.example.com/"
               },
               {
                  "tag": "at",
                  "user_id": "ou_18eac8********17ad4f02e8bbbb"
               }
            ]
         ]
      }
   }
   ```

1. 消息卡片 --msg_type interactive
   ```json
   {
      "elements": [
         {
            "tag": "div",
            "text": {
               "content": "**西湖**，位于浙江省杭州市西湖区龙井路1号，杭州市区西部，景区总面积49平方千米，汇水面积为21.22平方千米，湖面面积为6.38平方千米。",
               "tag": "lark_md"
            }
         },
         {
            "actions": [
               {
                  "tag": "button",
                  "text": {
                     "content": "更多景点介绍 :玫瑰:",
                     "tag": "lark_md"
                  },
                  "url": "https://www.example.com",
                  "type": "default",
                  "value": {}
               }
            ],
            "tag": "action"
         }
      ],
      "header": {
         "title": {
            "content": "今日旅游推荐",
            "tag": "plain_text"
         }
      }
   }
   ```

样例

linux

```shell
$ pmsg feishu bot -t access_token -m text 'HelloWorld'

ok
```

官方开发文档 [推送飞书自定义机器人消息](https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot)