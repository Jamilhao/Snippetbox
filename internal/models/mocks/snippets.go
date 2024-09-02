package mocks

import (
	"time"

	"snippetbox.jmorelli.dev/internal/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}

func (m *SnippetModel) GetByIDAndUserID(id int, userID int) (*models.Snippet, error) {
    // Implementação mockada do método GetByIDAndUserID
    // Retorne valores mockados ou configuráveis conforme necessário para os testes
    return &models.Snippet{ID: id, Title: "Mock Title", Content: "Mock Content"}, nil
} // teste mockado pro cmd\web\testutils_test.go