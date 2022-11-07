package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	noteMocks "github.com/MaksMalf/testGrpc/internal/app/storage/mocks"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func TestUpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id    = gofakeit.Int64()
		title = sql.NullString{
			String: gofakeit.JobTitle(),
			Valid:  true,
		}
		text = sql.NullString{
			String: gofakeit.BeerStyle(),
			Valid:  true,
		}
		author = sql.NullString{
			String: gofakeit.Name(),
			Valid:  true,
		}
		storeErrText = gofakeit.Phrase()

		req = &pb.UpdateNoteRequest{
			Id: id,
			UpdateInfo: &pb.UpdateNoteInfo{
				Title:  wrapperspb.String(title.String),
				Text:   wrapperspb.String(text.String),
				Author: wrapperspb.String(author.String),
			},
		}

		storeReq = &model.UpdateNoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		storeErr = errors.New(storeErrText)
	)

	noteMock := noteMocks.NewMockNoteStorage(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().UpdateNote(ctx, id, storeReq).Return(nil),
		noteMock.EXPECT().UpdateNote(ctx, id, storeReq).Return(storeErr),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		_, err := api.UpdateNote(ctx, req)
		require.Nil(t, err)
	})

	t.Run("note store err", func(t *testing.T) {
		_, err := api.UpdateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, storeErrText, err.Error())
	})
}
