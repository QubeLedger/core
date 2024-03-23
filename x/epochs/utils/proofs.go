package utils

import (
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	ibcKeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	ibctm "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

/* #nosec */
func ValidateProofOps(ctx sdk.Context, ibcKeeper *ibcKeeper.Keeper, connectionID string, chainID string, height int64, module string, key []byte, data []byte, proofOps *crypto.ProofOps) error {
	if proofOps == nil {
		return fmt.Errorf("unable to validate proof. No proof submitted")
	}

	connection, found := ibcKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return fmt.Errorf("connection %s not found", connectionID)
	}

	csHeight := clienttypes.NewHeight(clienttypes.ParseChainID(chainID), uint64(height)+1)
	consensusState, found := ibcKeeper.ClientKeeper.GetClientConsensusState(ctx, connection.ClientId, csHeight)

	if !found {
		return fmt.Errorf("unable to fetch consensus state")
	}

	tmConsState, ok := consensusState.(*ibctm.ConsensusState)
	if !ok {
		return fmt.Errorf("error unmarshaling consensus state")
	}

	clientState, found := ibcKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return fmt.Errorf("unable to fetch client state")
	}

	path := commitmenttypes.NewMerklePath([]string{module, url.PathEscape(string(key))}...)

	merkleProof, err := commitmenttypes.ConvertProofs(proofOps)
	if err != nil {
		return fmt.Errorf("error converting proofs")
	}

	tmClientState, ok := clientState.(*ibctm.ClientState)
	if !ok {
		return fmt.Errorf("error unmarshaling client state")
	}

	if len(data) != 0 {
		err = merkleProof.VerifyMembership(tmClientState.ProofSpecs, tmConsState.GetRoot(), path, data)
		err = sdkerrors.Wrap(err, "unable to verify inclusion proof")
	} else {
		// if we got a nil response, verify non inclusion proof.
		err = merkleProof.VerifyNonMembership(tmClientState.ProofSpecs, tmConsState.GetRoot(), path)
		err = sdkerrors.Wrap(err, "unable to verify non-inclusion proof")
	}

	return err
}
