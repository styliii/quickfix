package collateralinquiry

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//CollateralInquiry is the fix50sp2 CollateralInquiry type, MsgType = BB
type CollateralInquiry struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a CollateralInquiry from a quickfix.Message instance
func FromMessage(m quickfix.Message) CollateralInquiry {
	return CollateralInquiry{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m CollateralInquiry) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a CollateralInquiry initialized with the required fields for CollateralInquiry
func New(collinquiryid field.CollInquiryIDField) (m CollateralInquiry) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BB"))
	m.Set(collinquiryid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg CollateralInquiry, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BB", r
}

//SetAccount sets Account, Tag 1
func (m CollateralInquiry) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m CollateralInquiry) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCurrency sets Currency, Tag 15
func (m CollateralInquiry) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m CollateralInquiry) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m CollateralInquiry) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetPrice sets Price, Tag 44
func (m CollateralInquiry) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m CollateralInquiry) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetQuantity sets Quantity, Tag 53
func (m CollateralInquiry) SetQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewQuantity(value, scale))
}

//SetSide sets Side, Tag 54
func (m CollateralInquiry) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m CollateralInquiry) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m CollateralInquiry) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m CollateralInquiry) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m CollateralInquiry) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetNoDlvyInst sets NoDlvyInst, Tag 85
func (m CollateralInquiry) SetNoDlvyInst(f NoDlvyInstRepeatingGroup) {
	m.SetGroup(f)
}

//SetIssuer sets Issuer, Tag 106
func (m CollateralInquiry) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m CollateralInquiry) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetNoExecs sets NoExecs, Tag 124
func (m CollateralInquiry) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m CollateralInquiry) SetAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestAmt(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m CollateralInquiry) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetStandInstDbType sets StandInstDbType, Tag 169
func (m CollateralInquiry) SetStandInstDbType(v int) {
	m.Set(field.NewStandInstDbType(v))
}

//SetStandInstDbName sets StandInstDbName, Tag 170
func (m CollateralInquiry) SetStandInstDbName(v string) {
	m.Set(field.NewStandInstDbName(v))
}

//SetStandInstDbID sets StandInstDbID, Tag 171
func (m CollateralInquiry) SetStandInstDbID(v string) {
	m.Set(field.NewStandInstDbID(v))
}

//SetSettlDeliveryType sets SettlDeliveryType, Tag 172
func (m CollateralInquiry) SetSettlDeliveryType(v int) {
	m.Set(field.NewSettlDeliveryType(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m CollateralInquiry) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m CollateralInquiry) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m CollateralInquiry) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m CollateralInquiry) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m CollateralInquiry) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m CollateralInquiry) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetSpread sets Spread, Tag 218
func (m CollateralInquiry) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m CollateralInquiry) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m CollateralInquiry) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m CollateralInquiry) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m CollateralInquiry) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m CollateralInquiry) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m CollateralInquiry) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m CollateralInquiry) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m CollateralInquiry) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m CollateralInquiry) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m CollateralInquiry) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m CollateralInquiry) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m CollateralInquiry) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m CollateralInquiry) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m CollateralInquiry) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m CollateralInquiry) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m CollateralInquiry) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m CollateralInquiry) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m CollateralInquiry) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m CollateralInquiry) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m CollateralInquiry) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m CollateralInquiry) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m CollateralInquiry) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPriceType sets PriceType, Tag 423
func (m CollateralInquiry) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m CollateralInquiry) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m CollateralInquiry) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m CollateralInquiry) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m CollateralInquiry) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m CollateralInquiry) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m CollateralInquiry) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m CollateralInquiry) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m CollateralInquiry) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m CollateralInquiry) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m CollateralInquiry) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m CollateralInquiry) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAccountType sets AccountType, Tag 581
func (m CollateralInquiry) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m CollateralInquiry) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m CollateralInquiry) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m CollateralInquiry) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m CollateralInquiry) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m CollateralInquiry) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m CollateralInquiry) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m CollateralInquiry) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m CollateralInquiry) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetSettlSessID sets SettlSessID, Tag 716
func (m CollateralInquiry) SetSettlSessID(v string) {
	m.Set(field.NewSettlSessID(v))
}

//SetSettlSessSubID sets SettlSessSubID, Tag 717
func (m CollateralInquiry) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

//SetResponseTransportType sets ResponseTransportType, Tag 725
func (m CollateralInquiry) SetResponseTransportType(v int) {
	m.Set(field.NewResponseTransportType(v))
}

//SetResponseDestination sets ResponseDestination, Tag 726
func (m CollateralInquiry) SetResponseDestination(v string) {
	m.Set(field.NewResponseDestination(v))
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m CollateralInquiry) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m CollateralInquiry) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768
func (m CollateralInquiry) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTerminationType sets TerminationType, Tag 788
func (m CollateralInquiry) SetTerminationType(v int) {
	m.Set(field.NewTerminationType(v))
}

//SetQtyType sets QtyType, Tag 854
func (m CollateralInquiry) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m CollateralInquiry) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m CollateralInquiry) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m CollateralInquiry) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m CollateralInquiry) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m CollateralInquiry) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetNoTrades sets NoTrades, Tag 897
func (m CollateralInquiry) SetNoTrades(f NoTradesRepeatingGroup) {
	m.SetGroup(f)
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m CollateralInquiry) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

//SetMarginExcess sets MarginExcess, Tag 899
func (m CollateralInquiry) SetMarginExcess(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginExcess(value, scale))
}

//SetTotalNetValue sets TotalNetValue, Tag 900
func (m CollateralInquiry) SetTotalNetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalNetValue(value, scale))
}

//SetCashOutstanding sets CashOutstanding, Tag 901
func (m CollateralInquiry) SetCashOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOutstanding(value, scale))
}

//SetCollInquiryID sets CollInquiryID, Tag 909
func (m CollateralInquiry) SetCollInquiryID(v string) {
	m.Set(field.NewCollInquiryID(v))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m CollateralInquiry) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m CollateralInquiry) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m CollateralInquiry) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m CollateralInquiry) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m CollateralInquiry) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m CollateralInquiry) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m CollateralInquiry) SetDeliveryType(v int) {
	m.Set(field.NewDeliveryType(v))
}

//SetEndAccruedInterestAmt sets EndAccruedInterestAmt, Tag 920
func (m CollateralInquiry) SetEndAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndAccruedInterestAmt(value, scale))
}

//SetStartCash sets StartCash, Tag 921
func (m CollateralInquiry) SetStartCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartCash(value, scale))
}

//SetEndCash sets EndCash, Tag 922
func (m CollateralInquiry) SetEndCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndCash(value, scale))
}

