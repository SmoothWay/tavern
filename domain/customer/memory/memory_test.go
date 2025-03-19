package memory

import (
	"errors"
	"testing"

	"github.com/SmoothWay/tavern/domain/customer"

	"github.com/google/uuid"
)

func TestMemory_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := customer.NewCustomer("bob")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	var testCases = []testCase{
		{
			name:        "no customer by uuid",
			id:          uuid.MustParse("22696126-db2c-48e3-a3b5-d1efbb615721"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, but got: %v", tc.expectedErr, err)
			}
		})
	}
}
