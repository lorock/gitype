// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"testing"
	"time"

	"github.com/caixw/typing/path"
	"github.com/issue9/assert"
)

func TestLoadConfig(t *testing.T) {
	a := assert.New(t)
	p := path.New("../testdata/")

	conf, err := Load(p)
	a.NotError(err).NotNil(conf)
	a.Equal(conf.Port, ":8080")
	a.Equal(conf.Webhook.Frequency, time.Minute)
}