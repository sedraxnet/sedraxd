package rpchandlers

import (
	"github.com/sedraxnet/sedraxd/app/appmessage"
	"github.com/sedraxnet/sedraxd/app/rpc/rpccontext"
	"github.com/sedraxnet/sedraxd/domain/consensus/utils/constants"
	"github.com/sedraxnet/sedraxd/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kaspad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSompiSupply, err := context.UTXOIndex.GetCirculatingSompiSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSompi,
		circulatingSompiSupply,
	)

	return response, nil
}
