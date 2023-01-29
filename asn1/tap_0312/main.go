package tap_0312

import "encoding/asn1"

type DataInterChange = asn1.RawValue
type TransferBatch struct {
	BatchControlInfo       BatchControlInfo           `asn1:"optional"`
	AccountingInfo         AccountingInfo             `asn1:"optional"`
	NetworkInfo            NetworkInfo                `asn1:"optional"`
	MessageDescriptionInfo MessageDescriptionInfoList `asn1:"optional"`
	CallEventDetails       CallEventDetailList        `asn1:"optional"`
	AuditControlInfo       AuditControlInfo           `asn1:"optional"`
}
type Notification struct {
	Sender                     Sender                     `asn1:"optional"`
	Recipient                  Recipient                  `asn1:"optional"`
	FileSequenceNumber         FileSequenceNumber         `asn1:"optional"`
	RapFileSequenceNumber      RapFileSequenceNumber      `asn1:"optional"`
	FileCreationTimeStamp      FileCreationTimeStamp      `asn1:"optional"`
	FileAvailableTimeStamp     FileAvailableTimeStamp     `asn1:"optional"`
	TransferCutOffTimeStamp    TransferCutOffTimeStamp    `asn1:"optional"`
	SpecificationVersionNumber SpecificationVersionNumber `asn1:"optional"`
	ReleaseVersionNumber       ReleaseVersionNumber       `asn1:"optional"`
	FileTypeIndicator          FileTypeIndicator          `asn1:"optional"`
	OperatorSpecInformation    OperatorSpecInfoList       `asn1:"optional"`
}
type CallEventDetailList = []CallEventDetail
type CallEventDetail = asn1.RawValue
type BatchControlInfo struct {
	Sender                     Sender                     `asn1:"optional"`
	Recipient                  Recipient                  `asn1:"optional"`
	FileSequenceNumber         FileSequenceNumber         `asn1:"optional"`
	FileCreationTimeStamp      FileCreationTimeStamp      `asn1:"optional"`
	TransferCutOffTimeStamp    TransferCutOffTimeStamp    `asn1:"optional"`
	FileAvailableTimeStamp     FileAvailableTimeStamp     `asn1:"optional"`
	SpecificationVersionNumber SpecificationVersionNumber `asn1:"optional"`
	ReleaseVersionNumber       ReleaseVersionNumber       `asn1:"optional"`
	FileTypeIndicator          FileTypeIndicator          `asn1:"optional"`
	RapFileSequenceNumber      RapFileSequenceNumber      `asn1:"optional"`
	OperatorSpecInformation    OperatorSpecInfoList       `asn1:"optional"`
}
type AccountingInfo struct {
	Taxation               TaxationList           `asn1:"optional"`
	Discounting            DiscountingList        `asn1:"optional"`
	LocalCurrency          LocalCurrency          `asn1:"optional"`
	TapCurrency            TapCurrency            `asn1:"optional"`
	CurrencyConversionInfo CurrencyConversionList `asn1:"optional"`
	TapDecimalPlaces       TapDecimalPlaces       `asn1:"optional"`
}
type NetworkInfo struct {
	UtcTimeOffsetInfo UtcTimeOffsetInfoList `asn1:"optional"`
	RecEntityInfo     RecEntityInfoList     `asn1:"optional"`
}
type MessageDescriptionInfoList = []MessageDescriptionInformation
type MobileOriginatedCall struct {
	BasicCallInformation    MoBasicCallInformation `asn1:"optional"`
	LocationInformation     LocationInformation    `asn1:"optional"`
	EquipmentIdentifier     ImeiOrEsn              `asn1:"optional"`
	BasicServiceUsedList    BasicServiceUsedList   `asn1:"optional"`
	SupplServiceCode        SupplServiceCode       `asn1:"optional"`
	ThirdPartyInformation   ThirdPartyInformation  `asn1:"optional"`
	CamelServiceUsed        CamelServiceUsed       `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList   `asn1:"optional"`
}
type MobileTerminatedCall struct {
	BasicCallInformation    MtBasicCallInformation `asn1:"optional"`
	LocationInformation     LocationInformation    `asn1:"optional"`
	EquipmentIdentifier     ImeiOrEsn              `asn1:"optional"`
	BasicServiceUsedList    BasicServiceUsedList   `asn1:"optional"`
	CamelServiceUsed        CamelServiceUsed       `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList   `asn1:"optional"`
}
type SupplServiceEvent struct {
	ChargeableSubscriber    ChargeableSubscriber  `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber `asn1:"optional"`
	LocationInformation     LocationInformation   `asn1:"optional"`
	EquipmentIdentifier     ImeiOrEsn             `asn1:"optional"`
	SupplServiceUsed        SupplServiceUsed      `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList  `asn1:"optional"`
}
type ServiceCentreUsage struct {
	BasicInformation        ScuBasicInformation   `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber `asn1:"optional"`
	ServingNetwork          ServingNetwork        `asn1:"optional"`
	RecEntityCode           RecEntityCode         `asn1:"optional"`
	ChargeInformation       ChargeInformation     `asn1:"optional"`
	ScuChargeType           ScuChargeType         `asn1:"optional"`
	ScuTimeStamps           ScuTimeStamps         `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList  `asn1:"optional"`
}
type GprsCall struct {
	GprsBasicCallInformation GprsBasicCallInformation `asn1:"optional"`
	GprsLocationInformation  GprsLocationInformation  `asn1:"optional"`
	EquipmentIdentifier      ImeiOrEsn                `asn1:"optional"`
	GprsServiceUsed          GprsServiceUsed          `asn1:"optional"`
	CamelServiceUsed         CamelServiceUsed         `asn1:"optional"`
	OperatorSpecInformation  OperatorSpecInfoList     `asn1:"optional"`
}
type ContentTransaction struct {
	ContentTransactionBasicInfo ContentTransactionBasicInfo `asn1:"optional"`
	ChargedPartyInformation     ChargedPartyInformation     `asn1:"optional"`
	ServingPartiesInformation   ServingPartiesInformation   `asn1:"optional"`
	ContentServiceUsed          ContentServiceUsedList      `asn1:"optional"`
	OperatorSpecInformation     OperatorSpecInfoList        `asn1:"optional"`
}
type LocationService struct {
	RapFileSequenceNumber       RapFileSequenceNumber       `asn1:"optional"`
	RecEntityCode               RecEntityCode               `asn1:"optional"`
	CallReference               CallReference               `asn1:"optional"`
	TrackingCustomerInformation TrackingCustomerInformation `asn1:"optional"`
	LCSSPInformation            LCSSPInformation            `asn1:"optional"`
	TrackedCustomerInformation  TrackedCustomerInformation  `asn1:"optional"`
	LocationServiceUsage        LocationServiceUsage        `asn1:"optional"`
	OperatorSpecInformation     OperatorSpecInfoList        `asn1:"optional"`
}
type MessagingEvent struct {
	MessagingEventService   MessagingEventService `asn1:"optional"`
	ChargedParty            ChargedParty          `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber `asn1:"optional"`
	SimToolkitIndicator     SimToolkitIndicator   `asn1:"optional"`
	GeographicalLocation    GeographicalLocation  `asn1:"optional"`
	EventReference          EventReference        `asn1:"optional"`
	RecEntityCodeList       RecEntityCodeList     `asn1:"optional"`
	NetworkElementList      NetworkElementList    `asn1:"optional"`
	LocationArea            LocationArea          `asn1:"optional"`
	CellId                  CellId                `asn1:"optional"`
	ServiceStartTimestamp   ServiceStartTimestamp `asn1:"optional"`
	NonChargedParty         NonChargedParty       `asn1:"optional"`
	ExchangeRateCode        ExchangeRateCode      `asn1:"optional"`
	CallTypeGroup           CallTypeGroup         `asn1:"optional"`
	Charge                  Charge                `asn1:"optional"`
	TaxInformationList      TaxInformationList    `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList  `asn1:"optional"`
}
type MobileSession struct {
	MobileSessionService    MobileSessionService   `asn1:"optional"`
	ChargedParty            ChargedParty           `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber  `asn1:"optional"`
	SimToolkitIndicator     SimToolkitIndicator    `asn1:"optional"`
	GeographicalLocation    GeographicalLocation   `asn1:"optional"`
	LocationArea            LocationArea           `asn1:"optional"`
	CellId                  CellId                 `asn1:"optional"`
	EventReference          EventReference         `asn1:"optional"`
	RecEntityCodeList       RecEntityCodeList      `asn1:"optional"`
	ServiceStartTimestamp   ServiceStartTimestamp  `asn1:"optional"`
	CauseForTerm            CauseForTerm           `asn1:"optional"`
	TotalCallEventDuration  TotalCallEventDuration `asn1:"optional"`
	NonChargedParty         NonChargedParty        `asn1:"optional"`
	RequestedDestination    RequestedDestination   `asn1:"optional"`
	SessionChargeInfoList   SessionChargeInfoList  `asn1:"optional"`
	OperatorSpecInformation OperatorSpecInfoList   `asn1:"optional"`
}
type AuditControlInfo struct {
	EarliestCallTimeStamp       EarliestCallTimeStamp       `asn1:"optional"`
	LatestCallTimeStamp         LatestCallTimeStamp         `asn1:"optional"`
	TotalCharge                 TotalCharge                 `asn1:"optional"`
	TotalChargeRefund           TotalChargeRefund           `asn1:"optional"`
	TotalTaxRefund              TotalTaxRefund              `asn1:"optional"`
	TotalTaxValue               TotalTaxValue               `asn1:"optional"`
	TotalDiscountValue          TotalDiscountValue          `asn1:"optional"`
	TotalDiscountRefund         TotalDiscountRefund         `asn1:"optional"`
	TotalAdvisedChargeValueList TotalAdvisedChargeValueList `asn1:"optional"`
	CallEventDetailsCount       CallEventDetailsCount       `asn1:"optional"`
	OperatorSpecInformation     OperatorSpecInfoList        `asn1:"optional"`
}
type AccessPointNameNI = AsciiString
type AccessPointNameOI = AsciiString
type ActualDeliveryTimeStamp = DateTime
type AddressStringDigits = BCDString
type AdvisedCharge = Charge
type AdvisedChargeCurrency = Currency
type AdvisedChargeInformation struct {
	PaidIndicator         PaidIndicator         `asn1:"optional"`
	PaymentMethod         PaymentMethod         `asn1:"optional"`
	AdvisedChargeCurrency AdvisedChargeCurrency `asn1:"optional"`
	AdvisedCharge         AdvisedCharge         `asn1:"optional"`
	Commission            Commission            `asn1:"optional"`
}
type AgeOfLocation = int64
type BasicService struct {
	ServiceCode           BasicServiceCode      `asn1:"optional"`
	TransparencyIndicator TransparencyIndicator `asn1:"optional"`
	Fnur                  Fnur                  `asn1:"optional"`
	UserProtocolIndicator UserProtocolIndicator `asn1:"optional"`
	GuaranteedBitRate     GuaranteedBitRate     `asn1:"optional"`
	MaximumBitRate        MaximumBitRate        `asn1:"optional"`
}
type BasicServiceCode = interface {
}
type BasicServiceCodeList = []BasicServiceCode
type BasicServiceUsed struct {
	BasicService          BasicService          `asn1:"optional"`
	ChargingTimeStamp     ChargingTimeStamp     `asn1:"optional"`
	ChargeInformationList ChargeInformationList `asn1:"optional"`
	HSCSDIndicator        HSCSDIndicator        `asn1:"optional"`
}
type BasicServiceUsedList = []BasicServiceUsed
type BearerServiceCode = HexString
type CalledNumber = AddressStringDigits
type CalledPlace = AsciiString
type CalledRegion = AsciiString
type CallEventDetailsCount = int64
type CallEventStartTimeStamp = DateTime
type CallingNumber = AddressStringDigits
type CallOriginator struct {
	CallingNumber CallingNumber `asn1:"optional"`
	ClirIndicator ClirIndicator `asn1:"optional"`
	SMSOriginator SMSOriginator `asn1:"optional"`
}
type CallReference = []byte
type CallTypeGroup struct {
	CallTypeLevel1 CallTypeLevel1 `asn1:"optional"`
	CallTypeLevel2 CallTypeLevel2 `asn1:"optional"`
	CallTypeLevel3 CallTypeLevel3 `asn1:"optional"`
}
type CallTypeLevel1 = int64
type CallTypeLevel2 = int64
type CallTypeLevel3 = int64
type CamelDestinationNumber = AddressStringDigits
type CamelInvocationFee = AbsoluteAmount
type CamelServiceKey = int64
type CamelServiceLevel = int64
type CamelServiceUsed struct {
	CamelServiceLevel      CamelServiceLevel            `asn1:"optional"`
	CamelServiceKey        CamelServiceKey              `asn1:"optional"`
	DefaultCallHandling    DefaultCallHandlingIndicator `asn1:"optional"`
	ExchangeRateCode       ExchangeRateCode             `asn1:"optional"`
	TaxInformation         TaxInformationList           `asn1:"optional"`
	DiscountInformation    DiscountInformation          `asn1:"optional"`
	CamelInvocationFee     CamelInvocationFee           `asn1:"optional"`
	ThreeGcamelDestination ThreeGcamelDestination       `asn1:"optional"`
	CseInformation         CseInformation               `asn1:"optional"`
}
type CauseForTerm = int64
type CellId = int64
type Charge = AbsoluteAmount
type ChargeableSubscriber = asn1.RawValue
type ChargeableUnits = int64
type ChargeDetail struct {
	ChargeType            ChargeType            `asn1:"optional"`
	Charge                Charge                `asn1:"optional"`
	ChargeableUnits       ChargeableUnits       `asn1:"optional"`
	ChargedUnits          ChargedUnits          `asn1:"optional"`
	ChargeDetailTimeStamp ChargeDetailTimeStamp `asn1:"optional"`
}
type ChargeDetailList = []ChargeDetail
type ChargeDetailTimeStamp = ChargingTimeStamp
type ChargedItem = AsciiString
type ChargedParty struct {
	Imsi                    Imsi                    `asn1:"optional"`
	Msisdn                  Msisdn                  `asn1:"optional"`
	PublicUserId            PublicUserId            `asn1:"optional"`
	HomeBid                 HomeBid                 `asn1:"optional"`
	HomeLocationDescription HomeLocationDescription `asn1:"optional"`
	Imei                    Imei                    `asn1:"optional"`
}
type ChargedPartyEquipment struct {
	EquipmentIdType EquipmentIdType `asn1:"optional"`
	EquipmentId     EquipmentId     `asn1:"optional"`
}
type ChargedPartyHomeIdentification struct {
	HomeIdType     HomeIdType     `asn1:"optional"`
	HomeIdentifier HomeIdentifier `asn1:"optional"`
}
type ChargedPartyHomeIdList = []ChargedPartyHomeIdentification
type ChargedPartyIdentification struct {
	ChargedPartyIdType     ChargedPartyIdType     `asn1:"optional"`
	ChargedPartyIdentifier ChargedPartyIdentifier `asn1:"optional"`
}
type ChargedPartyIdentifier = AsciiString
type ChargedPartyIdList = []ChargedPartyIdentification
type ChargedPartyIdType = int64
type ChargedPartyInformation struct {
	ChargedPartyIdList       ChargedPartyIdList       `asn1:"optional"`
	ChargedPartyHomeIdList   ChargedPartyHomeIdList   `asn1:"optional"`
	ChargedPartyLocationList ChargedPartyLocationList `asn1:"optional"`
	ChargedPartyEquipment    ChargedPartyEquipment    `asn1:"optional"`
}
type ChargedPartyLocation struct {
	LocationIdType     LocationIdType     `asn1:"optional"`
	LocationIdentifier LocationIdentifier `asn1:"optional"`
}
type ChargedPartyLocationList = []ChargedPartyLocation
type ChargedPartyStatus = int64
type ChargedUnits = int64
type ChargeInformation struct {
	ChargedItem         ChargedItem         `asn1:"optional"`
	ExchangeRateCode    ExchangeRateCode    `asn1:"optional"`
	CallTypeGroup       CallTypeGroup       `asn1:"optional"`
	ChargeDetailList    ChargeDetailList    `asn1:"optional"`
	TaxInformation      TaxInformationList  `asn1:"optional"`
	DiscountInformation DiscountInformation `asn1:"optional"`
}
type ChargeInformationList = []ChargeInformation
type ChargeRefundIndicator = int64
type ChargeType = NumberString
type ChargingId = int64
type ChargingPoint = AsciiString
type ChargingTimeStamp = DateTime
type ClirIndicator = int64
type Commission = Charge
type CompletionTimeStamp = DateTime
type ContentChargingPoint = int64
type ContentProvider struct {
	ContentProviderIdType     ContentProviderIdType     `asn1:"optional"`
	ContentProviderIdentifier ContentProviderIdentifier `asn1:"optional"`
}
type ContentProviderIdentifier = AsciiString
type ContentProviderIdList = []ContentProvider
type ContentProviderIdType = int64
type ContentProviderName = AsciiString
type ContentServiceUsed struct {
	ContentTransactionCode       ContentTransactionCode       `asn1:"optional"`
	ContentTransactionType       ContentTransactionType       `asn1:"optional"`
	ObjectType                   ObjectType                   `asn1:"optional"`
	TransactionDescriptionSupp   TransactionDescriptionSupp   `asn1:"optional"`
	TransactionShortDescription  TransactionShortDescription  `asn1:"optional"`
	TransactionDetailDescription TransactionDetailDescription `asn1:"optional"`
	TransactionIdentifier        TransactionIdentifier        `asn1:"optional"`
	TransactionAuthCode          TransactionAuthCode          `asn1:"optional"`
	DataVolumeIncoming           DataVolumeIncoming           `asn1:"optional"`
	DataVolumeOutgoing           DataVolumeOutgoing           `asn1:"optional"`
	TotalDataVolume              TotalDataVolume              `asn1:"optional"`
	ChargeRefundIndicator        ChargeRefundIndicator        `asn1:"optional"`
	ContentChargingPoint         ContentChargingPoint         `asn1:"optional"`
	ChargeInformationList        ChargeInformationList        `asn1:"optional"`
	AdvisedChargeInformation     AdvisedChargeInformation     `asn1:"optional"`
}
type ContentServiceUsedList = []ContentServiceUsed
type ContentTransactionBasicInfo struct {
	RapFileSequenceNumber      RapFileSequenceNumber      `asn1:"optional"`
	OrderPlacedTimeStamp       OrderPlacedTimeStamp       `asn1:"optional"`
	RequestedDeliveryTimeStamp RequestedDeliveryTimeStamp `asn1:"optional"`
	ActualDeliveryTimeStamp    ActualDeliveryTimeStamp    `asn1:"optional"`
	TotalTransactionDuration   TotalTransactionDuration   `asn1:"optional"`
	TransactionStatus          TransactionStatus          `asn1:"optional"`
}
type ContentTransactionCode = int64
type ContentTransactionType = int64
type CseInformation = []byte
type CurrencyConversion struct {
	ExchangeRateCode      ExchangeRateCode      `asn1:"optional"`
	NumberOfDecimalPlaces NumberOfDecimalPlaces `asn1:"optional"`
	ExchangeRate          ExchangeRate          `asn1:"optional"`
}
type CurrencyConversionList = []CurrencyConversion
type CustomerIdentifier = AsciiString
type CustomerIdType = int64
type DataVolume = int64
type DataVolumeIncoming = DataVolume
type DataVolumeOutgoing = DataVolume
type DateTime struct {
	LocalTimeStamp    LocalTimeStamp    `asn1:"optional"`
	UtcTimeOffsetCode UtcTimeOffsetCode `asn1:"optional"`
}
type DateTimeLong struct {
	LocalTimeStamp LocalTimeStamp `asn1:"optional"`
	UtcTimeOffset  UtcTimeOffset  `asn1:"optional"`
}
type DefaultCallHandlingIndicator = int64
type DepositTimeStamp = DateTime
type Destination struct {
	CalledNumber         CalledNumber         `asn1:"optional"`
	DialledDigits        DialledDigits        `asn1:"optional"`
	CalledPlace          CalledPlace          `asn1:"optional"`
	CalledRegion         CalledRegion         `asn1:"optional"`
	SMSDestinationNumber SMSDestinationNumber `asn1:"optional"`
}
type DestinationNetwork = NetworkId
type DialledDigits = AsciiString
type Discount = DiscountValue
type DiscountableAmount = AbsoluteAmount
type DiscountApplied = interface {
}
type DiscountCode = int64
type DiscountInformation struct {
	DiscountCode       DiscountCode       `asn1:"optional"`
	Discount           Discount           `asn1:"optional"`
	DiscountableAmount DiscountableAmount `asn1:"optional"`
}
type Discounting struct {
	DiscountCode    DiscountCode    `asn1:"optional"`
	DiscountApplied DiscountApplied `asn1:"optional"`
}
type DiscountingList = []Discounting
type DiscountRate = PercentageRate
type DiscountValue = AbsoluteAmount
type DistanceChargeBandCode = AsciiString
type EarliestCallTimeStamp = DateTimeLong
type ElementId = AsciiString
type ElementType = int64
type EquipmentId = AsciiString
type EquipmentIdType = int64
type Esn = NumberString
type EventReference = AsciiString
type ExchangeRate = int64
type ExchangeRateCode = Code
type FileAvailableTimeStamp = DateTimeLong
type FileCreationTimeStamp = DateTimeLong
type FileSequenceNumber = NumberString
type FileTypeIndicator = AsciiString
type FixedDiscountValue = DiscountValue
type Fnur = int64
type GeographicalLocation struct {
	ServingNetwork             ServingNetwork             `asn1:"optional"`
	ServingBid                 ServingBid                 `asn1:"optional"`
	ServingLocationDescription ServingLocationDescription `asn1:"optional"`
}
type GprsBasicCallInformation struct {
	GprsChargeableSubscriber GprsChargeableSubscriber `asn1:"optional"`
	RapFileSequenceNumber    RapFileSequenceNumber    `asn1:"optional"`
	GprsDestination          GprsDestination          `asn1:"optional"`
	CallEventStartTimeStamp  CallEventStartTimeStamp  `asn1:"optional"`
	TotalCallEventDuration   TotalCallEventDuration   `asn1:"optional"`
	CauseForTerm             CauseForTerm             `asn1:"optional"`
	PartialTypeIndicator     PartialTypeIndicator     `asn1:"optional"`
	PDPContextStartTimestamp PDPContextStartTimestamp `asn1:"optional"`
	NetworkInitPDPContext    NetworkInitPDPContext    `asn1:"optional"`
	ChargingId               ChargingId               `asn1:"optional"`
}
type GprsChargeableSubscriber struct {
	ChargeableSubscriber    ChargeableSubscriber    `asn1:"optional"`
	PdpAddress              PdpAddress              `asn1:"optional"`
	NetworkAccessIdentifier NetworkAccessIdentifier `asn1:"optional"`
}
type GprsDestination struct {
	AccessPointNameNI AccessPointNameNI `asn1:"optional"`
	AccessPointNameOI AccessPointNameOI `asn1:"optional"`
}
type GprsLocationInformation struct {
	GprsNetworkLocation     GprsNetworkLocation     `asn1:"optional"`
	HomeLocationInformation HomeLocationInformation `asn1:"optional"`
	GeographicalLocation    GeographicalLocation    `asn1:"optional"`
}
type GprsNetworkLocation struct {
	RecEntity    RecEntityCodeList `asn1:"optional"`
	LocationArea LocationArea      `asn1:"optional"`
	CellId       CellId            `asn1:"optional"`
}
type GprsServiceUsed struct {
	IMSSignallingContext  IMSSignallingContext  `asn1:"optional"`
	DataVolumeIncoming    DataVolumeIncoming    `asn1:"optional"`
	DataVolumeOutgoing    DataVolumeOutgoing    `asn1:"optional"`
	ChargeInformationList ChargeInformationList `asn1:"optional"`
}
type GsmChargeableSubscriber struct {
	Imsi   Imsi   `asn1:"optional"`
	Msisdn Msisdn `asn1:"optional"`
}
type GuaranteedBitRate = []byte
type HomeBid = Bid
type HomeIdentifier = AsciiString
type HomeIdType = int64
type HomeLocationDescription = LocationDescription
type HomeLocationInformation struct {
	HomeBid                 HomeBid                 `asn1:"optional"`
	HomeLocationDescription HomeLocationDescription `asn1:"optional"`
}
type HorizontalAccuracyDelivered = int64
type HorizontalAccuracyRequested = int64
type HSCSDIndicator = AsciiString
type Imei = BCDString
type ImeiOrEsn = interface {
}
type Imsi = BCDString
type IMSSignallingContext = int64
type InternetServiceProvider struct {
	IspIdType     IspIdType     `asn1:"optional"`
	IspIdentifier IspIdentifier `asn1:"optional"`
}
type InternetServiceProviderIdList = []InternetServiceProvider
type IspIdentifier = AsciiString
type IspIdType = int64
type ISPList = []InternetServiceProvider
type NetworkIdType = int64
type NetworkIdentifier = AsciiString
type Network struct {
	NetworkIdType     NetworkIdType     `asn1:"optional"`
	NetworkIdentifier NetworkIdentifier `asn1:"optional"`
}
type NetworkList = []Network
type LatestCallTimeStamp = DateTimeLong
type LCSQosDelivered struct {
	LCSTransactionStatus        LCSTransactionStatus        `asn1:"optional"`
	HorizontalAccuracyDelivered HorizontalAccuracyDelivered `asn1:"optional"`
	VerticalAccuracyDelivered   VerticalAccuracyDelivered   `asn1:"optional"`
	ResponseTime                ResponseTime                `asn1:"optional"`
	PositioningMethod           PositioningMethod           `asn1:"optional"`
	TrackingPeriod              TrackingPeriod              `asn1:"optional"`
	TrackingFrequency           TrackingFrequency           `asn1:"optional"`
	AgeOfLocation               AgeOfLocation               `asn1:"optional"`
}
type LCSQosRequested struct {
	LCSRequestTimestamp         LCSRequestTimestamp         `asn1:"optional"`
	HorizontalAccuracyRequested HorizontalAccuracyRequested `asn1:"optional"`
	VerticalAccuracyRequested   VerticalAccuracyRequested   `asn1:"optional"`
	ResponseTimeCategory        ResponseTimeCategory        `asn1:"optional"`
	TrackingPeriod              TrackingPeriod              `asn1:"optional"`
	TrackingFrequency           TrackingFrequency           `asn1:"optional"`
}
type LCSRequestTimestamp = DateTime
type LCSSPIdentification struct {
	ContentProviderIdType     ContentProviderIdType     `asn1:"optional"`
	ContentProviderIdentifier ContentProviderIdentifier `asn1:"optional"`
}
type LCSSPIdentificationList = []LCSSPIdentification
type LCSSPInformation struct {
	LCSSPIdentificationList LCSSPIdentificationList `asn1:"optional"`
	ISPList                 ISPList                 `asn1:"optional"`
	NetworkList             NetworkList             `asn1:"optional"`
}
type LCSTransactionStatus = int64
type LocalCurrency = Currency
type LocalTimeStamp = NumberString
type LocationArea = int64
type LocationDescription = AsciiString
type LocationIdentifier = AsciiString
type LocationIdType = int64
type LocationInformation struct {
	NetworkLocation         NetworkLocation         `asn1:"optional"`
	HomeLocationInformation HomeLocationInformation `asn1:"optional"`
	GeographicalLocation    GeographicalLocation    `asn1:"optional"`
}
type LocationServiceUsage struct {
	LCSQosRequested       LCSQosRequested       `asn1:"optional"`
	LCSQosDelivered       LCSQosDelivered       `asn1:"optional"`
	ChargingTimeStamp     ChargingTimeStamp     `asn1:"optional"`
	ChargeInformationList ChargeInformationList `asn1:"optional"`
}
type MaximumBitRate = []byte
type Mdn = NumberString
type MessageDescription = AsciiString
type MessageDescriptionCode = Code
type MessageDescriptionInformation struct {
	MessageDescriptionCode MessageDescriptionCode `asn1:"optional"`
	MessageDescription     MessageDescription     `asn1:"optional"`
}
type MessageStatus = int64
type MessageType = int64
type MessagingEventService = int64
type Min = NumberString
type MinChargeableSubscriber struct {
	Min Min `asn1:"optional"`
	Mdn Mdn `asn1:"optional"`
}
type MoBasicCallInformation struct {
	ChargeableSubscriber    ChargeableSubscriber    `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber   `asn1:"optional"`
	Destination             Destination             `asn1:"optional"`
	DestinationNetwork      DestinationNetwork      `asn1:"optional"`
	CallEventStartTimeStamp CallEventStartTimeStamp `asn1:"optional"`
	TotalCallEventDuration  TotalCallEventDuration  `asn1:"optional"`
	SimToolkitIndicator     SimToolkitIndicator     `asn1:"optional"`
	CauseForTerm            CauseForTerm            `asn1:"optional"`
}
type MobileSessionService = int64
type Msisdn = BCDString
type MtBasicCallInformation struct {
	ChargeableSubscriber    ChargeableSubscriber    `asn1:"optional"`
	RapFileSequenceNumber   RapFileSequenceNumber   `asn1:"optional"`
	CallOriginator          CallOriginator          `asn1:"optional"`
	OriginatingNetwork      OriginatingNetwork      `asn1:"optional"`
	CallEventStartTimeStamp CallEventStartTimeStamp `asn1:"optional"`
	TotalCallEventDuration  TotalCallEventDuration  `asn1:"optional"`
	SimToolkitIndicator     SimToolkitIndicator     `asn1:"optional"`
	CauseForTerm            CauseForTerm            `asn1:"optional"`
}
type NetworkAccessIdentifier = AsciiString
type NetworkElement struct {
	ElementType ElementType `asn1:"optional"`
	ElementId   ElementId   `asn1:"optional"`
}
type NetworkElementList = []NetworkElement
type NetworkId = AsciiString
type NetworkInitPDPContext = int64
type NetworkLocation struct {
	RecEntityCode RecEntityCode `asn1:"optional"`
	CallReference CallReference `asn1:"optional"`
	LocationArea  LocationArea  `asn1:"optional"`
	CellId        CellId        `asn1:"optional"`
}
type NonChargedNumber = AsciiString
type NonChargedParty struct {
	NonChargedPartyNumber  NonChargedPartyNumber  `asn1:"optional"`
	NonChargedPublicUserId NonChargedPublicUserId `asn1:"optional"`
}
type NonChargedPartyNumber = AddressStringDigits
type NonChargedPublicUserId = AsciiString
type NumberOfDecimalPlaces = int64
type ObjectType = int64
type OperatorSpecInfoList = []OperatorSpecInformation
type OperatorSpecInformation = AsciiString
type OrderPlacedTimeStamp = DateTime
type OriginatingNetwork = NetworkId
type PacketDataProtocolAddress = AsciiString
type PaidIndicator = int64
type PartialTypeIndicator = AsciiString
type PaymentMethod = int64
type PdpAddress = PacketDataProtocolAddress
type PDPContextStartTimestamp = DateTime
type PlmnId = AsciiString
type PositioningMethod = int64
type PriorityCode = int64
type PublicUserId = AsciiString
type RapFileSequenceNumber = FileSequenceNumber
type RecEntityCode = Code
type RecEntityCodeList = []RecEntityCode
type RecEntityId = AsciiString
type RecEntityInfoList = []RecEntityInformation
type RecEntityInformation struct {
	RecEntityCode RecEntityCode `asn1:"optional"`
	RecEntityType RecEntityType `asn1:"optional"`
	RecEntityId   RecEntityId   `asn1:"optional"`
}
type RecEntityType = int64
type Recipient = PlmnId
type ReleaseVersionNumber = int64
type RequestedDeliveryTimeStamp = DateTime
type RequestedDestination struct {
	RequestedNumber       RequestedNumber       `asn1:"optional"`
	RequestedPublicUserId RequestedPublicUserId `asn1:"optional"`
}
type RequestedNumber = AddressStringDigits
type RequestedPublicUserId = AsciiString
type ResponseTime = int64
type ResponseTimeCategory = int64
type ScuBasicInformation struct {
	ChargeableSubscriber ScuChargeableSubscriber `asn1:"optional"`
	ChargedPartyStatus   ChargedPartyStatus      `asn1:"optional"`
	NonChargedNumber     NonChargedNumber        `asn1:"optional"`
	ClirIndicator        ClirIndicator           `asn1:"optional"`
	OriginatingNetwork   OriginatingNetwork      `asn1:"optional"`
	DestinationNetwork   DestinationNetwork      `asn1:"optional"`
}
type ScuChargeType struct {
	MessageStatus          MessageStatus          `asn1:"optional"`
	PriorityCode           PriorityCode           `asn1:"optional"`
	DistanceChargeBandCode DistanceChargeBandCode `asn1:"optional"`
	MessageType            MessageType            `asn1:"optional"`
	MessageDescriptionCode MessageDescriptionCode `asn1:"optional"`
}
type ScuTimeStamps struct {
	DepositTimeStamp    DepositTimeStamp    `asn1:"optional"`
	CompletionTimeStamp CompletionTimeStamp `asn1:"optional"`
	ChargingPoint       ChargingPoint       `asn1:"optional"`
}
type ScuChargeableSubscriber = asn1.RawValue
type Sender = PlmnId
type ServiceStartTimestamp = DateTime
type ServingBid = Bid
type ServingLocationDescription = LocationDescription
type ServingNetwork = AsciiString
type ServingPartiesInformation struct {
	ContentProviderName           ContentProviderName           `asn1:"optional"`
	ContentProviderIdList         ContentProviderIdList         `asn1:"optional"`
	InternetServiceProviderIdList InternetServiceProviderIdList `asn1:"optional"`
	NetworkList                   NetworkList                   `asn1:"optional"`
}
type SessionChargeInfoList = []SessionChargeInformation
type SessionChargeInformation struct {
	ChargedItem        ChargedItem        `asn1:"optional"`
	ExchangeRateCode   ExchangeRateCode   `asn1:"optional"`
	CallTypeGroup      CallTypeGroup      `asn1:"optional"`
	ChargeDetail       ChargeDetail       `asn1:"optional"`
	ChargeDetailList   ChargeDetailList   `asn1:"optional"`
	TaxInformationList TaxInformationList `asn1:"optional"`
}
type SimChargeableSubscriber struct {
	Imsi   Imsi   `asn1:"optional"`
	Msisdn Msisdn `asn1:"optional"`
}
type SimToolkitIndicator = AsciiString
type SMSDestinationNumber = AsciiString
type SMSOriginator = AsciiString
type SpecificationVersionNumber = int64
type SsParameters = AsciiString
type SupplServiceActionCode = int64
type SupplServiceCode = HexString
type SupplServiceUsed struct {
	SupplServiceCode       SupplServiceCode       `asn1:"optional"`
	SupplServiceActionCode SupplServiceActionCode `asn1:"optional"`
	SsParameters           SsParameters           `asn1:"optional"`
	ChargingTimeStamp      ChargingTimeStamp      `asn1:"optional"`
	ChargeInformation      ChargeInformation      `asn1:"optional"`
	BasicServiceCodeList   BasicServiceCodeList   `asn1:"optional"`
}
type TapCurrency = Currency
type TapDecimalPlaces = int64
type TaxableAmount = AbsoluteAmount
type Taxation struct {
	TaxCode      TaxCode      `asn1:"optional"`
	TaxType      TaxType      `asn1:"optional"`
	TaxRate      TaxRate      `asn1:"optional"`
	ChargeType   ChargeType   `asn1:"optional"`
	TaxIndicator TaxIndicator `asn1:"optional"`
}
type TaxationList = []Taxation
type TaxCode = int64
type TaxIndicator = AsciiString
type TaxInformation struct {
	TaxCode       TaxCode       `asn1:"optional"`
	TaxValue      TaxValue      `asn1:"optional"`
	TaxableAmount TaxableAmount `asn1:"optional"`
}
type TaxInformationList = []TaxInformation
type TaxRate = NumberString
type TaxType = AsciiString
type TaxValue = AbsoluteAmount
type TeleServiceCode = HexString
type ThirdPartyInformation struct {
	ThirdPartyNumber ThirdPartyNumber `asn1:"optional"`
	ClirIndicator    ClirIndicator    `asn1:"optional"`
}
type ThirdPartyNumber = AddressStringDigits
type ThreeGcamelDestination = asn1.RawValue
type TotalAdvisedCharge = AbsoluteAmount
type TotalAdvisedChargeRefund = AbsoluteAmount
type TotalAdvisedChargeValue struct {
	AdvisedChargeCurrency    AdvisedChargeCurrency    `asn1:"optional"`
	TotalAdvisedCharge       TotalAdvisedCharge       `asn1:"optional"`
	TotalAdvisedChargeRefund TotalAdvisedChargeRefund `asn1:"optional"`
	TotalCommission          TotalCommission          `asn1:"optional"`
	TotalCommissionRefund    TotalCommissionRefund    `asn1:"optional"`
}
type TotalAdvisedChargeValueList = []TotalAdvisedChargeValue
type TotalCallEventDuration = int64
type TotalCharge = AbsoluteAmount
type TotalChargeRefund = AbsoluteAmount
type TotalCommission = AbsoluteAmount
type TotalCommissionRefund = AbsoluteAmount
type TotalDataVolume = DataVolume
type TotalDiscountRefund = AbsoluteAmount
type TotalDiscountValue = AbsoluteAmount
type TotalTaxRefund = AbsoluteAmount
type TotalTaxValue = AbsoluteAmount
type TotalTransactionDuration = TotalCallEventDuration
type TrackedCustomerEquipment struct {
	EquipmentIdType EquipmentIdType `asn1:"optional"`
	EquipmentId     EquipmentId     `asn1:"optional"`
}
type TrackedCustomerHomeId struct {
	HomeIdType     HomeIdType     `asn1:"optional"`
	HomeIdentifier HomeIdentifier `asn1:"optional"`
}
type TrackedCustomerHomeIdList = []TrackedCustomerHomeId
type TrackedCustomerIdentification struct {
	CustomerIdType     CustomerIdType     `asn1:"optional"`
	CustomerIdentifier CustomerIdentifier `asn1:"optional"`
}
type TrackedCustomerIdList = []TrackedCustomerIdentification
type TrackedCustomerInformation struct {
	TrackedCustomerIdList     TrackedCustomerIdList     `asn1:"optional"`
	TrackedCustomerHomeIdList TrackedCustomerHomeIdList `asn1:"optional"`
	TrackedCustomerLocList    TrackedCustomerLocList    `asn1:"optional"`
	TrackedCustomerEquipment  TrackedCustomerEquipment  `asn1:"optional"`
}
type TrackedCustomerLocation struct {
	LocationIdType     LocationIdType     `asn1:"optional"`
	LocationIdentifier LocationIdentifier `asn1:"optional"`
}
type TrackedCustomerLocList = []TrackedCustomerLocation
type TrackingCustomerEquipment struct {
	EquipmentIdType EquipmentIdType `asn1:"optional"`
	EquipmentId     EquipmentId     `asn1:"optional"`
}
type TrackingCustomerHomeId struct {
	HomeIdType     HomeIdType     `asn1:"optional"`
	HomeIdentifier HomeIdentifier `asn1:"optional"`
}
type TrackingCustomerHomeIdList = []TrackingCustomerHomeId
type TrackingCustomerIdentification struct {
	CustomerIdType     CustomerIdType     `asn1:"optional"`
	CustomerIdentifier CustomerIdentifier `asn1:"optional"`
}
type TrackingCustomerIdList = []TrackingCustomerIdentification
type TrackingCustomerInformation struct {
	TrackingCustomerIdList     TrackingCustomerIdList     `asn1:"optional"`
	TrackingCustomerHomeIdList TrackingCustomerHomeIdList `asn1:"optional"`
	TrackingCustomerLocList    TrackingCustomerLocList    `asn1:"optional"`
	TrackingCustomerEquipment  TrackingCustomerEquipment  `asn1:"optional"`
}
type TrackingCustomerLocation struct {
	LocationIdType     LocationIdType     `asn1:"optional"`
	LocationIdentifier LocationIdentifier `asn1:"optional"`
}
type TrackingCustomerLocList = []TrackingCustomerLocation
type TrackingFrequency = int64
type TrackingPeriod = int64
type TransactionAuthCode = AsciiString
type TransactionDescriptionSupp = int64
type TransactionDetailDescription = AsciiString
type TransactionIdentifier = AsciiString
type TransactionShortDescription = AsciiString
type TransactionStatus = int64
type TransferCutOffTimeStamp = DateTimeLong
type TransparencyIndicator = int64
type UserProtocolIndicator = int64
type UtcTimeOffset = AsciiString
type UtcTimeOffsetCode = Code
type UtcTimeOffsetInfo struct {
	UtcTimeOffsetCode UtcTimeOffsetCode `asn1:"optional"`
	UtcTimeOffset     UtcTimeOffset     `asn1:"optional"`
}
type UtcTimeOffsetInfoList = []UtcTimeOffsetInfo
type VerticalAccuracyDelivered = int64
type VerticalAccuracyRequested = int64
type AbsoluteAmount = int64
type Bid = AsciiString
type Code = int64
type AsciiString = []byte
type BCDString = []byte
type Currency = []byte
type HexString = []byte
type NumberString = []byte
type PercentageRate = int64
