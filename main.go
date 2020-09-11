package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/app/index"
	"james.engineering/hello-go-gtk/app/settings"
	"james.engineering/hello-go-gtk/framework"
)

func main() {
	gtk.Init(nil)

	var err error

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Hello GTK")

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(640, 480)

	navigator := framework.NewNavigator(win)

	views := map[string]framework.View{
		"index":    index.NewView(index.NewPresenter(&navigator)),
		"settings": settings.NewView(settings.NewPresenter(&navigator)),
	}

	navigator.SetViews(views)

	navigator.Navigate("index")

	if err != nil {
		log.Fatal(err)
		gtk.MainQuit()
	}

	gtk.Main()
}