//SetNoCollInquiryQualifier sets NoCollInquiryQualifier, Tag 938
func (m CollateralInquiry) SetNoCollInquiryQualifier(f NoCollInquiryQualifierRepeatingGroup) {
	m.SetGroup(f)
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m CollateralInquiry) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m CollateralInquiry) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m CollateralInquiry) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m CollateralInquiry) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m CollateralInquiry) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m CollateralInquiry) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m CollateralInquiry) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m CollateralInquiry) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m CollateralInquiry) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m CollateralInquiry) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m CollateralInquiry) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m CollateralInquiry) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m CollateralInquiry) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m CollateralInquiry) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m CollateralInquiry) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m CollateralInquiry) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m CollateralInquiry) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m CollateralInquiry) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m CollateralInquiry) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m CollateralInquiry) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m CollateralInquiry) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m CollateralInquiry) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m CollateralInquiry) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayoutAmount sets OptPayoutAmount, Tag 1195
func (m CollateralInquiry) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m CollateralInquiry) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetValuationMethod sets ValuationMethod, Tag 1197
func (m CollateralInquiry) SetValuationMethod(v string) {
	m.Set(field.NewValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m CollateralInquiry) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m CollateralInquiry) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m CollateralInquiry) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m CollateralInquiry) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m CollateralInquiry) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m CollateralInquiry) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435
func (m CollateralInquiry) SetContractMultiplierUnit(v int) {
	m.Set(field.NewContractMultiplierUnit(v))
}

//SetFlowScheduleType sets FlowScheduleType, Tag 1439
func (m CollateralInquiry) SetFlowScheduleType(v int) {
	m.Set(field.NewFlowScheduleType(v))
}

//SetRestructuringType sets RestructuringType, Tag 1449
func (m CollateralInquiry) SetRestructuringType(v string) {
	m.Set(field.NewRestructuringType(v))
}

//SetSeniority sets Seniority, Tag 1450
func (m CollateralInquiry) SetSeniority(v string) {
	m.Set(field.NewSeniority(v))
}

//SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451
func (m CollateralInquiry) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

//SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452
func (m CollateralInquiry) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

//SetAttachmentPoint sets AttachmentPoint, Tag 1457
func (m CollateralInquiry) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

//SetDetachmentPoint sets DetachmentPoint, Tag 1458
func (m CollateralInquiry) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

//SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478
func (m CollateralInquiry) SetStrikePriceDeterminationMethod(v int) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

//SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479
func (m CollateralInquiry) SetStrikePriceBoundaryMethod(v int) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

//SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480
func (m CollateralInquiry) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

//SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481
func (m CollateralInquiry) SetUnderlyingPriceDeterminationMethod(v int) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

//SetOptPayoutType sets OptPayoutType, Tag 1482
func (m CollateralInquiry) SetOptPayoutType(v int) {
	m.Set(field.NewOptPayoutType(v))
}

//SetNoComplexEvents sets NoComplexEvents, Tag 1483
func (m CollateralInquiry) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

