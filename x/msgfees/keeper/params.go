package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/provenance-io/provenance/x/msgfees/types"
)

// GetParams returns the total set of distribution parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.Params{
		FloorGasPrice: uint32(int32(k.GetFloorGasPrice(ctx).Amount.Uint64())),
	}
}

// SetParams sets the account parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
