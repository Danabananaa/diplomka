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
	_, err := db.Exec("delete from users")
	if err != nil {
		t.Fail()
	}
}
