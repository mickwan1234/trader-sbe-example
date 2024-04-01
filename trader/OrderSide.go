// Generated SBE (Simple Binary Encoding) message codec

package trader

import (
	"fmt"
	"io"
	"reflect"
)

type OrderSideEnum byte
type OrderSideValues struct {
	BUY       OrderSideEnum
	SELL      OrderSideEnum
	NullValue OrderSideEnum
}

var OrderSide = OrderSideValues{66, 83, 0}

func (o OrderSideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderSideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderSideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderSide)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderSide, unknown enumeration value %d", o)
}

func (*OrderSideEnum) EncodedLength() int64 {
	return 1
}

func (*OrderSideEnum) BUYSinceVersion() uint16 {
	return 0
}

func (o *OrderSideEnum) BUYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BUYSinceVersion()
}

func (*OrderSideEnum) BUYDeprecated() uint16 {
	return 0
}

func (*OrderSideEnum) SELLSinceVersion() uint16 {
	return 0
}

func (o *OrderSideEnum) SELLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SELLSinceVersion()
}

func (*OrderSideEnum) SELLDeprecated() uint16 {
	return 0
}
