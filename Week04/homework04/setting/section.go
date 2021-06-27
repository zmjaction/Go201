package setting

// DatabaseSettingS 普通yaml配置文件中的Field
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}




func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *Setting) ReadSectionNacos(v interface{}) error {
	err := s.vp.Unmarshal(v)
	if err != nil {
		return err
	}
	return nil

}
