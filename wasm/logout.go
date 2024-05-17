package wasm

func (g *Global) Logout(path string) func() {
	return func() {
		go func() {
			code := DoDelete(path + "/logout")
			if code == 200 {
				g.Location.Set("href", "/")
				return
			}
			g.flashThree("error")
		}()
	}
}
