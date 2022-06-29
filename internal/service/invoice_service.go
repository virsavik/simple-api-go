package service

import (
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/repository"
	"sync"
	"time"
)

type InvoiceService struct {
	InvoiceRepository repository.IInvoiceRepository
}

func NewInvoiceService(invoiceRepository repository.IInvoiceRepository) InvoiceService {
	return InvoiceService{
		InvoiceRepository: invoiceRepository,
	}
}

func (i InvoiceService) GetAllByDuration(from time.Time, to time.Time) ([]model.Invoice, error) {
	// 1. Get all invoices
	invoices, err := i.InvoiceRepository.GetAllByDuration(from, to)
	if err != nil {
		return nil, err
	}

	if len(invoices) == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	return invoices, nil
}

func (i InvoiceService) GenCSVReportByDuration(from time.Time, to time.Time) ([]byte, error) {
	// 1. Get all invoices
	invoices, err := i.InvoiceRepository.GetAllByDuration(from, to)
	if err != nil {
		return nil, err
	}

	invoiceLen := len(invoices)
	if invoiceLen == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	buffSize := 20 // TODO: Find a better way to choose the buffer size
	workerNum := 4 // TODO: Find a better way to choose the worker number

	dataCh := make(chan []byte, buffSize)
	defer close(dataCh)

	wg := sync.WaitGroup{}
	wg.Add(workerNum)

	var csvData []byte

	for i := 0; i < workerNum; i++ {
		lo := i * invoiceLen / workerNum
		hi := (i + 1) * invoiceLen / workerNum
		go writeCsvWorker(invoices[lo:hi], dataCh, &wg)
	}

	// TODO: Is this goroutine ?
	for i := 0; i < invoiceLen; i++ {
		csvData = append(csvData, <-dataCh...)
	}

	wg.Wait()

	return csvData, nil
}

func toString(invoice model.Invoice) string {
	return invoice.ID + "," + invoice.StorekeeperID + "," + invoice.CreatedAt.Format("2006-01-02 15:04:05") + "," + invoice.UpdatedAt.Format("2006-01-02 15:04:05")
}

func writeCsvWorker(invoices []model.Invoice, dataCh chan<- []byte, wg *sync.WaitGroup) {
	for _, invoice := range invoices {
		dataCh <- []byte(toString(invoice) + "\n")
	}
	wg.Done()
}

func (i InvoiceService) GetAllByPaginate(offset, limit int) ([]model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceService) GetByID(id string) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceService) Create(invoice model.Invoice) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceService) Update(id string, invoice model.Invoice) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceService) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
