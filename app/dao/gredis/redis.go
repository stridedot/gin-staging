package gredis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go_code/gintest/bootstrap/gredis"
	"strconv"
	"time"
)

const (
	OneWeekExpired = 7 * 24 * 3600
	ScorePerVoted  = 432
)

// ZStorePostToRedis 添加帖子发布时间、添加到社区中
func ZStorePostToRedis(c *gin.Context, postID int64, communityID int64) {
	// 记录帖子发帖时间
	_, _ = gredis.Rdb.ZAdd(
		c,
		KeyPostTimeZSet,
		&redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: postID,
		}).Result()

	// 记录帖子发帖分数
	_, _ = gredis.Rdb.ZAdd(
		c,
		KeyPostScoreZSet,
		&redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: postID,
		}).Result()

	// 将 post_id 保存到社区中
	_, _ = gredis.Rdb.SAdd(
		c,
		KeyPostsForCommunity+strconv.FormatInt(communityID, 10),
		postID,
	).Result()
}

// GetPostCreatedAt 查询帖子的发布时间
func GetPostCreatedAt(c *gin.Context, postID string) float64 {
	return gredis.Rdb.ZScore(c, KeyPostTimeZSet, postID).Val()
}

// GetPostVoted 查询帖子指定用户的投票记录
func GetPostVoted(c *gin.Context, userID int64, postID string) float64 {
	return gredis.Rdb.ZScore(
		c,
		KeyPostVotedZSet+postID,
		strconv.FormatInt(userID, 10),
	).Val()
}

// ZIncrPostScore 更新帖子的分数
func ZIncrPostScore(c *gin.Context, postID string, score float64) (float64, error) {
	return gredis.Rdb.ZIncrBy(
		c,
		KeyPostScoreZSet,
		score,
		postID,
	).Result()
}

// ZRemovePostVoted 移除帖子的投票用户
func ZRemovePostVoted(c *gin.Context, postID string, userID int64) int64 {
	return gredis.Rdb.ZRem(
		c,
		KeyPostVotedZSet+postID,
		strconv.FormatInt(userID, 10),
	).Val()
}

// ZSavePostVoted 添加或更新帖子投票记录
func ZSavePostVoted(c *gin.Context, postID string, userID int64, score float64) (int64, error) {
	return gredis.Rdb.ZAdd(
		c,
		KeyPostVotedZSet+postID,
		&redis.Z{
			Score:  score,
			Member: userID,
		}).Result()
}

// GetPostVotes 帖子投赞成票的成员数量
func GetPostVotes(c *gin.Context, postID string) (int64, error) {
	return gredis.Rdb.ZCount(
		c,
		KeyPostVotedZSet+postID,
		"1",
		"1",
	).Result()
}
