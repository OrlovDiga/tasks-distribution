package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	client *openai.Client
}

func NewChatGPTClient(apiKey string) ChatGPTClient {
	return ChatGPTClient{
		client: openai.NewClient(apiKey),
	}
}

type TaskInfo struct {
	TaskName    string `json:"task_name"`
	TaskContent string `json:"task_content"`
	TaskType    string `json:"task_type"`
	Emoji       string `json:"emoji"`
}

func (c ChatGPTClient) GetTitle(taskDescription string) (TaskInfo, error) {
	promt := fmt.Sprintf(`В тройных кавычках ниже я предоставляю тебе описание задания.
Тебе необходимо проанализировать его и написать мне только ответ в виде json,
с полями "task_name", которое хранит краткое название задачи, "task_type", которое хранит тип работы: дом или работа,
а также "emoji", которое хранит подходящее эмодзи для этого задания и существует в Notion.
'''
%s
'''
`, taskDescription)

	message, err := c.SendSingleMessage(promt)
	if err != nil {
		return TaskInfo{}, err
	}

	var task TaskInfo
	err = json.Unmarshal([]byte(message), &task)
	if err != nil {
		return TaskInfo{}, fmt.Errorf("failed to unmasrhsla TaskINfo struct: %w", err)
	}

	return task, nil
}

func (c ChatGPTClient) SendSingleMessage(promt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: promt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to ChatCompletion: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
