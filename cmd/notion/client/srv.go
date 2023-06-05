package client

import (
	"context"
	"tasks-distribution/cmd/chatgpt/client"

	"github.com/jomei/notionapi"
)

const (
	TaskColumnName = "Task name"
)

type NotionTasksClient struct {
	SecretKey string
	DBId      string
	client    *notionapi.Client
}

func NewNotionTasksClient(secretKey string, dbId string) *NotionTasksClient {
	return &NotionTasksClient{
		DBId:   dbId,
		client: notionapi.NewClient(notionapi.Token(secretKey)),
	}
}

func (client *NotionTasksClient) AddNewTask(task client.TaskInfo) (*notionapi.Page, error) {
	curEmoji := notionapi.Emoji(task.Emoji)

	createdPage, err := client.client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       "",
			PageID:     "",
			DatabaseID: notionapi.DatabaseID(client.DBId),
			BlockID:    "",
			Workspace:  false,
		},
		Properties: notionapi.Properties{
			TaskColumnName: &notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						PlainText: task.TaskName,
						Text: &notionapi.Text{
							Content: task.TaskName,
						},
					},
				},
			},
		},
		Children: []notionapi.Block{
			&notionapi.ParagraphBlock{
				BasicBlock: notionapi.BasicBlock{
					Object: "block",
					Type:   "paragraph",
				},
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{
						{
							Type: "",
							Text: &notionapi.Text{
								Content: task.TaskContent,
							},
							PlainText: task.TaskContent,
						},
					},
					Color: "default",
				},
			},
		},
		Icon: &notionapi.Icon{
			Type:  "emoji",
			Emoji: &curEmoji,
		},
		Cover: nil,
	})
	if err != nil {
		return nil, err
	}

	return createdPage, nil
}
