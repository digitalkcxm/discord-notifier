package DiscordNotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var errorColor = 12389167

type notifier struct {
	Webhook     string
	ServiceName string
}

func Setup(ServiceName, Webhook string) *notifier {
	return &notifier{
		Webhook:     Webhook,
		ServiceName: ServiceName,
	}
}

func (n *notifier) NotifyRabbitDisconnection(err error) error {

	var message string = fmt.Sprintf("Conexão com o Rabbit perdida. %s", err)

	data := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       n.ServiceName,
				"description": message,
				"color":       errorColor,
			},
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return notify(n.Webhook, payload)
}

func notify(url string, payload []byte) (err error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
