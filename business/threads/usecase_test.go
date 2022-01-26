package threads_test

import (
	"context"
	"disspace/business/threads"
	_mockThreadRepository "disspace/business/threads/mocks"
	"disspace/helpers/messages"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadsRepository _mockThreadRepository.Repository

var threadUseCase threads.UseCase

func setup() {
	threadUseCase = threads.NewThreadUseCase(&threadsRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := threads.Domain{
			Username:   "forgetmenot7",
			Title:      "Can you recreate dinosaur??",
			ImageUrl:   "https://firebasestorage.googleapis.com/v0/b/disspace-76973.appspot.com/o/images%2Fpexels-chan-walrus-958545.jpg?alt=media&token=43d8db09-1074-446b-95b3-8a505a1751ac",
			CategoryID: "61c87cdada2db751926ee3ea",
		}

		threadsRepository.On("Create", mock.Anything, mock.Anything).Return(domain, nil).Once()

		result, err := threadUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Error Empty Title", func(t *testing.T) {
		domain := threads.Domain{
			Username:   "forgetmenot7",
			Title:      " ",
			ImageUrl:   "https://firebasestorage.googleapis.com/v0/b/disspace-76973.appspot.com/o/images%2Fpexels-chan-walrus-958545.jpg?alt=media&token=43d8db09-1074-446b-95b3-8a505a1751ac",
			CategoryID: "61c87cdada2db751926ee3ea",
		}

		threadsRepository.On("Create", mock.Anything, mock.Anything).Return(threads.Domain{}, messages.ErrEmptyTitle)

		result, err := threadUseCase.Create(context.TODO(), &domain)

		assert.Equal(t, messages.ErrEmptyTitle, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Internal Server Error", func(t *testing.T) {
		domain := threads.Domain{
			Username:   "forgetmenot7",
			Title:      "Can you recreate dinosaur??",
			ImageUrl:   "https://firebasestorage.googleapis.com/v0/b/disspace-76973.appspot.com/o/images%2Fpexels-chan-walrus-958545.jpg?alt=media&token=43d8db09-1074-446b-95b3-8a505a1751ac",
			CategoryID: "61c87cdada2db751926ee3ea",
		}

		threadsRepository.On("Create", mock.Anything, mock.Anything).Return(threads.Domain{}, messages.ErrInternalServerError).Once()

		result, err := threadUseCase.Create(context.Background(), &domain)

		assert.Equal(t, messages.ErrInternalServerError, err)
		assert.Empty(t, result)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		threadsDomain := []threads.Domain{
			{
				ID:          "61ee44afbe6ee82e98c03595",
				Username:    "patrickthestar",
				CategoryID:  "61c87cdada2db751926ee3ea",
				Title:       "These are the hardest kanji to write",
				Content:     "Even a native sometimes cannot write down and remember it",
				ImageUrl:    "https://static.wikia.nocookie.net/crowsxworst/images/6/6a/Suzuran_is_Going_Strong%2ng/revision/latest?cb=20200308023017",
				NumVotes:    50,
				NumComments: 5,
			},
			{
				ID:          "61ee44afbe6ee82e66c03676",
				Username:    "thesilencewang",
				CategoryID:  "61c87cdada2db751926ee3ea",
				Title:       "How whale sleep under the water?",
				Content:     "Let's see how it sleep",
				ImageUrl:    "https://static.wikia.nocookie.net/crowsxworst/images/6/6a/Suzuran_is_Going_Strong%2ng/revision/latest?cb=20200308023017",
				NumVotes:    100,
				NumComments: 27,
			},
		}

		threadsRepository.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return(threadsDomain, nil).Once()

		result, err := threadUseCase.GetAll(context.Background(), "created_at")

		assert.Nil(t, err)
		assert.Equal(t, threadsDomain, result)
	})

	t.Run("Test 2 | Invalid Query Param Sorting", func(t *testing.T) {
		threadsRepository.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{}, messages.ErrInvalidQueryParam).Once()

		result, err := threadUseCase.GetAll(context.Background(), "title")

		assert.Equal(t, messages.ErrInvalidQueryParam, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Not Found", func(t *testing.T) {
		threadsRepository.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{}, messages.ErrDataNotFound).Once()

		result, err := threadUseCase.GetAll(context.Background(), "num_votes")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := threads.Domain{
			ID:          "61ee44afbe6ee82e66c03676",
			Username:    "thesilencewang",
			CategoryID:  "61c87cdada2db751926ee3ea",
			Title:       "How whale sleep under the water?",
			Content:     "Let's see how it sleep",
			ImageUrl:    "https://static.wikia.nocookie.net/crowsxworst/images/6/6a/Suzuran_is_Going_Strong%2ng/revision/latest?cb=20200308023017",
			NumVotes:    100,
			NumComments: 27,
		}

		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := threadUseCase.GetByID(context.Background(), "61ee44afbe6ee82e66c03676")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("Test 2 | Invalid Thread ID Whitespace Only", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(threads.Domain{}, messages.ErrInvalidThreadID).Once()

		result, err := threadUseCase.GetByID(context.Background(), "       ")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Invalid Thread ID From Conversion", func(t *testing.T) {
		threadsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(threads.Domain{}, messages.ErrInvalidThreadID).Once()

		result, err := threadUseCase.GetByID(context.Background(), "thisisnotthecorrectid")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
		assert.Empty(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		threadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := threadUseCase.Delete(context.Background(), "61ee44afbe6ee82e66c03676")

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Invalid Thread ID From Database", func(t *testing.T) {
		threadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(messages.ErrInvalidThreadID).Once()

		err := threadUseCase.Delete(context.Background(), "61ee44afbe6ee82e66c03676123afd4")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
	})

	t.Run("Test 3 | Inavlid Thread ID Whitespace", func(t *testing.T) {
		threadsRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(messages.ErrInvalidThreadID).Once()

		err := threadUseCase.Delete(context.Background(), "         ")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
	})
}

func TestUpdate(t *testing.T) {
	update := threads.Domain{
		Content: "We gonna check the fact",
	}

	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		threadsRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := threadUseCase.Update(context.Background(), &update, "61dfd02a4203caaca6696b4d")

		assert.NoError(t, err)
	})

	t.Run("Test 2 | Invalid Thread ID Whitespace", func(t *testing.T) {
		threadsRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(messages.ErrInvalidThreadID).Once()

		err := threadUseCase.Update(context.Background(), &update, "       ")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
	})

	t.Run("Test 3 | Invalid Thread ID From Conversion", func(t *testing.T) {
		threadsRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(messages.ErrInvalidThreadID).Once()

		err := threadUseCase.Update(context.Background(), &update, "61dfd02a4203c")

		assert.Equal(t, messages.ErrInvalidThreadID, err)
	})
}

