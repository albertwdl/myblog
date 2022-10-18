package setting

type ServerSettingS struct {
	AppMode  string
	HttpPort string
}

type DatabaseSettingS struct {
	DBType     string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassWord string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return nil
	}
	return nil
}
