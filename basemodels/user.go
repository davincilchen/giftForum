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


func queryUserWithID(id int) (*sql.Rows, error) {
	return db.GetSlaveDB().Query(`
	 	SELECT users.email, users.password, users.rxpoint ,users.txpoint
	  	FROM users 
	  	WHERE id = $1`,id) 
	
}


func GetUserWithID(id int) (*User, error) {
	row, err := queryUserWithID(id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if !row.Next() {
		return nil, ErrNoMatchingData
	}

	ret := &User{
		BaseUser: BaseUser{
			ID: id,
		},
	}

	err = row.Scan(&ret.Email, &ret.Password, &ret.RxPoint, &ret.TxPoint)
	return ret, err

}


func CreateUserSendGift(fromID, toID int) error{
	d := db.GetMasterDB()
	tx, err := d.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	sqlStatement := `UPDATE users SET txpoint = txpoint+1 WHERE id = $1;`
	_, err = d.Exec(sqlStatement, fromID)
	if err != nil {
  		return err
	}
	
	sqlStatement = `UPDATE users SET rxpoint = rxpoint+1 WHERE id = $1;`
	_, err = d.Exec(sqlStatement, toID)
	if err != nil {
  		return err
	}

	tx.Commit()
	return nil
}


func GetTopTxUser() ([]User, error){

	rows, err := db.GetSlaveDB().Query(`
	 	SELECT users.id, users.email, users.rxpoint ,users.txpoint
	  	FROM users Order By users.txpoint DESC limit 10`) 
		

  	if err != nil {
	 	return nil, err
	}
	  
	defer rows.Close()
	  
	var ret []User
	for rows.Next() {
		var s User
		if err := rows.Scan(&s.ID, &s.Email,  &s.RxPoint, &s.TxPoint); err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}

	return ret, err		  
	
}


func GetTopRxUser() ([]User, error){

	rows, err := db.GetSlaveDB().Query(`
	 	SELECT users.id, users.email, users.rxpoint ,users.txpoint
	  	FROM users Order By users.rxpoint DESC limit 10`) 
		

  	if err != nil {
	 	return nil, err
	}
	  
	defer rows.Close()
	  
	var ret []User
	for rows.Next() {
		var s User
		if err := rows.Scan(&s.ID, &s.Email,  &s.RxPoint, &s.TxPoint); err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}

	return ret, err		  
	
}

