package code

const (
	SUCCESS = 200
	ERROR   = 500

	// User module errors
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	// Article module errors
	ErrorArticleNotExist = 2001

	// Category module errors
	ErrorCategoryNameUsed = 3001
	ErrorCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ErrorUsernameUsed:     "Username is already in use!",
	ErrorPasswordWrong:    "Wrong password!",
	ErrorUserNotExist:     "User does not exist!",
	ErrorTokenExist:       "Token does not exist!",
	ErrorTokenRuntime:     "Token expired!",
	ErrorTokenWrong:       "Wrong token!",
	ErrorTokenTypeWrong:   "Wrong token format!",
	ErrorUserNoRight:      "No permission!",
	ErrorArticleNotExist:  "Article not exist!",
	ErrorCategoryNameUsed: "Category name is already in use!",
	ErrorCategoryNotExist: "Category not exist!",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
