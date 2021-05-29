package internal

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestRedisClient_ReadString(t *testing.T) {
	addr := os.Getenv("REDIS_ADDRESS")
	pwd := os.Getenv("REDIS_PASSWORD")

	config := Config{
		RedisAddress:  addr,
		RedisPassword: pwd,
	}
	client, err := New(config)
	require.NoError(t, err)

	err = client.WriteString("tempKey", "tempValue", 10*time.Minute)
	require.NoError(t, err)

	value, err := client.ReadString("tempKey")
	require.NoError(t, err)
	require.Equal(t, "tempValue", value)
}
