package types

import "fmt"

func (a *Action) AddApprover(approver string) error {
	if a.Completed {
		return fmt.Errorf("action already completed")
	}

	for _, a := range a.Approvers {
		if a == approver {
			return fmt.Errorf("approver already added")
		}
	}

	a.Approvers = append(a.Approvers, approver)
	return nil
}
