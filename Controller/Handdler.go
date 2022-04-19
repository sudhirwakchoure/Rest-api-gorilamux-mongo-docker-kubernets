package Controller

import (
	"REST_API/Model"
	utility "REST_API/Utility"
	"os"

	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

//home page
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UNIVERCITY ")

}

// swagger:operation POST /students Student Poststudent
//
// Add new Student
//
// Returns new Student
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: Student
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"
//   '405':
//     description: Method not allowed

//Poststudent in DB
func Poststudent(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "New User creating ")

	//fmt.Println("Show result ", r.Body)
	var student Model.Student
	collection := utility.NewFunction()
	//log.Println("*********************", collection)
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	fmt.Println("Test", student)
	result, err := collection.InsertOne(ctx, student)
	//log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$", err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Println(result)
	log.Println(result)
	id, ok := result.InsertedID.(primitive.ObjectID)
	log.Println(id, ok)
	student.Id = id

	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(student)

}

// swagger:operation GET /students Student GetStudent
//
// Get Student
//
// Returns existing Student
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"

//GetStudent fetch all  from db
func GetStudent(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Display all student  ")

	var allstudent []Model.Student
	collection := utility.NewFunction()
	ctx := r.Context()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for cur.Next(ctx) {
		var student Model.Student
		err := cur.Decode(&student)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		allstudent = append(allstudent, student)
	}

	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(allstudent)

}

func GetStudentbyName(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	collection := utility.NewFunction()

	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	year := r.URL.Query().Get("yearofaddmision")
	params := []primitive.M{}

	filter := primitive.M{}

	if name != "" {
		params = append(params, primitive.M{"name": name})
		filter = primitive.M{"name": name}
	}
	if city != "" {
		params = append(params, primitive.M{"city": city})
		filter = primitive.M{"city": city}
	}
	if year != "" {
		params = append(params, primitive.M{"yearofaddmision": year})
		filter = primitive.M{"yearofaddmision": year}
	}
	if len(params) > 1 {
		filter = primitive.M{"$and": params}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return
	}
	students := []Model.Student{}
	for cur.Next(ctx) {
		var std Model.Student
		err := cur.Decode(&std)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		students = append(students, std)

	}
	w.Header().Add("content-type", "appllication/json")
	json.NewEncoder(w).Encode(students)

}

// swagger:operation GET /students/{id} Student GetStudentbyid
//
// Get Student
//
// Returns existing Student filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"
//   '400':
//     description: bad request
//   '404':
//     description: page not found

//GetStudentbyid fetch Student from db
func GetStudentbyid(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "GET BY ID STUDENT ")

	w.Header().Add("content-type", "application/json")
	var student Model.Student
	id := mux.Vars(r)["id"]

	ID, err := primitive.ObjectIDFromHex(id)
	log.Println(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(err)
		return

	}

	filter := primitive.M{"_id": ID}

	collection := utility.NewFunction()
	ctx := r.Context()
	err = collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
	}
	w.Header().Add("content-type", "appllication/json")
	json.NewEncoder(w).Encode(student)

}

// swagger:operation DELETE /students/{id} Student StudentDeletestudent
//
// Delete  student
//
// Delete existing Student filtered by id
//
// ---
//
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"
//   '410':
//     description: staus gone  already deleted

//Deletestudent fetch Student from db
func Deletestudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User ")
	id := mux.Vars(r)["id"]

	ctx := r.Context()

	collection := utility.NewFunction()

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := primitive.M{"_id": ID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
		//log.Fatal(err)
	}
	log.Println(result)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(result)

}

// swagger:operation PUT /students/{id} Student Updatestudent
//
// Update Student
//
// Update existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: name
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"
//   '200':
//     description: Student data
//   '204':
//     description: no content

//updatestudent----
func Updatestudent(w http.ResponseWriter, r *http.Request) {

	collection := utility.NewFunction()
	ctx := r.Context()

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		return
	}
	var student Model.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
	}
	filter := bson.M{"_id": id}

	update := primitive.M{
		"name":            student.Name,
		"city":            student.City,
		"country":         student.Country,
		"course":          student.Course,
		"yearofadmission": student.YearOfAdmission}

	err = collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
		return

	}
	log.Println(student)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(student)

}

// swagger:operation PATCH /students/{id} Student Updatestudent_patch
//
//  Parratially Update Student
//
// Patch existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: name
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"
//   '204':
//     description: no content

//Updatestudent_patch----
func Updatestudent_patch(w http.ResponseWriter, r *http.Request) {

	collection := utility.NewFunction()
	ctx := r.Context()

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	var student Model.Student
	filter := primitive.M{"_id": id}

	var update map[string]string

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}
	err = collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return

	}
	log.Println(student)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(student)

}

func Test(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)

}
