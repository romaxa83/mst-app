package tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	mock_repository "github.com/romaxa83/mst-app/library-app/internal/repositories/mocks"
	"github.com/romaxa83/mst-app/library-app/internal/services"
	"github.com/stretchr/testify/require"
	"testing"
)

var errInternalServErr = errors.New("test: internal server error")

func mockAuthorService(t *testing.T) (*services.AuthorService, *mock_repository.MockAuthor) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	authorRepo := mock_repository.NewMockAuthor(mockCtl)

	authorService := services.NewAuthorService(authorRepo)

	return authorService, authorRepo
}

func TestAuthorService_GetOne(t *testing.T) {
	authorService, authorRepo := mockAuthorService(t)

	id := 1
	m := models.Author{}

	authorRepo.EXPECT().GetOneById(id).Return(m, nil)

	_, err := authorService.GetOne(id)

	require.NoError(t, err)
}

func TestAuthorService_GetOneErr(t *testing.T) {
	authorService, authorRepo := mockAuthorService(t)

	authorRepo.EXPECT().GetOneById(gomock.Any()).Return(models.Author{}, errInternalServErr)

	res, err := authorService.GetOne(1)

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, models.Author{}, res)
}
