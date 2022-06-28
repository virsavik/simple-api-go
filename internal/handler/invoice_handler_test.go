package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"gokiosk/internal/errors"
	"gokiosk/internal/handler/testdata/invoice_handler/fakedata"
	"gokiosk/internal/model"
	"gokiosk/internal/service/mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestInvoiceHandler_GetAll(t *testing.T) {
	type GivenData struct {
		offset   int
		limit    int
		fakeData []model.Invoice
	}

	type ExpectedData struct {
		statusCode int
		resultPath string
	}

	tcs := map[string]struct {
		given  GivenData
		exp    ExpectedData
		expErr error
	}{
		"success": {
			given: GivenData{
				offset:   0,
				limit:    2,
				fakeData: fakedata.GetAllFakeData,
			},
			exp: ExpectedData{
				statusCode: http.StatusOK,
				resultPath: "testdata/invoice_handler/response/get_all-success.json",
			},
		},
		"error_offset_or_limit_negative": {
			given: GivenData{
				offset: -1,
				limit:  2,
			},
			exp: ExpectedData{
				statusCode: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("GetAllByPaginate", tc.given.offset, tc.given.limit).Return(tc.given.fakeData, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices?offset=%d&limit=%d", tc.given.offset, tc.given.limit), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("offset", strconv.Itoa(tc.given.offset))              // Add offset to chi route context
			rctx.URLParams.Add("limit", strconv.Itoa(tc.given.limit))                // Add limit to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.GetAll(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.exp.statusCode, w.Code, "Should equal status")
				require.EqualError(t, tc.expErr, w.Body.String(), "Should be error")
			} else {
				// Should be success
				require.Equal(t, tc.exp.statusCode, w.Code, "Should equal status")

				var actualResult []model.Invoice
				if err := json.Unmarshal(w.Body.Bytes(), &actualResult); err != nil {
					t.Fatal(err)
				}

				resultBytes, rfErr := ioutil.ReadFile(tc.exp.resultPath)
				if rfErr != nil {
					t.Fatal(rfErr)
				}

				var expectedResult []model.Invoice
				if err := json.Unmarshal(resultBytes, &expectedResult); err != nil {
					t.Fatal(err)
				}

				require.Equal(t, expectedResult, actualResult, "Should equal expected result")
			}
		})
	}
}

func TestInvoiceHandler_GetByID(t *testing.T) {
	type GivenData struct {
		invID    string
		fakeData model.Invoice
	}

	type ExpectedData struct {
		statusCode int
		resultPath string
	}

	tcs := map[string]struct {
		given  GivenData
		exp    ExpectedData
		expErr error
	}{
		"success": {
			given: GivenData{
				invID:    "1",
				fakeData: fakedata.GetByIDFakeData,
			},
			exp: ExpectedData{
				statusCode: http.StatusOK,
				resultPath: "testdata/invoice_handler/response/get_by_id-success.json",
			},
		},
		"error_invoice_not_found": {
			given: GivenData{
				invID: "1",
			},
			exp: ExpectedData{
				statusCode: http.StatusNotFound,
			},
			expErr: fmt.Errorf(errors.ERR_NOT_FOUND),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("GetByID", tc.given.invID).Return(tc.given.fakeData, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices/%s", tc.given.invID), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.given.invID)                                 // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.GetByID(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.exp.statusCode, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.exp.statusCode, w.Code)

				var actualResult model.Invoice
				if err := json.Unmarshal(w.Body.Bytes(), &actualResult); err != nil {
					t.Fatal(err)
				}

				resultBytes, rfErr := ioutil.ReadFile(tc.exp.resultPath)
				if rfErr != nil {
					t.Fatal(rfErr)
				}
				var expectedResult model.Invoice
				if err := json.Unmarshal(resultBytes, &expectedResult); err != nil {
					t.Fatal(err)
				}

				require.Equal(t, expectedResult, actualResult, "Should equal expected result")
			}
		})
	}
}

func TestInvoiceHandler_Create(t *testing.T) {
	type GivenData struct {
		serviceMethodParam  model.Invoice
		serviceMethodResult model.Invoice
		reqBodyPath         string
	}

	type ExpectedData struct {
		statusCode int
		resultPath string
	}

	tcs := map[string]struct {
		given  GivenData
		exp    ExpectedData
		expErr error
	}{
		"success": {
			given: GivenData{
				serviceMethodParam:  fakedata.CreateSuccessParamFake,
				serviceMethodResult: fakedata.CreateSuccessResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/create-success.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusCreated,
				resultPath: "testdata/invoice_handler/response/create-success.json",
			},
		},
		"error_id_must_be_empty": {
			given: GivenData{
				serviceMethodParam: fakedata.CreateErrorIDMustBeEmptyFakeData,
				reqBodyPath:        "testdata/invoice_handler/request/create-error_id_must_be_empty.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_ID_MUST_BE_EMPTY),
		},
		"error_storekeeper_id_not_found": {
			given: GivenData{
				serviceMethodParam: fakedata.CreateErrorStorekeeperIDNotFoundFakeData,
				reqBodyPath:        "testdata/invoice_handler/request/create-error_storekeeper_id_not_found.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("Create", tc.given.serviceMethodParam).Return(tc.given.serviceMethodResult, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			reqBody, err := ioutil.ReadFile(tc.given.reqBodyPath)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("POST", "/invoices", strings.NewReader(string(reqBody)))
			r.Header.Set("Content-Type", "application/json") // Set content type application/json for request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.Create(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.exp.statusCode, w.Code)       // Check response code
				require.EqualError(t, tc.expErr, w.Body.String()) // check response body
			} else {
				// Should be success
				require.Equal(t, tc.exp.statusCode, w.Code) // Check response code

				var actualResult model.Invoice
				if err := json.Unmarshal(w.Body.Bytes(), &actualResult); err != nil {
					t.Fatal(err)
				}

				resultBytes, rfErr := ioutil.ReadFile(tc.exp.resultPath)
				if rfErr != nil {
					t.Fatal(rfErr)
				}

				var expectedResult model.Invoice
				if err = json.Unmarshal(resultBytes, &expectedResult); err != nil {
					t.Fatal(err)
				}

				require.Equal(t, expectedResult, actualResult, "Should be equal expected result") // check response body
			}
		})
	}
}

