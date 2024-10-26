package config

type LarkConfig struct {
	FeishuAppId                string `json:"app_id"                 mapstructure:"app_id"`
	FeishuAppSecret            string `json:"app_secret"             mapstructure:"app_secret"`
	FeishuAppEncryptKey        string `json:"app_encrypt_key"        mapstructure:"app_encrypt_key"`
	FeishuAppVerificationToken string `json:"app_verification_token" mapstructure:"app_verification_token"`
	FeishuBotName              string `json:"bot_name"               mapstructure:"bot_name"`

	WordTemplateAppToken  string `json:"word_template_app_token" mapstructure:"word_template_app_token"`
	WordAppFolderToken    string `json:"word_app_folder_token" mapstructure:"word_app_folder_token"`
	WordTemplateTableName string `json:"word_template_table_name" mapstructure:"word_template_table_name"`

	AnnualStatisticsAppToken string `json:"annual_statistics_app_token" mapstructure:"annual_statistics_app_token"`
	AnnualStatisticsTableId  string `json:"annual_statistics_table_id" mapstructure:"annual_statistics_table_id"`

	ManagerUserId string `json:"manager_user_id" mapstructure:"manager_user_id"`

	TimingSyncDocData bool `json:"timing_sync_doc_data" mapstructure:"timing_sync_doc_data"`

	SqlAgentApiUrl   string `json:"sql_agent_api_url" mapstructure:"sql_agent_api_url"`
	SqlAgentApiToken string `json:"sql_agent_api_token" mapstructure:"sql_agent_api_token"`
}
