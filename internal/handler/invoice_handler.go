package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/orm"
	"gokiosk/internal/service"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceService service.IInvoiceService
}

func NewInvoiceHandler(invoiceService service.IInvoiceService) InvoiceHandler {
	return InvoiceHandler{
		InvoiceService: invoiceService,
	}
}

func (h InvoiceHandler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate input data

	// 1. Get offset and limit from request
	page := getPaginationParams(r)

	// 2. Throw bad request if offset or limit is not valid
	if page.Offset < 0 || page.Limit < 0 {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE))
		return
	}

	// 3. Get all invoices by offset and limit
	invoices, err := h.InvoiceService.GetAllByPaginate(page)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Convert to Handler Response (invoices, pagination)
	type Response struct {
		Invoices []orm.Invoice  `json:"invoices"`
		Paginate model.Paginate `json:"paginate"`
	}

	// 4. Write response
	writeJsonResponse(w, http.StatusOK, Response{Invoices: invoices, Paginate: page})
}

func (h InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// 1. Get id from request
	id := chi.URLParam(r, "id")

	// 2. Get invoice by id
	invoice, err := h.InvoiceService.GetByID(id)
	if err != nil {
		if err.Error() == errors.ERR_NOT_FOUND {
			// Write not found response if error occurs
			writeErrorResponse(w, http.StatusNotFound, err)
			return
		}
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusOK, invoice)
}

func (h InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice data from request body
	var invoice orm.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 2. Write bad request if invoice id is set
	if invoice.ID != "" {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(errors.ERR_ID_MUST_BE_EMPTY))
		return
	}

	// 3. Create invoice by invoice data
	result, err := h.InvoiceService.Create(invoice)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 4. Write response
	writeJsonResponse(w, http.StatusCreated, result)
}

func (h InvoiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice id from request url param
	id := chi.URLParam(r, "id")

	// 2. Get invoice data from request body
	var invoice orm.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write bad request if "request url param "id" is not set
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(errors.ERR_ID_MUST_BE_SET))
		return
	}

	// 4. Write bad request if "invoice id" is not match with request url param "id"
	if invoice.ID != id {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH))
		return
	}

	// 4. Update invoice by invoice data and invoice id
	result, err := h.InvoiceService.Update(id, invoice)
	if err != nil {
		if err.Error() == errors.ERR_NOT_FOUND {
			// Write not found response if error occurs
			writeErrorResponse(w, http.StatusNotFound, err)
			return
		}
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 5. Write response
	writeJsonResponse(w, http.StatusOK, result)
}

func (h InvoiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice id from request url param
	id := chi.URLParam(r, "id")

	// 2. Delete invoice by invoice id
	if err := h.InvoiceService.DeleteByID(id); err != nil {
		if err.Error() == errors.ERR_NOT_FOUND {
			// Write not found response if error occurs
			writeErrorResponse(w, http.StatusNotFound, err)
			return
		}
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusNoContent, nil)
}
