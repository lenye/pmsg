package flags

import "errors"

var ErrWeixinAccessToken = errors.New("flags in the group [access_token app_id] required set one")
var ErrWeixinWorkAccessToken = errors.New("flags in the group [access_token corp_id] required set one")
