package legacyUser

import (
	"database/sql"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	"github.com/lib/pq"
	"os"
)

type repository struct {
	client *sql.DB
}

func NewRepository() (*repository, error) {
	url := os.Getenv("DATABASE_URL")

	connection, _ := pq.ParseURL(url)
	//connection += " sslmode=require"

	client, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	r := repository{
		client: client,
	}

	if err := r.Ping(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *repository) GetUserById(id int) (*model.User, error) {
	stmt, err := r.client.Prepare(
		"SELECT email, password, firstName, surname, joined, adFree, admin" +
			" FROM Users" +
			" WHERE id = $1",
	)
	if err != nil {
		return nil, err
	}

	return r.doQuery(stmt, id)
}

func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	stmt, err := r.client.Prepare(
		"SELECT email, password, firstName, surname, joined, adFree, admin" +
			" FROM Users" +
			" WHERE email = $1",
	)
	if err != nil {
		return nil, err
	}

	return r.doQuery(stmt, email)
}

func (r *repository) doQuery(stmt *sql.Stmt, args ...interface{}) (*model.User, error) {
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	var user *model.User
	for rows.Next() {
		user = &model.User{}

		err := rows.Scan(
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.Surname,
			&user.Joined,
			&user.AdFree,
			&user.Admin,
		)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, model.ErrLegacyUserNotFound
	}

	return user, nil
}

func (r *repository) Ping() error {
	return r.client.Ping()
}

func (r *repository) Close() error {
	return r.client.Close()
}
