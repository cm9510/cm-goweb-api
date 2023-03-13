package config

type Items struct {
	App struct {
		Env   string `yaml:"env"`
		RunAt int32  `yaml:"run_at"`
	}
	ApiStatus struct {
		Ok         int8   `yaml:"ok"`
		Fail       int8   `yaml:"fail"`
		MsgSuccess string `yaml:"msg_success"`
		MsgFail    string `yaml:"msg_fail"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int32  `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Charset  string `yaml:"charset"`
		Timeout  string `yaml:"timeout"`
	}
}
