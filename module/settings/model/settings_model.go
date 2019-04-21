package model

// Setting Struct
type Setting struct {
	SettingID string `json:"setting_id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	FieldName string
	Value     string `json:"value"`
}

// Settings list
type Settings []Setting
