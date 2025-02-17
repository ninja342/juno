package core2p2p

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/p2p/starknet/spec"
	"github.com/NethermindEth/juno/utils"
)

func AdaptStateDiff(addr, classHash, nonce *felt.Felt, diff []core.StorageDiff) *spec.StateDiff_ContractDiff {
	return &spec.StateDiff_ContractDiff{
		Address:   AdaptAddress(addr),
		Nonce:     AdaptFelt(nonce),
		ClassHash: AdaptFelt(classHash),
		Values:    AdaptStorageDiff(diff),
	}
}

func AdaptStorageDiff(diff []core.StorageDiff) []*spec.ContractStoredValue {
	return utils.Map(diff, func(item core.StorageDiff) *spec.ContractStoredValue {
		return &spec.ContractStoredValue{
			Key:   AdaptFelt(item.Key),
			Value: AdaptFelt(item.Value),
		}
	})
}
