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
					IsFork bool `json:"isFork"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}

func GetRepositories(username, token string) ([]Repository, error) {
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
				}
			}
		}
	}
	`

	// GraphQLクエリにユーザー名を挿入
	query = fmt.Sprintf(query, username)

	// GraphQL APIにリクエストを送信
	url := "https://api.github.com/graphql"
	reqBody := GraphQLQuery{Query: query}
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		fmt.Println("Request Error:", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスをパース
	var response GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("JSON Decode Error:", err)
		return nil, err
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

	return repositories, nil
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

// リポジトリの言語別のファイルサイズを取得するための構造体です

type RepositoryLanguage struct {
	Name string `json:"name"`
}

type RepositoryLanguages struct {
	Edges []struct {
		Node RepositoryLanguage `json:"node"`
		Size int                `json:"size"`
	} `json:"edges"`
	TotalSize int `json:"totalSize"`
}

type RepositoryDetail struct {
	Data struct {
		Repository struct {
			Languages RepositoryLanguages `json:"languages"`
		} `json:"repository"`
	} `json:"data"`
}

// リポジトリの言語別のファイルサイズを取得する関数
func GetRepositoryLanguage(repoName, repoOwner, token string) (map[string]int, int, error) {

	// 変更するクエリをここに入力してください
	query_frame := `
	{
		repository(owner: "%s",name: "%s") {
			languages(first: 10, orderBy: {field: SIZE, direction: DESC}) {
				edges {
					node {
						name
					}
					size
				}
				totalSize
			}
		}
	}
	`
	fmt.Printf(repoName, repoOwner)
	// GraphQLクエリを構築
	query := fmt.Sprintf(query_frame, repoOwner, repoName)

	// GraphQLクエリを実行して詳細情報を取得
	request := GraphQLQuery{Query: query}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, 0, err
	}

	url := "https://api.github.com/graphql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	var response RepositoryDetail
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, 0, err
	}

	fmt.Printf("response: %v\n", response.Data.Repository.Languages.Edges)

	// 言語別のファイルサイズをマップに集計
	languageSizes := make(map[string]int)
	for _, edge := range response.Data.Repository.Languages.Edges {
		languageName := edge.Node.Name
		size := edge.Size
		languageSizes[languageName] = size
	}

	totalSize := response.Data.Repository.Languages.TotalSize

	fmt.Printf("languageSizes: %v\n", languageSizes)

	return languageSizes, totalSize, nil
}
