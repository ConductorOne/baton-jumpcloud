package connector

import (
	"fmt"
	"strconv"

	"github.com/conductorone/baton-sdk/pkg/pagination"
)

func unmarshalSkipToken(token *pagination.Token) (int32, *pagination.Bag, error) {
	b := &pagination.Bag{}
	err := b.Unmarshal(token.Token)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to unmarshal pagination bag during skip token processing: %w", err)
	}
	current := b.Current()
	skip := int32(0)
	if current != nil && current.Token != "" {
		skip64, err := strconv.ParseInt(current.Token, 10, 32)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to parse skip value as integer during pagination: %w", err)
		}
		skip = int32(skip64)
	}
	return skip, b, nil
}

func marshalSkipToken(newObjects int, lastSkip int32, b *pagination.Bag) (string, error) {
	if newObjects == 0 {
		return nextToken(b, "")
	}
	nextSkip := int64(newObjects) + int64(lastSkip)
	pageToken, err := nextToken(b, strconv.FormatInt(nextSkip, 10))
	if err != nil {
		return "", fmt.Errorf("failed to get next token during skip token marshaling: %w", err)
	}
	return pageToken, nil
}

func nextToken(b *pagination.Bag, v string) (string, error) {
	if b.Current() == nil {
		b.Push(pagination.PageState{Token: v})
		return b.Marshal()
	} else {
		return b.NextToken(v)
	}
}
