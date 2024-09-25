// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package client

type Ledger struct {
	V1 *V1
	V2 *V2

	sdkConfiguration sdkConfiguration
}

func newLedger(sdkConfig sdkConfiguration) *Ledger {
	return &Ledger{
		sdkConfiguration: sdkConfig,
		V1:               newV1(sdkConfig),
		V2:               newV2(sdkConfig),
	}
}