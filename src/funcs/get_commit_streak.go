package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const query_frame = `
{
  user(login: "%s") {
    contributionsCollection {
      contributionCalendar {
        totalContributions
        weeks {
          contributionDays {
            date
            contributionCount
          }
        }
      }
    }
  }
}
`

type response struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
					Weeks              []struct {
						ContributionDays []struct {
							ContributionCount int `json:"contributionCount"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}

func calculateStreak(weeks []struct {
	ContributionDays []struct {
		ContributionCount int `json:"contributionCount"`
	} `json:"contributionDays"`
}) int {
	var currentStreak, maxStreak int
	for _, week := range weeks {
		for _, day := range week.ContributionDays {
			if day.ContributionCount > 0 {
				currentStreak++
				if currentStreak > maxStreak {
					maxStreak = currentStreak
				}
			} else {
				currentStreak = 0
			}
		}
	}
	return maxStreak
}

func GetLongestStreak(username string) (int, error) {
	// ctx := context.Background()
	query := fmt.Sprintf(query_frame, username)

	request := GraphQLQuery{Query: query}
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("json Marshal failed: %v", err)
	}

	url := "https://api.github.com/graphql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("GitHub GraphQL API request failed: %v", err)
	}

	req.Header.Set("Authorization", "bearer "+os.Getenv("GITHUB_TOKEN"))

	client := &http.Client{}
	resp, _ := client.Do(req)

	fmt.Println("response1: ", resp)

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatalf("Decoder failed: %v", err)
	}

	fmt.Println("response: ", res.Data.User.ContributionsCollection.ContributionCalendar.Weeks)

	// res, err := makeRequest(ctx, query)
	// if err != nil {
	//     return 0, err
	// }

	// コミットストリークを計算
	return calculateStreak(res.Data.User.ContributionsCollection.ContributionCalendar.Weeks), nil
}
