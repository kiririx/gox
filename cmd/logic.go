package cmd

import (
	"github.com/kiririx/krutils/confx"
)

func create() {

}

func add() {

}

func goxSet() {}

func goxRm() {}

func goxLs() error {
	config, err := confx.ResolveProperties("~/.gox/config.properties")
	if err != nil {
		return err
	}

}
