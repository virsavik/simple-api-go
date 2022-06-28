package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/service/mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type paginationTest struct {
	offset int
	limit  int
}

func TestInvoiceHandler_GetAll(t *testing.T) {
	tcs := map[string]struct {
		given        paginationTest
		expStatus    int
		expResultLoc string
		expErr       error
	}{
		"success": {
			given: paginationTest{
				offset: 0,
				limit:  2,
			},
			expStatus:    http.StatusOK,
			expResultLoc: "testdata/invoice_handler/response/get_all-success.json",
		},
		"error_offset_or_limit_negative": {
			given: paginationTest{
				offset: 0,
				limit:  -2,
			},
			expStatus: http.StatusBadRequest,
			expErr:    fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			var expResult []model.Invoice

			if tc.expResultLoc != "" {
				plan, err := ioutil.ReadFile(tc.expResultLoc)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(plan, &expResult)
				if err != nil {
					t.Fatal(err)
				}
			}

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("GetAllByPaginate", tc.given.offset, tc.given.limit).Return(expResult, tc.expErr)
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
				require.Equal(t, tc.expStatus, w.Code, "Should equal status")
				require.EqualError(t, tc.expErr, w.Body.String(), "Should be error")
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code, "Should equal status")
				var actualResult []model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &actualResult)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, expResult, actualResult, "Should equal expected result")
			}
		})
	}
}

func TestInvoiceHandler_GetByID(t *testing.T) {
	tcs := map[string]struct {
		given        string
		expStatus    int
		expResultLoc string
		expErr       error
	}{
		"success": {
			given:        "3",
			expResultLoc: "testdata/invoice_handler/response/get_by_id-success.json",
			expStatus:    http.StatusOK,
		},
		"error_invoice_not_found": {
			given:     "XXXX",
			expErr:    fmt.Errorf(errors.ERR_NOT_FOUND),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			var expResult model.Invoice

			if tc.expResultLoc != "" {
				plan, err := ioutil.ReadFile(tc.expResultLoc)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(plan, &expResult)
				if err != nil {
					t.Fatal(err)
				}
			}

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("GetByID", tc.given).Return(expResult, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices/%s", tc.given), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.given)                                       // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.GetByID(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code)
				var actualResult model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &actualResult)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, expResult, actualResult, "Should equal expected result")
			}
		})
	}
}

func TestInvoiceHandler_Create(t *testing.T) {
	tcs := map[string]struct {
		given        string
		expResultLoc string
		expStatus    int
		expErr       error
	}{
		"success": {
			given:        "testdata/invoice_handler/request/create-success.json",
			expResultLoc: "testdata/invoice_handler/response/create-success.json",
			expStatus:    http.StatusCreated,
		},
		"error_id_must_be_empty": {
			given:     "testdata/invoice_handler/request/create-error_id_must_be_empty.json",
			expErr:    fmt.Errorf(errors.ERR_ID_MUST_BE_EMPTY),
			expStatus: http.StatusBadRequest,
		},
		"error_storekeeper_id_not_found": {
			given:     "testdata/invoice_handler/request/create-error_storekeeper_id_not_found.json",
			expErr:    fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			reqBody, err := ioutil.ReadFile(tc.given)
			if err != nil {
				t.Fatal(err)
			}

			var invoice model.Invoice
			err = json.Unmarshal(reqBody, &invoice)
			if err != nil {
				t.Fatal(err)
			}

			var expResult model.Invoice
			if tc.expResultLoc != "" {
				plan, err := ioutil.ReadFile(tc.expResultLoc)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(plan, &expResult)
				if err != nil {
					t.Fatal(err)
				}
			}

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("Create", invoice).Return(expResult, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("POST", "/invoices", strings.NewReader(string(reqBody)))
			r.Header.Set("Content-Type", "application/json") // Set content type application/json for request
			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.Create(w, r)

			// 3. Then
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)            // Check response code
				require.EqualError(t, tc.expErr, w.Body.String()) // check response body
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code) // Check response code
				var actualResult model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &actualResult)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, expResult, actualResult, "Should be equal expected result") // check response body
			}
		})
	}
}

