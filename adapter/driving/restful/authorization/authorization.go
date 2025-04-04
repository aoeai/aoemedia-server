package authorization

import (
	repoUser "github.com/aoemedia-server/adapter/driven/repository/user"
	"github.com/aoemedia-server/domain/user"
	"github.com/gin-gonic/gin"
)

var (
	userRepo user.Repository
)

func init() {
	userRepo = repoUser.Inst()
}

// Authorization 用户认证信息
type Authorization struct {
	Token string
	// Missing true: 没有Token false: 有Token
	Missing bool
	UserId  int64
}

func NewAuth(ctx *gin.Context) Authorization {
	authToken := ctx.GetHeader("Authorization")
	missing := authToken == ""

	if missing {
		return Authorization{
			Token:   authToken,
			Missing: missing,
		}
	}

	return Authorization{
		Token:   authToken,
		Missing: missing,
		UserId:  userRepo.GetIdByToken(authToken),
	}
}
