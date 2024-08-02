package wasm

func (g *Global) Logout(path, to string) func() {
	return func() {
		go func() {
			code := DoDelete(path + "/logout")
			if code == 200 {
				g.Location.Set("href", "/"+to)
				return
			}
			g.flashThree("error")
		}()
	}
}