func TestInvoiceHandler_Update(t *testing.T) {
	type UpdateGiven struct {
		invID      string
		reqBodyLoc string
	}

	type UpdateExpected struct {
		expResultLoc string
		expStatus    int
	}

	tcs := map[string]struct {
		given     UpdateGiven
		expResult UpdateExpected
		expErr    error
	}{
		"success": {
			given: UpdateGiven{
				invID:      "2",
				reqBodyLoc: "testdata/invoice_handler/request/update-success.json",
			},
			expResult: UpdateExpected{
				expResultLoc: "testdata/invoice_handler/response/update-success.json",
				expStatus:    http.StatusOK,
			},
		},
		"error_id_must_be_set": {
			given: UpdateGiven{
				invID:      "",
				reqBodyLoc: "testdata/invoice_handler/request/update-error_id_must_be_set.json",
			},
			expResult: UpdateExpected{
				expStatus: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_ID_MUST_BE_SET),
		},
		"error_id_must_be_match": {
			given: UpdateGiven{
				invID:      "1",
				reqBodyLoc: "testdata/invoice_handler/request/update-error_id_must_be_match.json",
			},
			expResult: UpdateExpected{
				expStatus: http.StatusBadRequest,
			},
			expErr: fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH),
		},
		"error_invoice_not_found": {
			given: UpdateGiven{
				invID:      "2",
				reqBodyLoc: "testdata/invoice_handler/request/update-error_invoice_not_found.json",
			},
			expResult: UpdateExpected{
				expStatus: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_NOT_FOUND),
		},
		"error_storekeeper_id_not_found": {
			given: UpdateGiven{
				invID:      "2",
				reqBodyLoc: "testdata/invoice_handler/request/update-error_storekeeper_id_not_found.json",
			},
			expResult: UpdateExpected{
				expStatus: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given
			reqBody, err := ioutil.ReadFile(tc.given.reqBodyLoc)
			if err != nil {
				t.Fatal(err)
			}

			var invoice model.Invoice
			err = json.Unmarshal(reqBody, &invoice)
			if err != nil {
				t.Fatal(err)
			}

			var expResult model.Invoice
			if tc.expResult.expResultLoc != "" {
				plan, err := ioutil.ReadFile(tc.expResult.expResultLoc)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(plan, &expResult)
				if err != nil {
					t.Fatal(err)
				}
			}

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("Update", tc.given.invID, invoice).Return(expResult, tc.expErr)
			invoiceHandler := NewInvoiceHandler(serviceMock)

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
				require.Equal(t, tc.expResult.expStatus, w.Code)  // Check response code
				require.EqualError(t, tc.expErr, w.Body.String()) // check response body
			} else {
				require.Equal(t, tc.expResult.expStatus, w.Code) // Check response code
				var actualResult model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &actualResult)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, expResult, actualResult, "Should be equal expected result") // check response body
			}
		})
	}
}

func TestInvoiceHandler_Delete(t *testing.T) {
	type DeleteGiven struct {
		invID string
	}

	type DeleteExpected struct {
		expResultLoc string
		expStatus    int
	}

	tcs := map[string]struct {
		given     DeleteGiven
		expResult DeleteExpected
		expErr    error
	}{
		"success": {
			given: DeleteGiven{
				invID: "1",
			},
			expResult: DeleteExpected{
				expStatus: http.StatusNoContent,
			},
		},
		"error_invoice_not_found": {
			given: DeleteGiven{
				invID: "1",
			},
			expResult: DeleteExpected{
				expStatus: http.StatusInternalServerError,
			},
			expErr: fmt.Errorf(errors.ERR_NOT_FOUND),
		},
		"error_another_referenced_it": {
			given: DeleteGiven{
				invID: "1",
			},
			expResult: DeleteExpected{
				expStatus: http.StatusInternalServerError,
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
				require.Equal(t, tc.expResult.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expResult.expStatus, w.Code)
			}

		})
	}
}
