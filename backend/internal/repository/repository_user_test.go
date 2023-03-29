package repository

import (
	"context"
	"diplomka/internal/model"
	"diplomka/pkg/sqlite"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	log.Println("first")
	var err error
	db, err = sqlite.Connect("data_test.db")

	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func Test_user_AddUser(t *testing.T) {
	_, err := db.Exec("delete from sqlite_sequence;")
	if err != nil {
		t.Fail()
	}
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{

		{
			name: "Beka",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:     "Aldik",
					Surname:  "Gibadatov",
					Email:    "alike",
					Password: "qwqwewe",
				},
			},
			want: &model.User{
				ID:       1,
				Name:     "Aldik",
				Surname:  "Gibadatov",
				Email:    "alike",
				Password: "qwqwewe",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				DB: tt.fields.DB,
			}
			got, err := u.AddUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
	_, err = db.Exec("delete from users;")
	if err != nil {
		t.Fail()
	}
}

// func TestNewUserRepo(t *testing.T) {
// 	_, err := db.Exec("INSERT INTO users (name, surname, email, password) VALUES ('aldik','beka','sadds','asdasd')")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	type args struct {
// 		db *sqlx.DB
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *user
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewUserRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewUserRepo() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_user_GetUser(t *testing.T) {
	_, err := db.Exec("INSERT INTO users (name, surname, email, password) VALUES ('aldik','beka','alike','qwqwewe')")
	if err != nil {
		t.Fail()
	}
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Beka",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: &model.User{
				ID:       1,
				Name:     "aldik",
				Surname:  "beka",
				Email:    "alike",
				Password: "qwqwewe",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				DB: tt.fields.DB,
			}
			got, err := u.GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
	// _, err = db.Exec("delete from users")
	// if err != nil {
	// 	t.Fail()
	// }
	// _, err = db.Exec("delete from sqlite_sequence;")
	// if err != nil {
	// 	t.Fail()
	// }
}

// func Test_user_GetUserforAuth(t *testing.T) {
// 	type fields struct {
// 		DB *sqlx.DB
// 	}
// 	type args struct {
// 		ctx      context.Context
// 		email    string
// 		password string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *model.User
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := &user{
// 				DB: tt.fields.DB,
// 			}
// 			got, err := u.GetUserforAuth(tt.args.ctx, tt.args.email, tt.args.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("user.GetUserforAuth() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("user.GetUserforAuth() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
