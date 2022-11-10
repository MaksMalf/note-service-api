package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	noteMocks "github.com/MaksMalf/testGrpc/internal/app/storage/mocks"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id       = gofakeit.Int64()
		title    = gofakeit.JobTitle()
		text     = gofakeit.BeerStyle()
		author   = gofakeit.Name()
		createAt = gofakeit.Date()
		updateAt = sql.NullTime{
			Time:  gofakeit.Date(),
			Valid: true,
		}
		storeErrText = gofakeit.Phrase()

		req = &pb.GetNoteRequest{
			Id: id,
		}

		validRes = &pb.GetNoteResponce{
			Note: &pb.Note{
				Id: id,
				Info: &pb.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
				},
				CreatedAt: timestamppb.New(createAt),
				UpdateAt:  timestamppb.New(updateAt.Time),
			},
		}

		storeRes = &model.Note{
			ID: id,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: createAt,
			UpdateAt:  updateAt,
		}

		storeErr = errors.New(storeErrText)
	)

	noteMock := noteMocks.NewMockNoteStorage(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetNote(ctx, id).Return(storeRes, nil),
		noteMock.EXPECT().GetNote(ctx, id).Return(nil, storeErr),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note store err", func(t *testing.T) {
		_, err := api.GetNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, storeErrText, err.Error())
	})
}
