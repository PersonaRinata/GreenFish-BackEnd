namespace go aigc
include "base.thrift"


struct qingyu_aigc_question_request {
    1:i64 user_id
    2:string content
}

struct qingyu_aigc_question_response {
    1: base.qingyu_base_response base_resp
    2: string msg // User authentication token
}

service AIGCServer {
    qingyu_aigc_question_response UserAskQuestion(1: qingyu_aigc_question_request req),
}