package repository

import (
	"database/sql"
	"dating-apps/domain/users/model"
	"dating-apps/helpers"
	"strconv"
	"time"
)

func (r *Repo) Insert(req model.User) error {
	var (
		now  = time.Now().UTC()
		args []interface{}
	)

	args = append(args, req.Email, req.Password, req.Fullname, req.DOB, req.Occupation, helpers.SUBSCRIPTION_FREE, now)

	query := `
		INSERT INTO users (email, password, fullname, dob, occupation, subscription, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateProfile(req model.User) error {
	var (
		now  = time.Now().UTC()
		args []interface{}
		i    = 2
	)

	query := `UPDATE users SET updated_at = $1`
	args = append(args, now)

	if req.Fullname != nil {
		query += ` ,fullname = $` + strconv.Itoa(i)
		args = append(args, req.Fullname)
		i++
	}

	if req.DOB != nil {
		query += ` ,dob = $` + strconv.Itoa(i)
		args = append(args, req.DOB)
		i++
	}

	if req.Occupation != nil {
		query += ` ,occupation = $` + strconv.Itoa(i)
		args = append(args, req.Fullname)
		i++
	}

	if req.Subscription != "" {
		query += ` ,subscription = $` + strconv.Itoa(i)
		args = append(args, req.Subscription)
		i++
	}

	query += `WHERE deleted_at IS NULL AND email = $` + strconv.Itoa(i)
	args = append(args, req.Email)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) SelectByCondition(params *model.User) (users []model.User, err error) {
	var (
		args []interface{}
		i    = 1
	)

	query := `SELECT email, password, fullname, dob, occupation, subscription, created_at FROM users WHERE deleted_at IS NULL`

	if params.Email != "" {
		query += ` AND email = $` + strconv.Itoa(i)
		args = append(args, params.Email)
		i++
	}

	if params.Id != "" {
		query += ` AND id = $` + strconv.Itoa(i)
		args = append(args, params.Id)
		i++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var data model.User

		err = rows.Scan(
			&data.Email, &data.Password, &data.Fullname, &data.DOB, &data.Occupation, &data.Subscription, &data.CreatedAt,
		)
		if err != nil {
			return users, err
		}

		users = append(users, data)
	}

	if len(users) == 0 {
		return users, sql.ErrNoRows
	}

	return users, nil
}

func (r *Repo) SelectUsers(email string) (users []model.User, err error) {
	query := `SELECT id, email, fullname, dob, occupation, subscription, created_at 
				FROM users 
				WHERE deleted_at IS NULL
				AND email != $1`

	rows, err := r.db.Query(query, email)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var data model.User

		err = rows.Scan(
			&data.Id, &data.Email, &data.Fullname, &data.DOB,
			&data.Occupation, &data.Subscription, &data.CreatedAt,
		)
		if err != nil {
			return users, err
		}

		users = append(users, data)
	}

	if len(users) == 0 {
		return users, sql.ErrNoRows
	}

	return users, nil
}
