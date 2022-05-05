package oraclekeeper_test

import (
	"fmt"
	"github.com/ODIN-PROTOCOL/odin-core/x/common/testapp"
	"github.com/ODIN-PROTOCOL/odin-core/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCollectReward(t *testing.T) {
	_, ctx, k := testapp.CreateTestInput(true, true)
	//b := app.BankKeeper
	fmt.Println("Data provider reward before: ", k.GetDataProviderAccumulatedReward(ctx, testapp.Alice.Address))

	dataSource := types.NewDataSource(
		testapp.Alice.Address, "NAME1",
		"DESCRIPTION1", "filename1",
		sdk.NewCoins(sdk.NewInt64Coin("loki", 10), sdk.NewInt64Coin("minigeo", 10)),
	)
	k.SetDataSource(ctx, 1, dataSource)

	req := types.NewRequest(
		1, BasicCalldata,
		[]sdk.ValAddress{testapp.Validators[0].ValAddress, testapp.Validators[1].ValAddress},
		2, 0, testapp.ParseTime(0),
		BasicClientID, []types.RawRequest{
			types.NewRawRequest(1, 1, BasicCalldata),
		}, nil, 0,
	)
	k.SetRequest(ctx, 1, req)

	rawReport := []types.RawReport{types.NewRawReport(1, 0, []byte("data1/1"))}
	k.SetReport(ctx, 1, types.NewReport(testapp.Alice.ValAddress, true, rawReport))
	require.True(t, k.HasReport(ctx, 1, testapp.Alice.ValAddress))

	initialReward := k.GetDataProviderRewardPerByteParam(ctx)
	k.SetAccumulatedDataProvidersRewards(
		ctx,
		types.NewDataProvidersAccumulatedRewards(
			initialReward, sdk.NewCoins(),
		),
	)

	_, err := k.CollectReward(ctx, rawReport, req.RawRequests) // collected == nil
	require.NoError(t, err)
	require.True(t, k.HasDataProviderReward(ctx, testapp.Alice.Address))
	fmt.Println("Data provider reward after: ", k.GetDataProviderAccumulatedReward(ctx, testapp.Alice.Address))
}

func TestAllocateRewardsToDataProviders(t *testing.T) {
	app, ctx, k := testapp.CreateTestInput(true, true)
	b := app.BankKeeper

	dataSource := types.NewDataSource(
		testapp.Alice.Address, "NAME1",
		"DESCRIPTION1", "filename1",
		sdk.NewCoins(sdk.NewInt64Coin("loki", 10)),
	)
	k.SetDataSource(ctx, 1, dataSource)

	req := types.NewRequest(
		1, BasicCalldata,
		[]sdk.ValAddress{testapp.Validators[0].ValAddress, testapp.Validators[1].ValAddress},
		2, 0, testapp.ParseTime(0),
		BasicClientID, []types.RawRequest{
			types.NewRawRequest(1, 1, BasicCalldata),
		}, nil, 0,
	)
	k.SetRequest(ctx, 1, req)
	require.True(t, k.HasRequest(ctx, 1))

	k.AddPendingRequest(ctx, 1)
	k.SetDataProviderAccumulatedReward(ctx, testapp.Alice.Address, sdk.NewCoins(sdk.NewInt64Coin("loki", 10)))
	require.True(t, k.HasDataProviderReward(ctx, testapp.Alice.Address))
	require.Equal(t, "10loki", k.GetDataProviderAccumulatedReward(ctx, testapp.Alice.Address).String())

	k.AllocateRewardsToDataProviders(ctx, 1)
	require.Equal(t, "1000010loki", b.GetBalance(ctx, testapp.Alice.Address, "loki").String())
}
