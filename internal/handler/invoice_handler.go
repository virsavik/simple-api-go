package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"gokiosk/internal/model"
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

func (h InvoiceHandler) Route() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)

	r.Get("/{id}", h.GetByID)

	r.Post("/", h.Create)

	r.Put("/{id}", h.Update)

	r.Delete("/{id}", h.Delete)

	return r
}

func (h InvoiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// 1. Get offset and limit from request
	offset, limit := getPaginationParams(r)

	// 2. Get all invoices by offset and limit
	invoices, err := h.InvoiceService.GetAllInvoices(offset, limit)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusOK, invoices)
}

func (h InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// 1. Get id from request
	id := chi.URLParam(r, "id")

	// 2. Get invoice by id
	invoice, err := h.InvoiceService.GetInvoice(id)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusOK, invoice)
}

func (h InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice data from request body
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 2. Create invoice by invoice data
	result, err := h.InvoiceService.CreateInvoice(invoice)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusOK, result)
}

func (h InvoiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice id from request url param
	id := chi.URLParam(r, "id")

	// 2. Get invoice data from request body
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Update invoice by invoice data and invoice id
	result, err := h.InvoiceService.UpdateInvoice(id, invoice)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 4. Write response
	writeJsonResponse(w, http.StatusOK, result)
}

func (h InvoiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. Get invoice id from request url param
	id := chi.URLParam(r, "id")

	// 2. Delete invoice by invoice id
	err := h.InvoiceService.DeleteInvoice(id)
	if err != nil {
		// Write internal server error response if error occurs
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// 3. Write response
	writeJsonResponse(w, http.StatusOK, nil)
}
