package leveldbctl

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)

	dbdir := path.Join(tmpdir, "init")

	// init db
	err = Init(dbdir)
	if err != nil {
		t.Error(err)
	}

	// when db is initialized, "LOG" is exist under the dbdir
	if _, err := os.Stat(path.Join(dbdir, "LOG")); err != nil {
		t.Error(err)
	}
}

func TestCRUD(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)

	dbdir := path.Join(tmpdir, "crud")

	// init db
	err = Init(dbdir)
	if err != nil {
		t.Error(err)
	}

	// const
	v := "Foo"
	k := "asdf"
	k2 := "asdf2"

	// get from blank db
	value, ok, err := Get(dbdir, k)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Errorf("Key{%s} must not be exist.", k)
	}
	if value != "" {
		t.Error("value should be blank when key is not exist")
	}

	// set a value with key
	err = Put(dbdir, k, v)
	if err != nil {
		t.Error(err)
	}

	// get with key
	value, ok, err = Get(dbdir, k)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("Key{%s} must be exist.", k)
	}
	if value != v {
		t.Errorf("value should be %s but that is %s.", v, value)
	}

	// get with unset key
	value, ok, err = Get(dbdir, k2)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Errorf("Key{%s} must not be exist.", k2)
	}
	if value != "" {
		t.Error("value should be blank when key is not exist")
	}

	// delete
	err = Delete(dbdir, k)
	if err != nil {
		t.Error(err)
	}

	// key is deleted.
	value, ok, err = Get(dbdir, k)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Errorf("Key{%s} must not be exist.", k)
	}
	if value != "" {
		t.Error("value should be blank when key is not exist")
	}
}

func TestWalk(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)

	dbdir := path.Join(tmpdir, "walk")

	// initialize
	err = Init(dbdir)
	if err != nil {
		t.Error(err)
	}

	// const
	keyvalue := map[string]string{
		"k1": "asdf",
		"k2": "fsda",
		"k3": "aaaa",
	}

	// set value
	for k, v := range keyvalue {
		err := Put(dbdir, k, v)
		if err != nil {
			t.Error(err)
		}
	}

	// walk
	actual_keyvalue := map[string]string{}
	err = Walk(dbdir, func(k, v string) {
		actual_keyvalue[k] = v
	})
	if err != nil {
		t.Error(err)
	}

	// Test
	if !reflect.DeepEqual(keyvalue, actual_keyvalue) {
		t.Errorf("k,v found with walk are not equal. expected: %v, actual: %v", keyvalue, actual_keyvalue)
	}
}

func TestCheckingExistenceDB(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)

	dbdir := path.Join(tmpdir, "existence")

	// Uninitialized Get, Delete, Put, Walk
	err = Put(dbdir, "k", "v")
	if err == nil {
		t.Error("Put not check whether db is initialized")
	}

	err = Delete(dbdir, "k")
	if err == nil {
		t.Error("Delete not check whether db is initialized")
	}

	value, ok, err := Get(dbdir, "k")
	if err == nil {
		t.Error("Get not check whether db is initialized")
	}

	if ok {
		t.Error("Get is returing wrong value that 'ok' should be false")
	}

	if value != "" {
		t.Error("Get is returing wrong value that 'value' should be blank string")
	}

	err = Walk(dbdir, func(k, v string) {
		t.Error("handling functions should be called when db is not initialized")
	})
	if err == nil {
		t.Error("Walk not check whether db is initialized")
	}

	// Initialized DB with init
	err = Init(dbdir)
	if err != nil {
		t.Error(err)
	}

	// Check
	err = Init(dbdir)
	if err == nil {
		t.Error("Init not check whether db is initialized")
	}

}
