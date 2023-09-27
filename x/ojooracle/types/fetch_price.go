package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFetchPriceData = "fetch_price_data"

var (
	_ sdk.Msg = &MsgFetchPriceData{}

	// FetchPriceResultStoreKeyPrefix is a prefix for storing result.
	FetchPriceResultStoreKeyPrefix = "fetch_price_result"

	// LastFetchPriceIDKey is the key for the last request id.
	LastFetchPriceIDKey = "fetch_price_last_id"

	TempFetchPriceIDKey = "fetch_price_temp_id"

	// FetchPriceClientIDKey is query request identifier.
	FetchPriceClientIDKey = "fetch_price_id"

	LastBlockHeightKey = "last_block_height"

	OracleValidationResultKey = "Oracle_Validation_Result"
)

// NewMsgFetchPriceData creates a new FetchPrice message.
func NewMsgFetchPriceData(
	creator string,
	sourceChannel string,
	callData *FetchPriceCallData,
	twaBatch uint64,
	acceptedHeightDiff int64,
) *MsgFetchPriceData {
	return &MsgFetchPriceData{
		ClientID:           FetchPriceClientIDKey,
		Creator:            creator,
		SourceChannel:      sourceChannel,
		Calldata:           callData,
		TwaBatchSize:       twaBatch,
		AcceptedHeightDiff: acceptedHeightDiff,
	}
}

// Route returns the message route.
func (m *MsgFetchPriceData) Route() string {
	return RouterKey
}

// Type returns the message type.
func (m *MsgFetchPriceData) Type() string {
	return TypeMsgFetchPriceData
}

// GetSigners returns the message signers.
func (m *MsgFetchPriceData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message.
func (m *MsgFetchPriceData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation.
func (m *MsgFetchPriceData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	if m.TwaBatchSize == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid batch size")
	}
	if m.AcceptedHeightDiff <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid height")
	}
	return nil
}

// FetchPriceResultStoreKey is a function to generate key for each result in store.
func FetchPriceResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(FetchPriceResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}
