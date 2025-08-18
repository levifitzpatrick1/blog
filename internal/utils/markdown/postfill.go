package markdown

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/levifitzpatrick1/blog/internal/utils"
	"gopkg.in/yaml.v3"
)

func LoadPosts(dir string, logger *log.Logger) ([]utils.Post, error) {
	var posts []utils.Post

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("Failed to read md format in %s: %w", path, err)
			}

			parts := bytes.SplitN(content, []byte("---"), 3)
			if len(parts) < 3 {
				return fmt.Errorf("Invalid md format in %s: missing header", path)
			}

			var tmpPost utils.TempPost
			var post utils.Post
			if err := yaml.Unmarshal(parts[1], &tmpPost); err != nil {
				return fmt.Errorf("Failed to parse header in %s: %w", path, err)
			}
			post.Title = tmpPost.Title
			post.Blurb = tmpPost.Blurb
			post.Slug = tmpPost.Slug
			post.CreateDate = tmpPost.CreateDate
			post.ModifyDate = tmpPost.ModifyDate
			post.Content = string(parts[2])

			post.GetTags(tmpPost.Tags)

			rendered, err := RenderHTML(post.Content)
			if err != nil {
				return fmt.Errorf("Failed to render html in %s: %w", path, err)
			}
			post.HTMLContent = template.HTML(rendered)

			posts = append(posts, post)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ModifyDate.After(posts[j].ModifyDate)
	})

	return posts, nil
}
