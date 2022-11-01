package setting

type ServerSettingS struct {
	AppMode  string
	HttpPort string
}

type DatabaseSettingS struct {
	DBType    string
	Host      string
	Port      string
	DBName    string
	UserName  string
	Password  string
	Charset   string
	ParseTime bool
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return nil
	}
	return nil
}
