package rewe

import (
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestFetchCategoriesForBill(t *testing.T) {
	var rs io.ReadSeeker = nil

	t.Run("calls fetcher for each bill position and combines results", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		bill := Bill{Positions: []Position{
			{
				Text:   "Apfelsaft",
				Amount: 2,
				Price:  100,
				Sum:    200,
				Tax:    "B",
			},
			{
				Text:   "Orangensaft",
				Amount: 3,
				Price:  100,
				Sum:    300,
				Tax:    "A",
			},
		}}

		billReader := NewMockBillReader(ctrl)
		billReader.EXPECT().Read(rs).Return(bill, nil)

		fetcher := NewMockCategoryFetcher(ctrl)
		fetcher.EXPECT().Fetch("Apfelsaft").Return(CategoryInfo{
			Product:    "Apfelsaft",
			Categories: []string{"trinken", "saft"},
		}, nil)
		fetcher.EXPECT().Fetch("Orangensaft").Return(CategoryInfo{
			Product:    "Orangensaft",
			Categories: []string{"trinken", "saft"},
		}, nil)

		// Act
		infos, err := FetchCategoriesForBill(rs, billReader, fetcher)

		// Assert
		assert.NilError(t, err)
		assert.DeepEqual(t, infos, []FullProductInfo{
			{
				Product:    "Apfelsaft",
				Categories: []string{"trinken", "saft"},
				Amount:     2,
				Price:      100,
				Sum:        200,
				Tax:        "B",
			},
			{
				Product:    "Orangensaft",
				Categories: []string{"trinken", "saft"},
				Amount:     3,
				Price:      100,
				Sum:        300,
				Tax:        "A",
			},
		})
	})

	t.Run("ignores errors for individual fetches", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		bill := Bill{Positions: []Position{{Text: "Apfelsaft"}, {Text: "Orangensaft"}}}

		billReader := NewMockBillReader(ctrl)
		billReader.EXPECT().Read(rs).Return(bill, nil)

		fetcher := NewMockCategoryFetcher(ctrl)
		fetcher.EXPECT().Fetch("Apfelsaft").Return(CategoryInfo{}, errors.New("mockError"))
		fetcher.EXPECT().Fetch("Orangensaft").Return(CategoryInfo{
			Product:    "Orangensaft",
			Categories: []string{"trinken", "saft"},
		}, nil)

		// Act
		infos, err := FetchCategoriesForBill(rs, billReader, fetcher)

		// Assert
		assert.NilError(t, err)
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0].Product, "Orangensaft")
	})
}
