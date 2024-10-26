package utils

import (
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

func NewCardTemplateCreateReqParam(templateId string, templateVariable map[string]interface{}) string {
	param := map[string]interface{}{
		"type": "template",
		"data": map[string]interface{}{
			"template_id":       templateId,
			"template_variable": templateVariable,
		},
	}

	return gconv.String(param)
}

func TableTimeField(fieldName string, fields map[string]interface{}) *time.Time {
	if fields == nil {
		return nil
	}

	v := fields[fieldName]
	if v == nil {
		return nil
	}

	if v, ok := v.(float64); ok {
		milli := time.UnixMilli(int64(v))
		return &milli
	}

	return nil
}

func TableIntField(fieldName string, fields map[string]interface{}) int {
	if fields == nil {
		return 0
	}

	v := fields[fieldName]
	if v == nil {
		return 0
	}

	if v, ok := v.(int); ok {
		return v
	}

	return 0
}

func TableBoolField(fieldName string, fields map[string]interface{}) bool {
	if fields == nil {
		return false
	}

	v := fields[fieldName]
	if v == nil {
		return false
	}

	if v, ok := v.(bool); ok {
		return v
	}

	return false
}

func TableStringField(fieldName string, fields map[string]interface{}) string {
	if fields == nil {
		return ""
	}

	v := fields[fieldName]
	if v == nil {
		return ""
	}

	if v, ok := v.(string); ok {
		return v
	}

	return ""
}
func TableFloatField(fieldName string, fields map[string]interface{}) float64 {
	if fields == nil {
		return 0
	}

	v := fields[fieldName]
	if v == nil {
		return 0
	}

	if v, ok := v.(float64); ok {
		return v
	}

	return 0
}

func TableTextField(fieldName string, fields map[string]interface{}) string {

	if fields == nil {
		return ""
	}

	v := fields[fieldName]
	if v == nil {
		return ""
	}

	if field, ok := v.([]interface{}); ok {
		if field == nil || len(field) == 0 {
			return ""
		}
		return field[0].(map[string]interface{})["text"].(string)
	}

	return ""
}
