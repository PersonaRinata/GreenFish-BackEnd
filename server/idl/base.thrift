namespace go base

struct qingyu_base_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct Comment {
    1: i64 id // Video comment id
    2: User user // Comment user information
    3: string content // Comment's content
    4: string create_date // Comment release date, format mm-dd
}

struct User {
    1: i64 id // User id
    2: string name // Username
    3: i64 follow_count // Total number of followings
    4: i64 follower_count // Total number of followers
    5: bool is_follow // true-followed, false-not followed
    6: string avatar,           // user avatar
    7: string background_image, // Image at the top of the user's personal page
    8: string signature,        // Personal signatrue
    9: i64 total_favorited,     // Number of Likes
    10: i64 work_count,         // Number of published videos
    11: i64 favorite_count,     // Total video likes
    12: bool is_doctor,         // Whether the user is a doctor
}

struct SocialInfo{
    1: i64 follow_count // Total number of followings
    2: i64 follower_count // Total number of followers
    3: bool is_follow // true-followed, false-not followed
}

struct UserInteractInfo{
    1: i64 total_favorited,     // Number of Likes
    2: i64 work_count,         // Number of published videos
    3: i64 favorite_count,     // Total video likes
}

struct Video {
    1: i64 id // Video unique identifier
    2: User author // Video author information
    3: string play_url // Video play URL
    4: string cover_url // Video cover address
    5: i64 favorite_count // Total number of likes for the video
    6: i64 comment_count // Total number of comments on the video
    7: bool is_favorite // true-liked, false-not liked
    8: string title // Video title
}

struct VideoInteractInfo{
    1: i64 favorite_count // Total number of likes for the video
    2: i64 comment_count // Total number of comments on the video
    3: bool is_favorite // true-liked, false-not liked
}

struct FriendUser {
    1: i64 id // User id
    2: string name // Username
    3: i64 follow_count // Total number of followings
    4: i64 follower_count // Total number of followers
    5: bool is_follow // true-followed, false-not followed
    6: string avatar,           // user avatar
    7: string background_image, // Image at the top of the user's personal page
    8: string signature,        // Personal signatrue
    9: i64 total_favorited,     // Number of Likes
    10: i64 work_count,         // Number of published videos
    11: i64 favorite_count,     // Total video likes
    12: string message // Latest chat messages with this friend
    13: i64 msgType // message type, 0 => the message received by the current requesting user, 1 => the message sent by the current requesting user
}

struct Message {
    1: i64 id // Message id
    2: i64 to_user_id // The id of the recipient of the message
    3: i64 from_user_id // The id of the sender of the message
    4: string content // Message content
    5: i64 create_time // Message creation time
}

struct LatestMsg{
    1: string message // Latest chat messages with this friend
    2: i64 msgType // message type, 0 => the message received by the current requesting user, 1 => the message sent by the current requesting user
}

struct medicalHistoryInfo{
    1:string symptom //患者的主要症状或原因
    2:string description //病情描述
    3:string history //过去的疾病、手术、药物过敏等历史记录。
    4:string familyInfo //患者家族中与当前疾病相关的疾病或遗传疾病的记录
} //病史信息

struct bodyInfo{
    1:string bloodPressure//血压
    2:string heartRate //心率
    3:string height //身高
    4:string weight //体重
    5:i64 create_time
    6:i64 update_time
} //身体指标

struct IssueList{
    1: string user_id
    2: string username
    3: bool gender
    4: i32 age
    5: i64 create_time
    6: i64 update_time
    7: list<string> department
    8: medicalHistoryInfo medicalHistoryInfo
    9: bodyInfo bodyInfo
    10: string introduction
    11: list<string> medicine
}