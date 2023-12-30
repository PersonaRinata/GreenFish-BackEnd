namespace go api
include "base.thrift"

struct qingyu_user_register_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // Password, up to 32 characters
}

struct qingyu_user_register_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: i64 user_id // User id
    4: string token // User authentication token
}

struct qingyu_user_login_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // Password, up to 32 characters
}

struct qingyu_user_login_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: i64 user_id // User id
    4: string token // User authentication token
}

struct qingyu_user_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_user_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: base.User user // User Information
}

struct qingyu_avatar_change_request {
    1: string token(api.form="token") // User authentication token
}

struct qingyu_avatar_change_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_feed_request {
    1: i64 latest_time(api.query="latest_time") // Optional parameter, limit the latest submission timestamp of the returned video, accurate to seconds, and leave it blank to indicate the current time
    2: string token(api.query="token") // Optional parameter, login user settings
}

struct qingyu_feed_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Video> video_list // Video list
    4: i64 next_time // In the video returned this time, publish the earliest time as the latest_time in the next request
}

struct qingyu_publish_action_request {
    1: string token(api.form="token") // User authentication token
    2: string title(api.form="title", api.vd="len($)>0 && len($)<33") // Video title
}

struct qingyu_publish_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_publish_list_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_publish_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Video> video_list // List of videos posted by the user
}

struct qingyu_favorite_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id") // Video Id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-like, 2-unlike
}

struct qingyu_favorite_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_favorite_list_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_favorite_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Video> video_list // List of videos posted by the user
}

struct qingyu_comment_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id") // Video Id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-like, 2-unlike
    4: string comment_text(api.query="comment_text") // The content of the comment filled by the user, used when action_type=1
    5: i64 comment_id(api.query="comment_id") // The comment id to be deleted is used when action_type=2
}

struct qingyu_comment_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: base.Comment comment // The comment successfully returns the comment content, no need to re-pull the entire list
}

struct qingyu_comment_list_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id") // Video Id
}

struct qingyu_comment_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Comment> comment_list // List of comments
}

struct qingyu_relation_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 to_user_id(api.query="to_user_id") // The other party's user id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-Follow, 2-Unfollow
}

struct qingyu_relation_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_relation_follow_list_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_relation_follow_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.User> user_list // List of user information
}

struct qingyu_relation_follower_list_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_relation_follower_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.User> user_list // List of user information
}

struct qingyu_relation_friend_list_request {
    1: i64 user_id(api.query="user_id") // User id
    2: string token(api.query="token") // User authentication token
}

struct qingyu_relation_friend_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.FriendUser> user_list,     // List of user information
}

struct qingyu_message_chat_request {
    1: string token(api.query="token") // User authentication token
    2: i64 to_user_id(api.query="to_user_id") // The other party's user id
    3: i64 pre_msg_time(api.query="pre_msg_time")// The time of time of last latest message
}

struct qingyu_message_chat_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Message> message_list // Message list
}

struct qingyu_message_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 to_user_id(api.query="to_user_id") // The other party's user id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1- Send a message
    4: string content(api.query="content", api.vd="len($)>0 && len($)<255") // Message content
}

struct qingyu_message_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_issue_list_update_request {
    1: string token(api.query="token") // User authentication token
    2: i64 userID
    3: base.IssueList issueList(api.query="issue_list")
}

struct qingyu_issue_list_update_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_issue_list_get_request {
    1: string token(api.query="token") // User authentication token
    2: i64 user_id(api.query="user_id")
}

struct qingyu_issue_list_get_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: base.IssueList issue_list
}

struct qingyu_search_user_request{
    1: string token(api.query="token") // User authentication token
    2: string content,  // content for search
    3: i64 offset
    4: i64 num
}

struct qingyu_search_user_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.User> user_list                     // User Information
}

struct qingyu_search_video_request{
    1: string token(api.query="token") // User authentication token
    2: string content,  // content for search
    3: i64 offset
    4: i64 num
}

struct qingyu_search_video_response{
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<base.Video> video_list                     // User Information
}

struct qingyu_aigc_question_request {
    1: string token(api.query="token") // User authentication token
    2: string content
}

struct qingyu_aigc_question_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string msg
}

struct qingyu_aigc_issueList_request {
    1: string token(api.query="token") // User authentication token
}

struct qingyu_aigc_issueList_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string msg
}

struct qingyu_aigc_choose_word_request {
    1:string token(api.query="token") // User authentication token
    2:string content
}

struct qingyu_aigc_choose_word_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string msg
}

struct qingyu_aigc_doctor_analyse_request {
    1:string token(api.query="token") // User authentication token
    2:list<string> content
}

struct qingyu_aigc_doctor_analyse_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string msg
}

