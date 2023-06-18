package qingflowapi

import (
	"fmt"
	"strconv"
	"time"
)

type ApplyApi struct {
	client Client
	appKey string
}

func (api ApplyApi) Page(query ApplyQuery) (PageResult[Apply], error) {
	path := fmt.Sprintf("app/%s/apply/filter", api.appKey)
	var result ApiResponse[PageResult[Apply]]
	err := api.client.post(path, query, &result)
	if err != nil {
		return PageResult[Apply]{}, err
	}
	return result.Result, nil
}

func (api ApplyApi) Update(applyId ID, answers []Answer) (RequestID, error) {
	endpoint := fmt.Sprintf("apply/%d", applyId)
	request := map[string]any{
		"answers": answers,
	}
	var result ApiResponse[struct {
		RequestId RequestID `json:"requestId"`
	}]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.RequestId, nil
}

func (api ApplyApi) Get(applyId ID) (Apply, error) {
	endpoint := fmt.Sprintf("apply/%d", applyId)
	var result ApiResponse[Apply]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return Apply{}, err
	}
	return result.Result, nil
}

type ApplyCreationRequest struct {
	ApplyUser struct {
		Email    string // 操作人的邮箱，如果不传就以“匿名用户”身份添加数据
		AreaCode string
		Mobile   string
	}
	Answers []Answer
}

func (api ApplyApi) Create(request ApplyCreationRequest) (RequestID, ID, error) {
	path := fmt.Sprintf("app/%s/apply", api.appKey)
	var result ApiResponse[struct {
		RequestId RequestID `json:"requestId"`
		ApplyId   ID        `json:"applyId"`
	}]
	err := api.client.post(path, request, &result)
	if err != nil {
		return "", 0, err
	}
	return result.Result.RequestId, result.Result.ApplyId, nil
}

func (api ApplyApi) DeletePage(query ApplyQuery) (RequestID, error) {
	endpoint := fmt.Sprintf("app/%s/apply", api.appKey)
	var result ApiResponse[struct {
		RequestId RequestID `json:"requestId"`
	}]
	err := api.client.deleteRequest(endpoint, query, &result)
	if err != nil {
		return "", err
	}
	return result.Result.RequestId, nil
}

func (api ApplyApi) GetAppApply(applyId ID) (Apply, error) {
	path := fmt.Sprintf("app/%s/apply/%d", api.appKey, applyId)
	var result ApiResponse[Apply]
	err := api.client.get(path, nil, &result)
	if err != nil {
		return Apply{}, err
	}
	return result.Result, nil
}

type QueRelationRole int

const (
	QUE_RELATION_ROLE_WORKSPACE_MANAGER QueRelationRole = 1 // 工作区管理员
	QUE_RELATION_ROLE_APPLICANT                         = 2
	QUE_RELATION_ROLE_PROCESSOR                         = 3
	QUE_RELATION_ROLE_QROBOT                            = 4
	QUE_RELATION_ROLE_SHARER                            = 5
	QUE_RELATION_ROLE_MAIL_PROCESSOR                    = 6 // 邮件处理人
)

type QueRelationQuery struct {
	Role         QueRelationRole
	Pass         string // 申请人身份请求时如果表单有密码需要传入
	AuditNodeId  ID
	SearchKey    string
	KeyQueValues []struct {
		KeyQueId int
		Ordinal  int
		Values   []string
	}
	QueryQuestions []struct {
		QueId   int
		Ordinal int
	}
}

/*
获取数据关联的内容
*/
func (api ApplyApi) GetQueRelation(query QueRelationQuery) ([]Apply, error) {
	endpoint := "data/queRel"
	var result ApiResponse[[]Apply]
	err := api.client.post(endpoint, query, &result)
	if err != nil {
		return []Apply{}, err
	}
	return result.Result, nil
}

type ApplyQuery struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
	Type     int `json:"type"`
	Sorts    []struct {
		QueId    int  `json:"queId"`
		IsAscend bool `json:"isAscend"`
	} `json:"sorts"`
	Queries []struct {
		QueId         int      `json:"queId"`
		SearchKey     string   `json:"searchKey"`
		SearchKeys    []string `json:"searchKeys"`
		MinValue      string   `json:"minValue"`
		MaxValue      string   `json:"maxValue"`
		Scope         int      `json:"scope"`
		SearchOptions []int    `json:"searchOptions"`
		SearchUserIds []string `json:"searchUserIds"`
	} `json:"queries"`
	QueryKey string `json:"queryKey"`
	ApplyIds []ID   `json:"applyIds"`
}

