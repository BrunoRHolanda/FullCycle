package cli_test

import (
	"fmt"
	"testing"

	"github.com/BrunoRHolanda/FullCycle/go-hexagonal/adapters/cli"
	mock_application "github.com/BrunoRHolanda/FullCycle/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product test"
	productPrice := 25.0
	productStatus := "enabled"
	productID := "ABC"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has benn created with the price %f and status %s",
		productID,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)

	result, err = cli.Run(productServiceMock, "enable", productID, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)

	result, err = cli.Run(productServiceMock, "disable", productID, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("%s|%s|%f|%s",
		productID,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(productServiceMock, "get", productID, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
