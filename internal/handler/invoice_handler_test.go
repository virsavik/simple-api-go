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
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type paginationTest struct {
	offset int
	limit  int
}

type ExpResult struct {
	code int
	body []byte
}

func TestInvoiceHandler_getAll(t *testing.T) {
	serviceMock := mocks.NewInvoiceServiceMock()
	invoiceHandler := NewInvoiceHandler(serviceMock)

	tcs := map[string]struct {
		input         paginationTest
		expResultCode int
		expResultLen  int
		expErr        error
	}{
		"success": {
			input: paginationTest{
				offset: 0,
				limit:  2,
			},
			expResultCode: http.StatusOK,
			expResultLen:  2,
		},
		"error_offset_or_limit_negative": {
			input: paginationTest{
				offset: 0,
				limit:  -2,
			},
			expErr: errors.NewAppError(
				errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE,
				"OFFSET_AND_LIMIT_MUST_BE_POSITIVE",
			),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// 1. Define http test request with query params for offset and limit
			r := httptest.NewRequest("GET", fmt.Sprintf("/invoices?offset=%d&limit=%d", tc.input.offset, tc.input.limit), nil)
			// Init chi route context
			rctx := chi.NewRouteContext()
			// Set offset and limit to chi route context
			rctx.URLParams.Add("offset", strconv.Itoa(tc.input.offset))
			rctx.URLParams.Add("limit", strconv.Itoa(tc.input.limit))
			// Add chi route context to request
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler getAll
			invoiceHandler.GetAll(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Should be error
				require.EqualError(t, tc.expErr, w.Body.String())
			} else {
				// Should be success
				require.Equal(t, tc.expResultCode, w.Code)
				var result []model.Invoice
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResultLen, len(result))
			}
		})
	}
}

//func TestInvoiceHandler_GetByID(t *testing.T) {
//	serviceMock := mocks.NewInvoiceServiceMock()
//	invoiceHandler := NewInvoiceHandler(serviceMock)
//
//	tcs := map[string]struct {
//		input string
//		expResult
//	}
//}
