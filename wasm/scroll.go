package wasm

import "fmt"

func (g *Global) IsBottom() bool {
	scrollTop := g.Document.Document.Get("documentElement").Get("scrollTop").Int()
	fmt.Println("scrollTop", scrollTop)

	return true
}

/*

function isBottomOfPage() {
  const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
  const windowHeight = window.innerHeight;
  const documentHeight = Math.max(
    document.body.scrollHeight, document.documentElement.scrollHeight,
    document.body.offsetHeight, document.documentElement.offsetHeight,
    document.body.clientHeight, document.documentElement.clientHeight
  );

  // Check if the user has scrolled to the bottom (with a small buffer)
  return scrollTop + windowHeight >= documentHeight - 100;
}
*/