type Apply struct {
	ApplyId       ID       `json:"applyId"`
	Ordinal       int      `json:"ordinal"` // 如果是表格子字段，ordinal表示行号
	Answers       []Answer `json:"answers"`
	ApplyBaseInfo string   `json:"applyBaseInfo"`
}

type Answer struct {
	QueId       ID            `json:"queId"`
	QueTitle    string        `json:"queTitle"`
	QueType     int           `json:"queType"`
	TableValues [][]SubAnswer `json:"tableValues"`
	Values      []struct {
		DataValue   string `json:"dataValue"`
		Id          ID     `json:"id"`
		Email       string `json:"email"`
		OptionId    string `json:"optionId"`
		Ordinal     string `json:"ordinal"`
		OtherInfo   string `json:"otherInfo"`
		PluginValue string `json:"pluginValue"`
		QueId       ID     `json:"queId"`
		Value       string `json:"value"`
	} `json:"values"`
}

type SubAnswer struct {
	QueId ID `json:"queId"`
}

type Assignment struct {
	AuditNodeId ID
	UserIdList  []ID
}

func (api ApplyApi) Reassign(applyId ID, assignments []Assignment) (ID, error) {
	endpoint := fmt.Sprintf("apply/%d/reassign", applyId)
	var result ApiResponse[Apply]
	request := map[string][]Assignment{
		"reassignmentInfo": assignments,
	}
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.ApplyId, nil
}

func (api ApplyApi) Rollback(applyId ID, userId string, auditNodeId ID, targetAuditNodeId ID, auditFeedback string) error {
	endpoint := fmt.Sprintf("%d/audit/rollback", applyId)
	request := map[string]any{
		"userId":            userId,
		"auditNodeId":       auditNodeId,
		"targetAuditNodeId": targetAuditNodeId,
		"auditFeedback":     auditFeedback,
	}
	var result ApiResponse[any]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return err
	}
	return nil
}

type AssociatedQueType int

type AuditRecord struct {
	AuditRcdId    ID `json:"auditRcdId"`
	AuditModifies []struct {
		QueID        ID      `json:"queId"`
		QueTitle     string  `json:"queTitle"`
		QueType      QueType `json:"queType"`
		BeforeAnswer Answer  `json:"beforeAnswer"`
		AfterAnswer  Answer  `json:"afterAnswer"`
	} `json:"auditModifies"`
}

func (api ApplyApi) GetAuditRecord(applyId ID, AuditRcdId ID) (AuditRecord, error) {
	endpoint := fmt.Sprintf("apply/%d/auditRecord/%d", applyId, AuditRcdId)
	var result ApiResponse[AuditRecord]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return AuditRecord{}, err
	}
	return result.Result, nil
}

type ApplyStatus int

const (
	APPLY_STATUS_DRAFT             ApplyStatus = 1 // 草稿
	APPLY_STATUS_PROCESSING                    = 2 // 流程中（已经有用户处理过）
	APPLY_STATUS_PASSED_WITH_FLOW              = 3 // 已通过（有流程）
	APPLY_STATUS_REFUSED                       = 4 // 已拒绝
	APPLY_STATUS_DRAWBACK                      = 5 // 待完善（退回申请人）
	APPLY_STATUS_PASS_WITHOUT_FLOW             = 6 // 已通过（无流程）
	APPLY_STATUS_PROCESSING_PURE               = 7 // 流程中（没有用户处理过）
)

// 负责人类型
type AuditUserType int

const (
	AUDIT_USER_NORMAL     AuditUserType = 0 // 普通处理人
	AUDIT_USER_EMAIL_USER               = 1 // 邮箱处理人（处理人为动态，且选择为邮箱字段时）
)

type AuditUser struct {
	Type   AuditUserType
	UserId string // userType为1时，userId会返回空值
	Email  string
	// 仅用于会签节点，标记这个处理人是否完成处理，true：处理完成
	FinishAudit bool
}

type ApplyAuditRecord struct {
	AuditFeedback     string
	AuditNodeId       ID
	AuditNodeName     string
	AuditRcdId        ID
	AuditResult       AuditResultType
	AuditTime         time.Time
	AuditUser         []AuditUser
	WaitAuditUserList []AuditUser
}

type ApplyAuditRecordDetail struct {
	ApplyStatus  ApplyStatus
	AuditRecords []ApplyAuditRecord
	CurrentNodes []ApplyAuditRecord
}

func (api ApplyApi) GetAllAuditRecord(applyId ID) (ApplyAuditRecord, error) {
	endpoint := fmt.Sprintf("apply/%d/auditRecord", applyId)
	var result ApiResponse[ApplyAuditRecord]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return ApplyAuditRecord{}, err
	}
	return result.Result, nil
}

