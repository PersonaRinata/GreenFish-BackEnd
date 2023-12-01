// Code generated by hertz generator.

package api

import (
	"GreenFish/server/common/middleware"
	"GreenFish/server/service/api/config"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/gzip"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".jpg", ".mp4", ".png"})),
	}
}

func _qingyuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _comment0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.LogFeedInfo(),
	}
}

func _feed0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sentmessageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chathistoryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _action2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _videolistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _action3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action4Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followinglistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list4Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _register0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _issuelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getissuelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateissuelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action5Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _searchuserlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _searchvideolistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _videoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _changeavatarMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcdoctoranalyseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcgethistoryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcissuelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _questionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcaskquestionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _wordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aigcchoosewordMw() []app.HandlerFunc {
	// your code...
	return nil
}
