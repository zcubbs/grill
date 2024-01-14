package util

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"github.com/zcubbs/grill/internal/config"
	db "github.com/zcubbs/grill/internal/db/sqlc"
	"github.com/zcubbs/log"
	"github.com/zcubbs/x/password"
	"github.com/zcubbs/x/random"
)

func InitAdminUser(store db.Store, cfg config.Configuration) error {
	// Check if the admin user already exists.
	_, err := store.GetUserByUsername(context.Background(), "admin")
	if err == nil {
		return nil
	}

	var hashedPass string
	var randPass string

	// check if init password is set
	if cfg.InitAdminPassword != "" {
		hashedPass, err = password.Hash(cfg.InitAdminPassword)
		if err != nil {
			return err
		}
	} else {
		randPass = random.String(32)
		hashedPass, err = password.Hash(randPass)
		if err != nil {
			return err
		}
	}

	_, err = store.CreateUser(context.Background(), db.CreateUserParams{
		Username:       "admin",
		HashedPassword: hashedPass,
		FullName:       "Administrator",
		Email:          "",
		Role:           pb.Role_ROLE_ADMIN.String(),
	})
	if err != nil {
		errCode := ErrorCode(err)
		if errCode == ForeignKeyViolation || errCode == UniqueViolation {
			return nil
		}
		return err
	}

	if cfg.InitAdminPassword == "" && randPass != "" {
		log.Info("Admin user created",
			"username", "admin",
			"password", randPass)
	}

	return nil
}
