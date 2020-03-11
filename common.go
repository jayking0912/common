package common

//发送类型
type SendType int32

const (
	ConnectInfo SendType = 0
	ReceiveData SendType = 1
	ToolPset    SendType = 2
	MqttReceive SendType = 3
	//Policy_MID      SendType = 2
	//Policy_AVG      SendType = 3
)

type RecipeInfo struct {
	TargetPoint [4]byte //目标点位
	TargetPset  int     //目标pset
	Xtolerance  int     //x公差
	Ytolerance  int     //y公差

}

type ToolData struct {
	TightingID               string //螺丝标识，防止重复
	ToolIP                   string //工具ip/com
	PsetIndex                string //pset
	FinalTorque              string //最终扭矩 nm
	FinalAngle               string //最终角度 nm
	TorqueMax                string //最大扭矩
	TorqueMin                string //最小扭矩
	AngleMax                 string //最大角度
	AngleMin                 string //最小角度
	TargetTorque             string //目标扭矩
	Status                   string //拧紧状态
	Tighteningbitfieldstatus string //es only 错误标识
}

type CommonData struct{
	ClientID string
	Data string
	State bool
	Backup1 string
	Backup2 string
	Backup3 string
}

type MqttMessage struct{
	Topic string
	Message string
}
