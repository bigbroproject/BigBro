package data

type ServiceData struct {
	Name      string
	Err       error // if nil, the status is OK
	Protocols []ProtocolData
}