func TestInvoiceHandler_Update(t *testing.T) {
	type GivenData struct {
		invID               string
		serviceMethodParam  model.Invoice
		serviceMethodResult model.Invoice
		reqBodyPath         string
	}

	type ExpectedData struct {
		resultPath string
		statusCode int
	}

	tcs := map[string]struct {
		given  GivenData
		exp    ExpectedData
		expErr error
	}{
		"success": {
			given: GivenData{
				invID:               "1",
				serviceMethodParam:  fakedata.UpdateSuccessParamFake,
				serviceMethodResult: fakedata.UpdateSuccessResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/update-success.json",
			},
			exp: ExpectedData{
				resultPath: "testdata/invoice_handler/response/update-success.json",
				statusCode: http.StatusOK,
			},
		},
		"error_id_must_be_set": {
			given: GivenData{
				invID:               "",
				serviceMethodParam:  fakedata.UpdateSuccessParamFake,
				serviceMethodResult: fakedata.UpdateSuccessResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/update-error_id_must_be_set.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_ID_MUST_BE_SET),
		},
		"error_id_must_be_match": {
			given: GivenData{
				invID:               "2",
				serviceMethodParam:  fakedata.UpdateSuccessParamFake,
				serviceMethodResult: fakedata.UpdateSuccessResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/update-error_id_must_be_match.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH),
		},
		"error_invoice_not_found": {
			given: GivenData{
				invID:               "X",
				serviceMethodParam:  fakedata.UpdateErrorNotFoundParamFake,
				serviceMethodResult: fakedata.UpdateErrorNotFoundResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/update-error_invoice_not_found.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusNotFound,
			},
			expErr: fmt.Errorf(errors.ERR_NOT_FOUND),
		},
		"error_storekeeper_id_not_found": {
			given: GivenData{
				invID:               "1",
				serviceMethodParam:  fakedata.UpdateErrorStoreKeeperIDNotFoundParamFake,
				serviceMethodResult: fakedata.UpdateErrorStoreKeeperIDNotFoundResultFake,
				reqBodyPath:         "testdata/invoice_handler/request/update-error_storekeeper_id_not_found.json",
			},
			exp: ExpectedData{
				statusCode: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("Update", tc.given.invID, tc.given.serviceMethodParam).Return(tc.given.serviceMethodResult, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			reqBody, err := ioutil.ReadFile(tc.given.reqBodyPath)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("PUT", fmt.Sprintf("/invoices/%s", tc.given.invID), strings.NewReader(string(reqBody)))
			r.Header.Set("Content-Type", "application/json")                         // Set content type application/json for request
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.given.invID)                                 // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.Update(w, r)

			// 3. Then
			if tc.expErr != nil {
				require.Equal(t, tc.exp.statusCode, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				require.Equal(t, tc.exp.statusCode, w.Code)

				resultBytes, rfErr := ioutil.ReadFile(tc.exp.resultPath)
				if rfErr != nil {
					t.Fatal(rfErr)
				}
				var expectedResult model.Invoice
				if err = json.Unmarshal(resultBytes, &expectedResult); err != nil {
					t.Fatal(err)
				}

				var actualResult model.Invoice
				if err = json.Unmarshal(w.Body.Bytes(), &actualResult); err != nil {
					t.Fatal(err)
				}

				require.Equal(t, expectedResult, actualResult, "Should be equal expected result") // check response body
			}
		})
	}
}

func TestInvoiceHandler_Delete(t *testing.T) {
	type GivenData struct {
		invID string
	}

	type ExpectedData struct {
		statusCode int
	}

	tcs := map[string]struct {
		given     GivenData
		expResult ExpectedData
		expErr    error
	}{
		"success": {
			given: GivenData{
				invID: "1",
			},
			expResult: ExpectedData{
				statusCode: http.StatusNoContent,
			},
		},
		"error_invoice_not_found": {
			given: GivenData{
				invID: "1",
			},
			expResult: ExpectedData{
				statusCode: http.StatusNotFound,
			},
			expErr: fmt.Errorf(errors.ERR_NOT_FOUND),
		},
		"error_another_referenced_it": {
			given: GivenData{
				invID: "1",
			},
			expResult: ExpectedData{
				statusCode: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_RELATION_EXISTS),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("DeleteByID", tc.given.invID).Return(tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("DELETE", fmt.Sprintf("/invoices/%s", tc.given.invID), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.given.invID)                                 // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.Delete(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expResult.statusCode, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expResult.statusCode, w.Code)
			}
		})
	}
}
