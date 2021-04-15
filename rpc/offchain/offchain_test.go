package offchain

import (
	"os"
	"testing"

	"github.com/rjman-self/centrifuge-substrate-rpc/v3/client"
	"github.com/rjman-self/centrifuge-substrate-rpc/v3/config"
)

var offchain *Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	offchain = NewOffchain(cl)
	os.Exit(m.Run())
}
