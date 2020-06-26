package udp_test

import (
	"github.com/aleitner/rollback/udp"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMmap(t *testing.T) {

	server, err := udp.NewServer("localhost:8080")
	require.NoError(t, err)
	defer func() {
		err := server.Close()
		require.NoError(t, err)
	}()

	client, err := udp.NewClient("localhost:8080")
	require.NoError(t, err)
	defer func() {
		err := client.Close()
		require.NoError(t, err)
	}()

	writtenBytes := []byte("Hello!")
	readBytes := make([]byte, 1024)

	go func() {
		for {
			n, err := server.Read(readBytes)
			require.NoError(t, err)
			require.Equal(t, len(writtenBytes), n)
			require.Equal(t, writtenBytes, readBytes)
		}
	}()

	n, err := client.Write(writtenBytes)
	require.NoError(t, err)
	require.Equal(t, len(writtenBytes), n)
}