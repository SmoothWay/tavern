package customer_test

import (
	"errors"
	"testing"

	"github.com/SmoothWay/tavern/domain/customer"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrinvalidPerson,
		},
		{
			test:        "valid name",
			name:        "BOB",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, but got: %v", tc.expectedErr, err)
			}
		})
	}
}