//GetAccount gets Account, Tag 1
func (m CollateralInquiry) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m CollateralInquiry) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m CollateralInquiry) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m CollateralInquiry) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m CollateralInquiry) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m CollateralInquiry) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m CollateralInquiry) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuantity gets Quantity, Tag 53
func (m CollateralInquiry) GetQuantity() (f field.QuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m CollateralInquiry) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m CollateralInquiry) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m CollateralInquiry) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m CollateralInquiry) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m CollateralInquiry) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoDlvyInst gets NoDlvyInst, Tag 85
func (m CollateralInquiry) GetNoDlvyInst() (NoDlvyInstRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDlvyInstRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetIssuer gets Issuer, Tag 106
func (m CollateralInquiry) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m CollateralInquiry) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoExecs gets NoExecs, Tag 124
func (m CollateralInquiry) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m CollateralInquiry) GetAccruedInterestAmt() (f field.AccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m CollateralInquiry) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbType gets StandInstDbType, Tag 169
func (m CollateralInquiry) GetStandInstDbType() (f field.StandInstDbTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbName gets StandInstDbName, Tag 170
func (m CollateralInquiry) GetStandInstDbName() (f field.StandInstDbNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbID gets StandInstDbID, Tag 171
func (m CollateralInquiry) GetStandInstDbID() (f field.StandInstDbIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDeliveryType gets SettlDeliveryType, Tag 172
func (m CollateralInquiry) GetSettlDeliveryType() (f field.SettlDeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m CollateralInquiry) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m CollateralInquiry) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m CollateralInquiry) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m CollateralInquiry) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m CollateralInquiry) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m CollateralInquiry) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m CollateralInquiry) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m CollateralInquiry) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m CollateralInquiry) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m CollateralInquiry) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m CollateralInquiry) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m CollateralInquiry) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m CollateralInquiry) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m CollateralInquiry) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m CollateralInquiry) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m CollateralInquiry) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m CollateralInquiry) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m CollateralInquiry) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m CollateralInquiry) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m CollateralInquiry) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m CollateralInquiry) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m CollateralInquiry) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m CollateralInquiry) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m CollateralInquiry) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m CollateralInquiry) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m CollateralInquiry) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m CollateralInquiry) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m CollateralInquiry) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m CollateralInquiry) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m CollateralInquiry) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m CollateralInquiry) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m CollateralInquiry) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m CollateralInquiry) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m CollateralInquiry) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m CollateralInquiry) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m CollateralInquiry) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m CollateralInquiry) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m CollateralInquiry) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m CollateralInquiry) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m CollateralInquiry) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m CollateralInquiry) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccountType gets AccountType, Tag 581
func (m CollateralInquiry) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m CollateralInquiry) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m CollateralInquiry) GetBenchmarkPrice() (f field.BenchmarkPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m CollateralInquiry) GetBenchmarkPriceType() (f field.BenchmarkPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m CollateralInquiry) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m CollateralInquiry) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m CollateralInquiry) GetBenchmarkSecurityID() (f field.BenchmarkSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m CollateralInquiry) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m CollateralInquiry) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessID gets SettlSessID, Tag 716
func (m CollateralInquiry) GetSettlSessID() (f field.SettlSessIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessSubID gets SettlSessSubID, Tag 717
func (m CollateralInquiry) GetSettlSessSubID() (f field.SettlSessSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetResponseTransportType gets ResponseTransportType, Tag 725
func (m CollateralInquiry) GetResponseTransportType() (f field.ResponseTransportTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetResponseDestination gets ResponseDestination, Tag 726
func (m CollateralInquiry) GetResponseDestination() (f field.ResponseDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m CollateralInquiry) GetBenchmarkSecurityIDSource() (f field.BenchmarkSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m CollateralInquiry) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768
func (m CollateralInquiry) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTerminationType gets TerminationType, Tag 788
func (m CollateralInquiry) GetTerminationType() (f field.TerminationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m CollateralInquiry) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m CollateralInquiry) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m CollateralInquiry) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m CollateralInquiry) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m CollateralInquiry) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m CollateralInquiry) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTrades gets NoTrades, Tag 897
func (m CollateralInquiry) GetNoTrades() (NoTradesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m CollateralInquiry) GetMarginRatio() (f field.MarginRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarginExcess gets MarginExcess, Tag 899
func (m CollateralInquiry) GetMarginExcess() (f field.MarginExcessField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalNetValue gets TotalNetValue, Tag 900
func (m CollateralInquiry) GetTotalNetValue() (f field.TotalNetValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOutstanding gets CashOutstanding, Tag 901
func (m CollateralInquiry) GetCashOutstanding() (f field.CashOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCollInquiryID gets CollInquiryID, Tag 909
func (m CollateralInquiry) GetCollInquiryID() (f field.CollInquiryIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m CollateralInquiry) GetAgreementDesc() (f field.AgreementDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m CollateralInquiry) GetAgreementID() (f field.AgreementIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m CollateralInquiry) GetAgreementDate() (f field.AgreementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartDate gets StartDate, Tag 916
func (m CollateralInquiry) GetStartDate() (f field.StartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndDate gets EndDate, Tag 917
func (m CollateralInquiry) GetEndDate() (f field.EndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m CollateralInquiry) GetAgreementCurrency() (f field.AgreementCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m CollateralInquiry) GetDeliveryType() (f field.DeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndAccruedInterestAmt gets EndAccruedInterestAmt, Tag 920
func (m CollateralInquiry) GetEndAccruedInterestAmt() (f field.EndAccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartCash gets StartCash, Tag 921
func (m CollateralInquiry) GetStartCash() (f field.StartCashField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndCash gets EndCash, Tag 922
func (m CollateralInquiry) GetEndCash() (f field.EndCashField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoCollInquiryQualifier gets NoCollInquiryQualifier, Tag 938
func (m CollateralInquiry) GetNoCollInquiryQualifier() (NoCollInquiryQualifierRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCollInquiryQualifierRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m CollateralInquiry) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m CollateralInquiry) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m CollateralInquiry) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m CollateralInquiry) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m CollateralInquiry) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m CollateralInquiry) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m CollateralInquiry) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m CollateralInquiry) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m CollateralInquiry) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m CollateralInquiry) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m CollateralInquiry) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m CollateralInquiry) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m CollateralInquiry) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m CollateralInquiry) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m CollateralInquiry) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m CollateralInquiry) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m CollateralInquiry) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m CollateralInquiry) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m CollateralInquiry) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m CollateralInquiry) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m CollateralInquiry) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m CollateralInquiry) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m CollateralInquiry) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutAmount gets OptPayoutAmount, Tag 1195
func (m CollateralInquiry) GetOptPayoutAmount() (f field.OptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m CollateralInquiry) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValuationMethod gets ValuationMethod, Tag 1197
func (m CollateralInquiry) GetValuationMethod() (f field.ValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m CollateralInquiry) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m CollateralInquiry) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m CollateralInquiry) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m CollateralInquiry) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m CollateralInquiry) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m CollateralInquiry) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435
func (m CollateralInquiry) GetContractMultiplierUnit() (f field.ContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlowScheduleType gets FlowScheduleType, Tag 1439
func (m CollateralInquiry) GetFlowScheduleType() (f field.FlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRestructuringType gets RestructuringType, Tag 1449
func (m CollateralInquiry) GetRestructuringType() (f field.RestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSeniority gets Seniority, Tag 1450
func (m CollateralInquiry) GetSeniority() (f field.SeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451
func (m CollateralInquiry) GetNotionalPercentageOutstanding() (f field.NotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452
func (m CollateralInquiry) GetOriginalNotionalPercentageOutstanding() (f field.OriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAttachmentPoint gets AttachmentPoint, Tag 1457
func (m CollateralInquiry) GetAttachmentPoint() (f field.AttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDetachmentPoint gets DetachmentPoint, Tag 1458
func (m CollateralInquiry) GetDetachmentPoint() (f field.DetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478
func (m CollateralInquiry) GetStrikePriceDeterminationMethod() (f field.StrikePriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479
func (m CollateralInquiry) GetStrikePriceBoundaryMethod() (f field.StrikePriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480
func (m CollateralInquiry) GetStrikePriceBoundaryPrecision() (f field.StrikePriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481
func (m CollateralInquiry) GetUnderlyingPriceDeterminationMethod() (f field.UnderlyingPriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutType gets OptPayoutType, Tag 1482
func (m CollateralInquiry) GetOptPayoutType() (f field.OptPayoutTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEvents gets NoComplexEvents, Tag 1483
func (m CollateralInquiry) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAccount returns true if Account is present, Tag 1
func (m CollateralInquiry) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m CollateralInquiry) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m CollateralInquiry) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m CollateralInquiry) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m CollateralInquiry) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasPrice returns true if Price is present, Tag 44
func (m CollateralInquiry) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m CollateralInquiry) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasQuantity returns true if Quantity is present, Tag 53
func (m CollateralInquiry) HasQuantity() bool {
	return m.Has(tag.Quantity)
}

//HasSide returns true if Side is present, Tag 54
func (m CollateralInquiry) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m CollateralInquiry) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m CollateralInquiry) HasText() bool {
	return m.Has(tag.Text)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m CollateralInquiry) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m CollateralInquiry) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasNoDlvyInst returns true if NoDlvyInst is present, Tag 85
func (m CollateralInquiry) HasNoDlvyInst() bool {
	return m.Has(tag.NoDlvyInst)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m CollateralInquiry) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m CollateralInquiry) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasNoExecs returns true if NoExecs is present, Tag 124
func (m CollateralInquiry) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m CollateralInquiry) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m CollateralInquiry) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasStandInstDbType returns true if StandInstDbType is present, Tag 169
func (m CollateralInquiry) HasStandInstDbType() bool {
	return m.Has(tag.StandInstDbType)
}

//HasStandInstDbName returns true if StandInstDbName is present, Tag 170
func (m CollateralInquiry) HasStandInstDbName() bool {
	return m.Has(tag.StandInstDbName)
}

//HasStandInstDbID returns true if StandInstDbID is present, Tag 171
func (m CollateralInquiry) HasStandInstDbID() bool {
	return m.Has(tag.StandInstDbID)
}

//HasSettlDeliveryType returns true if SettlDeliveryType is present, Tag 172
func (m CollateralInquiry) HasSettlDeliveryType() bool {
	return m.Has(tag.SettlDeliveryType)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m CollateralInquiry) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m CollateralInquiry) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m CollateralInquiry) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m CollateralInquiry) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m CollateralInquiry) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m CollateralInquiry) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasSpread returns true if Spread is present, Tag 218
func (m CollateralInquiry) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m CollateralInquiry) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m CollateralInquiry) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m CollateralInquiry) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m CollateralInquiry) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m CollateralInquiry) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m CollateralInquiry) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m CollateralInquiry) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m CollateralInquiry) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m CollateralInquiry) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m CollateralInquiry) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m CollateralInquiry) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m CollateralInquiry) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m CollateralInquiry) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m CollateralInquiry) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m CollateralInquiry) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m CollateralInquiry) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m CollateralInquiry) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m CollateralInquiry) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m CollateralInquiry) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m CollateralInquiry) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m CollateralInquiry) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m CollateralInquiry) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m CollateralInquiry) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m CollateralInquiry) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m CollateralInquiry) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m CollateralInquiry) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m CollateralInquiry) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m CollateralInquiry) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m CollateralInquiry) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m CollateralInquiry) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m CollateralInquiry) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m CollateralInquiry) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m CollateralInquiry) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m CollateralInquiry) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m CollateralInquiry) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m CollateralInquiry) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m CollateralInquiry) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m CollateralInquiry) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m CollateralInquiry) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m CollateralInquiry) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m CollateralInquiry) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m CollateralInquiry) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m CollateralInquiry) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasSettlSessID returns true if SettlSessID is present, Tag 716
func (m CollateralInquiry) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

//HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717
func (m CollateralInquiry) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

//HasResponseTransportType returns true if ResponseTransportType is present, Tag 725
func (m CollateralInquiry) HasResponseTransportType() bool {
	return m.Has(tag.ResponseTransportType)
}

//HasResponseDestination returns true if ResponseDestination is present, Tag 726
func (m CollateralInquiry) HasResponseDestination() bool {
	return m.Has(tag.ResponseDestination)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m CollateralInquiry) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m CollateralInquiry) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768
func (m CollateralInquiry) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m CollateralInquiry) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m CollateralInquiry) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m CollateralInquiry) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m CollateralInquiry) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m CollateralInquiry) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m CollateralInquiry) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m CollateralInquiry) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasNoTrades returns true if NoTrades is present, Tag 897
func (m CollateralInquiry) HasNoTrades() bool {
	return m.Has(tag.NoTrades)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m CollateralInquiry) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasMarginExcess returns true if MarginExcess is present, Tag 899
func (m CollateralInquiry) HasMarginExcess() bool {
	return m.Has(tag.MarginExcess)
}

//HasTotalNetValue returns true if TotalNetValue is present, Tag 900
func (m CollateralInquiry) HasTotalNetValue() bool {
	return m.Has(tag.TotalNetValue)
}

//HasCashOutstanding returns true if CashOutstanding is present, Tag 901
func (m CollateralInquiry) HasCashOutstanding() bool {
	return m.Has(tag.CashOutstanding)
}

//HasCollInquiryID returns true if CollInquiryID is present, Tag 909
func (m CollateralInquiry) HasCollInquiryID() bool {
	return m.Has(tag.CollInquiryID)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m CollateralInquiry) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m CollateralInquiry) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m CollateralInquiry) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m CollateralInquiry) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m CollateralInquiry) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m CollateralInquiry) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m CollateralInquiry) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasEndAccruedInterestAmt returns true if EndAccruedInterestAmt is present, Tag 920
func (m CollateralInquiry) HasEndAccruedInterestAmt() bool {
	return m.Has(tag.EndAccruedInterestAmt)
}

//HasStartCash returns true if StartCash is present, Tag 921
func (m CollateralInquiry) HasStartCash() bool {
	return m.Has(tag.StartCash)
}

//HasEndCash returns true if EndCash is present, Tag 922
func (m CollateralInquiry) HasEndCash() bool {
	return m.Has(tag.EndCash)
}

//HasNoCollInquiryQualifier returns true if NoCollInquiryQualifier is present, Tag 938
func (m CollateralInquiry) HasNoCollInquiryQualifier() bool {
	return m.Has(tag.NoCollInquiryQualifier)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m CollateralInquiry) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m CollateralInquiry) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m CollateralInquiry) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m CollateralInquiry) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m CollateralInquiry) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m CollateralInquiry) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m CollateralInquiry) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m CollateralInquiry) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m CollateralInquiry) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m CollateralInquiry) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m CollateralInquiry) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m CollateralInquiry) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m CollateralInquiry) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m CollateralInquiry) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m CollateralInquiry) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m CollateralInquiry) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m CollateralInquiry) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m CollateralInquiry) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m CollateralInquiry) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m CollateralInquiry) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m CollateralInquiry) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m CollateralInquiry) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m CollateralInquiry) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195
func (m CollateralInquiry) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m CollateralInquiry) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasValuationMethod returns true if ValuationMethod is present, Tag 1197
func (m CollateralInquiry) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m CollateralInquiry) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m CollateralInquiry) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m CollateralInquiry) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m CollateralInquiry) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m CollateralInquiry) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m CollateralInquiry) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435
func (m CollateralInquiry) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

//HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439
func (m CollateralInquiry) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

//HasRestructuringType returns true if RestructuringType is present, Tag 1449
func (m CollateralInquiry) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

//HasSeniority returns true if Seniority is present, Tag 1450
func (m CollateralInquiry) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

//HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451
func (m CollateralInquiry) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

//HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452
func (m CollateralInquiry) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

//HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457
func (m CollateralInquiry) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

//HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458
func (m CollateralInquiry) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

//HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478
func (m CollateralInquiry) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

//HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479
func (m CollateralInquiry) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

//HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480
func (m CollateralInquiry) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

//HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481
func (m CollateralInquiry) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

//HasOptPayoutType returns true if OptPayoutType is present, Tag 1482
func (m CollateralInquiry) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

//HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483
func (m CollateralInquiry) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
}

//NoDlvyInst is a repeating group element, Tag 85
type NoDlvyInst struct {
	quickfix.Group
}

//SetSettlInstSource sets SettlInstSource, Tag 165
func (m NoDlvyInst) SetSettlInstSource(v string) {
	m.Set(field.NewSettlInstSource(v))
}

//SetDlvyInstType sets DlvyInstType, Tag 787
func (m NoDlvyInst) SetDlvyInstType(v string) {
	m.Set(field.NewDlvyInstType(v))
}

//SetNoSettlPartyIDs sets NoSettlPartyIDs, Tag 781
func (m NoDlvyInst) SetNoSettlPartyIDs(f NoSettlPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlInstSource gets SettlInstSource, Tag 165
func (m NoDlvyInst) GetSettlInstSource() (f field.SettlInstSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDlvyInstType gets DlvyInstType, Tag 787
func (m NoDlvyInst) GetDlvyInstType() (f field.DlvyInstTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartyIDs gets NoSettlPartyIDs, Tag 781
func (m NoDlvyInst) GetNoSettlPartyIDs() (NoSettlPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlInstSource returns true if SettlInstSource is present, Tag 165
func (m NoDlvyInst) HasSettlInstSource() bool {
	return m.Has(tag.SettlInstSource)
}

//HasDlvyInstType returns true if DlvyInstType is present, Tag 787
func (m NoDlvyInst) HasDlvyInstType() bool {
	return m.Has(tag.DlvyInstType)
}

//HasNoSettlPartyIDs returns true if NoSettlPartyIDs is present, Tag 781
func (m NoDlvyInst) HasNoSettlPartyIDs() bool {
	return m.Has(tag.NoSettlPartyIDs)
}

//NoSettlPartyIDs is a repeating group element, Tag 781
type NoSettlPartyIDs struct {
	quickfix.Group
}

//SetSettlPartyID sets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) SetSettlPartyID(v string) {
	m.Set(field.NewSettlPartyID(v))
}

//SetSettlPartyIDSource sets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) SetSettlPartyIDSource(v string) {
	m.Set(field.NewSettlPartyIDSource(v))
}

//SetSettlPartyRole sets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) SetSettlPartyRole(v int) {
	m.Set(field.NewSettlPartyRole(v))
}

//SetNoSettlPartySubIDs sets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) SetNoSettlPartySubIDs(f NoSettlPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlPartyID gets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) GetSettlPartyID() (f field.SettlPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyIDSource gets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) GetSettlPartyIDSource() (f field.SettlPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyRole gets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) GetSettlPartyRole() (f field.SettlPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartySubIDs gets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) GetNoSettlPartySubIDs() (NoSettlPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlPartyID returns true if SettlPartyID is present, Tag 782
func (m NoSettlPartyIDs) HasSettlPartyID() bool {
	return m.Has(tag.SettlPartyID)
}

//HasSettlPartyIDSource returns true if SettlPartyIDSource is present, Tag 783
func (m NoSettlPartyIDs) HasSettlPartyIDSource() bool {
	return m.Has(tag.SettlPartyIDSource)
}

//HasSettlPartyRole returns true if SettlPartyRole is present, Tag 784
func (m NoSettlPartyIDs) HasSettlPartyRole() bool {
	return m.Has(tag.SettlPartyRole)
}

//HasNoSettlPartySubIDs returns true if NoSettlPartySubIDs is present, Tag 801
func (m NoSettlPartyIDs) HasNoSettlPartySubIDs() bool {
	return m.Has(tag.NoSettlPartySubIDs)
}

//NoSettlPartySubIDs is a repeating group element, Tag 801
type NoSettlPartySubIDs struct {
	quickfix.Group
}

//SetSettlPartySubID sets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) SetSettlPartySubID(v string) {
	m.Set(field.NewSettlPartySubID(v))
}

