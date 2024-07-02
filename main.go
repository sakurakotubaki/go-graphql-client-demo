package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

func main() {
	// GraphQLクライアントの初期化
	client := graphql.NewClient("https://rickandmortyapi.com/graphql")

	// GraphQLクエリ
	query := `
    query GetCharactersAndLocation {
        characters(page: 2, filter: { name: "rick" }) {
            info {
                count
            }
            results {
                name
            }
        }
        location(id: 1) {
            id
        }
        episodesByIds(ids: [1, 2]) {
            id
        }
    }
    `

	// GraphQLリクエストの作成
	request := graphql.NewRequest(query)

	// クエリの実行結果を格納するための構造体
	var response struct {
		Characters struct {
			Info struct {
				Count int
			}
			Results []struct {
				Name string
			}
		}
		Location struct {
			Id string
		}
		EpisodesByIds []struct {
			Id string
		}
	}

	// クエリの実行
	if err := client.Run(context.Background(), request, &response); err != nil {
		log.Fatal(err)
	}

	// 結果をJSON形式で出力
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(response); err != nil {
		log.Fatalf("Failed to encode response: %v", err)
	}
}
