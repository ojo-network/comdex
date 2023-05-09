package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPlaceMarketBidRequest{}, "comdex/auctionsV2/MsgPlaceMarketBidRequest", nil)
	cdc.RegisterConcrete(&MsgPlaceLimitBidRequest{}, "comdex/auctionsV2/MsgPlaceLimitBidRequest", nil)
	cdc.RegisterConcrete(&MsgCancelLimitBidRequest{}, "comdex/auctionsV2/MsgCancelLimitBidRequest", nil)
	cdc.RegisterConcrete(&MsgWithdrawLimitBidRequest{}, "comdex/auctionsV2/MsgWithdrawLimitBidRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgPlaceMarketBidRequest{},
		&MsgPlaceLimitBidRequest{},
		&MsgCancelLimitBidRequest{},
		&MsgWithdrawLimitBidRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
