package main

import (
	"wechat-bot-server/models"
	"wechat-bot-server/pkg/setting"
	"wechat-bot-server/routers"
	"wechat-bot-server/telegram"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	msg := telegram.SendMsg{
		ChatId:    "-604087776",
		Text:      "<b>bold</b>, <strong>bold</strong>\n<i>italic</i>, <em>italic</em>\n<u>underline</u>, <ins>underline</ins>\n<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>\n<span class=\"tg-spoiler\">spoiler</span>, <tg-spoiler>spoiler</tg-spoiler>\n<b>bold <i>italic bold <s>italic bold strikethrough <span class=\"tg-spoiler\">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>\n<a href=\"http://www.example.com/\">inline URL</a>\n<a href=\"tg://user?id=123456789\">inline mention of a user</a>\n<code>inline fixed-width code</code>\n<pre>pre-formatted fixed-width code block</pre>\n<pre><code class=\"language-python\">pre-formatted fixed-width code block written in the Python programming language</code></pre>",
		ParseMode: "HTML",
	}
	telegram.SendMessage(msg)
	r := routers.InitRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
