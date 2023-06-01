package models

import "log"

type Admin_User struct {
	ID       int
	UUID     string
	Name     string
	Password string
}

type Session struct {
	ID            int
	UUID          string
	Name          string
	Admin_User_ID int
}

func (au *Admin_User) CreateUser() (err error) {
	cmd := `insert into admin_users(
		uuid,
		name,
		password)
		values (?,?,?)`

	_, err = Db.Exec(cmd, createUUID(), au.Name, Encrypt(au.Password))

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (au Admin_User, err error) {
	au = Admin_User{}
	cmd := `select 
	id,uuid,name,password 
	from 
	admin_users
	where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&au.ID,
		&au.UUID,
		&au.Name,
		&au.Password,
	)
	return au, err

}

func (au *Admin_User) UpdateUser() (err error) {

	cmd := `update admin_users 
	set uuid = ?, name = ?, password = ?
	where id = ?`

	_, err = Db.Exec(cmd, createUUID(), au.Name, Encrypt(au.Password), au.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (au *Admin_User) DeleteUser() (err error) {

	cmd := `delete from admin_users
	where id = ?`
	_, err = Db.Exec(cmd, au.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByName(name string) (au Admin_User, err error) {
	au = Admin_User{}
	cmd := `select id, uuid, name, password 
	from admin_users
	where name = ?`
	err = Db.QueryRow(cmd, name).Scan(
		&au.ID,
		&au.UUID,
		&au.Name,
		&au.Password,
	)
	return au, err
}

func (au *Admin_User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (uuid,name,admin_user_id)
	values(?,?,?)`
	_, err = Db.Exec(cmd1, createUUID(), au.Name, au.ID)
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id,uuid,name,admin_user_id
	from sessions
	where admin_user_id = ? and name = ?`
	err = Db.QueryRow(cmd2, au.ID, au.Name).Scan(
		&session.ID,
		&session.UUID,
		&session.Name,
		&session.Admin_User_ID,
	)
	if err != nil {
		log.Println(err)
	}
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id,uuid,name,admin_user_id
	from sessions where uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Name,
		&sess.Admin_User_ID,
	)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
