package devearn

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/common"
	"sidechain/x/devearn/keeper"
	"sidechain/x/devearn/types"
	"strconv"
)

func NewDevEarnProposalHanodler(k *keeper.Keeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		switch c := content.(type) {
		case *types.RegisterDevEarnInfoProposal:
			return handleRegisterDevEarnProposal(ctx, k, c)
		case *types.CancelDevEarnInfoProposal:
			return handleCancelDevEarnProposal(ctx, k, c)
		default:
			return errorsmod.Wrapf(
				errortypes.ErrUnknownRequest,
				"unrecognized %s proposal content type: %T", types.ModuleName, c,
			)
		}
	}
}

func handleRegisterDevEarnProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterDevEarnInfoProposal) error {
	in, err := k.RegisterDevEarn(ctx, common.HexToAddress(p.Contract), p.Epochs)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterDevEarn,
			sdk.NewAttribute(types.AttributeKeyContract, in.Contract),
			sdk.NewAttribute(
				types.AttributeKeyEpochs,
				strconv.FormatUint(uint64(in.Epochs), 10),
			),
		),
	)
	return nil
}

func handleCancelDevEarnProposal(ctx sdk.Context, k *keeper.Keeper, p *types.CancelDevEarnInfoProposal) error {
	err := k.CancelDevEarn(ctx, common.HexToAddress(p.Contract))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCancelDevEarn,
			sdk.NewAttribute(types.AttributeKeyContract, p.Contract),
		),
	)
	return nil
}
