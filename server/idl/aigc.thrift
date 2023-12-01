namespace go aigc
include "base.thrift"


struct qingyu_aigc_question_request {
    1:i64 user_id
    2:string content
}

struct qingyu_aigc_question_response {
    1: base.qingyu_base_response base_resp
    2: string msg
}

struct qingyu_aigc_issueList_request {
    1:i64 user_id
}

struct qingyu_aigc_issueList_response {
    1: base.qingyu_base_response base_resp
    2: string msg
}

struct qingyu_aigc_choose_word_request {
    1:i64 user_id
    2:string content
}

struct qingyu_aigc_choose_word_response {
    1: base.qingyu_base_response base_resp
    2: string msg
}

struct qingyu_aigc_doctor_analyse_request {
    1:i64 user_id
    2:list<string> content
}

struct qingyu_aigc_doctor_analyse_response {
    1: base.qingyu_base_response base_resp
    2: string msg
}

struct qingyu_aigc_get_history_request {
    1:i64 user_id
}

struct qingyu_aigc_get_history_response {
    1: base.qingyu_base_response base_resp
    2: list<string> msg
}


service AIGCServer {
    qingyu_aigc_question_response UserAskQuestion(1: qingyu_aigc_question_request req),
    qingyu_aigc_issueList_response AnalyseIssueList(1:qingyu_aigc_issueList_request req),
    qingyu_aigc_choose_word_response ChooseWord(1:qingyu_aigc_choose_word_request req),
    qingyu_aigc_doctor_analyse_response DoctorAnalyse(1:qingyu_aigc_doctor_analyse_request req),
    qingyu_aigc_get_history_response GetAIGCHistory(1:qingyu_aigc_get_history_request req),
}