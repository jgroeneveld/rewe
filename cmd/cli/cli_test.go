package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
