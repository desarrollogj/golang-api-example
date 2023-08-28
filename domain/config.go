package domain

type MongoRepositoryConfiguration struct {
	Database        string `mapstructure:"database"`
	UsersCollection string `mapstructure:"usersCollection"`
}
