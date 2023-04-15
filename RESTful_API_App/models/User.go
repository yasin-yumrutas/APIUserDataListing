package models

// import "Test/Interest"
// import "./Interest"
// // import 	"C:\Users\yasin\OneDrive\Masaüstü\l\yedek0\Data\Go\GoOn\RESTful_API_App\models\Interest"
// import "C:\Users\yasin\OneDrive\Masaüstü\l\yedek0\Data\Go\GoOn\RESTful_API_App\models\Interest"

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Profile   string
	Interest  []Interest
}
