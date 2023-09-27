package ojooracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"

	"github.com/comdex-official/comdex/x/ojooracle/types"
)

func (im IBCModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData types.OracleResponsePacketData
	fetchPriceMsg := im.keeper.GetFetchPriceMsg(ctx)
	if modulePacket.DestinationChannel != fetchPriceMsg.SourceChannel {
		return channeltypes.Acknowledgement{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
			"Module packet destination channel and source channel mismatch")
	}

	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}

	switch modulePacketData.GetClientID() {
	case types.FetchPriceClientIDKey:
		var priceResult types.Price
		if err := types.ModuleCdc.UnmarshalJSON(modulePacketData.Result, &priceResult); err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				"cannot unmarshall price result packet")
		}
		ctx.Logger().Info("Received oracle packet this is here", "packet", modulePacketData.String())

		im.keeper.SetFetchPriceResult(ctx, types.OracleRequestID(modulePacketData.RequestID), priceResult)
	// TODO: FetchPrice market data reception logic //nolint:godox

	default:
		err := sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			"market received packet not found: %s", modulePacketData.GetClientID())
		ack = channeltypes.NewErrorAcknowledgement(err)
		return ack, err
	}
	//TODO: fix this
	//ack = channeltypes.NewResultAcknowledgement(
	//	types.ModuleCdc.MustMarshalJSON(
	//	//types.NewOracleRequestPacketAcknowledgement(modulePacketData.RequestID),
	//	),
	//)
	return ack, nil
}

func (im IBCModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck types.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		var data types.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}
		requestID := types.OracleRequestID(oracleAck.RequestID)

		switch data.GetClientID() {
		case types.FetchPriceClientIDKey:
			var fetchPriceData types.FetchPriceCallData
			if err = types.ModuleCdc.UnmarshalJSON(data.GetCalldata(), &fetchPriceData); err != nil {
				return nil, sdkerrors.Wrap(err,
					"cannot decode the fetchPrice market acknowledgment packet")
			}

			//TODO: fix this
			im.keeper.SetLastFetchPriceID(ctx, requestID)
			return &sdk.Result{}, nil

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"market acknowledgment packet not found: %s", data.GetClientID())
		}
	}
	return nil, nil
}
