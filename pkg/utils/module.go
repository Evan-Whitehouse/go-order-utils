package utils

import (
	"fmt"
	"math/big"

	"github.com/Evan-Whitehouse/go-order-utils/pkg/config"
	"github.com/Evan-Whitehouse/go-order-utils/pkg/model"
	"github.com/ethereum/go-ethereum/common"
)

func GetVerifyingContractAddress(chainId *big.Int, contract model.VerifyingContract) (common.Address, error) {
	contracts, err := config.GetContracts(chainId.Int64())
	if err != nil {
		return common.Address{}, err
	}

	switch contract {
	case model.CTFExchange:
		return contracts.Exchange, nil
	case model.NegRiskCTFExchange:
		return contracts.NegRiskExchange, nil
	}

	return common.Address{}, fmt.Errorf("invalid contract")
}
