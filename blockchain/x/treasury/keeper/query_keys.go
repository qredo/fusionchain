package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/qredo/fusionchain/x/treasury/types"
)

func (k Keeper) Keys(goCtx context.Context, req *types.QueryKeysRequest) (*types.QueryKeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	keyStore := prefix.NewStore(store, types.KeyPrefix(types.KeyKey))
	var keyResponse []*types.KeyResponse
	var pageResp *query.PageResponse
	wTypes := []types.WalletRequestType{types.WalletRequestType_WALLET_REQUEST_TYPE_QRDO, types.WalletRequestType_WALLET_REQUEST_TYPE_ETH_SEPOLIA, types.WalletRequestType_WALLET_REQUEST_TYPE_ETH}

	if req.Type != types.WalletRequestType_WALLET_REQUEST_TYPE_ALL {
		wTypes = []types.WalletRequestType{req.Type}
	}

	for _, wType := range wTypes {
		keys, page, err := query.GenericFilteredPaginate(k.cdc, keyStore, req.Pagination, func(key []byte, value *types.Key) (*types.KeyResponse, error) {
			if value.WorkspaceAddr != "" && value.WorkspaceAddr != req.WorkspaceAddr {
				return nil, nil
			}

			response := &types.KeyResponse{
				Key: value,
			}

			if wType != types.WalletRequestType_WALLET_REQUEST_TYPE_UNSPECIFIED {
				var (
					address string
					err     error
				)
				switch wType {
				case types.WalletRequestType_WALLET_REQUEST_TYPE_QRDO:
					address, err = types.FusionChainAddress(value)
					response.Type = types.WalletType_WALLET_TYPE_QRDO
				case types.WalletRequestType_WALLET_REQUEST_TYPE_ETH_SEPOLIA:
					address, err = types.EthereumAddress(value)
					response.Type = types.WalletType_WALLET_TYPE_ETH_SEPOLIA
				case types.WalletRequestType_WALLET_REQUEST_TYPE_ETH:
					address, err = types.EthereumAddress(value)
					response.Type = types.WalletType_WALLET_TYPE_ETH
				}
				if err != nil {
					return nil, err
				}
				response.Address = address
			}

			return response, nil
		}, func() *types.Key { return &types.Key{} })

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		keyResponse = append(keyResponse, keys...)
		pageResp = page
	}
	return &types.QueryKeysResponse{
		Keys:       keyResponse,
		Pagination: pageResp,
	}, nil
}
