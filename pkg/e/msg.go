package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "Invalid params",
	RECORD_NOT_FOUND:                "Record not found",
	ERROR_EXIST_TAG:                 "Exist tag",
	ERROR_EXIST_TAG_FAIL:            "Exit tag fail",
	ERROR_NOT_EXIST_TAG:             "Not exit tag",
	ERROR_GET_TAGS_FAIL:             "Get tags fail",
	ERROR_GET_TAG_FAIL:              "Get tag fail",
	ERROR_COUNT_TAG_FAIL:            "Count tag fail",
	ERROR_ADD_TAG_FAIL:              "Add tag fail",
	ERROR_EDIT_TAG_FAIL:             "Edit tag fail",
	ERROR_DELETE_TAG_FAIL:           "Delete tag fail",
	ERROR_EXPORT_TAG_FAIL:           "Export tag fail",
	ERROR_IMPORT_TAG_FAIL:           "Import tag fail",
	ERROR_NOT_EXIST_ARTICLE:         "Not exist article",
	ERROR_ADD_ARTICLE_FAIL:          "Add article fail",
	ERROR_DELETE_ARTICLE_FAIL:       "Delete article fail",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "Check axist article fail",
	ERROR_EDIT_ARTICLE_FAIL:         "Edit article fail",
	ERROR_COUNT_ARTICLE_FAIL:        "Count article fail",
	ERROR_GET_ARTICLES_FAIL:         "Get articles fail",
	ERROR_GET_ARTICLE_FAIL:          "Get article fail",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "Get article poster fail",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Auth check token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Auth check token timeout",
	ERROR_AUTH_TOKEN:                "Auth token error",
	ERROR_AUTH:                      "Auth error",
	ERROR_USERNAME_EXISTS:           "Username exists",
	ERROR_REGISTRATION_FAIL:         "Registration failure",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Save image fail",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Check image fail",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Upload check image fail",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
