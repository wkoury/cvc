package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	model        = "codestral"
	domain       = "http://localhost"
	port         = "11434"
	path         = "/api/generate"
	systemPrompt = `You are an expert in writing conventional commit messages. Your task is to analyze the provided Git diff and generate a properly formatted conventional commit message that accurately reflects the changes made.

### Conventional Commit Guidelines:
1. **Type**: The type should be one of the following (case-insensitive):
   - feat: A new feature.
   - fix: A bug fix.
   - ci: Changes to CI configuration files and scripts.
   - refactor: Code refactoring that doesn't add a new feature or fix a bug.
   - test: Adding missing tests or correcting existing tests.
   - chore: Maintenance tasks that don't fall into other specific categories.
   - build: Changes that affect the build system or external dependencies.
   - docs: Documentation changes.
   - perf: A performance improvement.
   - style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc).

2. **Subject Line**:
   - Use the imperative mood (e.g., "Fix" instead of "Fixed").
   - Keep it concise, ideally under 72 characters.
   - Avoid empty lines.

3. **Other Best Practices**:
   - Keep each line in the body under 72 characters.

### Instructions:
- Analyze the provided Git diff and identify the changes made.
- Determine the appropriate type, scope (if applicable), and subject line based on conventional commit guidelines.
- Ensure all formatting rules are followed.

### Example:
fix: nil pointer dereference

Provide your response as a properly formatted conventional commit message.`
)

const URL = domain + ":" + port + path

// A struct representation of the Ollama API response. Some fields are omitted.
type ResponseBody struct {
	Model     string
	CreatedAt time.Time
	Response  string
}

func Request(diff string) (string, error) {
	payload := map[string]any{
		"model":  model,
		"stream": false,
		"prompt": diff,
		"system": systemPrompt,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	response := ResponseBody{}

	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return "", err
	}

	return response.Response, nil
}
