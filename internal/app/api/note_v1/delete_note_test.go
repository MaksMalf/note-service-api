package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	noteMocks "github.com/MaksMalf/testGrpc/internal/app/storage/mocks"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id           = gofakeit.Int64()
		storeErrText = gofakeit.Phrase()

		req = &pb.DeleteNoteRequest{
			Id: id,
		}

		storeErr = errors.New(storeErrText)
	)

	noteMock := noteMocks.NewMockNoteStorage(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().DeleteNote(ctx, id).Return(nil),
		noteMock.EXPECT().DeleteNote(ctx, id).Return(storeErr),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		_, err := api.DeleteNote(ctx, req)
		require.Nil(t, err)
	})

	t.Run("note store err", func(t *testing.T) {
		_, err := api.DeleteNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, storeErrText, err.Error())
	})
}