func (api ApplyApi) SetUrge(applyId ID) (ID, error) {
	endpoint := fmt.Sprintf("apply/%d/urge", applyId)
	var result ApiResponse[Apply]
	err := api.client.post(endpoint, nil, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.ApplyId, nil
}

func (api ApplyApi) SetAuditResult(applyId ID, auditResult AuditResultType, auditNodeId ID, auditFeedback string) (bool, error) {
	endpoint := fmt.Sprintf("%d/audit", applyId)
	request := map[string]any{
		"auditResult":   auditResult,
		"auditNodeId":   auditNodeId,
		"auditFeedback": auditFeedback,
	}
	var result ApiResponse[bool]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return false, err
	}
	return result.Result, nil
}

type ApplyComment struct {
	CommentId      ID
	CommentMsg     string
	MentionUserIds []ID
}

func (api ApplyApi) PageComment(applyId ID, page int, size int) (PageResult[ApplyComment], error) {
	endpoint := fmt.Sprintf("apply/%d/comment", applyId)
	params := map[string]string{
		"pageSize": strconv.Itoa(page),
		"pageNum":  strconv.Itoa(size),
	}
	var result ApiResponse[PageResult[ApplyComment]]
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return PageResult[ApplyComment]{}, err
	}
	return result.Result, nil
}

func (api ApplyApi) CreateComment(applyId ID, comment ApplyComment) (ApplyComment, error) {
	endpoint := fmt.Sprintf("apply/%d/comment", applyId)
	var result ApiResponse[ApplyComment]
	err := api.client.post(endpoint, comment, &result)
	if err != nil {
		return ApplyComment{}, err
	}
	return result.Result, nil
}

type TaskType int

const (
	TASK_TYPE_TODO        TaskType = 1 // 待办事项
	TASK_TYPE_APPLY                = 2 // 我的申请
	TASK_TYPE_CARBON_COPY          = 3 // 抄送
	TASK_TYPE_DRAFT                = 4 // 草稿
	TASK_TYPE_DONE                 = 5 // 已办事项
)

type TodoTaskStatus int

const (
	TODO_TASK_STATUS_ALL     TodoTaskStatus = 1
	TODO_TASK_STATUS_URGE                   = 4
	TODO_TASK_STATUS_DDL                    = 5 // 即将超时
	TODO_TASK_STATUS_TIMEOUT                = 6 // 已超时
)

type ApplyTaskStatus int

const (
	APPLY_TASK_STATUS_ALL          ApplyTaskStatus = 1
	APPLY_TASK_STATUS_PASS                         = 2
	APPLY_TASK_STATUS_REFUSE                       = 3
	APPLY_TASK_STATUS_NEED_IMPROVE                 = 4
	APPLY_TASK_STATUS_PROCESS                      = 5
)

type CarbonCopyTaskStatus int

const (
	CARBON_COPY_TASK_ALL    CarbonCopyTaskStatus = 1
	CARBON_COPY_TASK_UNREAD                      = 3
)

type DraftTaskStatus int

const (
	DRAFT_TASK_STATUS_ALL = 1
)

type DoneTaskStatus int

const (
	DONE_TASK_STATUS_ALL     DoneTaskStatus = 1
	DONE_TASK_STATUS_PASS                   = 2
	DONE_TASK_STATUS_REFUEE                 = 3
	DONE_TASK_STATUS_PROCESS                = 5
)

type ApplyUserAuth int

const (
	APPLY_USER_AUTH_NORMAL            ApplyUserAuth = 1
	APPLY_USER_AUTH_MANAGER                         = 2
	APPLY_USER_AUTH_WORKSPACE_CREATOR               = 3
)

type ApplyTask struct {
	ApplyId     ID     `json:"applyId"`
	AppKey      string `json:"appKey"`
	ApplyNum    int    `json:"applyNum"`
	FormTitle   string `json:"formTitle"`
	CurAuditors []struct {
		AuditorName    string    `json:"auditorName"`
		TimeoutDate    time.Time `json:"timeoutDate"`
		PreTimeoutDate time.Time `json:"preTimeoutDate"`
	} `json:"cur_auditors"`
	ApplyUser struct {
		Uid       ID            `json:"uid"`
		Status    bool          `json:"status"` // 是否已经激活
		Email     string        `json:"email"`
		NickName  string        `json:"nickName"`
		HeadImage string        `json:"headImage"`
		Accepted  bool          `json:"accepted"` // 是否接受当前工作区的邀请
		Auth      ApplyUserAuth `json:"auth"`
		Remark    string        `json:"remark"`
		MobileNum string        `json:"mobileNum"`
	} `json:"apply_user"`
	ApplyTime   time.Time   `json:"applyTime"`
	ApplyStatus ApplyStatus `json:"applyStatus"`
	BeingUrged  bool        `json:"beingUrged"`
	BeingUnread bool        `json:"beingUnread"`
	FormId      string      `json:"formId"`
	TagIds      []string    `json:"tagIds"`
	UpdateTime  time.Time   `json:"updateTime"`
}

