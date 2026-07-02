package markdown

import (
	"log"
	"sync"
	"time"

	"github.com/levifitzpatrick1/blog/internal/utils"
)

type PostStore struct {
	dir      string
	logger   *log.Logger
	mu       sync.Mutex
	posts    []utils.Post
	lastLoad time.Time
}

func NewPostStore(dir string, logger *log.Logger) *PostStore {
	return &PostStore{
		dir:    dir,
		logger: logger,
	}
}

// GetPosts returns all published posts (CreateDate <= time.Now())
// It reloads files from disk if the cache is older than 5 minutes.
func (s *PostStore) GetPosts() ([]utils.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Reload from disk if cache is expired (5 minutes)
	if s.posts == nil || time.Since(s.lastLoad) > 5*time.Minute {
		posts, err := LoadPosts(s.dir, s.logger)
		if err != nil {
			return nil, err
		}
		s.posts = posts
		s.lastLoad = time.Now()
	}

	// Filter out future posts
	now := time.Now()
	var published []utils.Post
	for _, post := range s.posts {
		if !post.CreateDate.After(now) {
			published = append(published, post)
		}
	}

	return published, nil
}
