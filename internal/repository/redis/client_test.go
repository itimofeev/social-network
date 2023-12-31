package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	c, err := New(Config{
		RedisDSN: "redis://localhost:6379/0",
	})
	require.NoError(t, err)

	res := c.client.LPush(context.Background(), "foo2", "bar2")
	require.NoError(t, res.Err())
}
