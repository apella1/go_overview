package intermediate

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

// handling temporary resources created by programs e.g files and network connections
// * a simple version of the linux `cat` command
func Defer() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func DeferExample() int {
	a := 10
	defer func (val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func (val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

// * You can supply a function that returns values to defer
// * but there is no way to read those values

func SupplyReturnFunc() {
	// ! defer expressions must be function calls like JavaScript's IIFEs
	defer func () int {
		return 2
	}()
}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(){
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// * use tx to do more db inserts
	return nil
}