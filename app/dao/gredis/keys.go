package gredis

const (
	KeyPrefix            = "gintest:"
	KeyPostTimeZSet      = KeyPrefix + "posts:time"    // 帖子及发帖时间
	KeyPostScoreZSet     = KeyPrefix + "posts:score"   // 帖子及投票分数
	KeyPostVotedZSet     = KeyPrefix + "posts:voted:"  // 用户及投票类型，后接 post_id
	KeyPostsForCommunity = KeyPrefix + "community:"
)

// GetRedisKey 获取指定 key
func getRedisKey(key string) string {
	return KeyPrefix + key
}