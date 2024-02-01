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
func GetUserInfo(username, token string) (UserStats, error) {
	userStats := UserStats{}

	// クエリを生成
	queries := map[string]func([]byte) (int, error){
		fmt.Sprintf(`
		{
		  user(login: "%s") {
		    repositories(affiliations: OWNER, first: 100) {
		      nodes {
		        stargazers {
		          totalCount
		        }
		      }
		    }
		  }
		}`, username): parseTotalStars,

		fmt.Sprintf(`
		{
		  user(login: "%s") {
		    contributionsCollection {
		      totalCommitContributions
		    }
		  }
		}`, username): parseTotalCommits,

		fmt.Sprintf(`
		{
		  user(login: "%s") {
		    pullRequests(first: 100) {
		      totalCount
		    }
		  }
		}`, username): parseTotalPRs,

		fmt.Sprintf(`
		{
		  user(login: "%s") {
		    issues(first: 100) {
		      totalCount
		    }
		  }
		}`, username): parseTotalIssues,

		fmt.Sprintf(`
		{
		  user(login: "%s") {
		    repositoriesContributedTo(first: 100) {
		      totalCount
		    }
		  }
		}`, username): parseContributedTo,
	}

	// 各クエリを実行し結果を集計
	for query, parser := range queries {
		data, err := executeQuery(query, token)
		if err != nil {
			return userStats, fmt.Errorf("query execution failed: %w", err)
		}

		result, err := parser(data)
		if err != nil {
			return userStats, fmt.Errorf("parsing query result failed: %w", err)
		}

		switch parser {
		case parseTotalStars:
			userStats.TotalStars = result
		case parseTotalCommits:
			userStats.TotalCommits = result
		case parseTotalPRs:
			userStats.TotalPRs = result
		case parseTotalIssues:
			userStats.TotalIssues = result
		case parseContributedTo:
			userStats.ContributedTo = result
		}
	}

	return userStats, nil
}

// executeQuery は指定されたGraphQLクエリを実行し、結果をバイト配列で返します。
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
