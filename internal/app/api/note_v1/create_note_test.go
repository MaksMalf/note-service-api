package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	noteMocks "github.com/MaksMalf/testGrpc/internal/app/storage/mocks"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id           = gofakeit.Int64()
		title        = gofakeit.JobTitle()
		text         = gofakeit.BeerStyle()
		author       = gofakeit.Name()
		storeErrText = gofakeit.Phrase()

		req = &pb.CreateNoteRequest{
			Info: &pb.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		storeReq = &model.NoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		validRes = &pb.CreateNoteResponce{
			Id: id,
		}

		storeErr = errors.New(storeErrText)
	)

	noteMock := noteMocks.NewMockNoteStorage(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().CreateNote(ctx, storeReq).Return(id, nil),
		noteMock.EXPECT().CreateNote(ctx, storeReq).Return(int64(0), storeErr),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.CreateNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note store err", func(t *testing.T) {
		_, err := api.CreateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, storeErrText, err.Error())
	})
}
