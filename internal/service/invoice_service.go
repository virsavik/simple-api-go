package service

import (
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/repository"
	"strings"
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

func (inv InvoiceService) GetAllByDuration(from time.Time, to time.Time) ([]model.Invoice, error) {
	// 1. Get all invoices
	invoices, err := inv.InvoiceRepository.GetAllByDuration(from, to)
	if err != nil {
		return nil, err
	}

	if len(invoices) == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	return invoices, nil
}

type Input struct {
	Idx     int
	Invoice model.Invoice
}

type Record struct {
	Idx  int
	Data string
}

func (inv InvoiceService) GenCSVReportByDuration(from time.Time, to time.Time) ([]byte, error) {
	// 1. Get all invoices by duration sorted by created_at
	invoices, err := inv.InvoiceRepository.GetAllByDuration(from, to)
	if err != nil {
		return nil, err
	}

	invoiceLen := len(invoices)
	if invoiceLen == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	// TODO: Find a better way to choose the worker number and buffSize
	workerNum := 4
	buffSize := workerNum
	var wg sync.WaitGroup

	invCh := make(chan Input, buffSize)
	recordCh := make(chan Record, buffSize)

	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go writeCSVWorker(i, invCh, recordCh, &wg)
	}

	go func() {
		for i := 0; i < invoiceLen; i++ {
			invCh <- Input{i, invoices[i]}
		}
		// Close the channel when all invoices are sent
		close(invCh)
	}()

	sortedRecords := make([]string, invoiceLen, invoiceLen)
	for i := 0; i < invoiceLen; i++ {
		record := <-recordCh
		sortedRecords[record.Idx] = record.Data
	}

	wg.Wait()

	// Be careful big-O of strings.Join
	return []byte(strings.Join(sortedRecords, "\n")), nil
}

func toString(invoice model.Invoice) string {
	return invoice.ID + "," + invoice.StorekeeperID + "," + invoice.CreatedAt.Format("2006-01-02 15:04:05") + "," + invoice.UpdatedAt.Format("2006-01-02 15:04:05")
}

func writeCSVWorker(id int, inpCh <-chan Input, recordCh chan<- Record, wg *sync.WaitGroup) {
	for in := range inpCh {
		recordCh <- Record{in.Idx, toString(in.Invoice)}
	}
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func (inv InvoiceService) GetAllByPaginate(offset, limit int) ([]model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) GetByID(id string) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) Create(invoice model.Invoice) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) Update(id string, invoice model.Invoice) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
