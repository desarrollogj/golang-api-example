package domain

type User struct {
	GenericEntity
	FirstName string
	LastName  string
	Email     string
}

type UserCreateInput struct {
	FirstName string
	LastName  string
	Email     string
}

type UserUpdateInput struct {
	UserCreateInput
	Reference string
}

type UserSearchInput struct {
	SearchInput
	FirstName string
	LastName  string
	Email     string
}

type UserSearchOutput struct {
	SearchOutput
	Users []User
}
