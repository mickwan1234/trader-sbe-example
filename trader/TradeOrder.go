// Generated SBE (Simple Binary Encoding) message codec

package trader

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type TradeOrder struct {
	MessageHeader  MessageHeader
	OrderId        int64
	TradeDate      uint32
	OrderQty       int32
	Price          float32
	OrderType      OrderTypeEnum
	OrderSide      OrderSideEnum
	Symbol         []uint8
	ExecutionVenue []uint8
}

func (t *TradeOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := t.MessageHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, t.TradeDate); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat32(_w, t.Price); err != nil {
		return err
	}
	if err := t.OrderType.Encode(_m, _w); err != nil {
		return err
	}
	if err := t.OrderSide.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(t.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(t.ExecutionVenue))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.ExecutionVenue); err != nil {
		return err
	}
	return nil
}

func (t *TradeOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if t.MessageHeaderInActingVersion(actingVersion) {
		if err := t.MessageHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !t.OrderIdInActingVersion(actingVersion) {
		t.OrderId = t.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.OrderId); err != nil {
			return err
		}
	}
	if !t.TradeDateInActingVersion(actingVersion) {
		t.TradeDate = t.TradeDateNullValue()
	} else {
		if err := _m.ReadUint32(_r, &t.TradeDate); err != nil {
			return err
		}
	}
	if !t.OrderQtyInActingVersion(actingVersion) {
		t.OrderQty = t.OrderQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.OrderQty); err != nil {
			return err
		}
	}
	if !t.PriceInActingVersion(actingVersion) {
		t.Price = t.PriceNullValue()
	} else {
		if err := _m.ReadFloat32(_r, &t.Price); err != nil {
			return err
		}
	}
	if t.OrderTypeInActingVersion(actingVersion) {
		if err := t.OrderType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if t.OrderSideInActingVersion(actingVersion) {
		if err := t.OrderSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint32
		if err := _m.ReadUint32(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(t.Symbol) < int(SymbolLength) {
			t.Symbol = make([]uint8, SymbolLength)
		}
		t.Symbol = t.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, t.Symbol); err != nil {
			return err
		}
	}

	if t.ExecutionVenueInActingVersion(actingVersion) {
		var ExecutionVenueLength uint32
		if err := _m.ReadUint32(_r, &ExecutionVenueLength); err != nil {
			return err
		}
		if cap(t.ExecutionVenue) < int(ExecutionVenueLength) {
			t.ExecutionVenue = make([]uint8, ExecutionVenueLength)
		}
		t.ExecutionVenue = t.ExecutionVenue[:ExecutionVenueLength]
		if err := _m.ReadBytes(_r, t.ExecutionVenue); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradeOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.OrderIdInActingVersion(actingVersion) {
		if t.OrderId < t.OrderIdMinValue() || t.OrderId > t.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on t.OrderId (%v < %v > %v)", t.OrderIdMinValue(), t.OrderId, t.OrderIdMaxValue())
		}
	}
	if t.TradeDateInActingVersion(actingVersion) {
		if t.TradeDate < t.TradeDateMinValue() || t.TradeDate > t.TradeDateMaxValue() {
			return fmt.Errorf("Range check failed on t.TradeDate (%v < %v > %v)", t.TradeDateMinValue(), t.TradeDate, t.TradeDateMaxValue())
		}
	}
	if t.OrderQtyInActingVersion(actingVersion) {
		if t.OrderQty < t.OrderQtyMinValue() || t.OrderQty > t.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.OrderQty (%v < %v > %v)", t.OrderQtyMinValue(), t.OrderQty, t.OrderQtyMaxValue())
		}
	}
	if t.PriceInActingVersion(actingVersion) {
		if t.Price < t.PriceMinValue() || t.Price > t.PriceMaxValue() {
			return fmt.Errorf("Range check failed on t.Price (%v < %v > %v)", t.PriceMinValue(), t.Price, t.PriceMaxValue())
		}
	}
	if err := t.OrderType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := t.OrderSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for idx, ch := range t.Symbol {
		if ch > 127 {
			return fmt.Errorf("t.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	for idx, ch := range t.ExecutionVenue {
		if ch > 127 {
			return fmt.Errorf("t.ExecutionVenue[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func TradeOrderInit(t *TradeOrder) {
	return
}

func (*TradeOrder) SbeBlockLength() (blockLength uint16) {
	return 30
}

func (*TradeOrder) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*TradeOrder) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TradeOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*TradeOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TradeOrder) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*TradeOrder) MessageHeaderId() uint16 {
	return 0
}

func (*TradeOrder) MessageHeaderSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) MessageHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.MessageHeaderSinceVersion()
}

