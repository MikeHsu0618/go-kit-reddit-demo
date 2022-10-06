package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go-kit-reddit-demo/internal/pkg/config"
)

func TestNew(t *testing.T) {
	path := config.GetPath()
	conf, err := config.Load(path)
	require.NoError(t, err)

	userDB, err := NewUserDB(conf)
	require.NoError(t, err)
	require.NotNil(t, userDB)

	//postDB, err := NewPostDB(conf)
	//require.NoError(t, err)
	//require.NotNil(t, postDB)
}
