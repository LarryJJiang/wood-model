package bizcode

const (
	SUCCESS                    = 200   // 成功
	ERROR                      = 500   // 系统错误
	InvalidParams              = 400   // 参数错误
	NoneToken                  = 401   // token不存在
	NotPermission              = 403   // 没有权限
	NotExist                   = 404   // 请求资源不存在
	NotPair                    = 10005 // 未配对
	NotLeaderUkey              = 10006 // 未绑定主任密钥
	ServerNotAuthorize         = 10007 // 服务端未授权
	RepeatedAction             = 10008 // 你的操作太快了，请稍候再试
	ErrorRequestHeader         = 10009 // 请求头部信息错误
	ErrorLiteError             = 10010 // qc服务器未知错误
	Unknown                    = 10011 // 未知错误
	DialTimeout                = 10012 // 访问超时
	PluginsServerUnavailable   = 10013 // PAICS_Plugins服务未启动
	ReplacedByNewConnection    = 10014 // 当前连接被新连接替换
	ErrorGetHardwareInfo       = 10015 // 获取设备硬件信息错误
	ErrorAuthorizeDisable      = 10016 // 授权不可用
	ErrorIdentity              = 10017 // 不正确的身份
	ErrorHardwareInfoShutdown  = 10018 // 获取设备硬件信息错误(服务器即将关机)
	ErrorRemoteSrvErr          = 10019 // 调用服务器出现错误
	ErrorCloudSrvErr           = 10020 // 云端服务调用失败
	ErrorRemoteSrvNotActiveErr = 10021 // 服务器离线
	ErrorRemoteSrvTimeoutErr   = 10022 // 访问服务器超时或服务器离线

	ErrorUsername              = 10101 // 用户名错误
	ErrorPassword              = 10102 // 用户密码错误
	ErrorAuthCheckTokenFail    = 10103 // 用户token验证失败
	ErrorAuthCheckTokenTimeout = 10104 // 用户token验证超时
	ErrorUsernameExists        = 10105 // 用户名已存在
	ErrorPhoneExists           = 10106 // 手机号已被绑定
	ErrorUser                  = 10107 // 用户不存在
	ErrorUserDescTooLong       = 10108 // 用户简介字数过长(最多800字符)
	ErrorUserRepeatLogin       = 10109 // 用户已登录，请勿重复登录
	ErrorUserOtherLogin        = 10110 // 账号已在另一部设备登录，请确认是否本人操作
	ErrorEmptyToken            = 10111 // 头部token为空

	ErrorPhoneCheckCode          = 10201 // 短信验证码错误
	ErrorPhoneCheckCodeTooOften  = 10202 // 短信验证码发送频率过于频繁
	ErrorPhoneCheckCodeOverLimit = 10203 // 短信验证码发送超过单日限制
	ErrorLicenceCheckCode        = 10204 // 证书验证码错误
	ErrorLoginCheckCode          = 10205 // 登录验证码错误

	ErrorUkeyNotExists          = 10301 // U-key不存在
	ErrorUkeyEncrypt            = 10302 // U-key加密错误
	ErrorUkeyDecrypt            = 10303 // U-key解密错误
	ErrorUkeyPublickeyNotExists = 10304 // U-key公钥不存在
	ErrorUkeyIdentifierExists   = 10305 // U-key代号已存在
	ErrorUkeyAvailableTimes     = 10306 // U-key没有可授权次数
	ErrorUkeyNotAuthor          = 10307 // U-key没有绑定授权用户
	ErrorUkeyExists             = 10308 // U-key已存在

	ErrorLicenceNotExists             = 10401 // 证书不存在
	ErrorLicenceDeviceNumberOverLimit = 10402 // 证书设备编号超限
	ErrorLicenceAuthorIdOverLimit     = 10403 // 证书授权id超限
	ErrorLicenceRelateUkey            = 10404 // 证书对应ukey不匹配
	ErrorLicenceVersionHigher         = 10405 // 证书对应版本号高于申请版本号（不予授权）
	ErrorLicenceVersionLower          = 10406 // 证书对应版本号低于申请版本号（重新授权）
	ErrorLicenceNotAuthorize          = 10407 // 未被授权过
	ErrorLicenceCheck                 = 10408 // 证书验证失败

	ErrorRedis      = 10501 // redis相关错误
	ErrorDb         = 10502 // 数据库操作相关错误
	ErrorImgBase64  = 10503 // 图片base64不正确
	ErrorImgExt     = 10504 // 仅支持png/jpg/jpeg/gif图片
	ErrorImgTooMax  = 10505 // 图片大小超过限制
	ErrorSystemLock = 10506 // 系统锁不存在BizError404900 = BizError{"404900", "系统锁不存在!"}

	ErrorNameExists = 10601 // 名称已存在

	ErrorLeaderUkeyNotExists           = 10701 // 主任密钥不存在
	ErrorLeaderUkeyDataFileRead        = 10702 // 主任密钥数据文件解析错误
	ErrorLeaderUkeyDataFileGenerate    = 10703 // 主任密钥数据文件生成错误
	ErrorLeaderUkeyNotPair             = 10704 // 此密钥非绑定主任密钥
	ErrorLeaderUkeyIdentifier          = 10705 // 主任密钥编号错误
	ErrorLeaderUkeyNotBind             = 10706 // 未绑定主任密钥, 不予更新
	ErrorLeaderUkeyVersionLower        = 10707 // 主任密钥用户版本不高于本地用户，不予更新
	ErrorLeaderUkeyRepeat              = 10710 // 不能重复绑定
	ErrorLeaderCustomerSettingParse    = 10711 // 主任密钥用户设置数据解析错误
	ErrorLeaderCustomerSettingGenerate = 10712 // 主任密钥用户设置数据生成错误

	ErrorFileFormat    = 10801 // 错误的文件格式
	ErrorFileSave      = 10802 // 文件保存失败
	ErrorFileNotExists = 10803 // 文件不存在

	ErrorPairAuthorNumber = 11001 // 客户端未配对本地服务器(旧称：证书授权id不在预配对客户端里）
	ErrorPairNotConfirm   = 11002 // 配对请求未被确认
	ErrorPairClient       = 11003 // 客户端未配对
	ErrorPairClientLogin  = 11004 // 有客户端在登录状态，无法取消配对证书授权id不在预配对客户端里
	ErrorPairClientCancel = 11005 // 取消配对

	// screening 检查设置
	ErrorScreeningUuid       = 11101 // 检查设置uuid不存在
	ErrorScreeningNameLength = 11102 // 请输入2～30位的检查名称
	ErrorScreeningActive     = 11103 // 必须有一个检查设置为启用状态
	ErrorScreeningNameRepeat = 11104 // 检查设置名称不能重复使用
	ErrorScreeningKeyData    = 11105 // 检查设置密钥数据错误
	ErrorScreeningEmpty      = 11106 // 检查设置数据为空，不能同步
	ErrorScreeningSaving     = 11107 // 检查设置正在保存中，请勿重复操作
	ErrorScreeningNoticing   = 11108 // 检查设置正在通知中，请我重复操作
	ErrorScreeningModify     = 11109 // 检查设置修改数据有误，请基于最新数据修改

	ErrorCustomerDefaultDelete = 11201 // 本地用户之默认用户无法删除
	ErrorCustomerNotExist      = 11202 // 用户不存在

	ErrorServerState = 11301 // 远程服务器不在状态（请检查是否联网）

	ERROR_UPLOAD_CASE_EXIST = 30002

	// case
	ErrorCaseStartTimeEmpty                   = 11401 // 请填写病例检查的开始时间
	ErrorCaseSecondEmpty                      = 11402 // 请填写病例录制的时长
	ErrorCasePathEmpty                        = 11403 // 请填写病例录制的视频路径
	ErrorCaseVideoPathExist                   = 11404 // 病例信息已存在
	ErrorCaseRobotInfoEmpty                   = 11405 // 请填写超声机信息
	ErrorCaseDeviceNumberEmpty                = 11406 // 请填写设备号
	ErrorCaseCustomerEmpty                    = 11407 // 请填写医生名字或客户端使用者的名字
	ErrorCaseSearchStartTime                  = 11408 // 请选择历史病例的开始时间
	ErrorCaseSearchEndTime                    = 11409 // 请选择历史病例的结束时间
	ErrorCaseInfoNotExist                     = 11410 // 病例不存在
	ErrorCaseInfoExist                        = 11411 // 历史病例已存在
	ErrorCaseUuidEmpty                        = 11412 // 请填写历史病例的id
	ErrorCaseDoctorNameEmpty                  = 11413 // 请填写历史病例的检查医生名字
	ErrorCasePath                             = 11414 // 病例路径错误或病例目录不存在
	ErrorCaseIsUploadedError                  = 11415 // 重复上传：病例已上传完成
	ErrorCaseIsUploading                      = 11416 // 重复操作：病例正在分析中
	ErrorCaseSharedUserNotExist               = 11417 // 共享用户不存在
	ErrorCaseSharedUserAddFail                = 11418 // 共享用户添加失败
	ErrorCaseNeedSetUploadingError            = 11419 // 请先设置为上传中，再执行上传完成操作
	ErrorCaseCannotSetUploading               = 11420 // 当前病例的上传状态不允许设置为上传中
	ErrorCaseSaveDiskFreeSize                 = 11421 // 病例存储盘空间不足
	ErrorCaseScreeningUuidEmpty               = 11422 // 请填写病例的检查类型UUID
	ErrorCaseClientPatientIdEmpty             = 11423 // 请填写病例的客户端PatientID
	ErrorCaseFetusNumErr                      = 11424 // 请填写病例的检查是第几胎(单胎：1；多胎检查不允许重复胎数；单胎不允许大于1)
	ErrorCasePatientMultipleSerialErr         = 11425 // 请填写病例的多胎检查的孕妇的序号(单胎为0，多胎必填；同一孕妇的多胎序号必须一致)
	ErrorCaseCaseSerialErr                    = 11426 // 请填写客户端的病例序号
	ErrorCasePatientIDIsUploading             = 11427 // 正在处理当前孕妇的case中，请稍后再试
	ErrorCaseScreeningUuidError               = 11428 // 病例的检查类型不存在
	ErrorCaseMultipleFetusNumErr              = 11429 // 客户端数据错误：存在相同的客户端PatientID且不是单胎检查第一胎
	ErrorCaseMultipleFetusNumRepeated         = 11430 // 客户端数据错误：存在相同的客户端PatientID且多胎胎数重复
	ErrorCaseMultiplePatientMultipleSerialErr = 11431 // 客户端数据错误：存在相同的客户端PatientID且多胎检查的孕妇的序号不一致
	ErrorCasePatientMultipleSerialDiff        = 11432 // 客户端数据错误：多胎检查的孕妇的序号不一致
	ErrorCaseFetusNumSameErr                  = 11433 // 客户端数据错误：多胎检查出现相同的胎儿序号或PatientID重复
	ErrorCaseFetusNumFirstErr                 = 11434 // 客户端数据错误：多胎检查出现多次第一胎单胎检查
	ErrorCaseClientPatientIDSameErr           = 11435 // 客户端数据错误：出现相同的PatientID且不是多胎检查
	ErrorCaseJoinSecondErr                    = 11436 // 客户端数据错误：续接病例的时长不能小于或等于之前的病例时长
	ErrorCaseJoinIsPlayingErr                 = 11437 // 此历史病例正在播放中，请稍后再试
	ErrorCaseFileNotExist                     = 11438 // 病例文件未上传完成(文件不存在)，请先上传文件再进行确认上传完成
	ErrorCaseImageNotExist                    = 11439 // 截图文件未上传完成(文件不存在)，请先上传文件再进行确认上传完成
	ErrorCaseAloneClientPatientIDSameErr      = 11440 // 单胎检查的客户端PatientID与其他case的客户端PatientID相同
	ErrorCaseMultipleClientPatientIDSameErr   = 11441 // 多胎检查的客户端PatientID与其他单胎的客户端PatientID相同
	ErrorCaseSnapshootError                   = 11442 // 病例之镜像错误
	ErrorCasePatientExist                     = 11443 // 已存在此病人信息
	ErrorCaseScreenShotExist                  = 11444 // 已存在相同时间的截图
	ErrorCaseScreenShotNotExist               = 11445 // 不存在此病例的截图
	ErrorAssessExist                          = 11446 // 不能重复提交答卷
	ErrorMeasuredExist                        = 11447 // 测量值已存在
	ErrorCaseFetusGroupNotExist               = 11448 // 病例分组不存在

	// 标签
	ErrorCaseLabelExist       = 11450 // 已存在同名自定义标签
	ErrorCaseLabelNotExist    = 11451 // 标签不存在
	ErrorCaseLabelCountError  = 11452 // 自定义标签数量已达100个限额
	ErrorCaseRelateLabelExist = 11453 // 病例已关联了此标签

	// 病例报告
	ErrorCaseReportNotExists = 11460 // 病例报告不存在

	// 病例备份
	ErrorCaseBackupCopyError           = 11501 // 病例备份复制错误
	ErrorCaseBackupWriteRecordError    = 11502 // 病例写记录错误
	ErrorCaseBackupDiskNotExists       = 11503 // 备份盘不存在
	ErrorCaseBackupRecordFileNotExists = 11504 // 未找到备份成功记录

	// 知识图谱错误码序号 12001 - 13999
	// 知识图谱 特征
	ErrFeaturePartNotExist             = 12001 // 特征部位不存在
	ErrFeaturePartNameError            = 12002 // 特征部位名称长度必须2～20字符
	ErrFeaturePartNameExist            = 12003 // 特征部位名称重复(已存在)
	ErrFeaturePartNumNotEnough         = 12004 // 特征ID编号已无空余
	ErrFeatureInUsedSection            = 12005 // 特征被切面使用中
	ErrFeatureInUsedSyndrome           = 12006 // 特征被综合征使用中
	ErrFeatureInUsedSectionAndSyndrome = 12007 // 特征被切面与综合征使用中
	ErrFeatureNameError                = 12008 // 特征名称长度必须2～20字符
	ErrFeatureDefineLongError          = 12012 // 特征的定义最多2000字符
	ErrFeatureDiagnoseLongError        = 12009 // 特征的超声诊断要点最多2000字符
	ErrFeatureConsultLongError         = 12010 // 特征的预后咨询最多2000字符
	ErrFeatureOtherLongError           = 12011 // 特征的其他说明最多2000字符
	ErrFeatureLegendOriginNameError    = 12013 // 请填写特征的图例名称
	ErrFeatureLegendFilePathError      = 12014 // 特征图例文件格式只能是jpg/png/mp4
	ErrFeatureLegendTypeError          = 12015 // 请选择特征图例类型
	ErrFeatureExist                    = 12016 // 特征名称重复(已存在)
	ErrFeaturePartUUIDMustNeed         = 12017 // 请传参：特征部位uuid
	ErrFeatureNotExist                 = 12018 // 特征不存在
	ErrFeaturePartNameEnError          = 12019 // 特征部位英文名称长度必须2～100字符
	ErrFeatureNameEnError              = 12020 // 特征英文名称长度必须2～100字符
	ErrFeatureDefineEnLongError        = 12021 // 特征的英文定义最多6000字符
	ErrFeatureDiagnoseEnLongError      = 12022 // 特征的英文超声诊断要点最多6000字符
	ErrFeatureConsultEnLongError       = 12023 // 特征的英文预后咨询最多6000字符
	ErrFeatureOtherEnLongError         = 12024 // 特征的英文其他说明最多6000字符
	ErrFeatureDefineCnEnError          = 12025 // 请输入特征的定义(相应的中英文说明)
	ErrFeatureDiagnoseCnEnError        = 12026 // 请输入特征的超声诊断要点(相应的中英文说明)
	ErrFeatureConsultCnEnError         = 12027 // 请输入特征的预后咨询(相应的中英文说明)
	ErrFeatureOtherCnEnError           = 12028 // 请输入特征的其他说明(相应的中英文说明)
	ErrFeaturePartNameEnExist          = 12029 // 特征部位英文名称重复(已存在)
	ErrFeatureEnExist                  = 12030 // 特征英文名称重复(已存在)
	ErrFeatureLegendThumbError         = 12031 // 请选择特征图例缩略图不正确

	// 知识图谱 切面
	ErrSectionNotExist               = 12101 // 切面不存在
	ErrSectionFeatureNotExist        = 12102 // 切面特征不存在
	ErrSectionUUIDEmpty              = 12103 // 请填写切面uuid
	ErrSectionNameError              = 12104 // 切面名称长度必须3～30字符
	ErrSectionCodeError              = 12105 // 请选择切面程序编码
	ErrSectionExist                  = 12106 // 切面已存在(切面名称或程序编码已被使用)
	ErrSectionNumNotEnough           = 12107 // 切面ID编号已无空余
	ErrSectionFeatureNameError       = 12108 // 切面特征名称长度必须2～30字符
	ErrSectionFeatureNormalCodeError = 12109 // 请选择切面正常特征程序编码
	ErrSectionFeatureNormalRepeated  = 12110 // 部位特征有重复(特征重复或特征的程序编号重复)
	ErrSectionFeatureNormalNotExist  = 12111 // 对立特征不存在(部位特征不存在或被删除)
	ErrSectionFeatureExist           = 12112 // 切面特征已存在(切面特征名称或程序编码已被使用)
	ErrSectionUUIDMustNeed           = 12113 // 请传参：切面uuid

	// 知识图谱 综合征
	ErrSyndromeNotExist                  = 12201 // 综合征不存在
	ErrSyndromeNumNotEnough              = 12202 // 综合征ID编号已无空余
	ErrSyndromeExist                     = 12203 // 综合征已存在(综合征名称已被使用)
	ErrSyndromeNameError                 = 12204 // 综合征名称长度必须2～30字符
	ErrSyndromeGeneticsDescError         = 12205 // 综合征的遗传类型最多30字符
	ErrSyndromeGeneLocationError         = 12206 // 综合征的基因位点最多30字符
	ErrSyndromeDiagnoseError             = 12207 // 综合征的超声诊断要点最多2000字符
	ErrSyndromeConsultError              = 12208 // 综合征的预后咨询最多2000字符
	ErrSyndromeClassCodeError            = 12209 // 请选择综合征分组类型或综合征分组
	ErrSyndromeUUIDMustNeed              = 12210 // 请传参：综合征uuid
	ErrSyndromeTypeNotRight              = 12211 // 综合征类型不正确
	ErrSyndromeUUIDRepeated              = 12212 // 综合征uuid重复
	ErrSyndromeFeatureUUIDRepeated       = 12213 // 综合征形态学uuid重复
	ErrSyndromeFeatureOptError           = 12214 // 请选择形态学特征常见状态
	ErrSyndromeClassNotGeneError         = 12215 // 遗传综合征排序类型错误(有不属于遗传综合征类型的综合征)
	ErrSyndromeNameEnError               = 12216 // 综合征英文名称长度必须2～100字符
	ErrSyndromeGeneticsDescEnError       = 12217 // 综合征的英文遗传类型最多200字符
	ErrSyndromeGeneticsDescCnEnError     = 12218 // 请输入综合征的遗传类型(相应的中英文说明)
	ErrSyndromeGeneticsLocationEnError   = 12219 // 综合征的英文基因位点最多200字符
	ErrSyndromeGeneticsLocationCnEnError = 12220 // 请输入综合征的基因位点(相应的中英文说明)
	ErrSyndromeDiagnoseEnError           = 12221 // 综合征的英文超声诊断要点最多6000字符
	ErrSyndromeDiagnoseCnEnError         = 12222 // 请输入综合征的超声诊断要点(相应的中英文说明)
	ErrSyndromeConsultEnError            = 12223 // 综合征的英文预后咨询最多6000字符
	ErrSyndromeConsultCnEnError          = 12224 // 请输入综合征的预后咨询(相应的中英文说明)
	ErrSyndromeEnExist                   = 12225 // 综合征已存在(综合征英文名称已被使用)
	ErrSyndromeNotExistOrRepeated        = 12226 // 综合征不存在或综合征重复

	// 知识图谱 版本
	ErrVersionNotExist                   = 12301 // 版本不存在
	ErrVersionNameError                  = 12302 // 版本名称需要4～30字符
	ErrVersionSupportCodeError           = 12303 // 请选择适配产品版本号(适配产品编号错误)
	ErrVersionExist                      = 12304 // 版本名称已存在(版本名称已被使用)
	ErrVersionNumNotEnough               = 12305 // 版本ID编号已无空余
	ErrVersionSupportRepeated            = 12306 // 适配产品版本号有重复
	ErrVersionSupportClientError         = 12307 // 客户端版本号错误(版本号格式：V_X.X.X.X)
	ErrVersionSupportServerError         = 12308 // 服务端版本号错误(版本号格式：V_X.X)
	ErrVersionSerialEmpty                = 12309 // 请填写版本编号
	ErrVersionSectionFeatureTypeError    = 12310 // 版本切面特征类型不正确
	ErrVersionSupportMustServerAndClient = 12311 // 适配产品版本号中必须有一个客户端和一个服务端
	ErrVersionSupportCodeMatchError      = 12312 // 适配产品编号错误

	// 知识图谱 导入导出
	ErrVersionExportImport         = 12401 // 知识图谱版本导出导入
	ErrVersionImportFileIncorrect  = 12402 // 知识图谱版本文件不正确(请重新导出文件)
	ErrVersionImportVersionExist   = 12403 // 知识图谱版本已存在，无需再次导入
	ErrVersionImportUnzipError     = 12404 // 知识图谱版本文件解压失败(请重新导出文件)
	ErrVersionImportFileNotExist   = 12405 // 知识图谱版本文件不存在，请重新上传
	ErrVersionSyncErrIsUsed        = 12406 // 知识图谱版本正在使用中，同步失败
	ErrVersionSyncErrDoNotSupport  = 12407 // 同步会造成已匹配客户端原匹配关系失效
	ErrVersionImportParamIncorrect = 12408 // websocket参数错误
	ErrVersionExportLockError      = 12409 // 知识图谱正在打包中，请3秒后再试
	ErrVersionImportDecryptError   = 12410 // 知识图谱解密失败

	// 知识图谱错误码序号 12001 - 13599

	// dll
	ErrDLLPasswordDecodeErr    = 13600 // 密码解码错误
	ErrDLLFindProcFailed       = 13601 // 查无此处理程序
	ErrDLLWinUserExists        = 13602 // 用户已存在
	ErrDLLWinUserPassShortErr  = 13603 // 用户密码过短
	ErrDLLWinUserNotExist      = 13604 // 用户不存在
	ErrDLLWinUsername          = 13605 // 用户名不规范
	ErrDLLWinUserCreatedFail   = 13606 // 用户创建失败
	ErrDLLWinUserDeletedFail   = 13607 // 用户删除失败
	ErrDLLWinUserSharedPathErr = 13608 // 用户的路径添加失败
	ErrDLLProcessErr           = 13609 // 调用成功，处理过程中发生错误
	ErrWMIProcessErr           = 13610 // 获取设备过程中发生错误

	// dll 13600 - 13999

	// 客户端
	ErrClientInteractiveParamIncorrect = 14001 // tcp参数错误
	ErrClientAuthorNumberEmpty         = 14002 // 请填写有效的客户端授权id
	ErrClientVersionSerialEmpty        = 14003 // 请填写版本序列号
	ErrClientVersionNotSupport         = 14004 // 知识图谱不适配当前客户端版本
	ErrPluginsServerDialTimeout        = 14005 // 访问插件服务超时
	ErrClientCaseUUIDEmpty             = 14006 // 请填写病例uuid
	ErrClientVoiceTextEmpty            = 14007 // 请输入搜索内容

	// 客户端 14000 - 14100

	// 设备管理
	ErrorDeviceNumberEmpty       = 14201 // 请填写设备号
	ErrorDeviceNeverSubmit       = 14202 // 此设备还未对接上报设备信息接口
	ErrorDeviceIPError           = 14203 // 设备IP异常
	ErrorDeviceFileSrvErr        = 14204 // 文件服务调用失败(文件不存在或无权限)
	ErrorDeviceFileTypeErr       = 14205 // 错误的文件类型
	ErrorDeviceLogStartTimeErr   = 14206 // 请选择日志开始时间
	ErrorDeviceLogEndTimeErr     = 14207 // 请选择日志结束时间
	ErrorDeviceAuthorNumberEmpty = 14208 // 请填写授权ID
	ErrorDeviceCannotGetIPErr    = 14209 // 目标已脱机，且无法获取IP
	ErrorDeviceLogQueryErr       = 14210 // ES查询日志错误

	// 图文报告
	ErrorCasePrintReportNotExist                = 14301 // 图文报告不存在
	ErrorCasePrintReportNameAlreadyExist        = 14302 // 报告名称已存在，请重新输入
	ErrorCasePrintReportNameEmpty               = 14303 // 请输入报告名称
	ErrorCasePrintReportHospitalNameEmpty       = 14304 // 请填写医院名称
	ErrorCasePrintReportHospitalDepartmentEmpty = 14305 // 请填写医院科室
	ErrorCasePrintReportDoctorNameEmpty         = 14306 // 请填写医生名称
	ErrorCasePrintReportPatientNameEmpty        = 14307 // 请输入姓名
	ErrorCasePrintReportPatientAgeEmpty         = 14308 // 请输入年龄
	ErrorCasePrintReportReportImagesEmpty       = 14309 // 请上传超声机图像
	ErrorCasePrintReportUltrasoundSeeEmpty      = 14310 // 请填写超声所见
	ErrorCasePrintReportUltrasoundTipEmpty      = 14311 // 请填写超声提示
	ErrorCasePrintReportClientReportSerialEmpty = 14312 // 请填写报告序号
	ErrorCasePrintReportImagesTooMany           = 14313 // 超声机图像最多添加3张
	ErrorCasePrintCaseUUIDEmpty                 = 14314 // 请填写病例id
	ErrorCasePrintCaseStartTimeEmpty            = 14315 // 请填写病例开始时间
	ErrorCasePrintCaseSecondEmpty               = 14316 // 请填写病例时长
	ErrorCasePrintReportTimeEmpty               = 14317 // 请填写报告时间
	ErrorCasePrintReportAlreadyExist            = 14318 // 病例已有报告，不可重复创建报告
)
