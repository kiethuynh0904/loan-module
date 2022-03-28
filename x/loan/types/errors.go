package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrDeadline       = sdkerrors.Register(ModuleName, 2, "deadline")
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 1, "wrong loan state")
)
