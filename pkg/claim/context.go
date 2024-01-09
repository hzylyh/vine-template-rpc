package claim

import (
	"context"
	"vine-template-rpc/pkg/vjwt"
)

func GetUserInfoFromCtx(ctx context.Context) *vjwt.Claims {
	return ctx.Value("claims").(*vjwt.Claims)
}
