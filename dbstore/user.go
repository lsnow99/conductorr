package dbstore

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func SetUser(ctx context.Context, username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(`DELETE FROM user WHERE true`); err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO user (username, password) VALUES (?, ?)`, username, hashed)

	err = tx.Commit()
	return err
}

func CheckUser(ctx context.Context, username, password string) error {
	tx, err := db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err != nil {
		return err
	}

	var dbUsername string
	var dbPassword []byte
	if err = tx.QueryRow(`SELECT username, password FROM user WHERE true`).Scan(&dbUsername, &dbPassword); err != nil {
		return err
	}

	passwordMatches := bcrypt.CompareHashAndPassword(dbPassword, []byte(password)) == nil
	if username != dbUsername || !passwordMatches {
		return errors.New("credentials don't match")
	}

	err = tx.Commit()
	return err
}
