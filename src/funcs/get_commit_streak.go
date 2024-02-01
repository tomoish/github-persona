package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func GetCommitHistory(username string) (int, []int, int, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
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

	// fmt.Println("response1: ", resp)

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatalf("Decoder failed: %v", err)
	}

	var dailyCommits []int
	maxCommit := 0
	for weeklyCommits := range res.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for dailyCommit := range res.Data.User.ContributionsCollection.ContributionCalendar.Weeks[weeklyCommits].ContributionDays {
			num_commits := res.Data.User.ContributionsCollection.ContributionCalendar.Weeks[weeklyCommits].ContributionDays[dailyCommit].ContributionCount
			dailyCommits = append(dailyCommits, num_commits)
			if num_commits > maxCommit {
				maxCommit = num_commits
			}
		}
	}

	// fmt.Println("response: ", res.Data.User.ContributionsCollection.ContributionCalendar.Weeks)

	// fmt.Println("dailyCommits: ", dailyCommits)
	// fmt.Println("length of dailyCommits: ", len(dailyCommits))

	return calculateStreak(res.Data.User.ContributionsCollection.ContributionCalendar.Weeks), dailyCommits, maxCommit, err
}
