package parser

import (
	"github.com/margostino/just/common"
	"golang.org/x/net/html"
	"strings"
)

func Parse(content string) []html.Token {

	var tokens = make([]html.Token, 0)
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		data := common.NewString(token.Data).
			ReplaceAll("\n", "").
			ReplaceAll("\t", "").
			TrimSpace().
			Value()

		//if true || isValidTokenType(tokenType) && isValidData(token) {
		//	attrs := make([]*Attributes, 0)
		//	for _, attr := range currentToken.Attr {
		//		att := &Attributes{
		//			Key:   attr.Key,
		//			Value: attr.Val,
		//		}
		//		attrs = append(attrs, att)
		//	}
		//	token := &Token{
		//		Type:       tokenType,
		//		Data:       currentToken.Data,
		//		Attributes: attrs,
		//	}
		//}

		if !shouldFilter(data) {
			token.Data = data
			tokens = append(tokens, token)
		}

		if tokenType == html.ErrorToken {
			return tokens
		}
	}
}
