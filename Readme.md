# GraphQl server and query


This project shows an example and create graphql server and how to query data from graphql server.
The Dependency Injection using in project is uber/fx and mockery for mock test.

## Start graphql server

$ go build -o cmd/graphql ./cmd/main.go

$ ./cmd/graphql --cmdid=GRAPHQL_SERVER
	
Open browser http://localhost:8080/ and execute query with data to see the result.

```sh
query last_projects($n: Int = 5) {
  projects(last:$n) {
    nodes {
      name
      description
      forksCount
    }
  }
}
```

## Get project info from gitlab:

$ ./cmd/graphql --cmdid=GET_PROJECTS
	
	
