package staking

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrMinStake              = sdkerrors.Register(ModuleName, 99, "not enough staking coins")
)