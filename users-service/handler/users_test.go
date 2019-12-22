package handler

import (
	"context"
	"github.com/lbrulet/GoMicroservices/users-service/database/postgreSQL"
	"github.com/lbrulet/GoMicroservices/users-service/models"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/lbrulet/GoMicroservices/users-service/database"
	users "github.com/lbrulet/GoMicroservices/users-service/proto/users"
	"github.com/sirupsen/logrus"
)

func TestUsers_Create(t *testing.T) {
	repository := initDB()
	type fields struct {
		ServiceName string
		Logger      *logrus.Entry
		Database    database.Repository
	}
	type args struct {
		ctx context.Context
		req *users.CreateRequest
		rsp *users.UserResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		before func()
	}{
		{
			name: "Create a new user",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.CreateRequest{
					Username: "username",
					Password: "ui",
					Email:    "exemple@exemple.com",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: false,
			before: func() {},
		},
		{
			name: "Try to create a new user without password",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.CreateRequest{
					Username: "username",
					Email:    "exemple@exemple.com",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: true,
			before: func() {},
		},
		{
			name: "Try to create a user with same username",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.CreateRequest{
					Username: "username",
					Password: "password",
					Email:    "exemple@exemple.com",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: true,
			before: func() {
				repository.CreateUser(models.User{
					Username: "username",
					Email:    "exemple@exemple.com",
					Password: "password",
				})
			},
		},
		{
			name: "Try to create a user that already exist",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.CreateRequest{
					Username: "username",
					Password: "password",
					Email:    "exemple@exemple.com",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: true,
			before: func() {
				repository.CreateUser(models.User{
					Username: "username",
					Email:    "exemple@exemple.com",
					Password: "password",
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// MIGRATE DATABASE
			tt.fields.Database.MigrateDatabase()

			// EXECUTE BEFORE TESTING
			tt.before()

			// TEST
			e := &Users{
				ServiceName: tt.fields.ServiceName,
				Logger:      tt.fields.Logger,
				Database:    tt.fields.Database,
			}
			if err := e.Create(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Users.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// RESET DATABASE
			tt.fields.Database.ResetDatabase()
		})
	}
}

func TestLogin_Create(t *testing.T) {
	repository := initDB()
	type fields struct {
		ServiceName string
		Logger      *logrus.Entry
		Database    database.Repository
	}
	type args struct {
		ctx context.Context
		req *users.LoginRequest
		rsp *users.UserResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		before func()
	}{
		{
			name: "Login with a user",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.LoginRequest{
					Username: "username",
					Password: "password",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: false,
			before: func() {
				hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
				repository.CreateUser(models.User{
					Username: "username",
					Email:    "exemple@exemple.com",
					Password: string(hashed),
				})
			},
		},
		{
			name: "Login with a user that does not exist",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.LoginRequest{
					Username: "username",
					Password: "password",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: true,
			before: func() {},
		},
		{
			name: "Try to login with a wrong password as a user",
			fields: fields{
				ServiceName: "Test",
				Logger: logrus.WithFields(logrus.Fields{
					"micro-service": "users-service",
				}),
				Database: repository,
			},
			args: args{
				ctx: context.Background(),
				req: &users.LoginRequest{
					Username: "username",
					Password: "wrong_password",
				},
				rsp: &users.UserResponse{
					User: &users.User{},
				},
			},
			wantErr: true,
			before: func() {
				hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
				repository.CreateUser(models.User{
					Username: "username",
					Email:    "exemple@exemple.com",
					Password: string(hashed),
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// MIGRATE DATABASE
			tt.fields.Database.MigrateDatabase()

			// EXECUTE BEFORE TESTING
			tt.before()

			// TEST
			e := &Users{
				ServiceName: tt.fields.ServiceName,
				Logger:      tt.fields.Logger,
				Database:    tt.fields.Database,
			}
			if err := e.Login(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Users.Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			// RESET DATABASE
			tt.fields.Database.ResetDatabase()
		})
	}
}

func initDB()  database.Repository {
	db, err := postgreSQL.PostgresConnexion()
	if err != nil {
		return nil
	}
	repository := postgreSQL.NewPostgresRepository(db)
	repository.MigrateDatabase()
	return repository
}
