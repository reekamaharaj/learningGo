package code_challenge

import (
	"testing"
)

func TestUnset(t *testing.T) {
	db := NewDatabase()

	db.Set("ex", 10)
	assertValue(t, db, "ex", 10)

	db.Unset("ex")
	assertUnset(t, db, "ex")
}

func TestRollback(t *testing.T) {
	db := NewDatabase()

	db.Begin()
	db.Set("a", 10)
	assertValue(t, db, "a", 10)

	db.Begin()
	db.Set("a", 20)
	assertValue(t, db, "a", 20)

	err := db.Rollback()
	assertNoError(t, err)
	assertValue(t, db, "a", 10)

	err = db.Rollback()
	assertNoError(t, err)
	assertUnset(t, db, "a")
}

func TestNestedCommit(t *testing.T) {
	db := NewDatabase()

	db.Begin()
	db.Set("a", 30)

	db.Begin()
	db.Set("a", 40)

	err := db.Commit()
	assertNoError(t, err)

	assertValue(t, db, "a", 40)

	err = db.Rollback()
	assertError(t, err)

	err = db.Commit()
	assertError(t, err)
}

func TestTransactionInterleavedKeys(t *testing.T) {
	db := NewDatabase()

	db.Set("a", 10)
	db.Set("b", 10)
	assertValue(t, db, "a", 10)
	assertValue(t, db, "b", 10)

	db.Begin()
	db.Set("a", 20)
	assertValue(t, db, "a", 20)
	assertValue(t, db, "b", 10)

	db.Begin()
	db.Set("b", 30)
	assertValue(t, db, "a", 20)
	assertValue(t, db, "b", 30)

	db.Rollback()
	assertValue(t, db, "a", 20)
	assertValue(t, db, "b", 10)

	db.Rollback()
	assertValue(t, db, "a", 10)
	assertValue(t, db, "b", 10)
}

func TestTransactionRollbackUnset(t *testing.T) {
	db := NewDatabase()

	db.Set("a", 10)
	assertValue(t, db, "a", 10)

	db.Begin()
	assertValue(t, db, "a", 10)
	db.Set("a", 20)
	assertValue(t, db, "a", 20)

	db.Begin()
	db.Unset("a")
	assertUnset(t, db, "a")

	err := db.Rollback()
	assertNoError(t, err)
	assertValue(t, db, "a", 20)

	err = db.Commit()
	assertNoError(t, err)
	assertValue(t, db, "a", 20)
}

func TestTransactionCommitUnset(t *testing.T) {
	db := NewDatabase()

	db.Set("a", 10)
	assertValue(t, db, "a", 10)

	db.Begin()
	assertValue(t, db, "a", 10)
	db.Unset("a")
	assertUnset(t, db, "a")

	err := db.Rollback()
	assertNoError(t, err)
	assertValue(t, db, "a", 10)

	db.Begin()
	db.Unset("a")
	assertUnset(t, db, "a")

	db.Commit()
	assertUnset(t, db, "a")

	db.Begin()
	assertUnset(t, db, "a")
	db.Set("a", 20)
	assertValue(t, db, "a", 20)

	db.Commit()
	assertValue(t, db, "a", 20)
}

func assertUnset(t *testing.T, db Database, key string) {
	t.Helper()
	_, ok := db.Get(key)
	if ok {
		t.Fatalf("key %q should not exist", key)
	}
}

func assertValue(t *testing.T, db Database, key string, value int) {
	t.Helper()
	v, ok := db.Get(key)
	if !ok || value != v {
		t.Fatalf("db.Get(%q) should return %d, true: got %d, %v", key, value, v, ok)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error, got: nil")
	}
}
