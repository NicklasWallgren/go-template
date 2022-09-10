package migration

type Migrator interface {
	Up() error
	Down() error
	Create(filename string) error
	Fix() error
	Status() error
}
