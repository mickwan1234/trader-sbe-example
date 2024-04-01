// Generated SBE (Simple Binary Encoding) message codec

package trader

import (
	"fmt"
	"io"
	"reflect"
)

type OrderTypeEnum uint8
type OrderTypeValues struct {
	MARKET     OrderTypeEnum
	LIMIT      OrderTypeEnum
	STOP       OrderTypeEnum
	STOP_LIMIT OrderTypeEnum
	NullValue  OrderTypeEnum
}

var OrderType = OrderTypeValues{1, 2, 3, 4, 255}

func (o OrderTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderType, unknown enumeration value %d", o)
}

func (*OrderTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderTypeEnum) MARKETSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) MARKETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MARKETSinceVersion()
}

func (*OrderTypeEnum) MARKETDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) LIMITSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) LIMITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LIMITSinceVersion()
}

func (*OrderTypeEnum) LIMITDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) STOPSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) STOPInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.STOPSinceVersion()
}

func (*OrderTypeEnum) STOPDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) STOP_LIMITSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) STOP_LIMITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.STOP_LIMITSinceVersion()
}

func (*OrderTypeEnum) STOP_LIMITDeprecated() uint16 {
	return 0
}
