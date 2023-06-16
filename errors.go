package qingflowapi

type ApiError struct {
	Code           int
	Message        string
	DefaultMessage string
}

func (e ApiError) Error() string {
	return e.Code + ":" + e.Message
}

func newApiError(code int, message string, defaultMessage string) ApiError {
	if len(message) == 0 {
		message = defaultMessage
	}
	var err ApiError
	err.Code = code
	err.Message = message
	err.DefaultMessage = defaultMessage
	return err
}

type ClientApiError struct {
	ApiError
}

type NotFoundApiError struct {
	ClientApiError
}

func notfound(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err NotFoundApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type InvalidArgumentApiError struct {
	ClientApiError
}

func invalidArgument(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err InvalidArgumentApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type UnauthorizedApiError struct {
	ClientApiError
}

func unauthorized(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err UnauthorizedApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type ForbiddenApiError struct {
	ClientApiError
}

func forbidden(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err ForbiddenApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type FailedPreconditionApiError struct {
	ClientApiError
}

func failedPrecondition(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err FailedPreconditionApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type TooManyRequestsApiError struct {
	ClientApiError
}

func tooManyRequests(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err TooManyRequestsApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type AlreadyExistsApiError struct {
	ClientApiError
}

func alreadyExists(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err AlreadyExistsApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type ServerApiError struct {
	ApiError
}

type InternalApiError struct {
	ServerApiError
}

func internal(defaultMessage string) func(int, string) error {
	return func(code int, message string) error {
		var err InvalidArgumentApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

func translateError(code int, message string) error {
	if code == 0 {
		return nil
	}
	convert, ok := codeMessage[code]
	if ok {
		return convert(code, message)
	}
	return ApiError{Code: code, Message: message}
}

type CodeMessage map[int]func(int, string) error

// ref: https://qingflow.com/help/docs/61de4377f6076f0041e858ad
var codeMessage = CodeMessage{
	40001: internal("登录信息失效，请重新登录"),
	40002: unauthorized("您没有权限进行该操作"),
	40003: invalidArgument("请输入正确的邮箱"),
	40004: notfound("该邮箱尚未注册轻流"),
	40005: invalidArgument("账号或密码错误"),
	40006: alreadyExists("该邮箱已经注册"),
	40007: invalidArgument("请输入正确的手机号"),
	40008: invalidArgument("该邮箱已经激活"),
	40009: invalidArgument("验证码已失效"),
	40010: invalidArgument("用户名不得超过20个字"),
	40011: notfound("用户不存在"),
	40012: notfound("标签不存在"),
	40013: notfound("模版不存在"),
	40014: invalidArgument("应用名不得超过50个字"),
	40015: invalidArgument("该应用尚未发布，无法提交"),
	40016: invalidArgument("应用申请通道已关闭，无法提交"),
	40017: invalidArgument("应用已关闭，无法提交"),
	40018: notfound("应用不存在"),
	40019: invalidArgument("有必填字段为空"),
	40020: invalidArgument("字段校验未通过"),
	40021: invalidArgument("日期与已有时间重叠"),
	40022: invalidArgument("所填内容与已有数据重叠"),
	40023: notfound("该申请不存在"),
	40024: notfound("该节点不存在"),
	40025: invalidArgument("存在未设置负责人的节点"),
	40026: invalidArgument("应用未发布"),
	40027: notfound("该记录不存在"),
	40028: alreadyExists("该邮箱已注册"),
	40029: invalidArgument("用户登录密码错误"),
	40030: invalidArgument("该邮箱未激活"),
	40031: unauthorized("当前节点没有审批权限"),
	40032: notfound("审批记录不存在"),
	40033: invalidArgument("自动判定条件无效"),
	40034: invalidArgument("节点负责人不能为空"),
	40035: tooManyRequests("催办过于频繁，请稍后重试"),
	40036: invalidArgument("需要分支节点"),
	40037: invalidArgument("需要自动判定节点"),
	40038: notfound("对象不存在"),
	40039: alreadyExists("对象已存在"),
	40040: invalidArgument("退回给匿名申请人"),
	40041: notfound("表单不存在"),
	40042: invalidArgument("应用名不得超过50个字"),
	40043: notfound("主题不存在"),
	40044: invalidArgument("字段关联设置信息出错"),
	40045: notfound("模版不存在"),
	40046: alreadyExists("重复的值"),
	40047: invalidArgument("请选择正确的日期"),
	40048: invalidArgument("请输入数字"),
	40049: invalidArgument("数据已被拒绝"),
	40050: invalidArgument("CTCT_TAG_NOT_EXISTS"),
	40051: notfound("该部门不存在"),
	40052: notfound("父部门不存在"),
	40053: invalidArgument("部门存在成员"),
	40054: invalidArgument("当前部门有子部门，无法删除"),
	40055: notfound("通讯录内无该成员"),
	40056: invalidArgument("CTCT_USER_IS_USER_LEADER"),
	40057: invalidArgument("CTCT_USER_IS_DEPT_LEADER"),
	40058: invalidArgument("EXCEL_ID_NOT_EXISTS"),
	40059: invalidArgument("当前工作区无法修改通讯录"),
	40060: invalidArgument("该工作区数据量已达上限，无法继续提交"),
	40061: invalidArgument("该工作区上传文件量已达上限，无法继续上传"),
	40062: invalidArgument("您已到达该应用申请上限，无法继续提交"),
	40063: invalidArgument("该应用已达数据量上限，无法继续提交"),
	40064: unauthorized("您没有权限访问该应用"),
	40065: unauthorized("您没有权限访问该应用"),
	40066: unauthorized("您没有权限访问该应用"),
	40067: failedPrecondition("该应用已达月数据量上限，无法继续提交"),
	40068: failedPrecondition("该应用已达数据量自定义上限，无法继续提交"),
	40069: failedPrecondition("该功能仅付费后可使用"),
	40070: invalidArgument("上传文件的信息失效"),
	40071: invalidArgument("上传文件太大了"),
	40072: invalidArgument("上传文件签名计算失败"),
	40073: invalidArgument("oos上传文件失败"),
	40074: invalidArgument("编辑应用版本号不符合要求"),
	40075: alreadyExists("不能重复审批"),
	40078: invalidArgument("所有人可见仪表盘可用数量上限"),
	40079: invalidArgument("版本降级"),
	40080: invalidArgument("请输入正确的手机号"),
	40081: invalidArgument("发送验证码过于频繁，请稍后重试"),
	40082: invalidArgument("验证码已失效"),
	40083: internal("出现未知错误"),
	40084: tooManyRequests("发送验证码过于频繁，请稍后重试"),
	40085: tooManyRequests("通讯录同步过于频繁，请稍后重试"),
	40086: invalidArgument("ip当日验证码发送次数达到上限"),
	40091: invalidArgument("qRobot节点成环"),
	40092: invalidArgument("该字段不存在"),
	40093: invalidArgument("字段类型错误"),
	40094: invalidArgument("用户角色错误"),
	40095: unauthorized("您没有编辑权限"),
	40096: invalidArgument("该字段不存在"),
	40097: invalidArgument("请先登录轻流"),
	40098: invalidArgument("节点负责人过多"),
	40099: invalidArgument("节点配置错误"),
	40100: invalidArgument("该聚合函数不存在"),
	40101: invalidArgument("该聚合函数只接受数字类型的列值"),
	40102: notfound("该自定义报表不存在"),
	40103: alreadyExists("该手机号已经注册"),
	40104: internal("未知错误"),
	40105: invalidArgument("甘特图设置错误"),
	40106: invalidArgument("无效的用户登录code"),
	40110: notfound("应用的关联报表不存在"),
	40111: invalidArgument("应用的关联报表设置不正确"),
	40150: invalidArgument("qrobot关联字段不存在或者被删除"),
	40151: invalidArgument("节点名太长"),
	40152: invalidArgument("节点粘贴到原位了"),
	40153: invalidArgument("节点粘贴位置错误（比如成环了）"),
	40200: invalidArgument("UID不正确"),
	40201: invalidArgument("UID不在选项内"),
	40202: invalidArgument("UID不在选项内"),
	40300: invalidArgument("选项份额到达上限"),
	40301: invalidArgument("邀请信息失效"),
	40500: invalidArgument("数据必填"),
	40501: alreadyExists("数据重复"),
	40502: invalidArgument("数据不符合格式"),
	40503: invalidArgument("申请时间限制"),
	40504: invalidArgument("超出附件字段数量限制"),
	40505: invalidArgument("超出表格子字段数量上限"),
	40506: invalidArgument("提交校验未通过"),
	41000: invalidArgument("微信登录信息失效，请重新登录"),
	41001: invalidArgument("微信登录信息失效，请重新登录"),
	41002: alreadyExists("该微信已绑定了轻流账号"),
	41003: alreadyExists("该账号已绑定了其他微信"),
	41004: unauthorized("该微信公众号未对轻流进行足够的授权"),
	41005: alreadyExists("此微信公众号已经授权在别的工作区"),
	41006: invalidArgument("无法设置消息模板，请检查行业设置"),
	41007: notfound("找不到该微信公众号"),
	41008: invalidArgument("绑定第三方微信公众号失败"),
	41049: invalidArgument("调用钉钉api失败"),
	41050: failedPrecondition("您的企业需要添加轻流应用，如果已添加，则联系管理员同步通讯录"),
	41051: notfound("该企业不存在"),
	41056: unauthorized("组织未授权轻流访问（用户没有安装轻流）"),
	41057: invalidArgument("企业微信客户未登录"),
	41061: invalidArgument("获取预授权码失败"),
	41062: invalidArgument("设置授权失败"),
	41100: invalidArgument("导出excel失败"),
	41101: failedPrecondition("导入数据超出数量限制"),
	41102: failedPrecondition("导入进入流程的数据超出数量限制"),
	41103: failedPrecondition("导出数据超出数量限制"),
	41110: invalidArgument("帐号密码匹配失败,仅用于登录"),
	41111: failedPrecondition("输入错误密码次数超过五次，锁定五分钟"),
	42001: notfound("工作区不存在"),
	42002: notfound("工作区vip信息不存在"),
	42003: notfound("工作区不在试用期"),
	42004: invalidArgument("工作区相同插件类型重复绑定"),
	42005: invalidArgument("工作区未绑定该插件"),
	42006: invalidArgument("插件中心未开启"),
	42101: invalidArgument("付费类型不存在"),
	42102: notfound("订单不存在"),
	42103: invalidArgument("支付宝支付异常"),
	42104: invalidArgument("支付宝支付异常"),
	42105: invalidArgument("订单查询失败"),
	42106: failedPrecondition("工作区支付延期试用已达最大天数"),
	42107: invalidArgument("订单数据非法"),
	42201: invalidArgument("文件格式错误"),
	42202: invalidArgument("导入文件太大"),
	42203: invalidArgument("请先设置密码"),
	42204: invalidArgument("导入文件太多"),
	43001: invalidArgument("FORMULA_FIELD_IS_EMPTY"),
	43002: invalidArgument("FORMULA_FIELD_IS_NOT_NUMBER"),
	43003: invalidArgument("FORMULA_FIELD_IS_NOT_INT"),
	43004: invalidArgument("FORMULA_TYPE_ERROR"),
	43005: invalidArgument("FORMULA_UNKNOWN_OPERATOR"),
	43006: invalidArgument("FORMULA_NOT_DATE_TYPE"),
	43007: invalidArgument("FORMULA_DATE_OUT_OF_RANGE"),
	43008: invalidArgument("FORMULA_ERROR_PARAMS"),
	43009: invalidArgument("FORMULA_UNDEFINED_FUNCTION"),
	44001: invalidArgument("DASHBOARD_NOT_EXIST"),
	// 44001: invalidArgument("获取坚果云文件链接失败"),
	45001: failedPrecondition("发送qrobot邮件时检测到多个发送人"),
	45002: invalidArgument("qrobot邮件发送地址不正确"),
	45003: invalidArgument("错误的发件人类型"),
	45004: invalidArgument("邮件title过长，超过200"),
	45005: invalidArgument("qrobot邮件内容设置不正确"),
	45011: invalidArgument("超时提醒设置错误"),
	45012: invalidArgument("超时预警设置错误"),
	45021: invalidArgument("批量处理设置的处理节点不合法"),
	45022: invalidArgument("不合法的批量处理类型"),
	45033: invalidArgument("流程外qrobot设置错误"),
	45100: invalidArgument("角色设置错误"),
	45101: notfound("角色不存在"),
	45102: invalidArgument("部门设置错误"),
	45103: failedPrecondition("工作区成员数量受到限制"),
	45105: unauthorized("成员禁止编辑"),
	45106: unauthorized("部门禁止编辑"),
	45107: unauthorized("角色禁止编辑"),
	45108: invalidArgument("不足够的成员信息"),
	45109: alreadyExists("角色已经存在"),
	45110: alreadyExists("部门已经存在"),
	45111: alreadyExists("用户已经存在"),
	45112: alreadyExists("外部userId已经存在"),
	45114: alreadyExists("该邮箱已经存在"),
	46000: invalidArgument("PLUGIN_CENTER_PLUGIN_NOT_EXIST"),
	46001: invalidArgument("申请已经被处理或处理人已无权限"),
	46003: invalidArgument("字段类型转换设置错误"),
	46004: invalidArgument("报表设置错误"),
	46011: invalidArgument("自定义编号设置错误"),
	46021: notfound("短信模板找不到"),
	46031: invalidArgument("远程查询设置错误"),
	46032: invalidArgument("远程查询失败"),
	46041: invalidArgument("传进来的模板Id为null"),
	46042: notfound("模板不存在"),
	46043: notfound("模板url不存在"),
	46051: invalidArgument("qrobot节点关联的应用已经被删除"),
	46061: invalidArgument("数据关联筛选条件里关联问题已经被删除"),
	46071: notfound("解决方案不存在"),
	46072: invalidArgument("解决方案信息不全"),
	46073: invalidArgument("解决方案类型错误"),
	47001: invalidArgument("字段默认值过长"),
	47021: invalidArgument("不能重复申请"),
	47022: invalidArgument("未提交手写签名"),
	47101: invalidArgument("权限组未找到"),
	47201: invalidArgument("启用人数超过限制"),
	47202: invalidArgument("创建人不可被禁用"),
	47203: invalidArgument("此用户已被禁用"),
	47204: invalidArgument("不能禁用自己"),
	47301: invalidArgument("非法的listType"),
	47302: invalidArgument("没有评论权限"),
	47303: invalidArgument("评论文本和附件不能都为空"),
	47304: invalidArgument("评论不存在"),
	47401: invalidArgument("电子签章未开启"),
	47402: invalidArgument("电子签章申请中,无法重复申请"),
	47403: invalidArgument("电子签章申请信息残缺"),
	47404: invalidArgument("电子签章申请不存在"),
	47405: invalidArgument("个人认证真实姓名未填写"),
	47406: invalidArgument("获取第三方登录token失败"),
	47407: invalidArgument("第三方api请求失败"),
	47408: invalidArgument("文件下载出错"),
	47409: invalidArgument("印章创建过程出错"),
	47410: invalidArgument("电子签章参数残缺"),
	47411: invalidArgument("个人实名认证未完成"),
	47412: unauthorized("当前用户无权使用该印章"),
	47413: invalidArgument("印章不存在"),
	47414: invalidArgument("签署文件类型错误"),
	47415: invalidArgument("电子签章已经申请过"),
	47416: invalidArgument("工作区个人实名认证用户已达上限(试用)"),
	47417: invalidArgument("电子签章签署次数不足"),
	47418: invalidArgument("印章名称不可重复"),
	47419: invalidArgument("签署文件体积过大(不超过50MB)"),
	47420: invalidArgument("文件未签署"),
	47421: alreadyExists("存在重复值"),
	47501: invalidArgument("参数残缺"),
	47502: invalidArgument("serverName不存在"),
	47503: invalidArgument("未配置单点登录"),
	47504: invalidArgument("单点登录请求用户授权部分未配置"),
	47505: invalidArgument("单点登录请求AccessToken部分未配置"),
	47506: invalidArgument("获取AccessToken请求失败"),
	47507: invalidArgument("单点登录请求Uid部分未配置"),
	47508: invalidArgument("单点登录请求Uid信息请求错误"),
	47509: invalidArgument("单点登录未解析到第三方userId"),
	47510: invalidArgument("参数缺少orgId"),
	47511: invalidArgument("配置错误"),
	47601: invalidArgument("Q_SOURCE_NOT_EXIST"),
	47602: invalidArgument("Q_SOURCE_CLOSED"),
	47603: invalidArgument("Q_SOURCE_TYPE_ERROR"),
	47604: invalidArgument("Q_SOURCE_FLAG_QUE_NOT_EXIST"),
	47605: invalidArgument("Q_SOURCE_METHOD_ERROR"),
	47606: invalidArgument("Q_SOURCE_URL_ERROR"),
	47607: invalidArgument("Q_SOURCE_BODY_TYPE_ERROR"),
	47608: invalidArgument("Q_SOURCE_INIT_PAGE_NUM_ERROR"),
	47609: invalidArgument("Q_SOURCE_PARSE_RESULT_ERROR"),
	47610: invalidArgument("Q_SOURCE_CREATE_JOB_ERROR"),
	47611: invalidArgument("Q_SOURCE_JOB_EXECUTING"),
	47612: invalidArgument("Q_SOURCE_JOB_DURATION_TOO_SHORT"),
	47613: invalidArgument("ACTIVE_Q_SOURCE_NUM_EXCEED"),
	48001: invalidArgument("创建用户失败"),
	48002: invalidArgument("无效的auth"),
	48003: invalidArgument("无效的memberName"),
	48004: invalidArgument("无效的memberId"),
	48005: invalidArgument("无效的wsId"),
	48006: invalidArgument("无效的appInfo"),
	48007: invalidArgument("无效的member信息"),
	48008: invalidArgument("无效的同步类型"),
	48009: invalidArgument("管理后台service请求失败"),
	48010: invalidArgument("无效的solutionId"),
	48011: invalidArgument("无效的promptId"),
	48012: invalidArgument("MANAGEMENT_INVALID_DEMOTE_CONFIG"),
	48013: invalidArgument("无效的newfunctionId"),
	48020: invalidArgument("功能暂不可用，请稍候再试"),
	48101: invalidArgument("分组名不存在"),
	48102: invalidArgument("无效操作"),
	48103: invalidArgument("分组不存在"),
	48104: invalidArgument("参数错误"),
	48105: invalidArgument("分组下有未删除的应用或仪表盘"),
	48130: invalidArgument("逐级审批暂不支持撤回"),
	48201: invalidArgument("发起支付触碰到了限流"),
	48202: invalidArgument("不完整的支付信息"),
	48203: invalidArgument("支付金额与数据不一致"),
	48204: invalidArgument("不合法的支付设备"),
	48205: invalidArgument("不合法的支付方式"),
	48206: invalidArgument("支付金额为负数"),
	48207: invalidArgument("支付内容与数据不一致"),
	48208: invalidArgument("不是基础版，不能确认协议"),
	48209: invalidArgument("没有确认协议，不能执行开启在线支付的操作"),
	48210: invalidArgument("商品内容过长"),
	48211: invalidArgument("AppId重复"),
	48212: invalidArgument("商户号重复"),
	48213: invalidArgument("url地址不能为空"),
	48301: invalidArgument("获取TeamBition成员错误"),
	48302: invalidArgument("禁止登录"),
	48303: invalidArgument("获取Teambition企业拥有者失败"),
	48304: invalidArgument("获取Teambition部门失败"),
	48305: invalidArgument("Teambition"),
	48306: invalidArgument("获取Teambition部门成员失败"),
	48307: invalidArgument("获取Teambition用户信息失败"),
	48308: invalidArgument("同步teambition项目信息出错"),
	48309: invalidArgument("同步teambition项目成员信息出错"),
	48310: invalidArgument("当前工作区不是teambition工作区"),
	48401: invalidArgument("企业检索失败"),
	48402: invalidArgument("企业检索关键字非法"),
	48403: invalidArgument("用户搜索次数到达上限"),
	49001: invalidArgument("共享服务token鉴权失败"),
	49002: invalidArgument("管理后台openApi鉴权失败"),
	49005: invalidArgument("数据不在对应的节点上"),
	49010: invalidArgument("系统繁忙"),
	49100: invalidArgument("ocr次数不足"),
	49101: invalidArgument("ocr模板被删除"),
	49200: invalidArgument("无可打印的内容"),
	49201: invalidArgument("打印模板不存在"),
	49202: invalidArgument("单日打印量超出上限"),
	49203: invalidArgument("单次打印量超出上限"),
	49205: invalidArgument("打印时数据不在对应的节点"),
	49206: invalidArgument("不明原因批量打印错误"),
	49207: invalidArgument("自定义文件名过长"),
	49300: invalidArgument("无效的accessToken"),
	49301: invalidArgument("达到每日调用次数上限"),
	49302: invalidArgument("IP不在可调用白名单内"),
	49303: notfound("不存在的appKey"),
	49304: notfound("不存在的apply"),
	49305: notfound("不存在的邮箱"),
	49306: notfound("不存在的邮箱"),
	49307: notfound("该用户不为工作区成员"),
	49308: notfound("不存在的requestId"),
	49400: notfound("自定义按钮不存在"),
	49401: failedPrecondition("报表配置的自定义按钮超过限额"),
	49402: failedPrecondition("报表与自定义按钮的关联已删除"),
	49500: invalidArgument("系统请求外部资源超时"),
	49501: invalidArgument("无效的secret"),
	49602: invalidArgument("PPLUGIN_ONLINEPAY_ESIGN_REFUSE_TRAIL"),
	49603: invalidArgument("密码不满足要求"),
	49604: notfound("state不存在"),
	49605: notfound("OAUTH不存在"),
	49606: notfound("OAUTH的请求用户授权的url不存在"),
	49607: invalidArgument("参数异常"),
	49608: invalidArgument("鉴权没有被激活"),
	49609: invalidArgument("鉴权没有被激活"),
}
