package repos

import "database/sql"

type Repos struct {
	Subjects SubjectRepo
}

func New(db *sql.DB) Repos {
	return Repos{
		Subjects: SubjectRepo{DB: db},
	}
}
