package types

import (
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/repo"
)

var _ repo.Object = (*Policy)(nil)

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming
func (a *Policy) SetId(id uint64) {
	a.Id = id
}

