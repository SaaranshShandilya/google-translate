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
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("round-mark-397813-9b79a82e9b9d.json"))
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

