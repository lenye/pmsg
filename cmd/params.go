package cmd

var (
	userAgent   string
	accessToken string

	appID     string
	appSecret string

	toUser          string
	toParty         string
	toTag           string
	toParentUserID  string
	toStudentUserID string

	toUsers          []string
	toPartys         []string
	toTags           []string
	toParentUserIDs  []string
	toStudentUserIDs []string

	templateID       string
	url              string
	clientMsgID      string
	color            string
	mini             map[string]string
	msgType          string
	kfAccount        string
	page             string
	miniProgramState string
	language         string
	scene            string
	title            string

	corpID     string
	corpSecret string

	agentID                int64
	safe                   int
	enableIDTrans          int
	enableDuplicateCheck   int
	duplicateCheckInterval int
	chatID                 string
	toAll                  int
	recvScope              int
	openKfID               string
	msgID                  string
)
