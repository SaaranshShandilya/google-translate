package main

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "cloud.google.com/go/translate/apiv3/translatepb"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("")) //credentials of your json file
	if err != nil {
    panic(err)
	}
	defer c.Close()

  req := &translatepb.TranslateTextRequest {
    Contents: []string{"apka naam kya h?" },
    Parent: "projects/round-mark-397813",
    SourceLanguageCode: "hi",
    TargetLanguageCode: "en",
  }

  resp, err := c.TranslateText(ctx, req)
  if err != nil {
    panic(err)
  }

  fmt.Println(resp)
}

