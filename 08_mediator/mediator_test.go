package mediator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMediator(t *testing.T) {
	instance := GetInstance()
	instance.CD.ReadData()
	require.Equal(t, "music,image", mediator.CD.Data)
	require.Equal(t, "music", mediator.CPU.Sound)
	require.Equal(t, "image", mediator.CPU.Video)
	require.Equal(t, "music", mediator.SoundCard.data)
	require.Equal(t, "image", mediator.VideoCard.data)
}
