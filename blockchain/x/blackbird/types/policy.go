package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/repo"
	"gitlab.qredo.com/edmund/blackbird/verifier/golang/impl"
	"gitlab.qredo.com/edmund/blackbird/verifier/golang/simple"
)

var _ repo.Object = (*Policy)(nil)

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming
func (a *Policy) SetId(id uint64) {
	a.Id = id
}

func UnpackPolicy(cdc codec.BinaryCodec, p *Policy) (policy.Policy, error) {
	var policy policy.Policy
	err := cdc.UnpackAny(p.Policy, &policy)
	if err != nil {
		return nil, fmt.Errorf("unpacking Any: %w", err)
	}

	return policy, nil
}

var _ (policy.Policy) = (*BlackbirdPolicy)(nil)

func (p *BlackbirdPolicy) Validate() error {
	participants := make(map[string]impl.Authority, len(p.Participants))
	for abbr, participant := range p.Participants {
		participants[abbr] = impl.ParticipantAsAuthority(participant)
	}
	formatted_policy, err := simple.InstallCheck(p.Data, nil, participants)
	p.Data = formatted_policy
	return err
}

func (p *BlackbirdPolicy) AddressToParticipant(addr string) (string, error) {
	for abbr, participant := range p.Participants {
		if participant == addr {
			return abbr, nil
		}
	}
	return "", fmt.Errorf("address not a participant of this policy")
}

func (p *BlackbirdPolicy) Verify(approvers policy.ApproverSet, policyPayload policy.PolicyPayload) error {
	payload, err := policy.UnpackPayload[*BlackbirdPolicyPayload](policyPayload)
	if err != nil {
		return err
	}

	var witness []byte
	if payload != nil {
		witness = payload.Witness
	}

	return simple.Verify(p.Data, witness, nil, nil, approvers)
}
