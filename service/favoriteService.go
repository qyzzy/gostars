package service

import (
	"gostars/utils"
	"gostars/utils/code"
	"strconv"
	"strings"
)

type FavoriteService struct {
}

func (favoriteService *FavoriteService) FavoriteCount(articleID int) {

}

func (favoriteService *FavoriteService) FavoriteAction(userID, articleID, action int) int {
	strUserID := strconv.Itoa(userID)
	strArticleID := strconv.Itoa(articleID)
	// message send to RabbitMQ
	message := strings.Builder{}
	message.WriteString(strUserID)
	message.WriteString(" ")
	message.WriteString(strArticleID)

	// like action
	if action == utils.LikeAction {
		if errCode := likeService.ExistsRedisLikeUserID(strUserID); errCode == code.SUCCESS {
			if errCode := likeService.SAddRedisLikeUserID(strUserID, articleID); errCode != code.SUCCESS {
				return errCode
			} else {
				RmqLikeAdd.Publish(message.String())
			}
		} else {
			if errCode := likeService.SAddRedisLikeUserID(strUserID, utils.DefaultRedisValue); errCode != code.SUCCESS {
				likeService.DelRedisLikeUserID(strUserID)
				return errCode
			}

			if errCode := likeService.ExpireRedisLikeUserID(strUserID); errCode != code.SUCCESS {
				likeService.DelRedisLikeUserID(strUserID)
				return errCode
			}
			articleList, errCode := likeService.GetLikeArticleList(userID)
			if errCode != code.SUCCESS {
				return errCode
			}

			// for range articleList, prevent dirty reads
			// if it failed, delete strUserID key in redis
			// ** maintain data consistency between mysql and redis **
			for _, likeArticleID := range articleList {
				sAddErrCode := likeService.SAddRedisLikeUserID(strUserID, likeArticleID)
				if sAddErrCode != code.SUCCESS {
					likeService.DelRedisLikeUserID(strUserID)
					return sAddErrCode
				}
			}

			if errCode := likeService.SAddRedisLikeUserID(strUserID, articleID); errCode != code.SUCCESS {
				return errCode
			} else {
				RmqLikeAdd.Publish(message.String())
			}
		}
		if errCode := likeService.ExistsRedisLikeArticleID(strArticleID); errCode == code.SUCCESS {
			if errCode = likeService.SAddRedisLikeArticleID(strArticleID, userID); errCode != code.SUCCESS {
				return errCode
			}
		} else {
			if errCode = likeService.SAddRedisLikeArticleID(strArticleID, userID); errCode != code.SUCCESS {
				likeService.DelRedisLikeArticleID(strArticleID)
				return errCode
			}

			// set expire time
			if errCode = likeService.ExpireRedisLikeArticleID(strArticleID); errCode != code.SUCCESS {
				likeService.DelRedisLikeArticleID(strArticleID)
				return errCode
			}

			userList, errCode := likeService.GetLikeUserIDList(articleID)
			if errCode != code.SUCCESS {
				return errCode
			}

			// for range userList, prevent dirty reads
			// if it failed, delete strArticleID key in redis
			// ** maintain data consistency between mysql and redis **
			for _, likeUserID := range userList {
				sAddErrCode := likeService.SAddRedisLikeArticleID(strArticleID, likeUserID)
				if sAddErrCode != code.SUCCESS {
					likeService.DelRedisLikeArticleID(strArticleID)
					return sAddErrCode
				}
			}

			if errCode = likeService.SAddRedisLikeArticleID(strArticleID, userID); errCode != code.SUCCESS {
				return errCode
			}
		}
	} else { // unlike action
		if errCode := likeService.ExistsRedisLikeUserID(strUserID); errCode == code.SUCCESS {
			// remove articleID in strUserID before sAdd articleID
			if errCode = likeService.SRemRedisLikeUserID(strUserID, articleID); errCode != code.SUCCESS {
				return errCode
			} else {
				RmqLikeDel.Publish(message.String())
			}
		} else {
			if errCode := likeService.SAddRedisLikeUserID(strUserID, utils.DefaultRedisValue); errCode != code.SUCCESS {
				likeService.DelRedisLikeUserID(strUserID)
				return errCode
			}
			if errCode := likeService.ExpireRedisLikeUserID(strUserID); errCode != code.SUCCESS {
				likeService.DelRedisLikeUserID(strUserID)
				return errCode
			}

			articleList, errCode := likeService.GetLikeArticleList(userID)
			if errCode != code.SUCCESS {
				return errCode
			}

			for _, likeArticleID := range articleList {
				sAddErrCode := likeService.SAddRedisLikeUserID(strUserID, likeArticleID)
				if sAddErrCode != code.SUCCESS {
					likeService.DelRedisLikeUserID(strUserID)
					return sAddErrCode
				}
			}

			if errCode = likeService.SRemRedisLikeUserID(strUserID, articleID); errCode != code.SUCCESS {
				return errCode
			} else {
				RmqLikeDel.Publish(message.String())
			}
		}

		if errCode := likeService.ExistsRedisLikeArticleID(strArticleID); errCode == code.SUCCESS {
			if errCode = likeService.SRemRedisLikeArticleID(strArticleID, userID); errCode != code.SUCCESS {
				return errCode
			}
		} else {
			if errCode = likeService.SAddRedisLikeArticleID(strArticleID, utils.DefaultRedisValue); errCode != code.SUCCESS {
				likeService.DelRedisLikeArticleID(strArticleID)
				return errCode
			}

			if errCode = likeService.ExpireRedisLikeArticleID(strArticleID); errCode != code.SUCCESS {
				likeService.DelRedisLikeArticleID(strArticleID)
				return errCode
			}

			userList, errCode := likeService.GetLikeUserIDList(articleID)
			if errCode != code.SUCCESS {
				return errCode
			}

			for _, likeUserID := range userList {
				sAddErrCode := likeService.SAddRedisLikeArticleID(strArticleID, likeUserID)
				if sAddErrCode != code.SUCCESS {
					likeService.DelRedisLikeArticleID(strArticleID)
					return sAddErrCode
				}
			}

			if errCode = likeService.SRemRedisLikeArticleID(strArticleID, userID); errCode != code.SUCCESS {
				return errCode
			}
		}
	}
	return code.SUCCESS
}

func (favoriteService *FavoriteService) GetFavoriteList(userID, curID int) {

}
