package repo

import "realjobs/models"

func CreateUser(u models.User) error {
	ConnectToDB()
	insertStmt := `INSERT INTO users(firstname, lastname, email, password, passwordhash) 
		VALUES($1,$2,$3,$4,$5)`

	_, err := db.Exec(insertStmt, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.PasswordHash)

	if err != nil {
		return err
	}

	return nil
}
