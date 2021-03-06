package dbutil

import (
	"testing"

	"github.com/recursionpharma/ghost-postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConnect(t *testing.T) {
	Convey("Given a dbURL", t, func() {
		Convey("If we can't get a driver", func() {
			db, err := Connect("foo")
			Convey("No driver should be returned", func() { So(db, ShouldBeNil) })
			Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
			Convey("The error message should contain...", func() { So(err.Error(), ShouldContainSubstring, "is missing '://'") })
		})
		Convey("If we can't open a connection to the DB", func() {
			db, err := Connect("foo://bar")
			Convey("No driver should be returned", func() { So(db, ShouldBeNil) })
			Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
			Convey("The error message should contain...", func() { So(err.Error(), ShouldContainSubstring, "unknown driver") })
		})
		Convey("If we can't connectto the DB", func() {
			db, err := Connect("postgres://bar")
			Convey("No driver should be returned", func() { So(db, ShouldBeNil) })
			Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
			Convey("The error message should contain...", func() { So(err.Error(), ShouldContainSubstring, "no such host") })
		})
		Convey("If we can connect to the DB", func() {
			gp := ghost_postgres.New()
			defer gp.Terminate()
			if err := gp.Prepare(); err != nil {
				t.Fatal(err)
			}
			db, err := Connect(gp.URL())
			Convey("A driver should be returned", func() { So(db, ShouldNotBeNil) })
			Convey("No error should be returned", func() { So(err, ShouldBeNil) })
		})
	})
}

func TestGetDriver(t *testing.T) {
	Convey("Given a dbURL", t, func() {
		Convey("If it doesn't contain '://'", func() {
			driver, err := GetDriver("foo")
			Convey("No driver should be returned", func() { So(driver, ShouldBeEmpty) })
			Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
			Convey("The error message should contain...", func() { So(err.Error(), ShouldContainSubstring, "is missing '://'") })
		})
		Convey("If the driver is empty", func() {
			driver, err := GetDriver("://")
			Convey("No driver should be returned", func() { So(driver, ShouldBeEmpty) })
			Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
			Convey("The error message should contain...", func() { So(err.Error(), ShouldContainSubstring, "is missing a driver") })
		})
		Convey("If the driver is valid", func() {
			driver, err := GetDriver("foo://")
			Convey("The driver should be returned", func() { So(driver, ShouldEqual, "foo") })
			Convey("No error should be returned", func() { So(err, ShouldBeNil) })
		})
	})
}

func TestExists(t *testing.T) {
	gp := ghost_postgres.New()
	defer gp.Terminate()
	if err := gp.Prepare(); err != nil {
		t.Fatal(err)
	}
	db, err := Connect(gp.URL())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`
		CREATE TABLE foo (
			bar VARCHAR(10)
		);
	`); err != nil {
		t.Fatal(err)
	}

	Convey("If the query returns rows", t, func() {
		if _, err := db.Exec("INSERT INTO foo ( bar ) VALUES ( 'baz' );"); err != nil {
			t.Fatal(err)
		}
		b, err := Exists(db, "SELECT 1 FROM foo WHERE bar = 'baz'")
		Convey("No error should be returned", func() { So(err, ShouldBeNil) })
		Convey("It returns true", func() { So(b, ShouldBeTrue) })
	})

	Convey("If the query doesn't return rows", t, func() {
		b, err := Exists(db, "SELECT 1 FROM foo WHERE bar = 'quux'")
		Convey("No error should be returned", func() { So(err, ShouldBeNil) })
		Convey("It returns false", func() { So(b, ShouldBeFalse) })
	})

	Convey("If the query errors", t, func() {
		b, err := Exists(db, "kaboom!")
		Convey("An error should be returned", func() { So(err, ShouldNotBeNil) })
		Convey("Result should be false", func() { So(b, ShouldBeFalse) })
	})
}
