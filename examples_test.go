package dbutil

import (
	"fmt"

	"github.com/recursionpharma/ghost-postgres"
)

func ExampleConnect() {
	gp := ghost_postgres.New()
	defer gp.Terminate()
	if err := gp.Prepare(); err != nil {
		fmt.Println(err)
		return
	}
	db, err := Connect(gp.URL())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db != nil)
	// Output: true
}

func ExampleGetDriver() {
	driver, err := GetDriver("postgres://")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(driver)
	// Output: postgres
}
