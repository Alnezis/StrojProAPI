package controllers

import (
	"StrojProAPI/etc"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func GetStagesByObjName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := r.URL.Query()

	if params["object_name"] == nil || params["object_name"][0] == "" {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: "object_name отсутствует в запросе",
				Code:    202,
			},
		})
		return
	}

	if !etc.ExistObjectName(params["object_name"][0]) {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: fmt.Sprintf("object_name (%s) отсутствует в базе", params["object_name"][0]),
				Code:    500,
			},
		})
		return
	}

	i := etc.GetStagesByObjName(params["object_name"][0])

	json.NewEncoder(w).Encode(&Response{
		Result: i,
	})
}

func GetSubStageListByObjNameStage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := r.URL.Query()

	if params["object_name"] == nil || params["object_name"][0] == "" {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: "object_name отсутствует в запросе",
				Code:    202,
			},
		})
		return
	}

	if !etc.ExistObjectName(params["object_name"][0]) {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: fmt.Sprintf("object_name (%s) отсутствует в базе", params["object_name"][0]),
				Code:    500,
			},
		})
		return
	}

	if params["stage_name"] == nil || params["stage_name"][0] == "" {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: "stage_name отсутствует в запросе",
				Code:    202,
			},
		})
		return
	}

	if !etc.ExistStageName(params["stage_name"][0]) {
		json.NewEncoder(w).Encode(&Response{
			Error: &Error{
				Message: fmt.Sprintf("stage_name (%s) отсутствует в базе", params["stage_name"][0]),
				Code:    500,
			},
		})
		return
	}

	i := etc.GetSubStageListByObjNameStage(params["object_name"][0], params["stage_name"][0])

	json.NewEncoder(w).Encode(&Response{
		Result: i,
	})
}

func GetObjList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	i := etc.GetObjList()

	json.NewEncoder(w).Encode(&Response{
		Result: i,
	})
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	//  Ensure our file does not exceed 5MB
	r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024)

	file, handler, err := r.FormFile("image")

	// Capture any errors that may arise
	if err != nil {
		fmt.Fprintf(w, "Error getting the file")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded file name: %+v\n", handler.Filename)
	fmt.Printf("Uploaded file size %+v\n", handler.Size)
	fmt.Printf("File mime type %+v\n", handler.Header)

	// Get the file content type and access the file extension
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[1]

	// Create the temporary file name
	fileName := fmt.Sprintf("upload-*.%s", fileType)
	// Create a temporary file with a dir folder
	tempFile, err := ioutil.TempFile("images", fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully uploaded file")
}
