package contract

type Dependency struct {
	Usecase    *Usecase
	Repository *Repository
}

func NewDependency() *Dependency {
	// db := NewDB()
	repo := NewRepository()
	return &Dependency{
		Usecase:    NewUsecase(repo),
		Repository: repo,
	}
}
