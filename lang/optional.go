package lang

import "errors"

// Optional 类java Optional
type Optional struct {
	value interface{}
}

// EMPTY 空, 全局唯一
var EMPTY *Optional = &Optional{nil}

// Of 传入的value必须非nil
func Of(value interface{}) (*Optional, error) {
	if value == nil {
		return nil, errors.New("can't create a instance with `value`")
	}

	return &Optional{value}, nil
}

// OfNilable 传入的value可以是nil
func OfNilable(value interface{}) *Optional {
	return &Optional{value}
}

// Get 返回内部的value值, nil报错
func (o *Optional) Get() (interface{}, error) {
	if o.value == nil {
		return nil, errors.New("can't get a nil `value`")
	}

	return o.value, nil
}

// IsPresent 判断值是否不为nil
func (o *Optional) IsPresent() bool {
	return o.value != nil
}

// IfPresent 如果值非nil, 则执行consumer逻辑
func (o *Optional) IfPresent(consumer func(interface{})) {
	if o.value != nil {
		consumer(o.value)
	}
}
