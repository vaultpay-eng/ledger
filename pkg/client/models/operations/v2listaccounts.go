// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/ledger/client/internal/utils"
	"github.com/formancehq/stack/ledger/client/models/components"
	"time"
)

type V2ListAccountsRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor      *string        `queryParam:"style=form,explode=true,name=cursor"`
	Expand      *string        `queryParam:"style=form,explode=true,name=expand"`
	Pit         *time.Time     `queryParam:"style=form,explode=true,name=pit"`
	RequestBody map[string]any `request:"mediaType=application/json"`
}

func (v V2ListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2ListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2ListAccountsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2ListAccountsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V2ListAccountsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V2ListAccountsRequest) GetExpand() *string {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *V2ListAccountsRequest) GetPit() *time.Time {
	if o == nil {
		return nil
	}
	return o.Pit
}

func (o *V2ListAccountsRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type V2ListAccountsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	V2AccountsCursorResponse *components.V2AccountsCursorResponse
}

func (o *V2ListAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V2ListAccountsResponse) GetV2AccountsCursorResponse() *components.V2AccountsCursorResponse {
	if o == nil {
		return nil
	}
	return o.V2AccountsCursorResponse
}
