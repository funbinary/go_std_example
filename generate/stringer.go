package generate

//go:generate stringer -type=BACnetComfirmedServiceChoice -linecomment
type BACnetComfirmedServiceChoice byte

const (
	//报警与事件服务
	ServiceConfirmedAcknowledgeAlarm           BACnetComfirmedServiceChoice = 0 //确认报警服务
	ServiceConfirmedComfirmedCOVNotification   BACnetComfirmedServiceChoice = 1 //证实COV通告服务
	ServiceConfirmedComfirmedEventNotification BACnetComfirmedServiceChoice = 2 //证实事件通过服务
)
