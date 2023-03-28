package auth

//
//import (
//	"context"
//	"crypto/sha1"
//	"fmt"
//	"time"
//
//	"github.com/dgrijalva/jwt-go/v4"
//	"github.com/fokurly/streaky-backend/users_info_api/models"
//	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
//)
//
//type Authorizer struct {
//	repo postgre.Db
//
//	hashSalt       string
//	signingKey     []byte
//	expireDuration time.Duration
//}
//
//func NewAuthorizer(repo postgre.Db, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
//	return &Authorizer{
//		repo:           repo,
//		hashSalt:       hashSalt,
//		signingKey:     signingKey,
//		expireDuration: expireDuration,
//	}
//}
//
//func (a *Authorizer) SignUp(ctx context.Context, user *models.UserAuth) error {
//	// Create password hash
//	pwd := sha1.New()
//	pwd.Write([]byte(user.Password))
//	pwd.Write([]byte(a.hashSalt))
//	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
//
//	return a.repo.Insert(ctx, user)
//}
//
//func (a *Authorizer) SignIn(ctx context.Context, user *models.UserAuth) (string, error) {
//	pwd := sha1.New()
//	pwd.Write([]byte(user.Password))
//	pwd.Write([]byte(a.hashSalt))
//	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
//
//	user, err := a.repo.Get(ctx, user.Username, user.Password)
//	if err != nil {
//		return "", err
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
//			IssuedAt:  jwt.At(time.Now()),
//		},
//		Username: user.Username,
//	})
//
//	return token.SignedString(a.signingKey)
//}
