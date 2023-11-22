namespace go user

include "base.thrift"

struct qingyu_user_register_request {
    1: string username, // Username, up to 32 characters
    2: string password, // Password, up to 32 characters
}

struct qingyu_user_register_response {
    1: base.qingyu_base_response base_resp,
    2: i64 user_id,                         // User id
    3: string token,                        // User authentication token
}

struct qingyu_user_login_request {
    1: string username, // Username, up to 32 characters
    2: string password, // Password, up to 32 characters
}

struct qingyu_user_login_response {
    1: base.qingyu_base_response base_resp,
    2: i64 user_id,                         // User id
    3: string token,                        // User authentication token
}

struct qingyu_get_user_request {
    1: i64 viewer_id, // User id of viewer,set to zero when unclear
    2: i64 owner_id,  // User id of owner.
}

struct qingyu_get_user_response {
    1: base.qingyu_base_response base_resp,
    2: base.User user,                      // User Information
}

struct qingyu_batch_get_user_request {
    1: i64 viewer_id,       // User id of viewer,set to zero when unclear
    2: list<i64> owner_id_list, // User id list of info owners.
}

struct qingyu_batch_get_user_resonse {
    1: base.qingyu_base_response base_resp,
    2: list<base.User> user_list,
}

struct qingyu_get_relation_follow_list_request {
    1: i64 viewer_id, // User id of viewer,set to zero when unclear
    2: i64 owner_id,  // User id of owner.
}

struct qingyu_get_relation_follow_list_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.User> user_list,           // List of user information
}

struct qingyu_get_relation_follower_list_request {
    1: i64 viewer_id, // User id of viewer,set to zero when unclear
    2: i64 owner_id,  // User id of owner.
}

struct qingyu_get_relation_follower_list_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.User> user_list,           // List of user information
}

struct qingyu_get_relation_friend_list_request {
    1: i64 viewer_id, // User id of viewer,set to zero when unclear
    2: i64 owner_id,  // User id of owner.
}

struct qingyu_get_relation_friend_list_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.FriendUser> user_list,     // List of user information
}

struct qingyu_update_issue_list_request{
    1: i64 user_id,
    2: base.IssueList issue_list,
}

struct qingyu_update_issue_list_response{
    1: base.qingyu_base_response base_resp,
}

struct qingyu_get_issue_list_request{
    1: i64 user_id,
}

struct qingyu_get_issue_list_response{
    1: base.qingyu_base_response base_resp,
    2: base.IssueList issue_list
}

struct qingyu_search_user_request{
    1: i64 viewer_id, // User id of viewer,set to zero when unclear
    2: string content,  // content for search
    3: i64 offset
    4: i64 num
}

struct qingyu_search_user_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.User> user_list,                      // User Information
}


service UserService {
    qingyu_user_register_response Register(1: qingyu_user_register_request req),
    qingyu_user_login_response Login(1: qingyu_user_login_request req),
    qingyu_get_user_response GetUserInfo(1: qingyu_get_user_request req),
    qingyu_batch_get_user_resonse BatchGetUserInfo(1: qingyu_batch_get_user_request req),
    qingyu_get_relation_follow_list_response GetFollowList(1: qingyu_get_relation_follow_list_request req),
    qingyu_get_relation_follower_list_response GetFollowerList(1: qingyu_get_relation_follower_list_request req),
    qingyu_get_relation_friend_list_response GetFriendList(1: qingyu_get_relation_friend_list_request req),
    qingyu_update_issue_list_response UpdateIssueList(1:qingyu_update_issue_list_request req)
    qingyu_get_issue_list_response GetIssueList(1:qingyu_get_issue_list_request req)
    qingyu_search_user_response SearchUserList(1:qingyu_search_user_request req)
}