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

package test

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"testing"
)

func TestCrossChain_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(conf.SERVER_POLY_SWAP, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	wrapperTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","SrcChainId":2,"BlockHeight":9329385,"Time":1608885420,"DstChainId":4,"FeeTokenHash":"0000000000000000000000000000000000000000","FeeToken":null,"FeeAmount":1000000000000000000000000000000,"Status":0}]`)
	err = json.Unmarshal(wrapperTransactionsData, &wrapperTransactions)
	if err != nil {
		panic(err)
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"State":1,"Time":1608885420,"Fee":11370800000000000000000000,"Height":9329385,"User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","DstChainId":4,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Key":"0000000000000000000000000000000000000000000000000000000000000abe","Param":"200000000000000000000000000000000000000000000000000000000000000abe20e9ef3fe2112e936772ea145dad804d2a33786fe6953ba56f294de9fab4406b0614d8ae73e06552e270340b63a8bcabf9277a1aac99040000000000000014961a23e727ea1beb5f2e2863ec427b7a99cc6f0c06756e6c6f636b4a14bf9c0fd26055ff19245c8080df06d97ae32db3d7146e43f9988f2771f1a2b140cb3faad424767d39fc0000c9ed85be3f01000000000000000000000000000000000000000000000000","SrcTransfer":{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"Time":1608885420,"Asset":"0000000000000000000000000000000000000000","From":"8bc7e7304120b88d111431f6a4853589d10e8132","To":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Amount":9000000000000000000000000000000,"DstChainId":4,"DstAsset":"bf9c0fd26055ff19245c8080df06d97ae32db3d7","DstUser":"ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"}}]`)
	err = json.Unmarshal(srcTransactionsData, &srcTransactions)
	if err != nil {
		panic(err)
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	polyTransactionsData := []byte(`[{"Hash":"d2e8e325265ed314d9f538c2cb3f8e0a71ca2adad8b31db98278a4af6aecc1df","ChainId":0,"State":1,"Time":1609143919,"Fee":0,"Height":1641497,"SrcChainId":2,"SrcHash":"0000000000000000000000000000000000000000000000000000000000000abe","DstChainId":3,"Key":"","SrcTransaction":null}]`)
	err = json.Unmarshal(polyTransactionsData, &polyTransactions)
	if err != nil {
		panic(err)
	}
	chain := new(models.Chain)
	chainData := []byte(`{"ChainId":2,"Name":"Ethereum","Height":9329385}`)
	err = json.Unmarshal(chainData, chain)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(chain, wrapperTransactions, srcTransactions, polyTransactions, nil)
	if err != nil {
		panic(err)
	}
}

func TestCrossChainSrc_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(conf.SERVER_POLY_SWAP, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransactionsData := []byte(`[{"Hash":"278628f397ce9ad92a26b1385cb8347f15971e54c0bd653e83cfe0bc67d61d1c","ChainId":2,"State":1,"Time":1610497597,"Fee":121754000000000,"Height":9455141,"User":"bb04292cbe99e2a1a7c37dffd96fe24dd5ddb1db","DstChainId":79,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Key":"0000000000000000000000000000000000000000000000000000000000001548","Param":"20000000000000000000000000000000000000000000000000000000000000154820139fe21b302d653008032c7b4eb3cc8a5f089314bcc28637af8efc3155ac256514d8ae73e06552e270340b63a8bcabf9277a1aac994f0000000000000014097ae585bfef78ddc8e266abcb840daf7265130c06756e6c6f636b4a140000000000000000000000000000000000000000145cd3143f91a13fe971043e1e4605c1c23b46bf444e1a010000000000000000000000000000000000000000000000000000000000","SrcTransfer":{"TxHash":"278628f397ce9ad92a26b1385cb8347f15971e54c0bd653e83cfe0bc67d61d1c","ChainId":2,"Time":1610497597,"Asset":"09c6a1b0b32a8b2c327532518c68f9b0c54255b8","From":"bb04292cbe99e2a1a7c37dffd96fe24dd5ddb1db","To":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Amount":72270,"DstChainId":79,"DstAsset":"0000000000000000000000000000000000000000","DstUser":"5cd3143f91a13fe971043e1e4605c1c23b46bf44"}}]
`)
	err = json.Unmarshal(srcTransactionsData, &srcTransactions)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(nil, nil, srcTransactions, nil, nil)
	if err != nil {
		panic(err)
	}
}

func TestQuerySrcTransaction_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	srcTransaction := new(models.SrcTransaction)
	db.Model(&models.SrcTransaction{}).Preload("SrcTransfer").First(srcTransaction)
	json, _ := json.Marshal(srcTransaction)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQuerySrcTransfer_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	srcTransaction := new(models.SrcTransaction)
	db.Model(&models.SrcTransaction{}).Preload("SrcTransfer", "chain_id = ?", 2).First(srcTransaction)
	json, _ := json.Marshal(srcTransaction)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQueryPolySrcRelation_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	polySrcRelations := make([]*models.PolySrcRelation, 0)
	db.Debug().Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ?", "00000000").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("left join src_transactions on poly_transactions.src_hash = src_transactions.key").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	json, _ := json.Marshal(polySrcRelations)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQuerySrcPolyDstRelation_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	addresses := []string{"8bc7e7304120b88d111431f6a4853589d10e8132", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	db.Debug().Table("(?) as u", db.Model(&models.SrcTransfer{}).Select("hash").Where("`from` in ? or dst_user in ?", addresses, addresses)).
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").
		Joins("left join src_transactions on u.hash = src_transactions.hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	json, _ := json.Marshal(srcPolyDstRelations)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQueryWrapperTransactionWithToken_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	wrapperTransactionWithToken := new(models.WrapperTransactionWithToken)
	db.Table("wrapper_transactions").Debug().Preload("FeeToken").Preload("FeeToken.TokenBasic").First(wrapperTransactionWithToken)
	json, _ := json.Marshal(wrapperTransactionWithToken)
	fmt.Printf("src Transaction: %s\n", json)
}