//SetSettlPartySubIDType sets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) SetSettlPartySubIDType(v int) {
	m.Set(field.NewSettlPartySubIDType(v))
}

//GetSettlPartySubID gets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) GetSettlPartySubID() (f field.SettlPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartySubIDType gets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) GetSettlPartySubIDType() (f field.SettlPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSettlPartySubID returns true if SettlPartySubID is present, Tag 785
func (m NoSettlPartySubIDs) HasSettlPartySubID() bool {
	return m.Has(tag.SettlPartySubID)
}

//HasSettlPartySubIDType returns true if SettlPartySubIDType is present, Tag 786
func (m NoSettlPartySubIDs) HasSettlPartySubIDType() bool {
	return m.Has(tag.SettlPartySubIDType)
}

//NoSettlPartySubIDsRepeatingGroup is a repeating group, Tag 801
type NoSettlPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartySubIDsRepeatingGroup returns an initialized, NoSettlPartySubIDsRepeatingGroup
func NewNoSettlPartySubIDsRepeatingGroup() NoSettlPartySubIDsRepeatingGroup {
	return NoSettlPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartySubID), quickfix.GroupElement(tag.SettlPartySubIDType)})}
}

//Add create and append a new NoSettlPartySubIDs to this group
func (m NoSettlPartySubIDsRepeatingGroup) Add() NoSettlPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartySubIDs{g}
}

//Get returns the ith NoSettlPartySubIDs in the NoSettlPartySubIDsRepeatinGroup
func (m NoSettlPartySubIDsRepeatingGroup) Get(i int) NoSettlPartySubIDs {
	return NoSettlPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlPartyIDsRepeatingGroup is a repeating group, Tag 781
type NoSettlPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartyIDsRepeatingGroup returns an initialized, NoSettlPartyIDsRepeatingGroup
func NewNoSettlPartyIDsRepeatingGroup() NoSettlPartyIDsRepeatingGroup {
	return NoSettlPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartyID), quickfix.GroupElement(tag.SettlPartyIDSource), quickfix.GroupElement(tag.SettlPartyRole), NewNoSettlPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoSettlPartyIDs to this group
func (m NoSettlPartyIDsRepeatingGroup) Add() NoSettlPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartyIDs{g}
}

//Get returns the ith NoSettlPartyIDs in the NoSettlPartyIDsRepeatinGroup
func (m NoSettlPartyIDsRepeatingGroup) Get(i int) NoSettlPartyIDs {
	return NoSettlPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoDlvyInstRepeatingGroup is a repeating group, Tag 85
type NoDlvyInstRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDlvyInstRepeatingGroup returns an initialized, NoDlvyInstRepeatingGroup
func NewNoDlvyInstRepeatingGroup() NoDlvyInstRepeatingGroup {
	return NoDlvyInstRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDlvyInst,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlInstSource), quickfix.GroupElement(tag.DlvyInstType), NewNoSettlPartyIDsRepeatingGroup()})}
}

//Add create and append a new NoDlvyInst to this group
func (m NoDlvyInstRepeatingGroup) Add() NoDlvyInst {
	g := m.RepeatingGroup.Add()
	return NoDlvyInst{g}
}

//Get returns the ith NoDlvyInst in the NoDlvyInstRepeatinGroup
func (m NoDlvyInstRepeatingGroup) Get(i int) NoDlvyInst {
	return NoDlvyInst{m.RepeatingGroup.Get(i)}
}

//NoExecs is a repeating group element, Tag 124
type NoExecs struct {
	quickfix.Group
}

//SetExecID sets ExecID, Tag 17
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//GetExecID gets ExecID, Tag 17
func (m NoExecs) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasExecID returns true if ExecID is present, Tag 17
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//NoExecsRepeatingGroup is a repeating group, Tag 124
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecID)})}
}

//Add create and append a new NoExecs to this group
func (m NoExecsRepeatingGroup) Add() NoExecs {
	g := m.RepeatingGroup.Add()
	return NoExecs{g}
}

//Get returns the ith NoExecs in the NoExecsRepeatinGroup
func (m NoExecsRepeatingGroup) Get(i int) NoExecs {
	return NoExecs{m.RepeatingGroup.Get(i)}
}

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v string) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (f field.StipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (f field.StipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasStipulationType returns true if StipulationType is present, Tag 233
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

//HasStipulationValue returns true if StipulationValue is present, Tag 234
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

//NoStipulationsRepeatingGroup is a repeating group, Tag 232
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

//Add create and append a new NoStipulations to this group
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

//Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v int) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (f field.PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	quickfix.Group
}

//SetLegSymbol sets LegSymbol, Tag 600
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

//SetLegSymbolSfx sets LegSymbolSfx, Tag 601
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

//SetLegSecurityID sets LegSecurityID, Tag 602
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

//SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

//SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegProduct sets LegProduct, Tag 607
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

//SetLegCFICode sets LegCFICode, Tag 608
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

//SetLegSecurityType sets LegSecurityType, Tag 609
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

//SetLegSecuritySubType sets LegSecuritySubType, Tag 764
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

//SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

//SetLegMaturityDate sets LegMaturityDate, Tag 611
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

//SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

//SetLegIssueDate sets LegIssueDate, Tag 249
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

//SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

//SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

//SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

//SetLegCreditRating sets LegCreditRating, Tag 257
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

//SetLegInstrRegistry sets LegInstrRegistry, Tag 599
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

//SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

//SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

//SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

//SetLegRedemptionDate sets LegRedemptionDate, Tag 254
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

//SetLegStrikePrice sets LegStrikePrice, Tag 612
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

//SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

//SetLegOptAttribute sets LegOptAttribute, Tag 613
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

//SetLegContractMultiplier sets LegContractMultiplier, Tag 614
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

//SetLegSecurityExchange sets LegSecurityExchange, Tag 616
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

//SetLegIssuer sets LegIssuer, Tag 617
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

//SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

//SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

//SetLegSecurityDesc sets LegSecurityDesc, Tag 620
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

//SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

//SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

//SetLegRatioQty sets LegRatioQty, Tag 623
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

//SetLegSide sets LegSide, Tag 624
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

//SetLegCurrency sets LegCurrency, Tag 556
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

//SetLegPool sets LegPool, Tag 740
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

//SetLegDatedDate sets LegDatedDate, Tag 739
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

//SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

//SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
}

