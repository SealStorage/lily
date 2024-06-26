// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"bytes"
	"fmt"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	market13 "github.com/filecoin-project/go-state-types/builtin/v13/market"
	adt13 "github.com/filecoin-project/go-state-types/builtin/v13/util/adt"
	markettypes "github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/go-state-types/manifest"

	lotusactors "github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
)

var _ State = (*state13)(nil)

func load13(store adt.Store, root cid.Cid) (State, error) {
	out := state13{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make13(store adt.Store) (State, error) {
	out := state13{store: store}

	s, err := market13.ConstructState(store)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state13 struct {
	market13.State
	store adt.Store
}

func (s *state13) StatesChanged(otherState State) (bool, error) {
	otherState13, ok := otherState.(*state13)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.States.Equals(otherState13.State.States), nil
}

func (s *state13) States() (DealStates, error) {
	stateArray, err := adt13.AsArray(s.store, s.State.States, market13.StatesAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealStates13{stateArray}, nil
}

func (s *state13) ProposalsChanged(otherState State) (bool, error) {
	otherState13, ok := otherState.(*state13)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.Proposals.Equals(otherState13.State.Proposals), nil
}

func (s *state13) Proposals() (DealProposals, error) {
	proposalArray, err := adt13.AsArray(s.store, s.State.Proposals, market13.ProposalsAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealProposals13{proposalArray}, nil
}

type dealStates13 struct {
	adt.Array
}

func (s *dealStates13) Get(dealID abi.DealID) (DealState, bool, error) {
	var deal13 market13.DealState
	found, err := s.Array.Get(uint64(dealID), &deal13)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	deal := fromV13DealState(deal13)
	return deal, true, nil
}

func (s *dealStates13) ForEach(cb func(dealID abi.DealID, ds DealState) error) error {
	var ds13 market13.DealState
	return s.Array.ForEach(&ds13, func(idx int64) error {
		return cb(abi.DealID(idx), fromV13DealState(ds13))
	})
}

func (s *dealStates13) decode(val *cbg.Deferred) (DealState, error) {
	var ds13 market13.DealState
	if err := ds13.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}
	ds := fromV13DealState(ds13)
	return ds, nil
}

func (s *dealStates13) array() adt.Array {
	return s.Array
}

func fromV13DealState(v13 market13.DealState) DealState {
	return dealStateV13{v13}
}

type dealStateV13 struct {
	ds13 market13.DealState
}

func (d dealStateV13) SectorStartEpoch() abi.ChainEpoch {
	return d.ds13.SectorStartEpoch
}

func (d dealStateV13) SectorNumber() abi.SectorNumber {

	return d.ds13.SectorNumber

}

func (d dealStateV13) LastUpdatedEpoch() abi.ChainEpoch {
	return d.ds13.LastUpdatedEpoch
}

func (d dealStateV13) SlashEpoch() abi.ChainEpoch {
	return d.ds13.SlashEpoch
}

func (d dealStateV13) Equals(other DealState) bool {
	if ov13, ok := other.(dealStateV13); ok {
		return d.ds13 == ov13.ds13
	}

	if d.SectorStartEpoch() != other.SectorStartEpoch() {
		return false
	}
	if d.LastUpdatedEpoch() != other.LastUpdatedEpoch() {
		return false
	}
	if d.SlashEpoch() != other.SlashEpoch() {
		return false
	}

	return true
}

var _ DealState = (*dealStateV13)(nil)

type dealProposals13 struct {
	adt.Array
}

func (s *dealProposals13) Get(dealID abi.DealID) (*DealProposal, bool, error) {
	var proposal13 market13.DealProposal
	found, err := s.Array.Get(uint64(dealID), &proposal13)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	proposal, err := fromV13DealProposal(proposal13)
	if err != nil {
		return nil, true, xerrors.Errorf("decoding proposal: %w", err)
	}

	return &proposal, true, nil
}

func (s *dealProposals13) ForEach(cb func(dealID abi.DealID, dp DealProposal) error) error {
	var dp13 market13.DealProposal
	return s.Array.ForEach(&dp13, func(idx int64) error {
		dp, err := fromV13DealProposal(dp13)
		if err != nil {
			return xerrors.Errorf("decoding proposal: %w", err)
		}

		return cb(abi.DealID(idx), dp)
	})
}

func (s *dealProposals13) decode(val *cbg.Deferred) (*DealProposal, error) {
	var dp13 market13.DealProposal
	if err := dp13.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}

	dp, err := fromV13DealProposal(dp13)
	if err != nil {
		return nil, err
	}

	return &dp, nil
}

func (s *dealProposals13) array() adt.Array {
	return s.Array
}

func fromV13DealProposal(v13 market13.DealProposal) (DealProposal, error) {

	label, err := fromV13Label(v13.Label)

	if err != nil {
		return DealProposal{}, xerrors.Errorf("error setting deal label: %w", err)
	}

	return DealProposal{
		PieceCID:     v13.PieceCID,
		PieceSize:    v13.PieceSize,
		VerifiedDeal: v13.VerifiedDeal,
		Client:       v13.Client,
		Provider:     v13.Provider,

		Label: label,

		StartEpoch:           v13.StartEpoch,
		EndEpoch:             v13.EndEpoch,
		StoragePricePerEpoch: v13.StoragePricePerEpoch,

		ProviderCollateral: v13.ProviderCollateral,
		ClientCollateral:   v13.ClientCollateral,
	}, nil
}

func (s *state13) DealProposalsAmtBitwidth() int {
	return market13.ProposalsAmtBitwidth
}

func (s *state13) DealStatesAmtBitwidth() int {
	return market13.StatesAmtBitwidth
}

func (s *state13) ActorKey() string {
	return manifest.MarketKey
}

func (s *state13) ActorVersion() actorstypes.Version {
	return actorstypes.Version13
}

func (s *state13) Code() cid.Cid {
	code, ok := lotusactors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func fromV13Label(v13 market13.DealLabel) (DealLabel, error) {
	if v13.IsString() {
		str, err := v13.ToString()
		if err != nil {
			return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert string label to string: %w", err)
		}
		return markettypes.NewLabelFromString(str)
	}

	bs, err := v13.ToBytes()
	if err != nil {
		return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert bytes label to bytes: %w", err)
	}
	return markettypes.NewLabelFromBytes(bs)
}

func (s *state13) GetProviderSectors() (map[abi.SectorID][]abi.DealID, error) {

	sectorDeals, err := adt13.AsMap(s.store, s.State.ProviderSectors, market13.ProviderSectorsHamtBitwidth)
	if err != nil {
		return nil, err
	}
	var sectorMapRoot cbg.CborCid
	providerSectors := make(map[abi.SectorID][]abi.DealID)
	err = sectorDeals.ForEach(&sectorMapRoot, func(providerID string) error {
		provider, err := abi.ParseUIntKey(providerID)
		if err != nil {
			return nil
		}

		sectorMap, err := adt13.AsMap(s.store, cid.Cid(sectorMapRoot), market13.ProviderSectorsHamtBitwidth)
		if err != nil {
			return err
		}

		var dealIDs market13.SectorDealIDs
		err = sectorMap.ForEach(&dealIDs, func(sectorID string) error {
			sectorNumber, err := abi.ParseUIntKey(sectorID)
			if err != nil {
				return err
			}

			dealIDsCopy := make([]abi.DealID, len(dealIDs))
			copy(dealIDsCopy, dealIDs)

			providerSectors[abi.SectorID{Miner: abi.ActorID(provider), Number: abi.SectorNumber(sectorNumber)}] = dealIDsCopy
			return nil
		})
		return err
	})
	return providerSectors, err

}

func (s *state13) GetProviderSectorsByDealID(dealIDMap map[abi.DealID]bool, sectorIDMap map[abi.SectorNumber]bool) (map[abi.DealID]abi.SectorID, error) {

	sectorDeals, err := adt13.AsMap(s.store, s.State.ProviderSectors, market13.ProviderSectorsHamtBitwidth)
	if err != nil {
		return nil, err
	}
	var sectorMapRoot cbg.CborCid
	dealIDSectorMap := make(map[abi.DealID]abi.SectorID)
	err = sectorDeals.ForEach(&sectorMapRoot, func(providerID string) error {
		provider, err := abi.ParseUIntKey(providerID)
		if err != nil {
			return nil
		}

		sectorMap, err := adt13.AsMap(s.store, cid.Cid(sectorMapRoot), market13.ProviderSectorsHamtBitwidth)
		if err != nil {
			return err
		}

		var dealIDs market13.SectorDealIDs
		err = sectorMap.ForEach(&dealIDs, func(sectorID string) error {
			sectorNumber, err := abi.ParseUIntKey(sectorID)
			if err != nil {
				return err
			}

			if _, found := sectorIDMap[abi.SectorNumber(sectorNumber)]; !found {
				return nil
			}

			dealIDsCopy := make([]abi.DealID, len(dealIDs))
			copy(dealIDsCopy, dealIDs)

			for _, dealID := range dealIDsCopy {
				_, found := dealIDMap[dealID]
				if found {
					dealIDSectorMap[dealID] = abi.SectorID{Miner: abi.ActorID(provider), Number: abi.SectorNumber(sectorNumber)}
				}
			}

			return nil
		})
		return err
	})
	return dealIDSectorMap, err

}
