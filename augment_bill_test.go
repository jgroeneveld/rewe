package rewe

import (
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestAugmentBill(t *testing.T) {
	var rs io.ReadSeeker = nil

	t.Run("calls fetcher for each bill position and combines results", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		bill := Bill{
			OrderDate: "2020-01-05",
			Positions: []Position{
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
			},
		}

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
		augmentedBill, err := AugmentBill(rs, billReader, fetcher)

		// Assert
		assert.NilError(t, err)
		assert.DeepEqual(t, augmentedBill, AugmentedBill{
			OrderDate: "2020-01-05",
			AugmentedPositions: []AugmentedPosition{
				{
					Text:       "Apfelsaft",
					Categories: []string{"trinken", "saft"},
					Amount:     2,
					Price:      100,
					Sum:        200,
					Tax:        "B",
				},
				{
					Text:       "Orangensaft",
					Categories: []string{"trinken", "saft"},
					Amount:     3,
					Price:      100,
					Sum:        300,
					Tax:        "A",
				},
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
		augmentedBill, err := AugmentBill(rs, billReader, fetcher)

		// Assert
		assert.NilError(t, err)
		assert.Equal(t, len(augmentedBill.AugmentedPositions), 1)
		assert.Equal(t, augmentedBill.AugmentedPositions[0].Text, "Orangensaft")
	})
}
