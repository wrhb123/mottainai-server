/*

Copyright (C) 2017-2018  Ettore Di Giacinto <mudler@gentoo.org>
Credits goes also to Gogs authors, some code portions and re-implemented design
are also coming from the Gogs project, which is using the go-macaron framework
and was really source of ispiration. Kudos to them!

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package tiedot

import (
	"os"
	"testing"

	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"
	webhook "github.com/MottainaiCI/mottainai-server/pkg/webhook"
)

var dbtest5 *Database

func TestInsertWebHook(t *testing.T) {

	setting.Configuration.DBPath = "./DB"
	db := New(setting.Configuration.DBPath)
	db.Init()
	dbtest5 = db
	u := &webhook.WebHook{}
	u.Key = "test"
	u.URL = "test2"

	id, err := db.InsertWebHook(u)

	if err != nil {
		t.Fatal("Failed insert")
	}

	uu, _ := db.GetWebHook(id)

	if uu.Key != u.Key {
		t.Fatal("Failed insert")
	}
	if uu.URL != u.URL {
		t.Fatal("Failed insert")
	}
	db.DeleteWebHook(id)

	err = db.DeleteWebHook(id)

	if err == nil {
		t.Fatal("Failed Remove")
	}

}

func TestGetWebHookByKey(t *testing.T) {

	db := dbtest5

	u := &webhook.WebHook{}
	u.Key = "test2"
	u.URL = "test2url"
	id, err := db.InsertWebHook(u)

	if err != nil {
		t.Fatal("Failed insert", err)
	}

	uu, _ := db.GetWebHook(id)

	if uu.Key != u.Key {
		t.Fatal("Failed insert (Key differs)")
	}

	uuu, err := db.GetWebHookByURL("test2url")

	if err != nil {
		t.Fatal(err)
	}
	if uuu.Key != "test2" {
		t.Fatal("Could not find the inserted webhook")
	}

}

func TestGetWebHookByUid(t *testing.T) {
	defer os.RemoveAll(setting.Configuration.DBPath)

	db := dbtest5

	u := &webhook.WebHook{}
	u.Key = "test2"
	u.OwnerId = "20"
	id, err := db.InsertWebHook(u)

	if err != nil {
		t.Fatal("Failed insert", err)
	}

	uu, _ := db.GetWebHook(id)

	if uu.Key != u.Key {
		t.Fatal("Failed insert (Key differs)")
	}

	uuu, err := db.GetWebHookByKey("test2")

	if err != nil {
		t.Fatal(err)
	}
	if uuu.Key != "test2" {
		t.Fatal("Could not find the inserted webhook")
	}

	uuuu, err := db.GetWebHookByUserID("20")

	if err != nil {
		t.Fatal(err)
	}
	if uuuu.Key != "test2" {
		t.Fatal("Could not find the inserted webhook")
	}

}
