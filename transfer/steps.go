package transfer

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/validate"
)

type StepResult struct {
	Name   string
	OK     bool
	Detail string
}

func (s *Service) CheckInput(readers []model.Reader) StepResult {
	issues := validate.ValidateBatch(readers)
	if len(issues) > 0 {
		return StepResult{"validate", false, fmt.Sprintf("%d issues", len(issues))}
	}
	return StepResult{"validate", true, "input accepted"}
}
func (s *Service) CheckDuplicate(readers []model.Reader) StepResult {
	issues := validate.ValidateBatch(readers)
	for _, i := range issues {
		if i.Message == "duplicate" {
			return StepResult{"duplicate", false, i.CardNumber}
		}
	}
	return StepResult{"duplicate", true, "no duplicates"}
}
func (s *Service) CheckDebt(readers []model.Reader) StepResult {
	for _, r := range readers {
		if r.DebtCents < 0 {
			return StepResult{"debt", false, r.CardNumber}
		}
	}
	return StepResult{"debt", true, "debt values valid"}
}
func (s *Service) Prepare(readers []model.Reader) []StepResult {
	return []StepResult{s.CheckInput(readers), s.CheckDuplicate(readers), s.CheckDebt(readers)}
}
func AllPassed(steps []StepResult) bool {
	for _, x := range steps {
		if !x.OK {
			return false
		}
	}
	return true
}
