package user_domain



type userRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id uint64) (*User, error)
}

type UserDomain struct {
	repo userRepository
	ctx context.Context
}

func NewUserDomain(repo repository.UserRepository, ctx context.Context) *UserDomain {
	return &UserDomain{repo: repo, ctx: ctx}
}

func (u *UserDomain) CreateUser(user *User) error {
	return u.repo.CreateUser(u.ctx, user)
}

func (u *UserDomain) GetUser(id uint64) (*User, error) {
	return u.repo.GetUser(u.ctx, id)
}