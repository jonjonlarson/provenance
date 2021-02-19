package v040

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/auth/types"

	v039marker "github.com/provenance-io/provenance/x/marker/legacy/v039"
	v040marker "github.com/provenance-io/provenance/x/marker/types"
)

// Migrate accepts exported x/marker genesis state from v0.39 and migrates it
// to v0.40 x/marker genesis state. The migration includes:
//
// - Convert addresses from bytes to bech32 strings.
// - Re-encode in v0.40 GenesisState.
func Migrate(oldGenState v039marker.GenesisState) *v040marker.GenesisState {
	var markerAccounts = make([]v040marker.MarkerAccount, 0, len(oldGenState.Markers))
	for _, mark := range oldGenState.Markers {
		markerType := v040marker.MarkerType_value["MARKER_TYPE_"+strings.ToUpper(mark.MarkerType)]
		if markerType == int32(v040marker.MarkerType_Unknown) {
			panic(fmt.Sprintf("unknown marker type %s", mark.MarkerType))
		}
		markerAccounts = append(markerAccounts, v040marker.MarkerAccount{
			BaseAccount: &types.BaseAccount{
				Address:       mark.Address.String(),
				AccountNumber: mark.AccountNumber,
				Sequence:      mark.Sequence,
			},
			Manager: mark.Manager.String(),
			Status:  v040marker.MustGetMarkerStatus(mark.GetStatus()),
			Denom:   mark.Denom,
			Supply:  mark.GetSupply().Amount,
			// TODO PORT ACCESS LIST
			// v039 only supported COIN type
			MarkerType: v040marker.MarkerType_Coin,
		})
	}
	return &v040marker.GenesisState{
		Params: v040marker.Params{
			EnableGovernance: v040marker.DefaultEnableGovernance,
			MaxTotalSupply:   v040marker.DefaultMaxTotalSupply,
		},
		Markers: markerAccounts,
	}
}