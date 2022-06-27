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
	"gokiosk/test/testdata"
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
		input        paginationTest
		expStatus    int
		expResultLen int
		expErr       error
	}{
		"success": {
			input: paginationTest{
				offset: 0,
				limit:  2,
			},
			expStatus:    http.StatusOK,
			expResultLen: 2,
		},
		"error_offset_or_limit_negative": {
			input: paginationTest{
				offset: 0,
				limit:  -2,
			},
			expErr: fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Given

			serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
			serviceMock.On("GetAllByPaginate", 0, 2).Return(testdata.Invoices[0:2], nil).Once()
			serviceMock.On("GetAllByPaginate", 0, -2).Return(nil, nil).Once()
			invoiceHandler := NewInvoiceHandler(serviceMock)

			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices?offset=%d&limit=%d", tc.input.offset, tc.input.limit), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("offset", strconv.Itoa(tc.input.offset))              // Add offset to chi route context
			rctx.URLParams.Add("limit", strconv.Itoa(tc.input.limit))                // Add limit to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			w := httptest.NewRecorder()

			// 2. When
			invoiceHandler.GetAll(w, r)

			// 3. Then

			if tc.expErr != nil {
				// Should be error
				require.EqualError(t, tc.expErr, w.Body.String(), "Should be error")
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code, "Should equal status")
				var result []model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResultLen, len(result), "Should equal result length")
			}
		})
	}
}

func TestInvoiceHandler_GetByID(t *testing.T) {
	serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
	invoiceHandler := NewInvoiceHandler(serviceMock)

	tcs := map[string]struct {
		input     string
		expResult model.Invoice
		expStatus int
		expErr    error
	}{
		"success": {
			input: "3",
			expResult: model.Invoice{
				ID:            "3",
				StorekeeperID: "KEEPER-0003",
			},
			expStatus: http.StatusOK,
		},
		"error_invoice_not_found": {
			input:     "XXXX",
			expErr:    fmt.Errorf(errors.ERR_NOT_FOUND),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Define http test request with query params for offset and limit
			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices/%s", tc.input), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.input)                                       // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler getByID
			invoiceHandler.GetByID(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code)
				var result model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResult, result)
			}
		})
	}
}

func TestInvoiceHandler_Create(t *testing.T) {
	serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
	invoiceHandler := NewInvoiceHandler(serviceMock)

	tcs := map[string]struct {
		input     model.Invoice
		expResult model.Invoice
		expStatus int
		expErr    error
	}{
		"success": {
			input: model.Invoice{
				StorekeeperID: "KEEPER-0004",
			},
			expResult: model.Invoice{
				ID:            "7",
				StorekeeperID: "KEEPER-0004",
			},
			expStatus: http.StatusCreated,
		},
		"error_id_must_be_empty": {
			input: model.Invoice{
				ID:            "7",
				StorekeeperID: "KEEPER-0004",
			},
			expErr:    fmt.Errorf(errors.ERR_ID_MUST_BE_EMPTY),
			expStatus: http.StatusBadRequest,
		},
		"error_storekeeper_id_not_found": {
			input: model.Invoice{
				StorekeeperID: "KEEPER-XXXX",
			},
			expErr:    fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			reqBody, err := json.Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			// 1. Define http test request with query params for offset and limit
			r := httptest.NewRequest("POST", "/invoices", strings.NewReader(string(reqBody)))
			r.Header.Set("Content-Type", "application/json") // Set content type application/json for request

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler Create
			invoiceHandler.Create(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)            // Check response code
				require.EqualError(t, tc.expErr, w.Body.String()) // check response body
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code) // Check response code
				var result model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResult, result) // check response body
			}
		})
	}
}

func TestInvoiceHandler_Update(t *testing.T) {
	serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
	invoiceHandler := NewInvoiceHandler(serviceMock)

	tcs := map[string]struct {
		inReqParam string
		inReqBody  model.Invoice
		expResult  model.Invoice
		expStatus  int
		expErr     error
	}{
		"success": {
			inReqParam: "2",
			inReqBody: model.Invoice{
				ID:            "2",
				StorekeeperID: "KEEPER-0002",
			},
			expResult: model.Invoice{
				ID:            "2",
				StorekeeperID: "KEEPER-0002",
			},
			expStatus: http.StatusOK,
		},
		"error_id_must_be_set": {
			inReqBody: model.Invoice{
				ID:            "3",
				StorekeeperID: "KEEPER-0005",
			},
			expErr:    fmt.Errorf(errors.ERR_ID_MUST_BE_SET),
			expStatus: http.StatusBadRequest,
		},
		"error_id_must_be_match": {
			inReqParam: "2",
			inReqBody: model.Invoice{
				ID:            "3",
				StorekeeperID: "KEEPER-0005",
			},
			expErr:    fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH),
			expStatus: http.StatusBadRequest,
		},
		"error_invoice_not_found": {
			inReqParam: "8",
			inReqBody: model.Invoice{
				ID:            "8",
				StorekeeperID: "KEEPER-0003",
			},
			expErr:    fmt.Errorf(errors.ERR_NOT_FOUND),
			expStatus: http.StatusInternalServerError,
		},
		"error_storekeeper_id_not_found": {
			inReqParam: "2",
			inReqBody: model.Invoice{
				ID:            "2",
				StorekeeperID: "KEEPER-XXXX",
			},
			expErr:    fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			reqBody, err := json.Marshal(tc.inReqBody)
			if err != nil {
				t.Fatal(err)
			}

			// 1. Define http test request with query params for offset and limit
			r := httptest.NewRequest("PUT", fmt.Sprintf("/invoices/%s", tc.inReqParam), strings.NewReader(string(reqBody)))
			r.Header.Set("Content-Type", "application/json")                         // Set content type application/json for request
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.inReqParam)                                  // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler Create
			invoiceHandler.Update(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)            // Check response code
				require.EqualError(t, tc.expErr, w.Body.String()) // check response body
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code) // Check response code
				var result model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResult, result) // check response body
			}
		})
	}
}

func TestInvoiceHandler_Delete(t *testing.T) {
	serviceMock := new(mocks.InvoiceServiceMock) // Init mock service
	invoiceHandler := NewInvoiceHandler(serviceMock)

	tcs := map[string]struct {
		in        string
		expStatus int
		expResult string
		expErr    error
	}{
		"success": {
			in:        "6",
			expStatus: http.StatusNoContent,
			expResult: "",
		},
		"error_invoice_not_found": {
			in:        "8",
			expStatus: http.StatusInternalServerError,
			expErr:    fmt.Errorf(errors.ERR_NOT_FOUND),
		},
		"error_another_referenced_it": {
			in:        "1",
			expErr:    fmt.Errorf(errors.ERR_RELATION_EXISTS),
			expStatus: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Define http test request with query params for offset and limit
			r := httptest.NewRequest("DELETE", fmt.Sprintf("/invoices/%s", tc.in), nil)
			rctx := chi.NewRouteContext()                                            // Init chi route context
			rctx.URLParams.Add("id", tc.in)                                          // Set invoice id to chi route context
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx)) // Add chi route context to request

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler Delete
			invoiceHandler.Delete(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Should be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expStatus, w.Code)
				require.Equal(t, tc.expResult, w.Body.String())
			}
		})
	}
}
