package funcs


import (
	"fmt"
	"encoding/json"
    "net/http"
    "bytes"
)



// リポジトリの言語別のファイルサイズを取得するための構造体です


type RepositoryLanguage struct {
	Name string `json:"name"`
	
}

type RepositoryLanguages struct {
	Edges []struct {
		Node RepositoryLanguage `json:"node"`
		Size int `json:"size"`
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
func GetRepositoryLanguage(repoName, repoOwner, token string) (map[string]int,int, error) {

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
    query := fmt.Sprintf(query_frame,repoOwner,repoName)

    // GraphQLクエリを実行して詳細情報を取得
    request := GraphQLQuery{Query: query}
    requestBody, err := json.Marshal(request)
    if err != nil {
        return nil,0, err
    }

    url := "https://api.github.com/graphql"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil,0, err
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil,0, err
    }
    defer resp.Body.Close()

	var response RepositoryDetail
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil,0, err
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
	
	return languageSizes,totalSize, nil
}











