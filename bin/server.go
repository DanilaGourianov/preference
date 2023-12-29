package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Pictures struct {
	Pairs [45]string
}

type Values struct {
	Value1, Value2, Value3, Value4, Value5,
	Value6, Value7, Value8, Value9, Value10,
	Value11, Value12, Value13, Value14, Value15,
	Value16, Value17, Value18, Value19, Value20,
	Value21, Value22, Value23, Value24, Value25,
	Value26, Value27, Value28, Value29, Value30,
	Value31, Value32, Value33, Value34, Value35,
	Value36, Value37, Value38, Value39, Value40,
	Value41, Value42, Value43, Value44, Value45,
	Value46, Value47, Value48, Value49, Value50,
	Value51, Value52, Value53, Value54, Value55,
	Value56, Value57, Value58, Value59, Value60,
	Value61, Value62, Value63, Value64, Value65,
	Value66, Value67, Value68, Value69, Value70,
	Value71, Value72, Value73, Value74, Value75,
	Value76, Value77, Value78, Value79, Value80,
	Value81, Value82, Value83, Value84, Value85,
	Value86, Value87, Value88, Value89, Value90 int
}

var userIsAuthorized bool

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	return os.Getenv(key)
}

var address = goDotEnvVariable("DB_ADDRESS")

var adminPassword = goDotEnvVariable("ADMIN_PASSWORD")

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.ExecuteTemplate(w, "login", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/hello.html")
	tmpl.ExecuteTemplate(w, "hello", nil)
}

func forms(w http.ResponseWriter, r *http.Request) {
	Files := []string{"templates/pictures/1.html", "templates/pictures/2.html", "templates/pictures/3.html", "templates/pictures/4.html", "templates/pictures/5.html",
		"templates/pictures/6.html", "templates/pictures/7.html", "templates/pictures/8.html", "templates/pictures/9.html", "templates/pictures/10.html", "templates/forms.html"}
	tmpl, _ := template.ParseFiles(Files...)
	tmpl.ExecuteTemplate(w, "forms", nil)
}

