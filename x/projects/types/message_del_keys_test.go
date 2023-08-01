package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/lavanet/lava/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDelKeys_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDelKeys
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDelKeys{
				Creator: "invalid_address",
			},
			err: legacyerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDelKeys{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
