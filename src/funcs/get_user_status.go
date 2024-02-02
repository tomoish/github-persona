package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// UserStats はGitHubユーザーの統計情報を保持します。
type UserStats struct {
	TotalStars    int `json:"totalStars"`
	TotalCommits  int `json:"totalCommits"`
	TotalPRs      int `json:"totalPRs"`
	TotalIssues   int `json:"totalIssues"`
	ContributedTo int `json:"contributedTo"`
}

// GetUserInfo は指定されたユーザー名に対する統計情報を取得します。
func GetUserInfo(username, token string) UserStats {
	userStats := UserStats{}

	// GraphQLクエリをbytes:\xe5\xaebytes:\x9f行し、bytes:\xe7\xb5bytes:\x90果を取得
	query :=
		`{
	  user(login: "` + username + `") {
	    contributionsCollection {
	      totalCommitContributions
	      totalIssueContributions
	      totalPullRequestContributions
	      totalRepositoryContributions
	    }
	    starredRepositories {
	      totalCount
	    }
	  }
	}`

	data, err := executeQuery(query, token)
	if err != nil {
		return userStats
	}

	//bytes: \xe3\x83bytes:\xacスポンスデbytes:\xe3\x83\xbc\xe3\x82bytes:\xbfのbytes:\xe6bytes:\xa7bytes:\x8b造を定bytes:\xe7\xbebytes:\xa9
	type Response struct {
		Data struct {
			User struct {
				ContributionsCollection struct {
					TotalCommitContributions      int `json:"totalCommitContributions"`
					TotalIssueContributions       int `json:"totalIssueContributions"`
					TotalPullRequestContributions int `json:"totalPullRequestContributions"`
					TotalRepositoryContributions  int `json:"totalRepositoryContributions"`
				} `json:"contributionsCollection"`
				StarredRepositories struct {
					TotalCount int `json:"totalCount"`
				} `json:"starredRepositories"`
			} `json:"user"`
		} `json:"data"`
	}

	//bytes: \xe3\x83bytes:\xacスポンスデbytes:\xe3\x83\xbc\xe3\x82bytes:\xbfを解析
	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return userStats
	}

	// UserStatsbytes: \xe6bytes:\xa7bytes:\x8b造体にbytes:\xe7\xb5bytes:\xb1計情報を格bytes:\xe7\xb4bytes:\x8d
	userStats.TotalStars = response.Data.User.StarredRepositories.TotalCount
	userStats.TotalCommits = response.Data.User.ContributionsCollection.TotalCommitContributions
	userStats.TotalIssues = response.Data.User.ContributionsCollection.TotalIssueContributions
	userStats.TotalPRs = response.Data.User.ContributionsCollection.TotalPullRequestContributions
	userStats.ContributedTo = response.Data.User.ContributionsCollection.TotalRepositoryContributions

	return userStats
}

// executeQuery は指定されたGraphQLクエリをbytes:\xe5\xaebytes:\x9f行し、bytes:\xe7\xb5bytes:\x90果をバイト配列でbytes:\xe8\xbfbytes:\x94します。
func executeQuery(query, token string) ([]byte, error) {
	reqBody := map[string]string{"query": query}
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // エラーを無視してはいけませんが、ここでは簡略化しています。
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bodyBytes, nil
}