func theend(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/theend.html")
	tmpl.ExecuteTemplate(w, "theend", nil)
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !userIsAuthorized {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func formsadmin(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT `Count` FROM `art`")
	if err != nil {
		panic(err.Error())
	}
	var values Values

	i := 1
	for rows.Next() {
		var value int
		err := rows.Scan(&value)
		if err != nil {
			panic(err.Error())
		}

		switch i {
		case 1:
			values.Value1 = value
		case 2:
			values.Value2 = value
		case 3:
			values.Value3 = value
		case 4:
			values.Value4 = value
		case 5:
			values.Value5 = value
		case 6:
			values.Value6 = value
		case 7:
			values.Value7 = value
		case 8:
			values.Value8 = value
		case 9:
			values.Value9 = value
		case 10:
			values.Value10 = value
		case 11:
			values.Value11 = value
		case 12:
			values.Value12 = value
		case 13:
			values.Value13 = value
		case 14:
			values.Value14 = value
		case 15:
			values.Value15 = value
		case 16:
			values.Value16 = value
		case 17:
			values.Value17 = value
		case 18:
			values.Value18 = value
		case 19:
			values.Value19 = value
		case 20:
			values.Value20 = value
		case 21:
			values.Value21 = value
		case 22:
			values.Value22 = value
		case 23:
			values.Value23 = value
		case 24:
			values.Value24 = value
		case 25:
			values.Value25 = value
		case 26:
			values.Value26 = value
		case 27:
			values.Value27 = value
		case 28:
			values.Value28 = value
		case 29:
			values.Value29 = value
		case 30:
			values.Value30 = value
		case 31:
			values.Value31 = value
		case 32:
			values.Value32 = value
		case 33:
			values.Value33 = value
		case 34:
			values.Value34 = value
		case 35:
			values.Value35 = value
		case 36:
			values.Value36 = value
		case 37:
			values.Value37 = value
		case 38:
			values.Value38 = value
		case 39:
			values.Value39 = value
		case 40:
			values.Value40 = value
		case 41:
			values.Value41 = value
		case 42:
			values.Value42 = value
		case 43:
			values.Value43 = value
		case 44:
			values.Value44 = value
		case 45:
			values.Value45 = value
		case 46:
			values.Value46 = value
		case 47:
			values.Value47 = value
		case 48:
			values.Value48 = value
		case 49:
			values.Value49 = value
		case 50:
			values.Value50 = value
		case 51:
			values.Value51 = value
		case 52:
			values.Value52 = value
		case 53:
			values.Value53 = value
		case 54:
			values.Value54 = value
		case 55:
			values.Value55 = value
		case 56:
			values.Value56 = value
		case 57:
			values.Value57 = value
		case 58:
			values.Value58 = value
		case 59:
			values.Value59 = value
		case 60:
			values.Value60 = value
		case 61:
			values.Value61 = value
		case 62:
			values.Value62 = value
		case 63:
			values.Value63 = value
		case 64:
			values.Value64 = value
		case 65:
			values.Value65 = value
		case 66:
			values.Value66 = value
		case 67:
			values.Value67 = value
		case 68:
			values.Value68 = value
		case 69:
			values.Value69 = value
		case 70:
			values.Value70 = value
		case 71:
			values.Value71 = value
		case 72:
			values.Value72 = value
		case 73:
			values.Value73 = value
		case 74:
			values.Value74 = value
		case 75:
			values.Value75 = value
		case 76:
			values.Value76 = value
		case 77:
			values.Value77 = value
		case 78:
			values.Value78 = value
		case 79:
			values.Value79 = value
		case 80:
			values.Value80 = value
		case 81:
			values.Value81 = value
		case 82:
			values.Value82 = value
		case 83:
			values.Value83 = value
		case 84:
			values.Value84 = value
		case 85:
			values.Value85 = value
		case 86:
			values.Value86 = value
		case 87:
			values.Value87 = value
		case 88:
			values.Value88 = value
		case 89:
			values.Value89 = value
		case 90:
			values.Value90 = value

		}

		i++
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	Files := []string{"templates/pictures/1.html", "templates/pictures/2.html", "templates/pictures/3.html", "templates/pictures/4.html", "templates/pictures/5.html",
		"templates/pictures/6.html", "templates/pictures/7.html", "templates/pictures/8.html", "templates/pictures/9.html", "templates/pictures/10.html", "templates/formsadmin.html"}
	tmpl, _ := template.ParseFiles(Files...)
	tmpl.ExecuteTemplate(w, "formsadmin", values)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	pswrd := r.FormValue("password")
	if pswrd != adminPassword {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.ExecuteTemplate(w, "login", map[string]interface{}{
			"errorMsg": "Неправильный пароль",
		})
	} else {
		http.Redirect(w, r, "/formsadmin", http.StatusSeeOther)
		userIsAuthorized = true
	}
}

func (pic *Pictures) save(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 45; i++ {
		pic.Pairs[i] = r.FormValue(fmt.Sprintf("pair%d", i+1))
	}

	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	update, err := db.Query(fmt.Sprintf("UPDATE `art` SET `Count` = `Count` + 1 WHERE `Name` IN ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		pic.Pairs[0], pic.Pairs[1], pic.Pairs[2], pic.Pairs[3], pic.Pairs[4], pic.Pairs[5], pic.Pairs[6], pic.Pairs[7], pic.Pairs[8], pic.Pairs[9],
		pic.Pairs[10], pic.Pairs[11], pic.Pairs[12], pic.Pairs[13], pic.Pairs[14], pic.Pairs[15], pic.Pairs[16], pic.Pairs[17], pic.Pairs[18], pic.Pairs[19],
		pic.Pairs[20], pic.Pairs[21], pic.Pairs[22], pic.Pairs[23], pic.Pairs[24], pic.Pairs[25], pic.Pairs[26], pic.Pairs[27], pic.Pairs[28], pic.Pairs[29],
		pic.Pairs[30], pic.Pairs[31], pic.Pairs[32], pic.Pairs[33], pic.Pairs[34], pic.Pairs[35], pic.Pairs[36], pic.Pairs[37], pic.Pairs[38], pic.Pairs[39],
		pic.Pairs[40], pic.Pairs[41], pic.Pairs[42], pic.Pairs[43], pic.Pairs[44]))
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
	http.Redirect(w, r, "/theend", http.StatusSeeOther)
}

func generateExcel(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID, Count FROM art")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	f := excelize.NewFile()

	// Create a new sheet.
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

		// Write the headers.
	f.SetCellValue(sheetName, "A1", "ID")
	f.SetCellValue(sheetName, "B1", "Count")

	// Loop through the rows from the database.
	rowNumber := 2 // Start at row 2 since row 1 is for the headers.
	pairNumber := 1
	for rows.Next() {
		var id, count int
		err = rows.Scan(&id, &count)
		if err != nil {
			log.Fatal(err)
		}

		// If we've processed a pair of rows, add a new row with the pair number.
		if rowNumber%2 == 0 {
			f.InsertRow(sheetName, rowNumber+1)
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNumber+1), fmt.Sprintf("%d-я пара", pairNumber))
			pairNumber++
			rowNumber += 2 // Advance by two rows for the pair header
		}

		// Write the ID and count to the appropriate cells.
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNumber), id)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNumber), count)
		rowNumber++
	}


	// Save the file.
	err = f.SaveAs("data.xlsx")
	if err != nil {
		log.Fatal(err)

	}

	// Сохраняем файл Excel и отправляем его в ответе
	err = f.Write(w)
	if err != nil {
		panic(err.Error())
	}
}

func handleRequest() {
	var pic = &Pictures{}

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", hello).Methods("GET")
	rtr.HandleFunc("/login", login).Methods("POST", "GET")
	rtr.HandleFunc("/authorize", authorize).Methods("POST")
	rtr.HandleFunc("/forms", forms).Methods("GET", "POST")
	rtr.HandleFunc("/formsadmin", requireAuth(formsadmin)).Methods("GET", "POST")
	rtr.HandleFunc("/save", pic.save).Methods("POST")
	rtr.HandleFunc("/theend", theend).Methods("GET", "POST")
	rtr.HandleFunc("/generate-excel", generateExcel).Methods("GET", "POST")
	http.Handle("/", rtr)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
