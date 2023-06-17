package qingflowapi

type QueType int

// https://qingflow.com/help/docs/6239a267c7c55500418da348
const (
	// 描述文字
	QUE_TYPE_DESC QueType = 1
	// 单行文字
	QUE_TYPE_LINE = 2
	// 多行文字
	QUE_TYPE_MULTI_LINE = 3
	// 日期时间
	QUE_TYPE_DATE = 4
	// 成员字段
	QUE_TYPE_MEMBER = 5
	// 邮箱
	QUE_TYPE_EMAIL = 6
	// 手机
	QUE_TYPE_PHONE         = 7
	QUE_TYPE_NUMBER        = 8
	QUE_TYPE_LINK          = 9
	QUE_TYPE_SINGLE_OPTION = 10
	QUE_TYPE_DROP_OPTION   = 11
	QUE_TYPE_MULTI_OPTION  = 12
	QUE_TYPE_FILE          = 13
	// 起止时间
	QUE_TYPE_DURATION  = 14
	QUE_TYPE_IMAGE     = 15
	QUE_TYPE_RICH_TEXT = 16
	// 定位字段
	QUE_TYPE_LOCATION      = 17
	QUE_TYPE_TABLE         = 18
	QUE_TYPE_DATA_RELATION = 19
	QUE_TYPE_Q_LINKER      = 20
	QUE_TYPE_ADDRESS       = 21
	QUE_TYPE_DEPARTMENT    = 22
)

// https://qingflow.com/help/docs/6239a267c7c55500418da348
const (
	JUDGE_TYPE_NE            int = 1
	JUDGE_TYPE_IN                = 2
	JUDGE_TYPE_NOT_IN            = 3
	JUDGE_TYPE_GT                = 4
	JUDGE_TYPE_GTEQ              = 5
	JUDGE_TYPE_LT                = 6
	JUDGE_TYPE_LTEQ              = 7
	JUDGE_TYPE_DYNAMIC_RANGE     = 8
	// 等于任意一个
	JUDGE_TYPE_EQ_ANY = 9
	// 不等于任意一个
	JUDGE_TYPE_NEQ_ANY = 10
	// 为空值
	JUDGE_TYPE_EMPTY = 11
)

const (
	// 申请人
	SPECIAL_QUE_APPLICANT int = 1
	// 申请时间
	SPECIAL_QUE_APPLICATED_DATE = 2
	SPECIAL_QUE_LATEST_UPDATED  = 3
	// 数据当前所处节点
	SPECIAL_QUE_CURRENT_NODE = 4
	// 申请人的部门主管（不入库，question不存入）
	SPECIAL_QUE_MANAGER_OF_APPLICATION = 5
	// 当前用户（成员字段）（不入库，question不存入）
	SPECIAL_QUE_CURRENT_USER = 6
	// 当前节点负责人 （超时提醒）（不入库，question不存入）
	SPECIAL_QUE_CURRENT_NODE_MANAGER = 7
	// 编号
	SPECIAL_QUE_SERIAL_NUMBER = 0
	// 数据来源
	SPECIAL_QUE_DATA_SOURCE = -1
	// 当前用户所处部门（表单中部门字段的可选范围/默认值使用）（不入库，question不存入）
	SPECIAL_QUE_CURRENT_USER_MANAGER = -2
	// 流程节点（数据经过过的所有审批、填写节点）
	SPECIAL_QUE_FLOW_NODE = -3
	// 节点负责人 （处理过数据的所有负责人）
	SPECIAL_QUE_NODE_MANAGER = -4
	// 流程处理状态（数据在该节点最终的处理状态 正常处理、流程超时、超时预警）
	SPECIAL_QUE_FLOW_STATUS = -5
	// 流程处理时长
	SPECIAL_QUE_FLOW_PROCESS_DURATION = -6
	// 流程超时时长
	SPECIAL_QUE_FLOW_OVERTIME_DURATION = -7
	// 聚合时，后端特殊处理所占用，前端可忽略
	SPECIAL_QUE_BACKEND_SPECIAL = -8
	// 支付信息
	SPECIAL_QUE_PAYMENT = -9
	// 支付信息·支付状态（supid为-9）
	SPECIAL_QUE_PAYMENT_STATUS = -10
	// 支付信息·节点名（supid为-9）
	SPECIAL_QUE_PAYMENT_NODE_NAME = -11
	// 支付信息·支付内容（supid为-9）
	SPECIAL_QUE_PAYMENT_DETAIL = -12
	// 支付信息·支付金额（supid为-9）
	SPECIAL_QUE_PAYMENT_AMOUNT = -13
	// 支付信息·支付方式（supid为-9）
	SPECIAL_QUE_PAYMENT_METHOD = -14
	// 支付信息·订单号（supid为-9）
	SPECIAL_QUE_PAYMENT_ORDER_NUMBER = -15
	// 支付信息·支付时间（supid为-9）
	SPECIAL_QUE_PAYMENT_TIME = -16
	// 数据唯一标识(applyId)
	SPECIAL_QUE_DATA_UID = -17
)

