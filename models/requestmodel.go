package models

// IOSPayload _
type Phonetest struct {
	CommonPayload `structs:",flatten"`

	Phone string `binding:"-" form:"phone" structs:"phone"`
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
