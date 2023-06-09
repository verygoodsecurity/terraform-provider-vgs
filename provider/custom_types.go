package provider

// import (
// 	"context"
// 	"strings"

// 	"github.com/hashicorp/terraform-plugin-framework/attr"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// 	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
// 	"github.com/hashicorp/terraform-plugin-go/tftypes"
// )

// type ConditionType struct {
// 	types.ObjectType
// }

// // type ExpressionType struct {
// // 	types.ObjectType
// // }

// // func (c ExpressionType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
// // 	if in.IsNull() {
// // 		return ExpressionValue{}, nil
// // 	}
// // 	v, err := types.ObjectType{
// // 		AttrTypes: map[string]attr.Type{
// // 			"field":    types.StringType,
// // 			"operator": types.StringType,
// // 			"type":     types.StringType,
// // 			"values": types.ListType{
// // 				ElemType: types.StringType,
// // 			},
// // 		},
// // 	}.ValueFromTerraform(ctx, in)

// // 	attrs := v.(types.Object).Attributes()
// // 	return ExpressionValue{
// // 		Field:     attrs["field"].(basetypes.StringValue),
// // 		Operator:  attrs["operator"].(basetypes.StringValue),
// // 		ValueType: attrs["type"].(basetypes.StringValue),
// // 		Values:    attrs["values"].(basetypes.ListValue),
// // 	}, err
// // }

// var mvel = `(PathInfo ~= "/post" && (ContentType == "application/xml" || ContentType == "application/xml"))`

// type Condition struct {
// 	condition  string
// 	rules      []Condition
// 	expression struct {
// 		field    string
// 		operator string
// 		typ      string
// 		fields   []string
// 	}
// }

// func Parse(expr string) *Condition {

// 	extract := func(expr string) (before, inside, after string) {
// 		left := strings.Index(expr, "(")
// 		right := strings.LastIndex(expr, ")")
// 		return expr[0:left], expr[left:right], expr[right:len(expr)]
// 	}

// 	l, m, r := extract(expr)

// 	if l == "" && r == "" {
// 		if m == "" {
// 			return nil
// 		}
// 		return Parse(m)
// 	}


// 	condition := Condition {

// 	}
// 	condition := Parse(m)

// 	if l == "" {

// 	}

// 	if r == "" {

// 	}


// }

// func (c ConditionType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
// 	print("IN")
// 	if in.IsNull() {
// 		return ConditionValue{}, nil
// 	}
// 	v, err := types.ObjectType{
// 		AttrTypes: map[string]attr.Type{
// 			"condition": types.StringType,
// 			"expression": types.ObjectType{
// 				AttrTypes: map[string]attr.Type{
// 					"field":    types.StringType,
// 					"operator": types.StringType,
// 					"type":     types.StringType,
// 					"values": types.ListType{
// 						ElemType: types.StringType,
// 					},
// 				},
// 			},
// 			"rules": types.ListType{
// 				ElemType: types.ObjectType{},
// 			},
// 		},
// 	}.ValueFromTerraform(ctx, in)

// 	if err != nil {
// 		panic(err)
// 	}
// 	attrs := v.(types.Object).Attributes()

// 	return ConditionValue{
// 		Condition:  attrs["condition"].(basetypes.StringValue),
// 		Expression: attrs["expression"].(basetypes.ObjectValue),
// 		Rules:      attrs["rules"].(basetypes.ListValue),
// 	}, err
// }

// type ConditionValue struct {
// 	Condition  basetypes.StringValue
// 	Expression basetypes.ObjectValue
// 	Rules      basetypes.ListValue
// }

// // Equal implements attr.Value
// func (ConditionValue) Equal(attr.Value) bool {
// 	panic("unimplemented")
// }

// // IsNull implements attr.Value
// func (ConditionValue) IsNull() bool {
// 	panic("unimplemented")
// }

// // IsUnknown implements attr.Value
// func (ConditionValue) IsUnknown() bool {
// 	panic("unimplemented")
// }

// // String implements attr.Value
// func (ConditionValue) String() string {
// 	return "todo"
// }

// // ToTerraformValue implements attr.Value
// func (ConditionValue) ToTerraformValue(context.Context) (tftypes.Value, error) {
// 	panic("unimplemented")
// }

// // Type implements attr.Value
// func (ConditionValue) Type(context.Context) attr.Type {
// 	return ConditionType{}
// }

// // type ExpressionValue struct {
// // 	Field     basetypes.StringValue
// // 	Operator  basetypes.StringValue
// // 	ValueType basetypes.StringValue
// // 	Values    types.List
// // }

// // // Equal implements attr.Value
// // func (ExpressionValue) Equal(attr.Value) bool {
// // 	panic("unimplemented")
// // }

// // // IsNull implements attr.Value
// // func (ExpressionValue) IsNull() bool {
// // 	panic("unimplemented")
// // }

// // // IsUnknown implements attr.Value
// // func (ExpressionValue) IsUnknown() bool {
// // 	panic("unimplemented")
// // }

// // // String implements attr.Value
// // func (ExpressionValue) String() string {
// // 	return "todo"
// // }

// // // ToTerraformValue implements attr.Value
// // func (ExpressionValue) ToTerraformValue(context.Context) (tftypes.Value, error) {
// // 	panic("unimplemented")
// // }

// // // Type implements attr.Value
// // func (ExpressionValue) Type(context.Context) attr.Type {
// // 	return ExpressionType{}
// // }
