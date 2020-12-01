package basemodels


import (
	"fmt"
	"database/sql"
	"giftForum/db"
	
)

type BaseUser struct {
	ID      int
	Email   string
	RxPoint int
	TxPoint int
}
type User struct {
	BaseUser
	Password string
}


func CreateUser(email, password string) (*User, error) {

	d := db.GetMasterDB()
	
	sqlStatement := "INSERT INTO USERS (email, password) VALUES ($1, $2) RETURNING id;"
	
	var ID int
    err := d.QueryRow(sqlStatement,email,password).Scan(&ID)
	if err != nil {
		fmt.Println("CreateUser error =", err)
		return nil, err
	}

	
	u := &User{
		BaseUser: BaseUser{
			ID: ID,
			Email: email,
		},
		Password: password,
	}


	return u, nil
}


func queryUser(email string) (*sql.Rows, error) {
	return db.GetSlaveDB().Query(`
	 	SELECT users.id, users.password, users.rxpoint ,users.txpoint
	  	FROM users 
	  	WHERE email = $1`,email) 
	
}



func GetUser(email string) (*User, error) {
	row, err := queryUser(email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if !row.Next() {
		return nil, ErrNoMatchingData
	}

	ret := &User{
		BaseUser: BaseUser{
			Email: email,
		},
	}

	err = row.Scan(&ret.ID, &ret.Password, &ret.RxPoint, &ret.TxPoint)
	return ret, err

}