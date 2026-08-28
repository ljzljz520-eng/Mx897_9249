package validate

import (
	"librarytransfer/model"
	"regexp"
)

var cardPattern = regexp.MustCompile(`^[A-Za-z0-9-]+$`)

func IsCardFormat(card string) bool { return cardPattern.MatchString(card) }
func IsNameFormat(name string) bool {
	for _, r := range name {
		if r >= '0' && r <= '9' {
			return false
		}
	}
	return name != ""
}
func IsDepartmentFormat(dept string) bool { return len(dept) >= 2 }
func IsStatusFormat(status string) bool   { return status == "idle" || status == "borrowed" }
func ValidateCard(card string) Issue {
	if !IsCardFormat(card) {
		return Issue{CardNumber: card, Field: "card_number", Message: "invalid format"}
	}
	return Issue{}
}
func ValidateName(name string) Issue {
	if !IsNameFormat(name) {
		return Issue{Field: "name", Message: "invalid format"}
	}
	return Issue{}
}
func ValidateDepartment(dept string) Issue {
	if !IsDepartmentFormat(dept) {
		return Issue{Field: "department", Message: "invalid format"}
	}
	return Issue{}
}
func ValidateStatus(status string) Issue {
	if !IsStatusFormat(status) {
		return Issue{Field: "status", Message: "invalid format"}
	}
	return Issue{}
}
func ValidateAll(r model.Reader) []Issue {
	return append(ValidateReader(r), ValidateCard(r.CardNumber), ValidateName(r.Name), ValidateDepartment(r.Department), ValidateStatus(r.BorrowingStatus))
}
func ValidStrict(r model.Reader) bool {
	for _, i := range ValidateAll(r) {
		if i.Message != "" {
			return false
		}
	}
	return true
}
