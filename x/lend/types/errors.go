package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lend module sentinel errors
var (
	ErrInvalidAsset                    = sdkerrors.Register(ModuleName, 1100, "invalid asset")
	ErrInsufficientBalance             = sdkerrors.Register(ModuleName, 1101, "insufficient balance")
	ErrBorrowLimitLow                  = sdkerrors.Register(ModuleName, 1102, "borrow limit too low")
	ErrLendingPoolInsufficient         = sdkerrors.Register(ModuleName, 1103, "lending pool insufficient")
	ErrInvalidRepayment                = sdkerrors.Register(ModuleName, 1104, "invalid repayment")
	ErrInvalidAddress                  = sdkerrors.Register(ModuleName, 1105, "invalid address")
	ErrNegativeTotalBorrowed           = sdkerrors.Register(ModuleName, 1106, "total borrowed was negative")
	ErrInvalidUtilization              = sdkerrors.Register(ModuleName, 1107, "invalid token utilization")
	ErrLiquidationIneligible           = sdkerrors.Register(ModuleName, 1108, "borrower not eligible for liquidation")
	ErrBadValue                        = sdkerrors.Register(ModuleName, 1109, "bad USD value")
	ErrLiquidatorBalanceZero           = sdkerrors.Register(ModuleName, 1110, "liquidator base asset balance is zero")
	ErrNegativeTimeElapsed             = sdkerrors.Register(ModuleName, 1111, "negative time elapsed since last interest time")
	ErrInvalidOraclePrice              = sdkerrors.Register(ModuleName, 1112, "invalid oracle price")
	ErrNegativeAPY                     = sdkerrors.Register(ModuleName, 1113, "negative APY")
	ErrInvalidExchangeRate             = sdkerrors.Register(ModuleName, 1114, "exchange rate less than one")
	ErrInconsistentTotalBorrow         = sdkerrors.Register(ModuleName, 1115, "total adjusted borrow inconsistency")
	ErrInvalidInteresrScalar           = sdkerrors.Register(ModuleName, 1116, "interest scalar less than one")
	ErrEmptyAddress                    = sdkerrors.Register(ModuleName, 1117, "empty address")
	ErrLiquidationRewardRatio          = sdkerrors.Register(ModuleName, 1118, "requested liquidation reward not met")
	ErrorUnknownProposalType           = sdkerrors.Register(ModuleName, 1119, "unknown proposal type")
	ErrorEmptyProposalAssets           = sdkerrors.Register(ModuleName, 1120, "empty assets in proposal")
	ErrorAssetDoesNotExist             = sdkerrors.Register(ModuleName, 1121, "asset does not exist")
	ErrorDuplicateAsset                = sdkerrors.Register(ModuleName, 1122, "duplicate asset")
	ErrorPairDoesNotExist              = sdkerrors.Register(ModuleName, 1123, "pair does not exist")
	ErrorUnauthorized                  = sdkerrors.Register(ModuleName, 1124, "unauthorized")
	ErrBadOfferCoinAmount              = sdkerrors.Register(ModuleName, 1125, "invalid offer coin amount")
	ErrorDuplicateLendPair             = sdkerrors.Register(ModuleName, 1126, "Dublicate lend Pair")
	ErrorDuplicateLend                 = sdkerrors.Register(ModuleName, 1127, "Dublicate lend Position")
	ErrorLendOwnerNotFound             = sdkerrors.Register(ModuleName, 1128, "Lend Owner not found")
	ErrLendNotFound                    = sdkerrors.Register(ModuleName, 1129, "Lend Position not found")
	ErrWithdrawlAmountExceeds          = sdkerrors.Register(ModuleName, 1130, "Withdrawl Amount Exceeds")
	ErrLendAccessUnauthorised          = sdkerrors.Register(ModuleName, 1131, "Unauthorized user for the tx")
	ErrorPairNotFound                  = sdkerrors.Register(ModuleName, 1132, "Pair Not Found")
	ErrorInvalidCollateralizationRatio = sdkerrors.Register(ModuleName, 1133, "Error Invalid Collateralization Ratio")
	ErrorPriceInDoesNotExist           = sdkerrors.Register(ModuleName, 1134, "Error Price In Does Not Exist")
	ErrorPriceOutDoesNotExist          = sdkerrors.Register(ModuleName, 1135, "Error Price Out Does Not Exist")
	ErrorInvalidAmountIn               = sdkerrors.Register(ModuleName, 1136, "Error Invalid Amount In")
	ErrorInvalidAmountOut              = sdkerrors.Register(ModuleName, 1137, "Error Invalid Amount Out")
	ErrorDuplicateBorrow               = sdkerrors.Register(ModuleName, 1138, "Dublicate borrow Position")
	ErrBorrowingPoolInsufficient       = sdkerrors.Register(ModuleName, 1139, "borrowing pool insufficient")
	ErrBorrowNotFound                  = sdkerrors.Register(ModuleName, 1140, "Borrow Position not found")
	ErrBorrowingPositionOpen           = sdkerrors.Register(ModuleName, 1141, "borrowing position open")
	ErrAssetStatsNotFound              = sdkerrors.Register(ModuleName, 1142, "Asset Stats Not Found")
	ErrorDuplicateAssetRatesStats      = sdkerrors.Register(ModuleName, 1143, "Dublicate Asset Rates Stats")
	ErrorAssetStatsNotFound            = sdkerrors.Register(ModuleName, 1144, "Asset Stats Not Found")
	ErrInvalidAssetIdForPool           = sdkerrors.Register(ModuleName, 1145, "Asset Id not defined in the pool")
)
