package glumi

func Init() (err error) {
	err = DefaultShader.Load()
	if err != nil {
		return err
	}
	return err
}