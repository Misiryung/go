package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	// "fmt"
	"log"
	"math"
	"time"
	"youtube-clone/database/db"
	"youtube-clone/database/helpers"
	"youtube-clone/database/pbs/helper"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int64      `json:"-"`
	Username          string     `json:"username"`
	ChannelName       string     `json:"channel_name"`
	ProfilePhoto      string     `json:"ProfilePhoto"`
	ChannelPhoto      string     `json:"ChannelPhoto"`
	Email             string     `json:"email"`
	HashedPassword    string     `json:"-"`
	ApiKey            string     `json:"-"`
	EmailVerification string     `json:"-"`
	IsVerified        bool       `json:"-"`
	AboutMe           string     `json:"about_me"`
	Created_at        *time.Time `json:"created_at"`
	Updated_at        *time.Time `json:"updated_at,omitempty"`
	/// extra columns from other tables
	TotalViews          int64 `json:"total_views"`
	Subscribers         int64 `json:"subscribers_count"`
	Subscribings        int64 `json:"subscrings_count"`
	UploadCount         int64 `json:"upload_count"`
	IsCurrentUserSubbed bool  `json:"is_subscribed"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
	// return password, nil
}

func generateSecureToken(length int) string {
	buff := make([]byte, int(math.Ceil(float64(length)/2)))
	if _, err := rand.Read(buff); err != nil {
		log.Panicln(err.Error())
	}
	str := hex.EncodeToString(buff)
	return str[:length]
}

//// GetUsers

func SearchUsers(limit int, offset int, search string, sortType helper.SortType) (int64, []User, error) {
	st, err := helpers.SortToString(sortType)
	if err != nil {
		return 0, nil, err
	}
	var totalPages int64
	{
		query := `
		SELECT 
			COUNT(*)
		FROM 
			users U
		WHERE U.channel_name LIKE concat('%',$1::VARCHAR,'%') OR U.username LIKE concat('%',$1::VARCHAR,'%');
		`
		rows, err := db.Db.Query(query, search)
		if err != nil {
			return 0, nil, err
		}
		defer rows.Close()
		if !rows.Next() {
			return 0, nil, nil
		}
		err = rows.Scan(&totalPages)
		if err != nil {
			return 0, nil, err
		}
		totalPages = int64(math.Ceil(float64(totalPages) / float64(limit)))
	}
	query := `
	SELECT 
		U.channel_name,
		U.username,
		U.created_at,
		(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
		(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
		COALESCE(U.profile_photo,'') AS profile_photo,
		EXISTS(SELECT id FROM followings WHERE follower_id=$1 AND following_id=U.id) AS is_subbed
	FROM 
		users AS U
	WHERE  U.channel_name LIKE concat('%',$1::VARCHAR,'%') OR U.username LIKE concat('%',$1::VARCHAR,'%')
	ORDER BY ` + st + `
	LIMIT $2 OFFSET $3;
	`

	rows, err := db.Db.Query(query, search, limit, offset)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ChannelName, &u.Username, &u.Created_at, &u.TotalViews, &u.Subscribers, &u.ProfilePhoto)
		if err != nil {
			return 0, nil, err
		}
		users = append(users, u)
	}
	return totalPages, users, nil /// users will be nil if no user can be find
}

func AuthSearchUsers(limit int, offset int, search string, sortType helper.SortType, userId int64) (int64, []User, error) {
	st, err := helpers.SortToString(sortType)
	if err != nil {
		return 0, nil, err
	}
	var totalPages int64
	{
		query := `
		SELECT 
			COUNT(*)
		FROM 
			users U
		WHERE U.channel_name LIKE concat('%',$1::VARCHAR,'%') OR U.username LIKE concat('%',$1::VARCHAR,'%');
		`
		rows, err := db.Db.Query(query, search)
		if err != nil {
			return 0, nil, err
		}
		defer rows.Close()
		if !rows.Next() {
			return 0, nil, nil
		}
		err = rows.Scan(&totalPages)
		if err != nil {
			return 0, nil, err
		}
		totalPages = int64(math.Ceil(float64(totalPages) / float64(limit)))
	}
	query := `
	SELECT 
		U.channel_name,
		U.username,
		U.created_at,
		(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
		(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
		COALESCE(U.profile_photo,'') AS profile_photo,
		EXISTS(SELECT id FROM followings WHERE follower_id=$1 AND following_id=U.id) AS is_subbed
	FROM 
		users AS U
	WHERE  U.channel_name LIKE concat('%',$2::VARCHAR,'%') OR U.username LIKE concat('%',$2::VARCHAR,'%')
	ORDER BY ` + st + `
	LIMIT $3 OFFSET $4;
	`

	rows, err := db.Db.Query(query, userId, search, limit, offset)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ChannelName, &u.Username, &u.Created_at, &u.TotalViews, &u.Subscribers, &u.ProfilePhoto, &u.IsCurrentUserSubbed)
		if err != nil {
			return 0, nil, err
		}
		users = append(users, u)
	}
	return totalPages, users, nil /// users will be nil if no user can be find
}

func GetUsers(limit int, offset int, sortType helper.SortType) (int64, []User, error) {
	st, err := helpers.SortToString(sortType)
	if err != nil {
		return 0, nil, err
	}
	var totalPages int64
	{
		query := `
		SELECT 
			COUNT(*)
		FROM 
			users U;
		`
		rows, err := db.Db.Query(query)
		if err != nil {
			return 0, nil, err
		}
		defer rows.Close()
		if !rows.Next() {
			return 0, nil, nil
		}
		err = rows.Scan(&totalPages)
		if err != nil {
			return 0, nil, err
		}
		totalPages = int64(math.Ceil(float64(totalPages) / float64(limit)))
	}
	query := `
	SELECT 
		U.channel_name,
		U.username,
		U.created_at,
		(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
		(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
		COALESCE(U.profile_photo,'') AS profile_photo
	FROM 
		users AS U
	ORDER BY ` + st + `
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Db.Query(query, limit, offset)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ChannelName, &u.Username, &u.Created_at, &u.TotalViews, &u.Subscribers, &u.ProfilePhoto)
		if err != nil {
			return 0, nil, err
		}
		users = append(users, u)
	}
	return totalPages, users, nil /// users will be nil if no user can be find
}

func AuthGetUsers(limit int, offset int, sortType helper.SortType, userId int64) (int64, []User, error) {
	st, err := helpers.SortToString(sortType)
	if err != nil {
		return 0, nil, err
	}
	var totalPages int64
	{
		query := `
		SELECT 
			COUNT(*)
		FROM 
			users U;
		`
		rows, err := db.Db.Query(query)
		if err != nil {
			return 0, nil, err
		}
		defer rows.Close()
		if !rows.Next() {
			return 0, nil, nil
		}
		err = rows.Scan(&totalPages)
		if err != nil {
			return 0, nil, err
		}
		totalPages = int64(math.Ceil(float64(totalPages) / float64(limit)))
	}
	query := `
	SELECT 
		U.channel_name,
		U.username,
		U.created_at,
		(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
		(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
		COALESCE(U.profile_photo,'') AS profile_photo,
		EXISTS(SELECT id FROM followings WHERE follower_id=$1 AND following_id=U.id) AS is_subbed
	FROM 
		users AS U
	ORDER BY ` + st + `
	LIMIT $2 OFFSET $3;
	`

	rows, err := db.Db.Query(query, userId, limit, offset)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ChannelName, &u.Username, &u.Created_at, &u.TotalViews, &u.Subscribers, &u.ProfilePhoto, &u.IsCurrentUserSubbed)
		if err != nil {
			return 0, nil, err
		}
		users = append(users, u)
	}
	return totalPages, users, nil /// users will be nil if no user can be find
}

func GetUserByApikey(Apikey string) (*User, error) {
	query := "SELECT id,email,hashed_password,channel_name,username,email_verification IS NULL,api_key,COALESCE(profile_photo,'') AS profile_photo,COALESCE(about_me,'') AS about_me FROM users WHERE api_key=$1 ;"
	rows, err := db.Db.Query(query, Apikey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, NewModelError("incorrect api key", 400)
	}
	var u User
	err = rows.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.ChannelName, &u.Username, &u.IsVerified, &u.ApiKey, &u.ProfilePhoto, &u.AboutMe)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// func GetUserById(id int64) *User {
// 	query := `
// 	SELECT
// 		id,
// 		channel_name,
// 		username,
// 		created_at,
// 		-- (SELECT COUNT(*) FROM views WHERE user_id=users.id)AS views_count,
// 		(SELECT COUNT(*) FROM medias WHERE user_id=users.id) AS upload_count,
// 		(SELECT COUNT(*) FROM followings WHERE followings.following_id=users.id) AS subscribers,
// 		about_me
// 	FROM users WHERE id=$1;
// 	`
// 	rows, err := db.Db.Query(query, id)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return nil
// 	}
// 	var u User
// 	err = rows.Scan(&u.Id, &u.ChannelName, &u.Username, &u.Created_at, &u.UploadCount, &u.Subscribers, &u.AboutMe)
// 	if err != nil {
// 		log.Panicln(err.Error())
// 	}
// 	return &u
// }

//// GetUser

func GetUserByUsername(username string) (*User, error) {
	query := `
	SELECT 
		U.channel_name,
		U.username,
		U.email,
		email_verification IS NULL,
		COALESCE(U.profile_photo,'') AS profile_photo,
		COALESCE(U.channel_photo,'') AS channel_photo,
		COALESCE(U.about_me,'') AS about_me
	FROM users U WHERE U.username=$1;
	`
	rows, err := db.Db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, NewModelError(fmt.Sprintf("user with username:`%s` not found! ", username), 404)
	}
	var u User
	err = rows.Scan(&u.ChannelName, &u.Username, &u.Email, &u.IsVerified, &u.ProfilePhoto, &u.ChannelPhoto, &u.AboutMe)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func AuthGetUserByUsername(username string, userId int64) (*User, error) {
	query := `
	SELECT 
		U.id,
		U.channel_name,
		U.username,
		U.email,
		U.created_at,
		U.updated_at,
		(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
		(SELECT COUNT(*) FROM medias WHERE user_id=U.id) AS upload_count,
		(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
		(SELECT COUNT(*) FROM followings WHERE followings.follower_id=U.id) AS subscrings,
		COALESCE(U.profile_photo,'') AS profile_photo,
		COALESCE(U.channel_photo,'') AS channel_photo,
		EXISTS(SELECT id FROM followings WHERE follower_id=$1 AND following_id=U.id) AS is_subbed,
		COALESCE(U.about_me,'') AS about_me
	FROM users U WHERE U.username=$2;
	`
	rows, err := db.Db.Query(query, userId, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, NewModelError(fmt.Sprintf("user with username:`%s` not found! ", username), 404)
	}
	var u User
	err = rows.Scan(
		&u.Id, &u.ChannelName, &u.Username, &u.Email, &u.Created_at, &u.Updated_at, &u.TotalViews,
		&u.UploadCount, &u.Subscribers, &u.Subscribings, &u.ProfilePhoto,
		&u.ChannelPhoto, &u.IsCurrentUserSubbed, &u.AboutMe)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// sign in
func GetUserByUsernameAndPassword(usernameOrEmail string, password string) (*User, error) {
	query := `
		SELECT 
			U.channel_name,
			U.email,
			U.username,
			U.api_key,
			U.hashed_password,
			U.created_at,
			U.updated_at,
			(SELECT COUNT(*) FROM views V JOIN medias M ON V.media_id=M.id WHERE V.user_id=U.id) AS views_count,
			(SELECT COUNT(*) FROM medias WHERE user_id=U.id) AS upload_count,
			(SELECT COUNT(*) FROM followings WHERE followings.following_id=U.id) AS subscribers,
			(SELECT COUNT(*) FROM followings WHERE followings.follower_id=U.id) AS subscrings,
			COALESCE(U.profile_photo,'') AS profile_photo,
			COALESCE(U.channel_photo,'') AS channel_photo,
			COALESCE(U.about_me,'') AS about_me
		FROM 
			users U
		WHERE (U.username=$1 OR U.email=$1);
	`
	rows, err := db.Db.Query(query, usernameOrEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, NewModelError("Invalid email/username or password", 400)
	}
	var u User
	err = rows.Scan(&u.ChannelName, &u.Email, &u.Username, &u.ApiKey, &u.HashedPassword,
		&u.Created_at, &u.Updated_at, &u.TotalViews, &u.UploadCount, &u.Subscribers, &u.Subscribings,
		&u.ProfilePhoto, &u.ChannelPhoto, &u.AboutMe)
	println(u.HashedPassword, password)
	if err = bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)); err != nil {
		return nil, NewModelError("Invalid email/username or password", 400)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

/// CreateUser

func CreateUser(email string, channelName string, username string, password string, aboutMe string) (*User, error) {
	query := "INSERT INTO users (email,username,hashed_password,email_verification,api_key,channel_name,about_me,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	apiKey := generateSecureToken(128)
	emailVerification := generateSecureToken(32)
	t := time.Now()
	res, err := db.Db.Exec(query, email, username, hashedPassword, emailVerification, apiKey, channelName, aboutMe, t)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return nil, &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return nil, err
	}
	id, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &User{
		Id:                id,
		Username:          username,
		Email:             email,
		ChannelName:       channelName,
		ApiKey:            apiKey,
		EmailVerification: emailVerification,
	}, nil
}

//// EditUser

func EditUserInfo(id int64, email string, username string, aboutMe string, channelName string, hashedPassword string) error {
	query := "UPDATE users SET email=$1,username=$2,updated_at=$3,about_me=$4,channel_name=$5,hashed_password=$6 WHERE id=$7;"
	res, err := db.Db.Exec(query, email, username, time.Now(), aboutMe, channelName, hashedPassword, id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("authenticated user `id` was incorrect")
	}
	return nil
}

func VerifyUserEmail(username, emailCode string) error {
	id := 0
	var isNull bool
	query := "SELECT id,email_verification IS NULL FROM users WHERE username=$1;"
	res1, err := db.Db.Query(query, username)
	if err != nil {
		return err
	}
	if !res1.Next() {
		return NewModelError("Invalid username!", 400)
	}
	err = res1.Scan(&id, &isNull)
	if err != nil {
		return err
	}
	if isNull {
		return NewModelError("email already verified!", 400)
	}

	query = "UPDATE users SET email_verification=NULL WHERE id=$1 AND email_verification=$2;"
	res2, err := db.Db.Exec(query, id, emailCode)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return err
	}
	n, err := res2.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return NewModelError("Invalid email verfication code!", 400)
	}
	return nil
}

