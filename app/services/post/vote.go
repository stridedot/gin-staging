package post

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/dao"
	daoRedis "go_code/gintest/app/dao/gredis"
	validatorPost "go_code/gintest/app/validators/post"
	"math"
	"time"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加 432 分   86400/200  --> 200 张赞成票可以给你的帖子续一天

/* 投票的几种情况：
   direction=1 时，有 2 种情况：
   	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2  +432*2
   direction=0 时，有 2 种情况：
   	1. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  -432
   direction=-1 时，有 2 种情况：
   	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录  差值的绝对值：2  -432*2

   投票的限制：
   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
   	1. 到期之后将 redis 中保存的赞成票数及反对票数存储到 mysql 表中
   	2. 到期之后删除那个 KeyPostVotedZSet
*/

// VotePost 帖子投票
func VotePost(
	c *gin.Context,
	request *validatorPost.VotePostRequest,
) (err error) {
	// 1. 判断投票的时间限制
	createdAt := daoRedis.GetPostCreatedAt(c, request.PostID)
	if float64(time.Now().Unix())-createdAt >= daoRedis.OneWeekExpired {
		return dao.VoteTimeExpired
	}

	// 2. 更新帖子的分数
	var direction float64
	// 获取当前用户对该帖子的投票记录
	value := daoRedis.GetPostVoted(c, request.UserID, request.PostID)
	if float64(request.Direction) == value {
		return dao.VoteRepeated
	}
	if float64(request.Direction) > value {
		direction = 1
	} else {
		direction = -1
	}
	// 计算分数系数，取绝对值
	absValue := math.Abs(value - float64(request.Direction))

	// 3. 更新帖子相关数据：投票分数、用户投票记录
	_, err = daoRedis.ZIncrPostScore(c, request.PostID, direction * absValue * daoRedis.ScorePerVoted)
	if err != nil {
		return
	}
	if request.Direction == 0 {
		// 移除帖子的投票记录
		_ = daoRedis.ZRemovePostVoted(c, request.PostID, request.UserID)
	} else {
		// 添加或更新帖子的投票记录
		_, err = daoRedis.ZSavePostVoted(c, request.PostID, request.UserID, direction)
		if err != nil {
			return
		}
	}
	return nil
}
