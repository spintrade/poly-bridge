/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package controllers

import (
	"bytes"
	"fmt"
	"math/big"
	"poly-bridge/chainsdk"
	"strings"

	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
)

type ItemController struct {
	beego.Controller
}

// todo: cache url and token ids
func (c *ItemController) Items() {
	var req ItemsOfAddressReq
	if !input(&c.Controller, &req) {
		return
	}

	sdk := selectNode(req.ChainId)
	if sdk == nil {
		customInput(&c.Controller, ErrCodeRequest, "chain id not exist")
		return
	}
	wrapper := selectWrapper(req.ChainId)
	if wrapper == emptyAddr {
		customInput(&c.Controller, ErrCodeRequest, "chain id not exist")
		return
	}

	owner := common.HexToAddress(req.Address)
	asset := common.HexToAddress(req.Asset)
	totalCnt, err := sdk.NFTBalance(asset, owner)
	if err != nil {
		nodeInvalid(&c.Controller)
		return
	}

	totalPage := getPageNo(totalCnt, req.PageSize)
	start := req.PageNo * req.PageSize

	empty := func() {
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{})
		output(&c.Controller, data)
	}

	if start >= totalCnt {
		empty()
		return
	}

	tokenIdStr := strings.Trim(req.TokenId, " ")
	if tokenIdStr != "" {
		item, err := findItem(sdk, asset, owner, tokenIdStr)
		if err != nil {
			empty()
			return
		}
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{item})
		output(&c.Controller, data)
		return
	}

	res, err := sdk.GetNFTs(wrapper, asset, owner, start, req.PageSize)
	if err != nil || len(res) == 0 {
		empty()
		return
	}

	items := make([]*Item, 0)
	for tokenId, url := range res {
		items = append(items, &Item{
			TokenId: tokenId.String(),
			Url:     url,
		})
	}

	data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, items)
	output(&c.Controller, data)
}

func findItem(sdk *chainsdk.EthereumSdkPro, asset, owner common.Address, tokenIdString string) (*Item, error) {
	tokenId, ok := new(big.Int).SetString(tokenIdString, 10)
	if !ok {
		return nil, fmt.Errorf("invalid params")
	}

	addr, err := sdk.GetNFTOwner(asset, tokenId)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(addr.Bytes(), owner.Bytes()) {
		return nil, fmt.Errorf("user is not token's owner")
	}

	data, err := sdk.GetNFTURLs(asset, []*big.Int{tokenId})
	if err != nil {
		return nil, err
	}
	url, ok := data[tokenId.String()]
	if !ok {
		return nil, fmt.Errorf("token %s url not exist", tokenId.String())
	}

	return &Item{
		TokenId: tokenId.String(),
		Url:     url,
	}, nil
}
