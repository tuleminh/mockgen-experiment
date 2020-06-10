package main

import (
	"errors"
	mock_main "github.com/tuleminh/mockgen-experiment/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateSomething(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := mock_main.NewMockRepository(mockCtrl)
	r.EXPECT().Insert(1).Return(1, nil)
	r.EXPECT().Insert(2).Return(-1, errors.New("inserting failed"))
	r.EXPECT().Insert(gomock.Any()).Return(-1, nil)

	testCases := []struct {
		testCaseName string
		r            Repository
		num          int
		expectedErr  bool
	}{
		{"TestCase01", r, 1, false},
		{"TestCase02", r, 2, true},
		{"TestCase03", r, 3, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testCaseName, func(t *testing.T) {
			got := NewService(tc.r).CreateSomething(&CreateSomethingRequest{Num: tc.num})
			assert.Equal(t, tc.expectedErr, got != nil)
		})
	}
}
