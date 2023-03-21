package redis

const UserTokenKeyPrefix = "usertoken-"

func GetUserTokenKey(token string) string {
	return UserTokenKeyPrefix + token
}
