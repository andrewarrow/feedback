package router

func (r *Router) StartChannels() {
	go r.StartPathChannel()
	go r.StartAfterChannel()
}

func (r *Router) StartPathChannel() {
	for {
		<-r.PathChan
		copyMap := make(map[string]func(*Context, string, string))
		for k, v := range r.Paths {
			copyMap[k] = v
		}
		r.PathChan <- copyMap
	}
}

func (r *Router) StartAfterChannel() {
	for {
		<-r.AfterChan
		copyMap := make(map[string]func(*Context, string))
		for k, v := range r.AfterCreate {
			copyMap[k] = v
		}
		r.AfterChan <- copyMap
	}
}
