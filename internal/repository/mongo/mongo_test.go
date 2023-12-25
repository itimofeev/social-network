//go:build integration

package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSendListMessages(t *testing.T) {
	ctx := context.Background()
	repo, err := New(ctx, Config{
		MongoDSN: "mongodb://admin:admin@localhost:27017",
	})
	require.NoError(t, err)

	user1 := uuid.MustParse("21913822-d04f-4823-9713-da015bb72848")
	user2 := uuid.MustParse("7f35ca5e-531b-4e8d-be04-7bf9b656ee6e")

	now := time.Now()
	require.NoError(t, repo.SendMessage(ctx, user1, user2, "hello", now))

	messages, err := repo.ListMessages(ctx, user1, user2, now.Add(-time.Second))
	require.NoError(t, err)

	fmt.Println(messages)
}
