// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"fmt"
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var adapter *Adapter
var CasdoorOrganization string
var CasdoorApplication string

type Session struct {
	SessionKey    string  `xorm:"char(64) notnull pk"`
	SessionData   []uint8 `xorm:"blob"`
	SessionExpiry int     `xorm:"notnull"`
}

func InitAdapter() {
	adapter = NewAdapter(beego.AppConfig.String("driverName"), beego.AppConfig.String("dataSourceName"), beego.AppConfig.String("dbName"))
	adapter.createTable()

	CasdoorOrganization = beego.AppConfig.String("casdoorOrganization")
	CasdoorApplication = beego.AppConfig.String("casdoorApplication")
}

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	Engine         *xorm.Engine
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	err := a.Engine.Close()
	if err != nil {
		panic(err)
	}
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(driverName string, dataSourceName string, dbName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	a.dbName = dbName

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) createDatabase() error {
	Engine, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}
	defer Engine.Close()

	_, err = Engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci", a.dbName))
	return err
}

func (a *Adapter) open() {
	if a.driverName != "postgres" {
		if err := a.createDatabase(); err != nil {
			panic(err)
		}
	}

	Engine, err := xorm.NewEngine(a.driverName, a.dataSourceName+a.dbName)
	if err != nil {
		panic(err)
	}

	a.Engine = Engine
}

func (a *Adapter) close() {
	a.Engine.Close()
	a.Engine = nil
}

func (a *Adapter) createTable() {
	err := a.Engine.Sync2(new(Session))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Topic))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Reply))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Poster))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Translator))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Node))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Favorites))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Tab))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Notification))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(BasicInfo))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(Plane))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(ConsumptionRecord))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(BrowseRecord))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(UploadFileRecord))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(SensitiveWord))
	if err != nil {
		panic(err)
	}

	err = a.Engine.Sync2(new(FrontConf))
	if err != nil {
		panic(err)
	}
}