func SetEmailCode(email string) (string, error) {
	emailCode := generateSecureToken(32)
	// t := time.Now()
	query := "UPDATE users SET email_verification=$1 WHERE email=$2;"
	res, err := db.Db.Exec(query, emailCode, email)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return "", &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return "", err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if n == 0 {
		return "", fmt.Errorf("incorrect email") // never should happen
	}
	return emailCode, nil
}

func SetUserProfilePhoto(id int64, photo string) error {
	query := "UPDATE users SET profile_photo=$1,updated_at=$2 WHERE id=$3;"
	res, err := db.Db.Exec(query, photo, time.Now(), id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("authenticated user `id` was incorrect")
		// panic("authenticated user `id` was incorrect !!!") // never should happen
	}
	return nil
}

func SetUserChannelPhoto(id int64, photo string) error {
	query := "UPDATE users SET channel_photo=$1,updated_at=$2 WHERE id=$3;"
	res, err := db.Db.Exec(query, photo, time.Now(), id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("authenticated user `id` was incorrect")
		// panic("authenticated user `id` was incorrect !!!") // never should happen
	}
	return nil
}

func EditUserApikey(id int64) (*User, error) {
	query := "UPDATE users SET api_key=$1,updated_at=$2 WHERE id=$3;"
	apikey := generateSecureToken(128)
	res, err := db.Db.Exec(query, apikey, time.Now(), id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("%v\n", err.Detail)
			if err.Code == "23505" {
				return nil, &ModelError{Message: err.Detail, Status: 400}
			}
		}
		return nil, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, fmt.Errorf("authenticated user `id` was incorrect")
		// panic("authenticated user `id` was incorrect !!!") // never should happen
	}
	return &User{ApiKey: apikey}, nil
}

//// DeleteUser

func DeleteUser(id int64) error {
	query := "DELETE FROM users WHERE id=$1;"
	res, err := db.Db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("authenticated user `id` was incorrect")
	}
	return nil
}