struct qingyu_aigc_get_history_request {
    1:string token(api.query="token") // User authentication token
}

struct qingyu_aigc_get_history_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<string> msg
}

struct qingyu_judge_doctor_request {
    1:string token(api.query="token") // User authentication token
}

struct qingyu_judge_doctor_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string department
}

struct qingyu_add_doctor_request {
    1:string token(api.query="token") // User authentication token
    2:string department
}

struct qingyu_add_doctor_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct qingyu_aigc_recommend_docotor_request {
    1:string token(api.query="token") // User authentication token
    2:string content
}

struct qingyu_aigc_recommend_docotor_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: string department
    4: list<base.User> doctor_list
}

service ApiService {
    qingyu_user_register_response Register(1: qingyu_user_register_request req)(api.post="/qingyu/user/register/");
    qingyu_user_login_response Login(1: qingyu_user_login_request req)(api.post="/qingyu/user/login/");
    qingyu_user_response GetUserInfo(1: qingyu_user_request req)(api.get="/qingyu/user/");
    qingyu_avatar_change_response ChangeAvatar(1:qingyu_avatar_change_request req)(api.post="/qingyu/user/avatar")
    qingyu_search_user_response SearchUserList(1: qingyu_search_user_request req)(api.post="/qingyu/user/search");
    qingyu_judge_doctor_response JudgeDoctor(1:qingyu_judge_doctor_request req)(api.get="/qingyu/user/judge/doctor");
    qingyu_add_doctor_response AddDoctor(1:qingyu_add_doctor_request req)(api.post="/qingyu/user/add/doctor");

    qingyu_feed_response Feed (1: qingyu_feed_request req)(api.get="/qingyu/feed/");
    qingyu_publish_action_response PublishVideo (1: qingyu_publish_action_request req)(api.post="/qingyu/publish/action/");
    qingyu_publish_list_response VideoList (1: qingyu_publish_list_request req)(api.get="/qingyu/publish/list/");
    qingyu_search_video_response SearchVideoList(1: qingyu_search_video_request req)(api.post="/qingyu/video/search");

    qingyu_favorite_action_response Favorite(1: qingyu_favorite_action_request req)(api.post="/qingyu/favorite/action/");
    qingyu_favorite_list_response FavoriteList(1: qingyu_favorite_list_request req)(api.get="/qingyu/favorite/list/");
    qingyu_comment_action_response Comment(1: qingyu_comment_action_request req)(api.post="/qingyu/comment/action/");
    qingyu_comment_list_response CommentList(1: qingyu_comment_list_request req)(api.get="/qingyu/comment/list/");

    qingyu_relation_action_response Action(1: qingyu_relation_action_request req)(api.post="/qingyu/relation/action/");
    qingyu_relation_follow_list_response FollowingList(1: qingyu_relation_follow_list_request req)(api.get="/qingyu/relation/follow/list/");
    qingyu_relation_follower_list_response FollowerList(1: qingyu_relation_follower_list_request req)(api.get="/qingyu/relation/follower/list/");
    qingyu_relation_friend_list_response FriendList(1: qingyu_relation_friend_list_request req)(api.get="/qingyu/relation/friend/list/");

    qingyu_message_chat_response ChatHistory(1: qingyu_message_chat_request req)(api.get="/qingyu/message/chat/");
    qingyu_message_action_response SentMessage(1: qingyu_message_action_request req)(api.post="/qingyu/message/action/");

    qingyu_issue_list_update_response UpdateIssueList(1:qingyu_issue_list_update_request req)(api.post="/qingyu/issuelist/action/");
    qingyu_issue_list_get_response GetIssueList(1:qingyu_issue_list_get_request req)(api.get="/qingyu/issuelist/");

    qingyu_aigc_question_response AIGCAskQuestion(1:qingyu_aigc_question_request req)(api.post="/qingyu/aigc/question/");
    qingyu_aigc_issueList_response AIGCIssueList(1:qingyu_aigc_issueList_request req)(api.get="/qingyu/aigc/issuelist");
    qingyu_aigc_choose_word_response AIGCChooseWord(1:qingyu_aigc_choose_word_request req)(api.post="/qingyu/aigc/word/");
    qingyu_aigc_doctor_analyse_response AIGCDoctorAnalyse(1:qingyu_aigc_doctor_analyse_request req)(api.post="/qingyu/aigc/doctor");
    qingyu_aigc_get_history_response AIGCGetHistory(1:qingyu_aigc_get_history_request req)(api.get="/qingyu/aigc/history")
    qingyu_aigc_recommend_docotor_response AIGCRecommendDoctor(1:qingyu_aigc_recommend_docotor_request req)(api.get="/qingyu/aigc/recommend/doctor")
}