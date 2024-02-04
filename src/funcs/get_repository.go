package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// / GraphQLQueries はGraphQLのクエリを格納するための構造体です
type GraphQLQueries struct {
	Query string `json:"query"`
}

// Repository はリポジトリの情報を格納するための構造体です
type Repository struct {
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	IsFork bool   `json:"isFork"`
}

// レポジトリ取得の際のGraphQLのレスポンスを格納するための構造体です
type GraphQLResponse struct {
	Data struct {
		User struct {
			RepositoriesContributedTo struct {
				Nodes []struct {
					Name  string `json:"name"`
					Owner struct {
						Login string `json:"login"`
					} `json:"owner"`
				} `json:"nodes"`
			} `json:"repositoriesContributedTo"`

			Repositories struct {
				Nodes []struct {
					Name  string `json:"name"`
					Owner struct {
						Login string `json:"login"`
					} `json:"owner"`
					IsFork     bool `json:"isFork"`
					Stargazers struct {
						TotalCount int `json:"totalCount"`
					} `json:"stargazers"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}

func GetRepositories(username string) ([]Repository, int, error) {
	// GraphQLクエリを定義
	query := `
	{
		user(login: "%s") {
			repositoriesContributedTo(first: 100) {
				nodes {
					name
					owner {
						login
					}
				}
			}
            repositories(first: 100) {
				nodes {
					name
					owner {
						login
					}
					isFork
					stargazers {
						totalCount
					  }
				}
			}
		}
	}

	`

	token, _ := GetTokens(0)
	// GraphQLクエリにユーザー名を挿入
	query = fmt.Sprintf(query, username)

	// GraphQL APIにリクエストを送信
	url := "https://api.github.com/graphql"
	reqBody := GraphQLQueries{Query: query}
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		fmt.Println("Request Error:", err)
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error:", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	// レスポンスをパース
	var response GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("JSON Decode Error:", err)
		return nil, 0, err
	}

	repositories := []Repository{}

	// コントリビュートしたリポジトリ情報を取得
	for _, repo := range response.Data.User.RepositoriesContributedTo.Nodes {
		repository := Repository{
			Name:   repo.Name,
			Owner:  repo.Owner.Login,
			IsFork: false,
		}
		repositories = append(repositories, repository)
	}

	// 残りのリポジトリ情報を取得
	for _, repo := range response.Data.User.Repositories.Nodes {
		// 重複していないかと、フォークでないかを確認
		if !containsRepository(repositories, repo.Name, repo.Owner.Login) && !repo.IsFork {
			repository := Repository{
				Name:   repo.Name,
				Owner:  repo.Owner.Login,
				IsFork: repo.IsFork,
			}
			repositories = append(repositories, repository)
		}
	}
	// スターの総数を計算
	totalStars := 0
	for _, node := range response.Data.User.Repositories.Nodes {
		totalStars += node.Stargazers.TotalCount
	}

	fmt.Printf("全リポジトリのスターの総数: %d\n", totalStars)
	return repositories, totalStars, nil
}

// スライス内にリポジトリが存在するかを確認
func containsRepository(repositories []Repository, name, owner string) bool {
	for _, repo := range repositories {
		if repo.Name == name && repo.Owner == owner {
			return true
		}
	}
	return false
}
