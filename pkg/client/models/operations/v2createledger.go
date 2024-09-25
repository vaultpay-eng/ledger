// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/ledger/client/models/components"
)

type V2CreateLedgerRequest struct {
	// Name of the ledger.
	Ledger                string                            `pathParam:"style=simple,explode=false,name=ledger"`
	V2CreateLedgerRequest *components.V2CreateLedgerRequest `request:"mediaType=application/json"`
}

func (o *V2CreateLedgerRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2CreateLedgerRequest) GetV2CreateLedgerRequest() *components.V2CreateLedgerRequest {
	if o == nil {
		return nil
	}
	return o.V2CreateLedgerRequest
}

type V2CreateLedgerResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *V2CreateLedgerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}