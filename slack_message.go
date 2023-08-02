package main

import (
	"bytes"
	"encoding/json"
	"io"
)

type SlackMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Icon     string `json:"icon_emoji"`
}

func (m *SlackMessage) Serialized() io.Reader {
	out, _ := json.Marshal(m)
	return bytes.NewBuffer(out)
}
