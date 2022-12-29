package services

import "testing"

func runTests(t *testing.T, f func(t *testing.T, accounts Accounts)) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f(t, test.accounts)
		})
	}
}
