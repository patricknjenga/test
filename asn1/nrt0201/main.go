package nrt0201

import (
	"encoding/asn1"
)

type Nrtrde struct {
	SpecificationVersionNumber SpecificationVersionNumber
	ReleaseVersionNumber       ReleaseVersionNumber
	Sender                     Sender
	Recipient                  Recipient
	SequenceNumber             SequenceNumber
	FileAvailableTimeStamp     FileAvailableTimeStamp
	UtcTimeOffset              UtcTimeOffset
	CallEvents                 CallEventList
	CallEventsCount            CallEventsCount
}
type CallEventList = []CallEvent
type CallEvent = asn1.RawValue
type Moc struct {
	Imsi                     Imsi
	Imei                     Imei
	CallEventStartTimeStamp  CallEventStartTimeStamp
	UtcTimeOffset            UtcTimeOffset
	CallEventDuration        CallEventDuration
	CauseForTermination      CauseForTermination
	ServiceCode              ServiceCode
	SupplementaryServiceCode SupplementaryServiceCode
	DialledDigits            DialledDigits
	ConnectedNumber          ConnectedNumber
	ThirdPartyNumber         ThirdPartyNumber
	RecEntityId              RecEntityId
	CallReference            CallReference
	ChargeAmount             ChargeAmount
	ServingNetwork           ServingNetwork
	Msisdn                   Msisdn
	LocationArea             LocationArea
	CellId                   CellId
}
type Mtc struct {
	Imsi                    Imsi
	Imei                    Imei
	CallEventStartTimeStamp CallEventStartTimeStamp
	UtcTimeOffset           UtcTimeOffset
	CallEventDuration       CallEventDuration
	CauseForTermination     CauseForTermination
	ServiceCode             ServiceCode
	CallingNumber           CallingNumber
	RecEntityId             RecEntityId
	CallReference           CallReference
	ChargeAmount            ChargeAmount
	ServingNetwork          ServingNetwork
	Msisdn                  Msisdn
	LocationArea            LocationArea
	CellId                  CellId
}
type Gprs struct {
	Imsi                    Imsi
	Imei                    Imei
	CallEventStartTimeStamp CallEventStartTimeStamp
	UtcTimeOffset           UtcTimeOffset
	CallEventDuration       CallEventDuration
	CauseForTermination     CauseForTermination
	AccessPointNameNI       AccessPointNameNI
	AccessPointNameOI       AccessPointNameOI
	DataVolumeIncoming      DataVolumeIncoming
	DataVolumeOutgoing      DataVolumeOutgoing
	SgsnAddress             SgsnAddress
	GgsnAddress             GgsnAddress
	ChargingId              ChargingId
	ChargeAmount            ChargeAmount
	ServingNetwork          ServingNetwork
	Msisdn                  Msisdn
	LocationArea            LocationArea
	CellId                  CellId
}
type AccessPointNameNI = AsciiString
type AccessPointNameOI = AsciiString
type BearerServiceCode = HexString
type CallEventDuration = int64
type CallEventsCount = int64
type CallEventStartTimeStamp = NumberString
type CallingNumber = NumberString
type CallReference = int64
type CauseForTermination = int64
type CellId = int64
type ChargeAmount = int64
type ChargingId = int64
type ConnectedNumber = NumberString
type DataVolumeIncoming = int64
type DataVolumeOutgoing = int64
type DialledDigits = AsciiString
type FileAvailableTimeStamp = NumberString
type GgsnAddress = AsciiString
type Imei = BCDString
type Imsi = BCDString
type LocationArea = int64
type Msisdn = BCDString
type RecEntityId = AsciiString
type Recipient = AsciiString
type ReleaseVersionNumber = int64
type Sender = AsciiString
type SequenceNumber = NumberString
type ServiceCode = interface {
}
type ServingNetwork = AsciiString
type SgsnAddress = AsciiString
type SpecificationVersionNumber = int64
type SupplementaryServiceCode = HexString
type TeleServiceCode = HexString
type ThirdPartyNumber = NumberString
type UtcTimeOffset = AsciiString
type AsciiString = string
type HexString = string
type NumberString = string
type BCDString = string
