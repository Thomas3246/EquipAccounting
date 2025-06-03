package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	_ "modernc.org/sqlite"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users(login, password, name, isAdmin, department) VALUES (?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.IsAdmin, user.DepartmentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	query := "SELECT id, password, name, isAdmin, department FROM users WHERE login = ?"

	row := r.db.QueryRowContext(ctx, query, login)

	user := &domain.User{Login: login}

	err := row.Scan(&user.Id, &user.Password, &user.Name, &user.IsAdmin, &user.DepartmentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) GetUsersView(ctx context.Context) (users []domain.ViewUser, err error) {
	query := `SELECT u.id, u.login, u.name, dep.name || " " || div.name
			  FROM users AS u
			  INNER JOIN department AS dep ON u.department = dep.id
			  INNER JOIN departmentDivisions AS div ON dep.division = div.id
			  WHERE u.isAdmin = 0`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := domain.ViewUser{}
		err := rows.Scan(&u.Id, &u.Login, &u.Name, &u.Department)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetById(ctx context.Context, id int) (domain.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return domain.User{}, row.Err()
	}

	u := domain.User{}

	err := row.Scan(&u.Id, &u.Login, &u.Password, &u.Name, &u.IsAdmin, &u.DepartmentId)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *UserRepo) New(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users(login, password, name, isAdmin, department) VALUES (?,?,?,0,?)"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.DepartmentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateUserData(ctx context.Context, user domain.User) error {
	query := `UPDATE users 
			  SET login = ?, name = ?, department = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, user.Login, user.Name, user.DepartmentId, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) ChangeUserPassword(ctx context.Context, id int, password string) error {
	query := `UPDATE users
			  SET password = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return err
	}
	return nil
}
