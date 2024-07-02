# Go + GraphQL

create project:
```bash
mkdir go-graphql-client-demo
cd go-graphql-client-demo
```

init:
```bash
go mod init go-graphql-client-demo
```

add graphql package:

[github link](https://github.com/graphql-go/graphql)

```bash
go get github.com/graphql-go/graphql
```

create file:
```bash
touch main.go
```

source code:
```go
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
```

start query:
```bash
go run main.go
```

log:
```
{
  "Characters": {
    "Info": {
      "Count": 107
    },
    "Results": [
      {
        "Name": "Mechanical Rick"
      },
      {
        "Name": "Mega Fruit Farmer Rick"
      },
      {
        "Name": "Morty Rick"
      },
      {
        "Name": "Pickle Rick"
      },
      {
        "Name": "Plumber Rick"
      },
      {
        "Name": "Quantum Rick"
      },
      {
        "Name": "Regional Manager Rick"
      },
      {
        "Name": "Reverse Rick Outrage"
      },
      {
        "Name": "Rick D. Sanchez III"
      },
      {
        "Name": "Rick Guilt Rick"
      },
      {
        "Name": "Rick Prime"
      },
      {
        "Name": "Rick D-99"
      },
      {
        "Name": "Rick D716"
      },
      {
        "Name": "Rick D716-B"
      },
      {
        "Name": "Rick D716-C"
      },
      {
        "Name": "Rick Sanchez"
      },
      {
        "Name": "Rick J-22"
      },
      {
        "Name": "Rick K-22"
      },
      {
        "Name": "Rick Sanchez"
      },
      {
        "Name": "Ricktiminus Sancheziminius"
      }
    ]
  },
  "Location": {
    "Id": "1"
  },
  "EpisodesByIds": [
    {
      "Id": "1"
    },
    {
      "Id": "2"
    }
  ]
}
```