func TestSearch(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		threadsDomain := []threads.Domain{
			{
				ID:          "61ee44afbe6ee82e98c03595",
				Username:    "patrickthestar",
				CategoryID:  "61c87cdada2db751926ee3ea",
				Title:       "These are the hardest kanji to write",
				Content:     "Even a native sometimes cannot write down and remember it",
				ImageUrl:    "https://static.wikia.nocookie.net/crowsxworst/images/6/6a/Suzuran_is_Going_Strong%2ng/revision/latest?cb=20200308023017",
				NumVotes:    50,
				NumComments: 5,
			},
			{
				ID:          "61ee44afbe6ee82e66c03676",
				Username:    "thesilencewang",
				CategoryID:  "61c87cdada2db751926ee3ea",
				Title:       "How are whales sleep under the water?",
				Content:     "Let's see how it sleep",
				ImageUrl:    "https://static.wikia.nocookie.net/crowsxworst/images/6/6a/Suzuran_is_Going_Strong%2ng/revision/latest?cb=20200308023017",
				NumVotes:    100,
				NumComments: 27,
			},
		}

		threadsRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(threadsDomain, nil).Once()

		result, err := threadUseCase.Search(context.Background(), "are", "num_votes")

		assert.Nil(t, err)
		assert.Equal(t, threadsDomain, result)
	})

	t.Run("Test 2 | Invalid Query Param Sorting", func(t *testing.T) {
		threadsRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{}, messages.ErrInvalidQueryParam).Once()

		result, err := threadUseCase.Search(context.Background(), "are", "title")

		assert.Equal(t, messages.ErrInvalidQueryParam, err)
		assert.Empty(t, result)
	})

	t.Run("Test 3 | Not Found", func(t *testing.T) {
		threadsRepository.On("Search", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]threads.Domain{}, messages.ErrDataNotFound).Once()

		result, err := threadUseCase.Search(context.Background(), "super", "num_votes")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