//SetLegMaturityTime sets LegMaturityTime, Tag 1212
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

//SetLegPutOrCall sets LegPutOrCall, Tag 1358
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

//SetLegExerciseStyle sets LegExerciseStyle, Tag 1420
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

//SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

//SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

//SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

//SetLegContractMultiplierUnit sets LegContractMultiplierUnit, Tag 1436
func (m NoLegs) SetLegContractMultiplierUnit(v int) {
	m.Set(field.NewLegContractMultiplierUnit(v))
}

//SetLegFlowScheduleType sets LegFlowScheduleType, Tag 1440
func (m NoLegs) SetLegFlowScheduleType(v int) {
	m.Set(field.NewLegFlowScheduleType(v))
}

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (f field.LegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (f field.LegSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (f field.LegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (f field.LegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (f field.LegProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (f field.LegCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (f field.LegSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (f field.LegSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (f field.LegMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (f field.LegMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (f field.LegCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (f field.LegIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (f field.LegRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (f field.LegRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (f field.LegRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (f field.LegFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (f field.LegCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (f field.LegInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (f field.LegCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (f field.LegStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (f field.LegLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (f field.LegRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (f field.LegStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (f field.LegStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (f field.LegOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (f field.LegContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (f field.LegCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (f field.LegSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (f field.LegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (f field.EncodedLegIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (f field.EncodedLegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (f field.LegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (f field.EncodedLegSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (f field.EncodedLegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (f field.LegRatioQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (f field.LegSideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (f field.LegCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (f field.LegPoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (f field.LegDatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (f field.LegContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (f field.LegInterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (f field.LegUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (f field.LegTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (f field.LegOptionRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (f field.LegPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityTime gets LegMaturityTime, Tag 1212
func (m NoLegs) GetLegMaturityTime() (f field.LegMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPutOrCall gets LegPutOrCall, Tag 1358
func (m NoLegs) GetLegPutOrCall() (f field.LegPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegExerciseStyle gets LegExerciseStyle, Tag 1420
func (m NoLegs) GetLegExerciseStyle() (f field.LegExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) GetLegUnitOfMeasureQty() (f field.LegUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) GetLegPriceUnitOfMeasure() (f field.LegPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (f field.LegPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractMultiplierUnit gets LegContractMultiplierUnit, Tag 1436
func (m NoLegs) GetLegContractMultiplierUnit() (f field.LegContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFlowScheduleType gets LegFlowScheduleType, Tag 1440
func (m NoLegs) GetLegFlowScheduleType() (f field.LegFlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSymbol returns true if LegSymbol is present, Tag 600
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

//HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

//HasLegSecurityID returns true if LegSecurityID is present, Tag 602
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

//HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

//HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

//HasLegProduct returns true if LegProduct is present, Tag 607
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

//HasLegCFICode returns true if LegCFICode is present, Tag 608
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

//HasLegSecurityType returns true if LegSecurityType is present, Tag 609
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

//HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

//HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

//HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

//HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

//HasLegIssueDate returns true if LegIssueDate is present, Tag 249
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

//HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

//HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

//HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

//HasLegFactor returns true if LegFactor is present, Tag 253
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

//HasLegCreditRating returns true if LegCreditRating is present, Tag 257
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

//HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

//HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

//HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

//HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

//HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

//HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

//HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

//HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

//HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

//HasLegCouponRate returns true if LegCouponRate is present, Tag 615
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

//HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

//HasLegIssuer returns true if LegIssuer is present, Tag 617
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

//HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

//HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

//HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

//HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

//HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

//HasLegRatioQty returns true if LegRatioQty is present, Tag 623
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

//HasLegSide returns true if LegSide is present, Tag 624
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

//HasLegCurrency returns true if LegCurrency is present, Tag 556
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

//HasLegPool returns true if LegPool is present, Tag 740
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

//HasLegDatedDate returns true if LegDatedDate is present, Tag 739
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

//HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

//HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

//HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

//HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

//HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

//HasLegPrice returns true if LegPrice is present, Tag 566
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

//HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

//HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

//HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

//HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

//HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

//HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

//HasLegContractMultiplierUnit returns true if LegContractMultiplierUnit is present, Tag 1436
func (m NoLegs) HasLegContractMultiplierUnit() bool {
	return m.Has(tag.LegContractMultiplierUnit)
}

//HasLegFlowScheduleType returns true if LegFlowScheduleType is present, Tag 1440
func (m NoLegs) HasLegFlowScheduleType() bool {
	return m.Has(tag.LegFlowScheduleType)
}

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	quickfix.Group
}

//SetLegSecurityAltID sets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

//SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

//GetLegSecurityAltID gets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) GetLegSecurityAltID() (f field.LegSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (f field.LegSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

//HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

//NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

//Add create and append a new NoLegSecurityAltID to this group
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

//Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegContractMultiplierUnit), quickfix.GroupElement(tag.LegFlowScheduleType)})}
}

//Add create and append a new NoLegs to this group
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

//Get returns the ith NoLegs in the NoLegsRepeatinGroup
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

//NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v string) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v string) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

//SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437
func (m NoUnderlyings) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

//SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441
func (m NoUnderlyings) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

//SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453
func (m NoUnderlyings) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

//SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454
func (m NoUnderlyings) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

//SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

//SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

//SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (f field.UnderlyingSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (f field.UnderlyingProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (f field.UnderlyingCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (f field.UnderlyingSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (f field.UnderlyingMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (f field.UnderlyingCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (f field.UnderlyingIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (f field.UnderlyingRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (f field.UnderlyingRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (f field.UnderlyingRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (f field.UnderlyingFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (f field.UnderlyingCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (f field.UnderlyingInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (f field.UnderlyingCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (f field.UnderlyingStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (f field.UnderlyingLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (f field.UnderlyingRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (f field.UnderlyingStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (f field.UnderlyingCPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (f field.UnderlyingCPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (f field.UnderlyingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (f field.UnderlyingPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (f field.UnderlyingDirtyPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (f field.UnderlyingEndPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (f field.UnderlyingStartValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (f field.UnderlyingCurrentValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (f field.UnderlyingEndValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (f field.UnderlyingAllocationPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (f field.UnderlyingSettlementTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (f field.UnderlyingCashAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (f field.UnderlyingCashTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (f field.UnderlyingUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (f field.UnderlyingTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (f field.UnderlyingCapValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (f field.UnderlyingSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (f field.UnderlyingAdjustedQuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (f field.UnderlyingFXRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (f field.UnderlyingFXRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) GetUnderlyingMaturityTime() (f field.UnderlyingMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) GetUnderlyingPutOrCall() (f field.UnderlyingPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (f field.UnderlyingExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (f field.UnderlyingUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (f field.UnderlyingPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (f field.UnderlyingPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437
func (m NoUnderlyings) GetUnderlyingContractMultiplierUnit() (f field.UnderlyingContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441
func (m NoUnderlyings) GetUnderlyingFlowScheduleType() (f field.UnderlyingFlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453
func (m NoUnderlyings) GetUnderlyingRestructuringType() (f field.UnderlyingRestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454
func (m NoUnderlyings) GetUnderlyingSeniority() (f field.UnderlyingSeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455
func (m NoUnderlyings) GetUnderlyingNotionalPercentageOutstanding() (f field.UnderlyingNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456
func (m NoUnderlyings) GetUnderlyingOriginalNotionalPercentageOutstanding() (f field.UnderlyingOriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459
func (m NoUnderlyings) GetUnderlyingAttachmentPoint() (f field.UnderlyingAttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460
func (m NoUnderlyings) GetUnderlyingDetachmentPoint() (f field.UnderlyingDetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

//HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437
func (m NoUnderlyings) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

//HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441
func (m NoUnderlyings) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

//HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453
func (m NoUnderlyings) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

//HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454
func (m NoUnderlyings) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

//HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455
func (m NoUnderlyings) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

//HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456
func (m NoUnderlyings) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

//HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459
func (m NoUnderlyings) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

//HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460
func (m NoUnderlyings) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
}

//NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

//SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

//GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (f field.UnderlyingSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (f field.UnderlyingSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

//HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

//NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

//Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	quickfix.Group
}

//SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

//SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

//GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (f field.UnderlyingStipTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (f field.UnderlyingStipValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

//HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

//NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

//Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

//Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentParties is a repeating group element, Tag 1058
type NoUndlyInstrumentParties struct {
	quickfix.Group
}

//SetUnderlyingInstrumentPartyID sets UnderlyingInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyID(v))
}

//SetUnderlyingInstrumentPartyIDSource sets UnderlyingInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyIDSource(v))
}

//SetUnderlyingInstrumentPartyRole sets UnderlyingInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.Set(field.NewUnderlyingInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUnderlyingInstrumentPartyID gets UnderlyingInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (f field.UnderlyingInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (f field.UnderlyingInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (f field.UnderlyingInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUnderlyingInstrumentPartyID returns true if UnderlyingInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyID() bool {
	return m.Has(tag.UnderlyingInstrumentPartyID)
}

//HasUnderlyingInstrumentPartyIDSource returns true if UnderlyingInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyIDSource() bool {
	return m.Has(tag.UnderlyingInstrumentPartyIDSource)
}

//HasUnderlyingInstrumentPartyRole returns true if UnderlyingInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyRole() bool {
	return m.Has(tag.UnderlyingInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUnderlyingInstrumentPartySubID sets UnderlyingInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartySubID(v))
}

//SetUnderlyingInstrumentPartySubIDType sets UnderlyingInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubIDType(v int) {
	m.Set(field.NewUnderlyingInstrumentPartySubIDType(v))
}

//GetUnderlyingInstrumentPartySubID gets UnderlyingInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (f field.UnderlyingInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (f field.UnderlyingInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingInstrumentPartySubID returns true if UnderlyingInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubID() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubID)
}

//HasUnderlyingInstrumentPartySubIDType returns true if UnderlyingInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubIDType() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartySubID), quickfix.GroupElement(tag.UnderlyingInstrumentPartySubIDType)})}
}

//Add create and append a new NoUndlyInstrumentPartySubIDs to this group
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

//Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoUndlyInstrumentParties to this group
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

//Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint)})}
}

//Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

//Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoTrdRegTimestamps is a repeating group element, Tag 768
type NoTrdRegTimestamps struct {
	quickfix.Group
}

//SetTrdRegTimestamp sets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time) {
	m.Set(field.NewTrdRegTimestamp(v))
}

//SetTrdRegTimestampType sets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) SetTrdRegTimestampType(v int) {
	m.Set(field.NewTrdRegTimestampType(v))
}

//SetTrdRegTimestampOrigin sets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) {
	m.Set(field.NewTrdRegTimestampOrigin(v))
}

//SetDeskType sets DeskType, Tag 1033
func (m NoTrdRegTimestamps) SetDeskType(v string) {
	m.Set(field.NewDeskType(v))
}

//SetDeskTypeSource sets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) SetDeskTypeSource(v int) {
	m.Set(field.NewDeskTypeSource(v))
}

//SetDeskOrderHandlingInst sets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) SetDeskOrderHandlingInst(v string) {
	m.Set(field.NewDeskOrderHandlingInst(v))
}

//GetTrdRegTimestamp gets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) GetTrdRegTimestamp() (f field.TrdRegTimestampField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRegTimestampType gets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) GetTrdRegTimestampType() (f field.TrdRegTimestampTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTrdRegTimestampOrigin gets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) GetTrdRegTimestampOrigin() (f field.TrdRegTimestampOriginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskType gets DeskType, Tag 1033
func (m NoTrdRegTimestamps) GetDeskType() (f field.DeskTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskTypeSource gets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) GetDeskTypeSource() (f field.DeskTypeSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskOrderHandlingInst gets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) GetDeskOrderHandlingInst() (f field.DeskOrderHandlingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTrdRegTimestamp returns true if TrdRegTimestamp is present, Tag 769
func (m NoTrdRegTimestamps) HasTrdRegTimestamp() bool {
	return m.Has(tag.TrdRegTimestamp)
}

//HasTrdRegTimestampType returns true if TrdRegTimestampType is present, Tag 770
func (m NoTrdRegTimestamps) HasTrdRegTimestampType() bool {
	return m.Has(tag.TrdRegTimestampType)
}

//HasTrdRegTimestampOrigin returns true if TrdRegTimestampOrigin is present, Tag 771
func (m NoTrdRegTimestamps) HasTrdRegTimestampOrigin() bool {
	return m.Has(tag.TrdRegTimestampOrigin)
}

//HasDeskType returns true if DeskType is present, Tag 1033
func (m NoTrdRegTimestamps) HasDeskType() bool {
	return m.Has(tag.DeskType)
}

//HasDeskTypeSource returns true if DeskTypeSource is present, Tag 1034
func (m NoTrdRegTimestamps) HasDeskTypeSource() bool {
	return m.Has(tag.DeskTypeSource)
}

//HasDeskOrderHandlingInst returns true if DeskOrderHandlingInst is present, Tag 1035
func (m NoTrdRegTimestamps) HasDeskOrderHandlingInst() bool {
	return m.Has(tag.DeskOrderHandlingInst)
}

//NoTrdRegTimestampsRepeatingGroup is a repeating group, Tag 768
type NoTrdRegTimestampsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTrdRegTimestampsRepeatingGroup returns an initialized, NoTrdRegTimestampsRepeatingGroup
func NewNoTrdRegTimestampsRepeatingGroup() NoTrdRegTimestampsRepeatingGroup {
	return NoTrdRegTimestampsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRegTimestamps,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRegTimestamp), quickfix.GroupElement(tag.TrdRegTimestampType), quickfix.GroupElement(tag.TrdRegTimestampOrigin), quickfix.GroupElement(tag.DeskType), quickfix.GroupElement(tag.DeskTypeSource), quickfix.GroupElement(tag.DeskOrderHandlingInst)})}
}

//Add create and append a new NoTrdRegTimestamps to this group
func (m NoTrdRegTimestampsRepeatingGroup) Add() NoTrdRegTimestamps {
	g := m.RepeatingGroup.Add()
	return NoTrdRegTimestamps{g}
}

//Get returns the ith NoTrdRegTimestamps in the NoTrdRegTimestampsRepeatinGroup
func (m NoTrdRegTimestampsRepeatingGroup) Get(i int) NoTrdRegTimestamps {
	return NoTrdRegTimestamps{m.RepeatingGroup.Get(i)}
}

//NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v int) {
	m.Set(field.NewEventType(v))
}

//SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

//SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

//GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (f field.EventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (f field.EventDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (f field.EventPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (f field.EventTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (f field.EventTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

//HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

//HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

//HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

//HasEventTime returns true if EventTime is present, Tag 1145
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

//Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

//Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

//NoTrades is a repeating group element, Tag 897
type NoTrades struct {
	quickfix.Group
}

//SetTradeReportID sets TradeReportID, Tag 571
func (m NoTrades) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

//SetSecondaryTradeReportID sets SecondaryTradeReportID, Tag 818
func (m NoTrades) SetSecondaryTradeReportID(v string) {
	m.Set(field.NewSecondaryTradeReportID(v))
}

//GetTradeReportID gets TradeReportID, Tag 571
func (m NoTrades) GetTradeReportID() (f field.TradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryTradeReportID gets SecondaryTradeReportID, Tag 818
func (m NoTrades) GetSecondaryTradeReportID() (f field.SecondaryTradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradeReportID returns true if TradeReportID is present, Tag 571
func (m NoTrades) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

//HasSecondaryTradeReportID returns true if SecondaryTradeReportID is present, Tag 818
func (m NoTrades) HasSecondaryTradeReportID() bool {
	return m.Has(tag.SecondaryTradeReportID)
}

//NoTradesRepeatingGroup is a repeating group, Tag 897
type NoTradesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradesRepeatingGroup returns an initialized, NoTradesRepeatingGroup
func NewNoTradesRepeatingGroup() NoTradesRepeatingGroup {
	return NoTradesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrades,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradeReportID), quickfix.GroupElement(tag.SecondaryTradeReportID)})}
}

//Add create and append a new NoTrades to this group
func (m NoTradesRepeatingGroup) Add() NoTrades {
	g := m.RepeatingGroup.Add()
	return NoTrades{g}
}

//Get returns the ith NoTrades in the NoTradesRepeatinGroup
func (m NoTradesRepeatingGroup) Get(i int) NoTrades {
	return NoTrades{m.RepeatingGroup.Get(i)}
}

//NoCollInquiryQualifier is a repeating group element, Tag 938
type NoCollInquiryQualifier struct {
	quickfix.Group
}

//SetCollInquiryQualifier sets CollInquiryQualifier, Tag 896
func (m NoCollInquiryQualifier) SetCollInquiryQualifier(v int) {
	m.Set(field.NewCollInquiryQualifier(v))
}

//GetCollInquiryQualifier gets CollInquiryQualifier, Tag 896
func (m NoCollInquiryQualifier) GetCollInquiryQualifier() (f field.CollInquiryQualifierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCollInquiryQualifier returns true if CollInquiryQualifier is present, Tag 896
func (m NoCollInquiryQualifier) HasCollInquiryQualifier() bool {
	return m.Has(tag.CollInquiryQualifier)
}

//NoCollInquiryQualifierRepeatingGroup is a repeating group, Tag 938
type NoCollInquiryQualifierRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoCollInquiryQualifierRepeatingGroup returns an initialized, NoCollInquiryQualifierRepeatingGroup
func NewNoCollInquiryQualifierRepeatingGroup() NoCollInquiryQualifierRepeatingGroup {
	return NoCollInquiryQualifierRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoCollInquiryQualifier,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.CollInquiryQualifier)})}
}

//Add create and append a new NoCollInquiryQualifier to this group
func (m NoCollInquiryQualifierRepeatingGroup) Add() NoCollInquiryQualifier {
	g := m.RepeatingGroup.Add()
	return NoCollInquiryQualifier{g}
}

//Get returns the ith NoCollInquiryQualifier in the NoCollInquiryQualifierRepeatinGroup
func (m NoCollInquiryQualifierRepeatingGroup) Get(i int) NoCollInquiryQualifier {
	return NoCollInquiryQualifier{m.RepeatingGroup.Get(i)}
}

//NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	quickfix.Group
}

//SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

//SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

//SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

//SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (f field.InstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (f field.InstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (f field.InstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

//HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

//HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

//HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

//NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

//SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

//GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (f field.InstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (f field.InstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

//HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

//NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

//Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

//Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

//Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoComplexEvents is a repeating group element, Tag 1483
type NoComplexEvents struct {
	quickfix.Group
}

//SetComplexEventType sets ComplexEventType, Tag 1484
func (m NoComplexEvents) SetComplexEventType(v int) {
	m.Set(field.NewComplexEventType(v))
}

//SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

//SetComplexEventPrice sets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

//SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v int) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

//SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
}

//SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) SetComplexEventPriceTimeType(v int) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

//SetComplexEventCondition sets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) SetComplexEventCondition(v int) {
	m.Set(field.NewComplexEventCondition(v))
}

//SetNoComplexEventDates sets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventType gets ComplexEventType, Tag 1484
func (m NoComplexEvents) GetComplexEventType() (f field.ComplexEventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) GetComplexOptPayoutAmount() (f field.ComplexOptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPrice gets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) GetComplexEventPrice() (f field.ComplexEventPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (f field.ComplexEventPriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (f field.ComplexEventPriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) GetComplexEventPriceTimeType() (f field.ComplexEventPriceTimeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventCondition gets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) GetComplexEventCondition() (f field.ComplexEventConditionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventDates gets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventType returns true if ComplexEventType is present, Tag 1484
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

//HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

//HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

//HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

//HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

//HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

//HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

//HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

//NoComplexEventDates is a repeating group element, Tag 1491
type NoComplexEventDates struct {
	quickfix.Group
}

//SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

//SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

//SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) GetComplexEventStartDate() (f field.ComplexEventStartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) GetComplexEventEndDate() (f field.ComplexEventEndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

//HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

//HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

//NoComplexEventTimes is a repeating group element, Tag 1494
type NoComplexEventTimes struct {
	quickfix.Group
}

//SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

//SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

//GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) GetComplexEventStartTime() (f field.ComplexEventStartTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) GetComplexEventEndTime() (f field.ComplexEventEndTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

//HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

//NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

//Add create and append a new NoComplexEventTimes to this group
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

//Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

//NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
}

//Add create and append a new NoComplexEventDates to this group
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

//Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

//NoComplexEventsRepeatingGroup is a repeating group, Tag 1483
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
}

//Add create and append a new NoComplexEvents to this group
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

//Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}
