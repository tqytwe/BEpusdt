package model

import (
	"testing"
)

func TestDefaultBSCRPCAvoidsSubscriptionOnlyEndpoint(t *testing.T) {
	endpoint := defaultConf[RpcEndpointBsc]

	if endpoint != "https://bsc-dataseed.binance.org/" {
		t.Fatalf("unexpected default BSC RPC endpoint: %q", endpoint)
	}
}