func (*TradeOrder) MessageHeaderDeprecated() uint16 {
	return 0
}

func (*TradeOrder) MessageHeaderMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) OrderIdId() uint16 {
	return 1
}

func (*TradeOrder) OrderIdSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderIdSinceVersion()
}

func (*TradeOrder) OrderIdDeprecated() uint16 {
	return 0
}

func (*TradeOrder) OrderIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradeOrder) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*TradeOrder) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*TradeOrder) TradeDateId() uint16 {
	return 2
}

func (*TradeOrder) TradeDateSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeDateSinceVersion()
}

func (*TradeOrder) TradeDateDeprecated() uint16 {
	return 0
}

func (*TradeOrder) TradeDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) TradeDateMinValue() uint32 {
	return 0
}

func (*TradeOrder) TradeDateMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*TradeOrder) TradeDateNullValue() uint32 {
	return math.MaxUint32
}

func (*TradeOrder) OrderQtyId() uint16 {
	return 3
}

func (*TradeOrder) OrderQtySinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderQtySinceVersion()
}

func (*TradeOrder) OrderQtyDeprecated() uint16 {
	return 0
}

func (*TradeOrder) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) OrderQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeOrder) OrderQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeOrder) OrderQtyNullValue() int32 {
	return math.MinInt32
}

func (*TradeOrder) PriceId() uint16 {
	return 4
}

func (*TradeOrder) PriceSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PriceSinceVersion()
}

func (*TradeOrder) PriceDeprecated() uint16 {
	return 0
}

func (*TradeOrder) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) PriceMinValue() float32 {
	return -math.MaxFloat32
}

func (*TradeOrder) PriceMaxValue() float32 {
	return math.MaxFloat32
}

func (*TradeOrder) PriceNullValue() float32 {
	return float32(math.NaN())
}

func (*TradeOrder) OrderTypeId() uint16 {
	return 5
}

func (*TradeOrder) OrderTypeSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderTypeSinceVersion()
}

func (*TradeOrder) OrderTypeDeprecated() uint16 {
	return 0
}

func (*TradeOrder) OrderTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) OrderSideId() uint16 {
	return 6
}

func (*TradeOrder) OrderSideSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) OrderSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderSideSinceVersion()
}

func (*TradeOrder) OrderSideDeprecated() uint16 {
	return 0
}

func (*TradeOrder) OrderSideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) SymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) SymbolSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.SymbolSinceVersion()
}

func (*TradeOrder) SymbolDeprecated() uint16 {
	return 0
}

func (TradeOrder) SymbolCharacterEncoding() string {
	return "ASCII"
}

func (TradeOrder) SymbolHeaderLength() uint64 {
	return 4
}

func (*TradeOrder) ExecutionVenueMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*TradeOrder) ExecutionVenueSinceVersion() uint16 {
	return 0
}

func (t *TradeOrder) ExecutionVenueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ExecutionVenueSinceVersion()
}

func (*TradeOrder) ExecutionVenueDeprecated() uint16 {
	return 0
}

func (TradeOrder) ExecutionVenueCharacterEncoding() string {
	return "ASCII"
}

func (TradeOrder) ExecutionVenueHeaderLength() uint64 {
	return 4
}
