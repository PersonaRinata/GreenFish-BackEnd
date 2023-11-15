namespace go interaction

include "base.thrift"

struct qingyu_favorite_action_request {
    1: i64 user_id,    // User Id
    2: i64 video_id,   // Video Id
    3: i8 action_type, // 1-like, 2-unlike
}

struct qingyu_favorite_action_response {
    1: base.qingyu_base_response base_resp,
}

struct qingyu_get_favorite_video_id_list_request {
    1: i64 user_id,
}

struct qingyu_get_favorite_video_id_list_response {
    1: base.qingyu_base_response base_resp,
    2: list<i64> video_id_list,                // List of videos posted by the user
}

struct qingyu_comment_action_request {
    1: i64 user_id,         // User Id
    2: i64 video_id,        // Video Id
    3: i8 action_type,      // 1-valid, 2-invalid
    4: string comment_text, // The content of the comment filled by the user, used when action_type=1
    5: i64 comment_id,      // The comment id to be deleted is used when action_type=2
}

struct qingyu_comment_action_response {
    1: base.qingyu_base_response base_resp,
    2: base.Comment comment,                // The comment successfully returns the comment content, no need to re-pull the entire list
}

struct qingyu_get_comment_list_request {
    1: i64 video_id, // Video Id
}

struct qingyu_get_comment_list_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.Comment> comment_list,     // List of comments
}

struct qingyu_get_video_interact_info_request{
    1: i64 video_id, // Video Id
    2: i64 viewer_id, // viewer_id,
}

struct qingyu_get_video_interact_info_response {
    1: base.qingyu_base_response base_resp,
    2: base.VideoInteractInfo interact_info,
}

struct qingyu_batch_get_video_interact_info_request{
    1: list<i64> video_id_list, // Video Id list.
    2: i64 viewer_id, // viewer_id,
}

struct qingyu_batch_get_video_interact_info_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.VideoInteractInfo> interact_info_list,
}

struct qingyu_get_user_interact_info_request{
    1: i64 user_id,
}

struct qingyu_get_user_interact_info_response {
    1: base.qingyu_base_response base_resp,
    2: base.UserInteractInfo interact_info,
}

struct qingyu_batch_get_user_interact_info_request{
    1: list<i64> user_id_list, // user Id list.
}

struct qingyu_batch_get_user_interact_info_response {
    1: base.qingyu_base_response base_resp,
    2: list<base.UserInteractInfo> interact_info_list,
}


service InteractionServer {
    qingyu_favorite_action_response Favorite(1: qingyu_favorite_action_request req),
    qingyu_get_favorite_video_id_list_response GetFavoriteVideoIdList(1: qingyu_get_favorite_video_id_list_request req),
    qingyu_comment_action_response Comment(1: qingyu_comment_action_request req),
    qingyu_get_comment_list_response GetCommentList(1: qingyu_get_comment_list_request req),
    qingyu_get_video_interact_info_response GetVideoInteractInfo (1: qingyu_get_video_interact_info_request req),
    qingyu_batch_get_video_interact_info_response BatchGetVideoInteractInfo (1: qingyu_batch_get_video_interact_info_request req),
    qingyu_get_user_interact_info_response GetUserInteractInfo (1: qingyu_get_user_interact_info_request req),
    qingyu_batch_get_user_interact_info_response BatchGetUserInteractInfo (1: qingyu_batch_get_user_interact_info_request req),
}