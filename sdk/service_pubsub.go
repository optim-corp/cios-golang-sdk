package ciossdk

type CiosPubSub struct {
	_instance
	debug          bool
	token          *string
	wsReadTimeout  int64
	wsWriteTimeout int64
}

func (self *CiosPubSub) setDebug(isDebug bool) {
	self.debug = isDebug
}

func (self *CiosPubSub) setToken(token string) {
	if token == "" {
		self.token = nil
	} else {
		self.token = &token
	}
}
