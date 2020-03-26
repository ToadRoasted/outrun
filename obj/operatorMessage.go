package obj

import (
	"strconv"
	"time"

	"github.com/Mtbcooler/outrun/enums"
)

type OperatorMessage struct {
	ID         string      `json:"messageId"`
	Content    string      `json:"contents"`
	Item       MessageItem `json:"item"`
	ExpireTime int64       `json:"expireTime,omitempty"`
}

func DefaultOperatorMessage() OperatorMessage {
	id := "1"
	content := "A test gift."
	item := NewMessageItem(
		enums.ItemIDRing,
		100,
		0,
		0,
	)
	expireTime := time.Now().Unix() + 86400
	return OperatorMessage{
		id,
		content,
		item,
		expireTime,
	}
}

func NewOperatorMessage(id int64, content string, item MessageItem, expiresAfter int64) OperatorMessage {
	expireTime := time.Now().Unix() + expiresAfter
	return OperatorMessage{
		strconv.Itoa(int(id)),
		content,
		item,
		expireTime,
	}
}
