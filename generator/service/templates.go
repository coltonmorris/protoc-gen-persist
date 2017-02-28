// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package service

import (
	"text/template"

	"github.com/Sirupsen/logrus"
	"github.com/tcncloud/protoc-gen-persist/persist"
)

var (
	UnaryTemplateString  = map[persist.PersistenceOptions]string{}
	ServerTemplateString = map[persist.PersistenceOptions]string{}
	ClientTmplateString  = map[persist.PersistenceOptions]string{}
	BidirTemplateString  = map[persist.PersistenceOptions]string{}

	UnaryTemplate        map[persist.PersistenceOptions]*template.Template
	ServerStreamTemplate map[persist.PersistenceOptions]*template.Template
	ClientStreamTemplate map[persist.PersistenceOptions]*template.Template
	BidirStreamTemplate  map[persist.PersistenceOptions]*template.Template
)

func init() {
	// initialize mongo templates
	SetupMongoTemplates()
	// initialize sql templates
	SetupSQLTemplates()

	var err error
	UnaryTemplate, err = InitTemplates("unary", UnaryTemplateString)
	if err != nil {
		logrus.WithError(err).Fatal("fatal error initializing unary templates")
	}
	ServerStreamTemplate, err = InitTemplates("server_stream", ServerTemplateString)
	if err != nil {
		logrus.WithError(err).Fatal("fatal error initializing server templates")
	}
	ClientStreamTemplate, err = InitTemplates("client_stream", ClientTmplateString)
	if err != nil {
		logrus.WithError(err).Fatal("fatal error initializing client templates")
	}
	BidirStreamTemplate, err = InitTemplates("bidir_stream", BidirTemplateString)
	if err != err {
		logrus.WithError(err).Fatal("fatal error initializing bidir templates")
	}

}

func InitTemplates(class string, templateStrings map[persist.PersistenceOptions]string) (map[persist.PersistenceOptions]*template.Template, error) {
	ret := make(map[persist.PersistenceOptions]*template.Template)
	for k, v := range templateStrings {
		r, err := template.New(class + "_" + k.String()).Parse(v)
		if err != nil {
			return nil, err
		}
		ret[k] = r
	}

	return ret, nil
}
