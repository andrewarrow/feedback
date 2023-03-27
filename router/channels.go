package router

func (r *Router) StartChannels() {
	go r.StartPathChannel()
	go r.StartAfterChannel()
}

func (r *Router) StartPathChannel() {
	for {
		key := <-r.PathChan
		r.PathChan <- r.Paths[key.(string)]
	}
}

func (r *Router) StartAfterChannel() {
	for {
		key := <-r.AfterChan
		r.AfterChan <- r.AfterCreate[key.(string)]
	}
}
