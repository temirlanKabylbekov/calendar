package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserGetFullName(t *testing.T) {
	cases := []struct {
		lastName         string
		firstName        string
		middleName       string
		expectedFullName string
	}{
		{"", "", "", ""},
		{"", "firstname", "middlename", "firstname middlename"},
		{"lastname", "", "middlename", "lastname middlename"},
		{"lastname", "firstname", "", "lastname firstname"},
		{"lastname", "firstname", "middlename", "lastname firstname middlename"},
	}
	for _, c := range cases {
		user := User{LastName: c.lastName, FirstName: c.firstName, MiddleName: c.middleName}
		require.Equal(t, user.GetFullName(), c.expectedFullName)
	}
}
