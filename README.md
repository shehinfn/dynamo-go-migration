# Migration Package

The `migration` package is a Go library that provides an easy way to create and manage AWS DynamoDB tables based on Go structs. It allows developers to define their DynamoDB schema using Go structs and `dynamo` tags, and automatically handles table creation.

## Installation

To install the `migration` package, run the following command:

go get github.com/shehinfn/dynamo-go-migration

Replace `github.com/shehinfn/dynamo-go-migration` with the actual path to the GitHub repository you created earlier.

## Usage

To use the `migration` package, first import it in your Go project:

import "github.com/shehinfn/dynamo-go-migration"

Then, define your DynamoDB schema using Go structs and `dynamo` tags:

type Lead struct {
	ID        string `dynamo:"S,dynamo_hash" json:"ID,omitempty"`
	FirstName string `dynamo:"S" json:"first_name"`
	LastName  string `dynamo:"S" json:"last_name"`
	Email     string `dynamo:"S" json:"email"`
	Phone     string `dynamo:"S" json:"phone"`
}

Finally, call the `Migrate` function to create the DynamoDB tables:

func main() {
	// Initialize the DynamoDB connection
	database.InitDynnamoDB()

	// Run migrations
	runMigrations()

	// ... rest of your main function
}

func runMigrations() {
	migration.Migrate(db *dynamodb.DynamoDB, migration.ModelInfo{Model: Lead{}, TableName: "leads"})
}

## Contributing

Contributions to the `migration` package are welcome. To contribute, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your changes.
3. Implement your changes and commit them to your forked repository.
4. Open a pull request to merge your changes into the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
