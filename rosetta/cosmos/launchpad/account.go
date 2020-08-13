package launchpad

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const AccountSdkHandler = "/bank/balances"

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {

	addr := fmt.Sprintf("%s/%s", l.cosmos(AccountSdkHandler), request.AccountIdentifier.Address)
	resp, err := l.c.Get(addr)
	if err != nil {
		return nil, ErrNodeConnection
	}
	defer resp.Body.Close()

	var res balanceResp

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, ErrInterpreting
	}

	return &types.AccountBalanceResponse{
		Balances: convertCoinsToRosettaBalances(res.Result),
	}, nil
}

func convertCoinsToRosettaBalances(coins sdk.Coins) []*types.Amount {
	var amounts []*types.Amount

	for _, coin := range coins {
		amounts = append(amounts, &types.Amount{
			Value: coin.Amount.String(),
			Currency: &types.Currency{
				Symbol: coin.Denom,
			},
		})
	}

	return amounts
}

type balanceResp struct {
	Result sdk.Coins
}
