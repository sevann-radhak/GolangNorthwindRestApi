package order

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

	addOrderHandler := kithttp.NewServer(
		makeAddOrderEndPoint(s),
		addOrderRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/", addOrderHandler)

	deleteOrderDetailHandler := kithttp.NewServer(
		makeDeleteOrderDetailEndPoint(s),
		deleteOrderDetailByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodDelete, "/{id}/detail/{orderDetailId}", deleteOrderDetailHandler)

	deleteOrderHandler := kithttp.NewServer(
		makeDeleteOrderEndPoint(s),
		deleteOrderByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodDelete, "/{id}", deleteOrderHandler)

	getOrderItemByIdHandler := kithttp.NewServer(
		makeGetOrderItemByIdEndPoint(s),
		getOrderItemByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getOrderItemByIdHandler)

	getOrdersHandler := kithttp.NewServer(
		makeGetOrdersEndPoint(s),
		getOrdersRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	return r
}

func addOrderRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func deleteOrderByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	orderId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return deleteOrderRequest{OrderId: orderId}, nil
}

func deleteOrderDetailByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderDetailRequest{OrderDetailId: chi.URLParam(r, "orderDetailId")}, nil
}

func getOrderItemByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	orderId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getOrderItemByIdRequest{OrderId: orderId}, nil
}

func getOrdersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
