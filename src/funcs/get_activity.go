package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// / GraphQLQuery はGraphQLのクエリを格納するための構造体です
type GraphQLQuery struct {
	Query string `json:"query"`
}

// GitHubのGraphQLクエリを定義
type response struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions      int `graphql:"totalCommitContributions"`
				TotalIssueContributions       int `graphql:"totalIssueContributions"`
				TotalPullRequestContributions int `graphql:"totalPullRequestContributions"`
				TotalRepositoryContributions  int `graphql:"totalRepositoryContributions"`
			} `graphql:"contributionsCollection"`
			StarredRepositories struct {
				TotalCount int `graphql:"totalCount"`
			} `graphql:"starredRepositories"`
		} `graphql:"user(login: $username)"`
	}
}

func FetchData(username string) (int, int, int, int, int, error) {
	token, _ := GetTokens(0)
	const query_frame = `
	{
		user(login: "%s") {
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
	  }
	`

	// GraphQLクエリを構築
	query := fmt.Sprintf(query_frame, username)

	// GraphQLクエリを実行して詳細情報を取得
	request := GraphQLQuery{Query: query}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return 0, 0, 0, 0, 0, err
	}

	url := "https://api.github.com/graphql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, 0, 0, 0, 0, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, 0, 0, 0, err
	}
	defer resp.Body.Close()

	var response response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, 0, 0, 0, 0, err
	}

	// 総コントリビューション数
	totalCommitContributions := response.Data.User.ContributionsCollection.TotalCommitContributions
	// 総スター数
	totalStarredRepositories := response.Data.User.StarredRepositories.TotalCount
	// 総Issue数
	totalIssueContributions := response.Data.User.ContributionsCollection.TotalIssueContributions
	// 総PullRequest数
	totalPullRequestContributions := response.Data.User.ContributionsCollection.TotalPullRequestContributions
	// 総コミット数
	totalRepositoryContributions := response.Data.User.ContributionsCollection.TotalRepositoryContributions

	return totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, nil
}

// func main() {
// 	// ユーザー名
// 	username := "kou7306"

// 	// 期間を指定
// 	// from := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
// 	// to := time.Date(2024, 2, 1, 23, 59, 59, 0, time.UTC)

// 	// GitHubのアクセストークンを設定
// 	token,_ :=funcs.GetTokens(0)

// 	// データを取得
// 	totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, err := fetchDataInTimeRange(token, username)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("totalCommitContributions: ", totalCommitContributions)
// 	fmt.Println("totalStarredRepositories: ", totalStarredRepositories)
// 	fmt.Println("totalIssueContributions: ", totalIssueContributions)
// 	fmt.Println("totalPullRequestContributions: ", totalPullRequestContributions)
// 	fmt.Println("totalRepositoryContributions: ", totalRepositoryContributions)
// }
