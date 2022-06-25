package code

const (
	SUCCESS = 200
	ERROR   = 500

	// User module errors
	ErrorUsernameUsed      = 1001
	ErrorPasswordWrong     = 1002
	ErrorUserNotExist      = 1003
	ErrorTokenExist        = 1004
	ErrorTokenRuntime      = 1005
	ErrorTokenWrong        = 1006
	ErrorTokenTypeWrong    = 1007
	ErrorUserNoRight       = 1008
	ErrorTokenCreateFailed = 1009
	ErrorTokenInBlacklist  = 1010

	// Article module errors
	ErrorArticleNotExist   = 2001
	ErrorLikeArticleFailed = 2002

	// Category module errors
	ErrorCategoryNameUsed = 3001
	ErrorCategoryNotExist = 3002

	ErrorRedisGetFailed    = 6001
	ErrorRedisSaveFailed   = 6002
	ErrorRedisDeleteFailed = 6003

	ErrorImageCreateFailed = 8001

	ErrorDataNotFound = 53011
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ErrorUsernameUsed:      "Username is already in use!",
	ErrorPasswordWrong:     "Wrong password!",
	ErrorUserNotExist:      "User does not exist!",
	ErrorTokenExist:        "Token does not exist!",
	ErrorTokenRuntime:      "Token expired!",
	ErrorTokenWrong:        "Wrong token!",
	ErrorTokenTypeWrong:    "Wrong token format!",
	ErrorUserNoRight:       "No permission!",
	ErrorArticleNotExist:   "Article not exist!",
	ErrorCategoryNameUsed:  "Category name is already in use!",
	ErrorCategoryNotExist:  "Category not exist!",
	ErrorRedisGetFailed:    "Redis get failed!",
	ErrorRedisSaveFailed:   "Redis save failed!",
	ErrorRedisDeleteFailed: "Redis delete failed!",
	ErrorTokenCreateFailed: "Token create failed!",
	ErrorTokenInBlacklist:  "Token in blacklist!",
	ErrorImageCreateFailed: "Image create failed",
	ErrorDataNotFound:      "Data could not found",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
