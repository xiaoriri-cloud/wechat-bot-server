package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"log"
	"net/http"
	"wechat-bot-server/models"
	"wechat-bot-server/pkg/response"
)

func GetConfig(c *gin.Context) {
	configName := c.Query("name")
	var result = response.Result{
		Code: 200,
		Msg:  "ok",
		Data: "",
	}
	if configName != "" {
		config, _ := models.GetConfig(configName)
		if config.Name == "" {
			result.Data = nil
		} else {
			result.Data = config
		}
	} else {
		result.Code = 10001
		result.Msg = "参数校验错误，name不能为空"
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(200, result)
}

type setTokenRequest struct {
	Name   string `form:"name" binding:"required"`
	Config string `form:"config" binding:"required"`
}

func SetConfig(c *gin.Context) {
	var params setTokenRequest
	err := c.Bind(&params)
	var result = response.Result{
		Code: 200,
		Msg:  "ok",
		Data: params,
	}
	if err != nil {
		result.Code = 10001
		result.Msg = "参数校验错误"
		result.Data = err.Error()
		c.JSON(http.StatusOK, result)
		return
	}
	var config models.Config
	config.Name = params.Name
	config.Config = params.Config
	err = models.SetConfig(config)

	if err != nil {
		result.Code = 10001
		result.Msg = "保存失败"
		result.Data = err.Error()
	}
	c.JSON(http.StatusOK, result)
}

type TgResponse struct {
	Ok     bool `json:"ok"`
	Result []struct {
		MyChatMember struct {
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date int `json:"date"`
			From struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			NewChatMember struct {
				CanBeEdited         bool   `json:"can_be_edited"`
				CanChangeInfo       bool   `json:"can_change_info"`
				CanDeleteMessages   bool   `json:"can_delete_messages"`
				CanEditMessages     bool   `json:"can_edit_messages"`
				CanInviteUsers      bool   `json:"can_invite_users"`
				CanManageChat       bool   `json:"can_manage_chat"`
				CanManageVoiceChats bool   `json:"can_manage_voice_chats"`
				CanPostMessages     bool   `json:"can_post_messages"`
				CanPromoteMembers   bool   `json:"can_promote_members"`
				CanRestrictMembers  bool   `json:"can_restrict_members"`
				IsAnonymous         bool   `json:"is_anonymous"`
				Status              string `json:"status"`
				User                struct {
					FirstName string `json:"first_name"`
					ID        int64  `json:"id"`
					IsBot     bool   `json:"is_bot"`
					Username  string `json:"username"`
				} `json:"user"`
			} `json:"new_chat_member"`
			OldChatMember struct {
				Status string `json:"status"`
				User   struct {
					FirstName string `json:"first_name"`
					ID        int64  `json:"id"`
					IsBot     bool   `json:"is_bot"`
					Username  string `json:"username"`
				} `json:"user"`
			} `json:"old_chat_member"`
		} `json:"my_chat_member,omitempty"`
		UpdateID    int `json:"update_id"`
		ChannelPost struct {
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date         int    `json:"date"`
			MessageID    int    `json:"message_id"`
			NewChatTitle string `json:"new_chat_title"`
			SenderChat   struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"sender_chat"`
		} `json:"channel_post,omitempty"`
		Message struct {
			Chat struct {
				AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
				ID                          int    `json:"id"`
				Title                       string `json:"title"`
				Type                        string `json:"type"`
			} `json:"chat"`
			Date int `json:"date"`
			From struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			GroupChatCreated bool `json:"group_chat_created"`
			MessageID        int  `json:"message_id"`
		} `json:"message,omitempty"`
	} `json:"result"`
}

func GetUpdates() TgResponse {
	url := "https://api.telegram.org/bot" + getToken() + "/getUpdates"
	log.Println("getUpdates：" + url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data TgResponse
	json.Unmarshal(body, &data)
	return data
}

func getToken() string {
	config, _ := models.GetConfig("bot-q-token")
	return config.Config
}

type SendMsg struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendMessage(msgContent SendMsg) {
	url := "https://api.telegram.org/bot" + getToken() + "/sendMessage"
	params, _ := json.Marshal(msgContent)
	log.Println(msgContent)
	resp, _ := http.Post(url, binding.MIMEJSON, bytes.NewBuffer(params))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
