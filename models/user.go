package models

import (
	"api/db"
	"api/utils"
	"errors"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User)Save()error{
	query := "INSERT INTO users(email, password) VALUES (? , ?)"
	stmt,err := db.DB.Prepare(query)
	if err!=nil{
		return err
	}

	defer stmt.Close()
	hashedPassword,err:=utils.HashPassword(u.Password)
	if err!=nil{
		return err
	}

	result,err := stmt.Exec(u.Email,hashedPassword)
	if err!=nil{
		return err
	}
	userId ,err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials()error{
	query:="SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query,u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID,&retrievedPassword)

	if err!=nil{
		return errors.New("Email Invalid!")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password,retrievedPassword)

	if !passwordIsValid{
		return errors.New("Password Invalid!")
	}

	return nil
}

func GetAllUsers()([]User,error){
	query := "SELECT * FROM users"
	rows,err := db.DB.Query(query)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	var Users []User
	for rows.Next(){
		var User User
		err := rows.Scan(&User.ID, &User.Email, &User.Password)

		if err != nil{
			return nil,err
		}

		Users = append(Users,User)
	}
	return Users,nil
}