package repository

import (
	"database/sql"
	"fmt"
	"log"
	"portservices/config"
	"portservices/model"
	"portservices/utils"
	"time"
)

// RegisterUser ...
func RegisterUser(user *model.UserRegister) (string, bool) {
	now := time.Now()
	code := utils.StringCode(6)
	sql := `
		INSERT INTO users (email, password, uuid, verify_code, created_at, modified_at, active_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, sqlError := config.DB.Exec(sql, user.Email, user.Password, user.UUID, code, now, now, now)
	if sqlError != nil {
		return "", false
	}
	return code, true
}

//LoginUser ...
func LoginUser(Email string) (model.UserLogin, error) {
	now := time.Now()
	sqlSelect := `
	SELECT _id, email, password, uuid FROM users
		WHERE email = $1;`

	rows, err := config.DB.Query(sqlSelect, Email)

	if err != nil {
		return model.UserLogin{}, err
	}
	defer rows.Close()
	sqlUpdate := `
	UPDATE users
	SET is_active=$1, active_at=$2
	WHERE email=$3;
`
	_, sqlError := config.DB.Exec(sqlUpdate, true, now, Email)
	if sqlError != nil {
		return model.UserLogin{}, sqlError
	}
	var user model.UserLogin
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.UUID,
		)
		if err != nil {
			return model.UserLogin{}, err
		}
	}
	fmt.Println("user", user)
	err = rows.Err()
	if err != nil {
		return model.UserLogin{}, err
	}
	return user, nil
}

//LogoutUser ...
func LogoutUser(Email string) (bool, error) {
	sqlUpdate := `
		UPDATE users
		SET is_active=$1
		WHERE email=$2
	`

	_, err := config.DB.Exec(sqlUpdate, false, Email)
	if err != nil {
		return false, err
	}
	return true, nil
}

//VerifyAccountUser ...
func VerifyAccountUser(Code string) (bool, error) {
	var code string
	sqlSelect := `
	SELECT verify_code FROM users
		WHERE verify_code = $1;`
	err := config.DB.QueryRow(sqlSelect, Code).Scan(&code)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("Not found.")
		return false, err
	case err != nil:
		log.Fatal(err)
		return false, err
	default:
		sqlUpdate := `
		UPDATE users
		SET verify_code=$1
		WHERE verify_code=$2;
	`
		_, sqlErr := config.DB.Exec(sqlUpdate, "", Code)
		if sqlErr != nil {
			return false, err
		}
		return true, nil
	}

}

//CheckEmail ...
func CheckEmail(user *model.UserRegister) bool {
	sqlSelect := `
	SELECT email FROM users
		WHERE email = $1;`

	_, err := config.DB.Query(sqlSelect, user.Email)
	if err != nil {
		return false
	}
	return true
}

//ChangePassword ...
// func ChangePassword(Email string) (bool, error) {

// }
