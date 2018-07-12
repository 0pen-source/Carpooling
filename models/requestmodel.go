package models

// Phonetest _
type Phonetest struct {
	CommonPayload `structs:",flatten"`

	Phone            string `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
	VerificationCode string `binding:"omitempty,len=6,numeric" form:"verification_code" structs:"verification_code"`
}

// Phonetest _
type Login struct {
	CommonPayload `structs:",flatten"`

	Phone        string `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
	Password     string `binding:"-" form:"password" structs:"password"`
	Nickname     string `binding:"-" form:"nickname" structs:"nickname"`
	LastLocation string `binding:"-" form:"last_location" structs:"last_location"`
	Sex          int    `binding:"-" form:"sex" structs:"sex"`
	Username     string `binding:"-" form:"username" structs:"username"`
}
type Code struct {
	Code   string `binding:"-" form:"code" structs:"code"`
	Msg    string `binding:"-" form:"msg" structs:"msg"`
	SmUuid string `binding:"-" form:"smUuid" structs:"smUuid"`
}

// CommonPayload _
// it shouldn't be exported but it fails to bind, not familiar with go 1.8.3 features
type CommonPayload struct {
	IDFA          string `binding:"omitempty,ne=00000000-0000-0000-0000-000000000000" form:"idfa" structs:"idfa"`
	MACAddress    string `binding:"omitempty,len=32" form:"mac_address" structs:"mac_address"`
	IMEI          string `binding:"omitempty,len=32" form:"imei" structs:"imei"`
	IMSI          string `binding:"omitempty,len=32" form:"imsi" structs:"imsi"`
	AndroidID     string `binding:"omitempty,len=32" form:"android_id" structs:"android_id"`
	AdvertisingID string `binding:"-" form:"advertising_id" structs:"advertising_id"`
	Location      string `binding:"-" form:"location" structs:"location"`
}
