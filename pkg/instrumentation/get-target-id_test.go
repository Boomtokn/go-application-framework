package instrumentation

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
)

func Test_GetTargetId(t *testing.T) {
	t.Run("handles a filesystem directory path", func(t *testing.T) {
		tempDir := t.TempDir()
		targetId, err := GetTargetId(tempDir, WithSubPath("myfile.go"))
		assert.NoError(t, err)

		pattern := `^pkg:filesystem/[a-fA-F0-9]{64}/001#myfile.go$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("handles a file directory path", func(t *testing.T) {
		tempDir := t.TempDir()
		tempFile1 := filepath.Join(tempDir, "test1.ts")
		targetId, err := GetTargetId(tempFile1, WithSubPath("test1.ts"))
		assert.NoError(t, err)

		pattern := `^pkg:filesystem/[a-fA-F0-9]{64}/001#test1.ts$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("handles paths with special characters", func(t *testing.T) {
		tempDir := t.TempDir()
		targetId, err := GetTargetId(tempDir, WithSubPath("filecontaining>specialcharacters123<.ts"))
		assert.NoError(t, err)

		pattern := `^pkg:filesystem/[a-fA-F0-9]{64}/001#filecontaining\%3Especialcharacters123\%3C.ts$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("handles a directory which has a .git file at the root", func(t *testing.T) {
		tempDir := clone(t)

		targetId, err := GetTargetId(tempDir)
		assert.NoError(t, err)

		pattern := `^pkg:git/github\.com/snyk-fixtures/shallow-goof-locked@[a-fA-F0-9]{40}\?branch=master$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("fails back to filesystem due to invalid repo", func(t *testing.T) {
		tempDir := clone(t)

		// Remove HEAD ref to break git
		headFile := filepath.Join(tempDir, ".git", "HEAD")
		err := os.Remove(headFile)
		assert.NoError(t, err)

		targetId, err := GetTargetId(tempDir)
		assert.NoError(t, err)

		pattern := `^pkg:filesystem/[a-fA-F0-9]{64}/001$`
		matched, err := regexp.MatchString(pattern, string(targetId))
		assert.NoError(t, err)
		assert.True(t, matched)
	})

	t.Run("fails back to filesystem if invalid git url configured", func(t *testing.T) {
		tempDir := clone(t)

		// Update .git/config file to include empty url
		err := updateFile(t, tempDir+"/.git/config", "https://github.com/snyk-fixtures/shallow-goof-locked.git", "")
		assert.NoError(t, err)

		targetId, err := GetTargetId(tempDir, WithSubPath("package.json"))
		assert.NoError(t, err)

		pattern := `^pkg:filesystem/[a-fA-F0-9]{64}/001#package.json$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("handles a git directory with a file location", func(t *testing.T) {
		tempDir := clone(t)

		targetId, err := GetTargetId(tempDir, WithSubPath("package.json"))
		assert.NoError(t, err)

		pattern := `^pkg:git/github\.com/snyk-fixtures/shallow-goof-locked@[a-fA-F0-9]{40}\?branch=master#package.json$`
		assert.Regexp(t, pattern, string(targetId))
	})

	t.Run("sanitize git url if it contains credentials", func(t *testing.T) {
		tempDir := clone(t)

		// Edit .git/config file to adjust remote url
		err := updateFile(t, tempDir+"/.git/config", "https://github.com/snyk-fixtures/shallow-goof-locked.git", "https://username:password@github.com/snyk-fixtures/shallow-goof-locked.git")
		assert.NoError(t, err)

		targetId, err := GetTargetId(tempDir, WithSubPath("package.json"))
		assert.NoError(t, err)

		pattern := `^pkg:git/github\.com/snyk-fixtures/shallow-goof-locked@[a-fA-F0-9]{40}\?branch=master#package.json$`
		assert.Regexp(t, pattern, string(targetId))
	})
}

func clone(t *testing.T) string {
	t.Helper()
	repoUrl := "https://github.com/snyk-fixtures/shallow-goof-locked.git"
	repoDir := t.TempDir()
	repo, err := git.PlainClone(repoDir, false, &git.CloneOptions{URL: repoUrl})
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	return repoDir
}

func updateFile(t *testing.T, filePath, target, replacement string) error {
	t.Helper()
	file, err := os.Open(filePath)
	assert.NoError(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var updatedLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			newLine := strings.Replace(line, target, replacement, -1)
			updatedLines = append(updatedLines, newLine)
		} else {
			updatedLines = append(updatedLines, line)
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	// Write updated content back to the file
	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := bufio.NewWriter(newFile)
	for _, line := range updatedLines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
