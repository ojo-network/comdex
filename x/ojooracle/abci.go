package ojooracle

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/comdex-official/comdex/x/ojooracle/keeper"
	"github.com/comdex-official/comdex/x/ojooracle/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)

	block := k.GetLastBlockHeight(ctx)
	if block != types.Int64Zero {
		if ctx.BlockHeight()%types.Int64One == types.Int64Zero {
			for _, request := range []types.PriceRequestType{
				types.PRICE_REQUEST_RATE,
				types.PRICE_REQUEST_MEDIAN,
				types.PRICE_REQUEST_DEVIATION,
			} {
				requestType := int32(request)
				if !k.GetCheckFlag(ctx) {
					_, err := k.FetchPrice(ctx, request)
					if err != nil {
						ctx.Logger().Error("Error in Fetch Price in 1st condition")
					}
					k.SetTempFetchPriceID(ctx, 0, requestType)
					k.SetCheckFlag(ctx, true)
					k.SetOracleValidationResult(ctx, false)

				} else {
					msg := k.GetFetchPriceMsg(ctx)
					_, err := k.FetchPrice(ctx, request)
					if err != nil {
						ctx.Logger().Error("Error in Fetch Price in 2nd condition")
					}

					id := k.GetLastFetchPriceID(ctx, requestType)
					req := k.GetTempFetchPriceID(ctx, requestType)
					res := k.OraclePriceValidationByRequestID(ctx, req, requestType)
					discardData := k.GetDiscardData(ctx)
					//By default discard height -1 - set while adding band proposal
					//addd new parameter in kvv store to save the accepted discard height
					//one more bool value to save the result of the operation---bydefault false
					if !res && discardData.BlockHeight < 0 {
						discardData.BlockHeight = ctx.BlockHeight()
					} else if res && discardData.BlockHeight > 0 {
						if (ctx.BlockHeight() - discardData.BlockHeight) < msg.AcceptedHeightDiff {
							// No issues
							discardData.BlockHeight = -1
						} else if (ctx.BlockHeight() - discardData.BlockHeight) >= msg.AcceptedHeightDiff {
							discardData.DiscardBool = true
							discardData.BlockHeight = -1
						}
					}
					k.SetDiscardData(ctx, discardData)
					k.SetOracleValidationResult(ctx, res)
					k.SetTempFetchPriceID(ctx, id, requestType)
				}
			}
		}
	}
}

// if discardBool true------> setCounter to 0 , price to false   ----should be a if condition at the top----> at the end of condition set Discard Bool to false
// if discardBool false-------> nothing to do.
