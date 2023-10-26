/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/26 21:42
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/26 21:42
 */

package middleware

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"vine-template-rpc/pkg/response"
	"vine-template-rpc/pkg/vjwt"
)

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrTokenInvalid = errors.New(response.TOKEN_INVALID_CODE, response.TOKEN_INVALID_MESSAGE, "Token is invalid")
	ErrWrongContext = errors.Unauthorized(response.TOKEN_INVALID_MESSAGE, "Wrong context for middleware")
)

func AuthMiddleware(enforcer *casbin.Enforcer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			withAuth := ""
			if header, ok := transport.FromServerContext(ctx); ok {
				token := header.RequestHeader().Get(authorizationKey)
				if token != "" {
					withAuth = "auth"
					claims, err := vjwt.ParseJwtToken(token)
					if err != nil {
						fmt.Println(err)
						return nil, ErrTokenInvalid
					}
					ctx = context.WithValue(ctx, "claims", claims)
				}

				// [role authFlag resource act]
				res, err := enforcer.Enforce("dd", withAuth, "read", header.Operation())
				if err != nil {
					fmt.Println(err)
					return nil, ErrTokenInvalid
				}

				if res {
					return handler(ctx, req)
				}
			}
			return nil, ErrTokenInvalid
		}
	}
}