// 获取个人在当前工作区全部事项信息
func (api ApplyApi) PageTask(page int, size int, userId ID, taskType TaskType, status int) (PageResult[ApplyTask], error) {
	endpoint := "dynamic/audits"
	request := map[string]string{
		"pageNum":  strconv.Itoa(page),
		"pageSize": strconv.Itoa(size),
		"userId":   strconv.Itoa(int(userId)),
		"type":     strconv.Itoa(int(taskType)),
		"status":   strconv.Itoa(status),
	}
	var result ApiResponse[PageResult[ApplyTask]]
	err := api.client.get(endpoint, request, &result)
	if err != nil {
		return PageResult[ApplyTask]{}, err
	}
	return result.Result, nil
}

type ApplyPrintTemplate struct {
	PrintKey  string `json:"printKey"`
	PrintName string `json:"printName"`
	Url       string `json:"url"`
}

/*
获取单条数据对应打印模版文件
*/
func (api ApplyApi) GetPrintTemplate(applyId ID, userId string, auditNodeId ID) (ApplyPrintTemplate, error) {
	endpoint := fmt.Sprintf("apply/%d/print", applyId)
	request := map[string]any{
		"userId":      userId,
		"auditNodeId": auditNodeId,
	}
	var result ApiResponse[ApplyPrintTemplate]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return ApplyPrintTemplate{}, err
	}
	return result.Result, nil
}

type QueIdSearchKey struct {
	QueId     string `json:"queId"`
	SearchKey string `json:"searchKey"`
}

/*
获取数据表列表数据
*/
func (api ApplyApi) PageChart(applyId ID, query []QueIdSearchKey) (PageResult[Apply], error) {
	endpoint := fmt.Sprintf("chart/%d/apply/filter", applyId)
	request := map[string]any{
		"accurateQuery": query,
	}
	var result ApiResponse[PageResult[Apply]]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return PageResult[Apply]{}, err
	}
	return result.Result, nil
}

type ApplyChart struct {
	BeingTime   time.Time `json:"beingTime"` // 折线图/柱状图 是否按照时间类型显示
	ChartSeries []struct {
		SeriesName   string `json:"seriesName"` // 分类（也叫系列）
		Value        int64  `json:"value"`      // 时间类型，返回对应的时间戳
		SeriesValues []struct {
			RowName      string `json:"rowName"`
			Value        int    `json:"value"`        // 行高（也叫值），可能是计数的值，也可能是求和、平均值等
			DisplayValue string `json:"displayValue"` // 格式化后的值
		} `json:"seriesValues"`
		// 汇总表的表头
		ChartHeaders []struct {
			Value string `json:"value"`
		} `json:"chartHeaders"`
		// 汇总表数据
		CollectValues []struct {
			Value       string `json:"value"`
			TargetValue string `json:"targetValue"` // 指标卡数据
		} `json:"collectValues"`
		GanttValue struct {
			// 甘特图的id，在最后一个维度按照申请去聚合,后端生成规则可以按照，第一级从1开始自增，第二级id为第一级+从1开始自增，第三级以此类推，哈夫曼编码
			PId string `json:"pId"`
			// 父级的pId，第一级维度，这里返回0
			PParent string `json:"pParent"`
			// 可以显示时分秒，也可以不显示，最后一级才会返回
			PStart string `json:"pStart"`
			// 最后一级才会返回
			PEnd string `json:"pEnd"`
			// 标题
			PName   string `json:"pName"`
			ApplyId ID     `json:"applyId"`
		} `json:"gantt_value"`
		CalendarValue []struct {
			Date     time.Time `json:"date"`
			DataList []struct {
				Title string `json:"title"`
			} `json:"dataList"`
		} `json:"calendar_value"`
	} `json:"chart_series"`
}

func (api ApplyApi) GetChart(charKey string, query []QueIdSearchKey) (ApplyChart, error) {
	endpoint := "chart/%s/chartData"
	request := map[string]any{
		"accurateQuery": query,
	}
	var result ApiResponse[ApplyChart]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return ApplyChart{}, err
	}
	return result.Result, nil
}
