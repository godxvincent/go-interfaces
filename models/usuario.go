package models

// 1. **Definición de estructuras (Clases)**
type User struct {
	name      string
	surname   string
	username  string
	email     string
	cellphone int64
}

func (user *User) GetUserName() string {
	return user.username
}

func (user *User) GetCellphone() int64 {
	return user.cellphone
}

func (user *User) GetEmail() string {
	return user.email
}

type UserBuilderInterface interface {
	WithName(string) UserBuilderInterface
	WithSurName(string) UserBuilderInterface
	WithCellphone(int) UserBuilderInterface
	WithUserName(string) UserBuilderInterface
	WithEmail(string) UserBuilderInterface
	Build() (*User, error)
}

type userBuilder struct {
	name      string
	surname   string
	cellPhone int
	username  string
	email     string
}

func NewUserBuilder() UserBuilderInterface {
	return &userBuilder{name: "nombreUsuario", surname: "apellidoUsuario"}
}

func NewUserBuilderWithNames(name, surname string) UserBuilderInterface {
	return &userBuilder{name: name, surname: surname}
}

func (userBuilder *userBuilder) WithName(name string) UserBuilderInterface {
	userBuilder.name = name
	return userBuilder
}
func (userBuilder *userBuilder) WithSurName(surname string) UserBuilderInterface {
	userBuilder.surname = surname
	return userBuilder
}

func (userBuilder *userBuilder) WithCellphone(cellphone int) UserBuilderInterface {
	userBuilder.cellPhone = cellphone
	return userBuilder
}

func (userBuilder *userBuilder) WithUserName(username string) UserBuilderInterface {
	userBuilder.username = username
	return userBuilder
}

func (userBuilder *userBuilder) WithEmail(email string) UserBuilderInterface {
	userBuilder.email = email
	return userBuilder
}

func (userBuilder *userBuilder) Build() (*User, error) {
	user := &User{}
	user.name = userBuilder.name
	user.surname = userBuilder.surname
	if userBuilder.username != "" {
		user.username = userBuilder.username
	} else {
		user.username = userBuilder.name + "." + userBuilder.surname
	}
	user.email = user.username + "@hotmail.com"
	user.cellphone = int64(userBuilder.cellPhone)

	return user, nil
}
