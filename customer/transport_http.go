package customer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getCustomersHandler := kithttp.NewServer(
		makeGetCustomersEndPoint(s),
		getCustomersRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getCustomersHandler)

	return r
}

func getCustomersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getCustomersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