const (
	// 提交处理
	AUDIT_RESULT_SUBMITTED = 1
	// 通过
	AUDIT_RESULT_PASS = 2
	// 拒绝
	AUDIT_RESULT_REFUSE = 3
	// 退回数据
	AUDIT_RESULT_REWORK = 4
	// 发起申请/申请人草稿提交
	AUDIT_RESULT_APPLY = 5
	// 退回申请人完善
	AUDIT_RESULT_REWORK_TO_APPLICANT = 6
	// 抄送
	AUDIT_RESULT_CARBON_COPY = 7
	// 处理中
	AUDIT_RESULT_PROCESSING = 8
	// 移交数据
	AUDIT_RESULT_TRANSFER = 9
	// 申请人撤销申请
	AUDIT_RESULT_RECALL = 10
	// 处理人撤销处理
	AUDIT_RESULT_CANCEL = 11
	// 创建人/处理人更新数据
	AUDIT_RESULT_UPDATED = 12
	// qrobot添加数据
	AUDIT_RESULT_QROBOT_ADD_DATA = 13
	// qrobot更新数据
	AUDIT_RESULT_QROBOT_UPDATE_DATA = 14
	// 管理员重新指派负责人
	AUDIT_RESULT_REASSIGN_MANAGER = 15
	// qrobot触发创建数据
	AUDIT_RESULT_QROBOT_TRIGGER_DATA_CREATED = 16
	// qrobot触发更新数据
	AUDIT_RESULT_QROBOT_TRIGGER_DATA_UPDATED = 17
	// qrobot触发发送邮件
	AUDIT_RESULT_QROBOT_TRIGGER_MAIL_SENT = 18
	// 创建人更新邮箱
	AUDIT_RESULT_CREATOR_UPDATE_MAIL = 19
	// qrobot触发发送短信
	AUDIT_RESULT_QROBOT_TRIGGER_SMS_SENT = 20
	// qrobot触发webhook
	AUDIT_RESULT_QROBOT_TRIGGER_WEBHOOK = 21
	// webhook更新数据
	AUDIT_RESULT_WEBHOOK_UPDATE_DATA = 22
	// 任务委托
	AUDIT_RESULT_DELEGATE_TASK = 23
	// 回退重新提交
	AUDIT_RESULT_REWORK_RESUBMITTED = 24
	// 节点无负责人转交给自定义负责人
	AUDIT_RESULT_NODE_TO_CUSTOMIZED_MANAGER = 25
	// 删除数据
	AUDIT_RESULT_DATA_DELETE = 26
	// 恢复数据
	AUDIT_RESULT_RECOVER_DATA = 27
	// 自动通过
	AUDIT_RESULT_AUTO_PASS = 28
	// 加签
	AUDIT_RESULT_COUNTERSIGN = 29
	// OPEN API重新指派负责人
	AUDIT_RESULT_REASSIGN_MANAGER_BY_OPENAPI = 30
	// Q-Robot执行发送提醒
	AUDIT_RESULT_QROBOT_SENT_REMIND = 31
)

const (
	ICON_COLOR_QING_ORANGE string = "qing-orange"
	ICON_COLOR_YELLOW             = "yellow"
	ICON_COLOR_GREEN              = "green"
	ICON_COLOR_EMERALD            = "emerald"
	ICON_COLOR_BLUE               = "blue"
	ICON_COLOR_AZURE              = "azure"
	ICON_COLOR_INDIGO             = "indigo"
	ICON_COLOR_QING_PURPLE        = "qing-purple"
	ICON_COLOR_PURPLE             = "purple"
	ICON_COLOR_PINK               = "pink"
	ICON_COLOR_RED                = "red"
	ICON_COLOR_ORANGE             = "orange"
)
