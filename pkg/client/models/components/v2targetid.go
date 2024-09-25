// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/formancehq/stack/ledger/client/internal/utils"
	"math/big"
)

type V2TargetIDType string

const (
	V2TargetIDTypeStr    V2TargetIDType = "str"
	V2TargetIDTypeBigint V2TargetIDType = "bigint"
)

type V2TargetID struct {
	Str    *string
	Bigint *big.Int

	Type V2TargetIDType
}

func CreateV2TargetIDStr(str string) V2TargetID {
	typ := V2TargetIDTypeStr

	return V2TargetID{
		Str:  &str,
		Type: typ,
	}
}

func CreateV2TargetIDBigint(bigint *big.Int) V2TargetID {
	typ := V2TargetIDTypeBigint

	return V2TargetID{
		Bigint: bigint,
		Type:   typ,
	}
}

func (u *V2TargetID) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = V2TargetIDTypeStr
		return nil
	}

	var bigint *big.Int = big.NewInt(0)
	if err := utils.UnmarshalJSON(data, &bigint, "", true, true); err == nil {
		u.Bigint = bigint
		u.Type = V2TargetIDTypeBigint
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for V2TargetID", string(data))
}

func (u V2TargetID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Bigint != nil {
		return utils.MarshalJSON(u.Bigint, "", true)
	}

	return nil, errors.New("could not marshal union type V2TargetID: all fields are null")
}