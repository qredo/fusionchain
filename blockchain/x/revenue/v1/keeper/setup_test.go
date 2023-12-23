// Copyright 2023 Qredo Ltd.
// This file is part of the Fusion library.
//
// The Fusion library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Fusion library. If not, see https://github.com/qredo/fusionchain/blob/main/LICENSE
package keeper_test

// import (
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"

// 	"github.com/cosmos/cosmos-sdk/crypto/keyring"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
// 	ethtypes "github.com/ethereum/go-ethereum/core/types"

// 	"github.com/qredo/fusionchain/app"
// 	utiltx "github.com/qredo/fusionchain/testutil/tx"
// 	"github.com/qredo/fusionchain/utils"
// 	evm "github.com/evmos/ethermint/x/evm/types"
// 	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
// 	"github.com/qredo/fusionchain/x/revenue/v1/types"

// 	"github.com/stretchr/testify/suite"
// )

// type KeeperTestSuite struct {
// 	suite.Suite

// 	ctx sdk.Context

// 	app            *app.Evmos
// 	queryClient    types.QueryClient
// 	queryClientEvm evm.QueryClient
// 	address        common.Address
// 	signer         keyring.Signer
// 	ethSigner      ethtypes.Signer
// 	consAddress    sdk.ConsAddress
// 	validator      stakingtypes.Validator
// 	denom          string
// }

// var s *KeeperTestSuite

// var (
// 	contract = utiltx.GenerateAddress()
// 	deployer = sdk.AccAddress(utiltx.GenerateAddress().Bytes())
// 	withdraw = sdk.AccAddress(utiltx.GenerateAddress().Bytes())
// )

// func TestKeeperTestSuite(t *testing.T) {
// 	s = new(KeeperTestSuite)
// 	suite.Run(t, s)

// 	// Run Ginkgo integration tests
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Keeper Suite")
// }

// func (suite *KeeperTestSuite) SetupTest() {
// 	chainID := utils.TestnetChainID + "-1"
// 	suite.app = app.Setup(false, feemarkettypes.DefaultGenesisState(), chainID)
// 	suite.SetupApp(chainID)
// }
