package comments_test

import (
	"context"
	"disspace/business/comments"
	_mockCommentRepository "disspace/business/comments/mocks"
	"disspace/helpers/messages"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var commentRepository _mockCommentRepository.Repository

var commentUseCase comments.UseCase

func setup() {
	commentUseCase = comments.NewCommentUseCase(&commentRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := comments.Domain{
			ThreadID: "61c3ef84cdece58dbf9d112a",
			ParentID: "61c3ef84cdece58dbf9d112a",
			Username: "blueflower1234",
			Text:     "This is a very nice thread, I mostly agree with what you've said",
		}

		commentRepository.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := commentUseCase.Create(context.Background(), &domain, "redrose123")

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Empty Text", func(t *testing.T) {
		domain := comments.Domain{
			ThreadID: "61c3ef84cdece58dbf9d112a",
			ParentID: "61c3ef84cdece58dbf9d112a",
			Username: "blueflower1234",
			Text:     "   ",
		}
		commentRepository.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(comments.Domain{}, messages.ErrTextCannotBeEmpty).Once()
		result, err := commentUseCase.Create(context.Background(), &domain, "redrose123")

		assert.Equal(t, err, messages.ErrTextCannotBeEmpty)
		assert.Equal(t, comments.Domain{}, result)
	})

	t.Run("Text 3 | Invalid User ID Cannot Empty", func(t *testing.T) {
		commentRepository.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(comments.Domain{}, messages.ErrInvalidUserID).Once()
		result, err := commentUseCase.Create(context.Background(), &comments.Domain{}, "  ")

		assert.Equal(t, messages.ErrInvalidUserID, err)
		assert.Empty(t, result)
	})

	t.Run("Test 4 | Internal Server Error", func(t *testing.T) {
		domain := comments.Domain{
			ThreadID: "61c3ef84cdece58dbf9d112a",
			ParentID: "61c3ef84cdece58dbf9d112a",
			Username: "blueflower1234",
			Text:     "This is a very nice thread, I mostly agree with what you've said",
		}

		commentRepository.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(comments.Domain{}, messages.ErrInternalServerError).Once()

		result, err := commentUseCase.Create(context.Background(), &domain, "redrose123")

		assert.Equal(t, messages.ErrInternalServerError, err)
		assert.Empty(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		commentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.Delete(context.Background(), "onetwothree123", "61c3ef84cdece58dbf9d112a")

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Invalid Comment ID", func(t *testing.T) {
		commentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(messages.ErrInvalidCommentID).Once()

		err := commentUseCase.Delete(context.Background(), "onetwothree123", "")

		assert.Equal(t, messages.ErrInvalidCommentID, err)
	})

	t.Run("Test 3 | Invalid User ID", func(t *testing.T) {
		commentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(messages.ErrInvalidUserID).Once()

		err := commentUseCase.Delete(context.Background(), "", "61e7d4ae2286ebd562e314d7")

		assert.Equal(t, messages.ErrInvalidUserID, err)
	})

	t.Run("Test 4 | Invalid Comment ID From Database", func(t *testing.T) {
		commentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(messages.ErrInvalidCommentID).Once()

		err := commentUseCase.Delete(context.Background(), "onetwothree123", "61e7d4ae2286ebd562e3")

		assert.Equal(t, messages.ErrInvalidCommentID, err)
	})
}

func TestSearch(t *testing.T) {
	t.Run("Test Case 1 | Valid Test", func(t *testing.T) {
		commentsDomain := []comments.Domain{
			{
				ID:       "61ef7e825b3bde4d5d8a5809",
				ThreadID: "61ee44afbe6ee82e98c03595",
				ParentID: "61e829422286ebd562e314da",
				Username: "milens1234",
				Text:     "<p>nested 1</p>",
			},
			{
				ID:       "61ef7fab5b3bde4d5d8a580b",
				ThreadID: "61ee44afbe6ee82e98c03595",
				ParentID: "61ef7e825b3bde4d5d8a5809",
				Username: "yupi123",
				Text:     "<p>nested 2</p>",
			},
		}
		commentRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(commentsDomain, nil).Once()

		result, err := commentUseCase.Search(context.Background(), "nest", "created_at")

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Invalid Query Params Sorting", func(t *testing.T) {
		commentRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]comments.Domain{}, messages.ErrInvalidQueryParam).Once()

		result, err := commentUseCase.Search(context.Background(), "nest", "title")

		assert.Equal(t, messages.ErrInvalidQueryParam, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Not Found", func(t *testing.T) {
		commentRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]comments.Domain{}, messages.ErrDataNotFound).Once()

		result, err := commentUseCase.Search(context.Background(), "marvel", "num_votes")

		assert.NotEmpty(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := comments.Domain{
			ID:       "61ef7e825b3bde4d5d8a5809",
			ThreadID: "61ee44afbe6ee82e98c03595",
			ParentID: "61e829422286ebd562e314da",
			Username: "milens1234",
			Text:     "<p>nested 1</p>",
		}

		commentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := commentUseCase.GetByID(context.Background(), "61ef7e825b3bde4d5d8a5809")

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("Test 2 | Data Not Found", func(t *testing.T) {
		commentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(comments.Domain{}, messages.ErrDataNotFound).Once()

		result, err := commentUseCase.GetByID(context.Background(), "61ef7e825b3bde4d5d8a5aaa")

		assert.Equal(t, messages.ErrDataNotFound, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Invalid Comment ID", func(t *testing.T) {
		commentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(comments.Domain{}, messages.ErrInvalidCommentID).Once()

		result, err := commentUseCase.GetByID(context.Background(), "  ")

		assert.Equal(t, messages.ErrInvalidCommentID, err)
		assert.Empty(t, result)
	})
}

func TestGetAllInThread(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		commentsDomain := []comments.Domain{
			{
				ID:       "61ef7e825b3bde4d5d8a5809",
				ThreadID: "61ee44afbe6ee82e98c03595",
				ParentID: "61ee44afbe6ee82e98c03595",
				Username: "milens1234",
				Text:     "<p>great thread</p>",
			},
			{
				ID:       "61ef7fab5b3bde4d5d8a580b",
				ThreadID: "61ee44afbe6ee82e98c03595",
				ParentID: "61ee44afbe6ee82e98c03595",
				Username: "yupi123",
				Text:     "<p>interesting...</p>",
			},
		}

		commentRepository.On("GetAllInThread", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(commentsDomain, nil).Once()

		result, err := commentUseCase.GetAllInThread(context.Background(), "61ee44afbe6ee82e98c03595", "61ee44afbe6ee82e98c03595", "")

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Invalid Option", func(t *testing.T) {
		commentRepository.On("GetAllInThread", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]comments.Domain{}, messages.ErrInvalidOption).Once()

		result, err := commentUseCase.GetAllInThread(context.Background(), "61ee44afbe6ee82e98c03595", "61ee44afbe6ee82e98c03595", "in")

		assert.Equal(t, messages.ErrInvalidOption, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Invalid Thread Or Parent ID", func(t *testing.T) {
		commentRepository.On("GetAllInThread", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]comments.Domain{}, messages.ErrInvalidThreadOrParent).Once()

		result, err := commentUseCase.GetAllInThread(context.Background(), " ", "    ", "")

		assert.Equal(t, messages.ErrInvalidThreadOrParent, err)
		assert.Empty(t, result)
	})

	t.Run("Test 4 | Internal Server Error or Data Not Found", func(t *testing.T) {
		commentRepository.On("GetAllInThread", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]comments.Domain{}, messages.ErrDataNotFound)

		result, err := commentUseCase.GetAllInThread(context.Background(), "61ee44afbe6ee82e98c03595", "61ee44afbe6ee82e98c03595", "")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
