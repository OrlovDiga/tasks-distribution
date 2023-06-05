package main

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func main() {
	client := notionapi.NewClient("token")

	createdPage, err := client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       "",
			PageID:     "",
			DatabaseID: "your_db_id",
			BlockID:    "",
			Workspace:  false,
		},
		Properties: notionapi.Properties{
			"Task name": &notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						PlainText: "Новая задача Plain Text",
						Text: &notionapi.Text{
							Content: "Новая задача Content",
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
								Content: "помыть кукуху",
							},
							PlainText: "помыть кукуху",
						},
					},
					Color: "default",
				},
			},
		},
		Icon:  nil,
		Cover: nil,
	})
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(createdPage)

	page, err := client.Page.Get(context.Background(), "your_page_id")
	if err != nil {
		fmt.Println("err", err)
	}

	for k, v := range page.Properties {
		fmt.Println(k, " ", v.GetType(), " ", v)
	}

	block, err := client.Block.Get(context.Background(), "your_block_id")
	if err != nil {
		fmt.Println("err", err)
	}

	paragraphblock := block.(*notionapi.ParagraphBlock)

	fmt.Println(paragraphblock.Paragraph.RichText)
	for _, v := range paragraphblock.Paragraph.RichText {
		fmt.Println("type", v.Type)
		fmt.Println("text", v.Text)
		fmt.Println("PlainText", v.PlainText)
		fmt.Println("Mention", v.Mention)
		fmt.Println("Equation", v.Equation)
	}
	//fmt.Println(page)
	//result, err := client.Page.Create(context.Background(),
	//	&notionapi.PageCreateRequest{
	//		Parent: notionapi.Parent{
	//			DatabaseID: notionapi.DatabaseID(database.ID),
	//		},
	//		Properties: map[string]notionapi.Property{"title": notionapi.TitleProperty{
	//			ID: "title",
	//			Title: []notionapi.RichText{
	//				{
	//					Text: &notionapi.Text{Content: "Новая задача"},
	//				},
	//			},
	//		}},
	//	},
	//)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(result)

}
