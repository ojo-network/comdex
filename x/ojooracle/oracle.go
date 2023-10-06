package ojooracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	protobuftypes "github.com/gogo/protobuf/types"

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
		//TODO: remove
		ctx.Logger().Info("Received oracle packet ", "packet", modulePacketData.String())

		var priceResult types.OracleRequestResult
		if err := types.ModuleCdc.Unmarshal(modulePacketData.GetResult(), &priceResult); err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				"cannot unmarshall price result packet")
		}

		//TODO: remove
		ctx.Logger().Info("price result oracle packet", "price result", priceResult.String())

		im.keeper.SetFetchPriceResult(ctx, types.OracleRequestID(modulePacketData.RequestID), priceResult)
	// TODO: FetchPrice market data reception logic //nolint:godox

	default:
		err := sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			"market received packet not found: %s", modulePacketData.GetClientID())
		ack = channeltypes.NewErrorAcknowledgement(err)
		return ack, err
	}

	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			&types.OracleRequestPacketAcknowledgement{RequestID: modulePacketData.RequestID},
		),
	)
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

		//TODO :remove after testing
		//ctx.Logger().Error("processing acknow response packet", "packet", oracleAck.String())

		var data types.OracleRequestPacketData
		if err = types.ModuleCdc.Unmarshal(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}

		//TODO :remove after testing
		//ctx.Logger().Error("processing acknow request packet data", "packet", data.String())
		//

		requestID := types.OracleRequestID(oracleAck.RequestID)
		ctx.Logger().Info("Received acknow oracle packet this is here", "packet", data.String(), "requestID", requestID)

		var requestType protobuftypes.Int32Value
		err = types.ModuleCdc.UnmarshalLengthPrefixed(modulePacket.GetData(), &requestType)
		if err != nil {
			return nil, err
		}

		switch data.GetClientID() {
		case types.FetchPriceClientIDKey:
			var fetchPriceData types.RequestPrice
			if err = types.ModuleCdc.Unmarshal(data.GetCalldata(), &fetchPriceData); err != nil {
				return nil, sdkerrors.Wrap(err,
					"cannot decode the fetchPrice market acknowledgment packet")
			}

			//TODO :remove after testing
			ctx.Logger().Info("Received acknow oracle packet", "fetch price data", fetchPriceData.String())

			im.keeper.SetLastFetchPriceID(ctx, requestID, requestType.Value)
			return &sdk.Result{}, nil

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"market acknowledgment packet not found: %s", data.GetClientID())
		}
	}

	return nil, nil
}
