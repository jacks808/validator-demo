package tiv

import "reflect"

type FieldLevel interface {

	// Top 如果有的话，返回上层的结构
	Top() reflect.Value

	// Parent 返回当前字段父结构（如果有）
	Parent() reflect.Value

	// Field 返回当前字段
	Field() reflect.Value

	// FieldName 返回字段名称，其中标签名称优先于字段实际名称。
	FieldName() string

	// StructFieldName 返回结构体字段的名称
	StructFieldName() string

	// Param 返回当前验证器tag的参数
	Param() string

	// GetTag 返回当前验证器的tag
	GetTag() string

	// ExtractType 获取字段值的实际基础类型。它将深入指针、customTypes 并返回基础值，它很友好。
	ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool)
}
