package domain

type ApplicationConfiguration struct {
	PagingDefaultPage int `mapstructure:"pagingDefaultPage"`
	PagingDefaultSize int `mapstructure:"pagingDefaultSize"`
}

type MongoRepositoryConfiguration struct {
	Database        string `mapstructure:"database"`
	UsersCollection string `mapstructure:"usersCollection"`
}
