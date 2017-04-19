// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/cryptocurrency.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type RegisterAddressRes struct {
	Type   string `codec:"type" json:"type"`
	Family string `codec:"family" json:"family"`
}

type RegisterAddressArg struct {
	SessionID    int    `codec:"sessionID" json:"sessionID"`
	Address      string `codec:"address" json:"address"`
	Force        bool   `codec:"force" json:"force"`
	WantedFamily string `codec:"wantedFamily" json:"wantedFamily"`
}

type CryptocurrencyInterface interface {
	RegisterAddress(context.Context, RegisterAddressArg) (RegisterAddressRes, error)
}

func CryptocurrencyProtocol(i CryptocurrencyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.cryptocurrency",
		Methods: map[string]rpc.ServeHandlerDescription{
			"registerAddress": {
				MakeArg: func() interface{} {
					ret := make([]RegisterAddressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RegisterAddressArg)
					if !ok {
						err = rpc.NewTypeError((*[]RegisterAddressArg)(nil), args)
						return
					}
					ret, err = i.RegisterAddress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type CryptocurrencyClient struct {
	Cli rpc.GenericClient
}

func (c CryptocurrencyClient) RegisterAddress(ctx context.Context, __arg RegisterAddressArg) (res RegisterAddressRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.cryptocurrency.registerAddress", []interface{}{__arg}, &res)
	return
}
