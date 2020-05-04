package employee

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEmployeeByIdHandler := kithttp.NewServer(
		makeGetEmployeeByIdEndPoint(s),
		getEmployeeByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getEmployeeByIdHandler)

	getEmployeesHandler := kithttp.NewServer(
		makeGetEmployeesEndPoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeTopHandler := kithttp.NewServer(
		makeGetEmployeeTopEndPoint(s),
		getEmployeeTopRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/best-seller", getEmployeeTopHandler)

	return r
}

func getEmployeeByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	employeeId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getEmployeeByIdRequest{EmployeId: employeeId}, nil
}

func getEmployeesRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeeTopRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeTopRequest{}, nil
}
