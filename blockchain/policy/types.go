package policy

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type ApproverSet map[string]bool

func BuildApproverSet(approvers []string) ApproverSet {
	approverSet := make(ApproverSet, len(approvers))
	for _, a := range approvers {
		approverSet[a] = true
	}
	return approverSet
}

type PolicyPayload struct {
	cdc codec.BinaryCodec
	any *cdctypes.Any
}

func NewPolicyPayload(cdc codec.BinaryCodec, any *cdctypes.Any) PolicyPayload {
	return PolicyPayload{
		cdc: cdc,
		any: any,
	}
}

func EmptyPolicyPayload() PolicyPayload {
	return NewPolicyPayload(nil, nil)
}

func (p PolicyPayload) Unpack(out any) error {
	return p.cdc.UnpackAny(p.any, out)
}

type Policy interface {
	// Verify tries to verify the current policy. The returned error is nil if
	// the policy is valid.
	Verify(approvers ApproverSet, payload PolicyPayload) error
}
