package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/logrus_conf"
	"os"
	"strings"
)

func main() {

	if err := func() error {
		p := os.Getenv("PATH")
		s := os.Getenv("SELECTOR")
		if len(s) == 0 {
			return errors.Errorf("missing selector")
		}
		if err := logrus_conf.PrepareFromEnv("translate-dropdown"); err != nil {
			return errors.WithStack(err)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return errors.WithStack(err)
		}
		r := bytes.NewReader(b)
		d, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			return errors.WithStack(err)
		}
		d.Find(s).Each(func(i int, selection *goquery.Selection) {
			fmt.Println(strings.TrimSpace(selection.Text()))
		})
		return nil
	}(); err != nil {
		logrus.Errorf("%+v", err)
	}

}