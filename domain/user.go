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
