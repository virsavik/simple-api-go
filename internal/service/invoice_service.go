package service

import (
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/repository"
	"gokiosk/internal/repository/orm"
	"strings"
	"sync"
)

type InvoiceService struct {
	InvoiceRepository repository.IInvoiceRepository
	MailService       IMailService
}

func NewInvoiceService(invoiceRepository repository.IInvoiceRepository, mailService IMailService) InvoiceService {
	return InvoiceService{
		InvoiceRepository: invoiceRepository,
		MailService:       mailService,
	}
}

// TODO: Using Struct for function argument
//type Input struct {
//	from time.Time
//	to   time.Time
//}

func (inv InvoiceService) GetInvoicesByDuration(duration model.Duration) ([]orm.Invoice, error) {
	// 1. Get all invoices
	invoices, err := inv.InvoiceRepository.GetInvoicesByDuration(duration)
	if err != nil {
		return nil, err
	}

	if len(invoices) == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	return invoices, nil
}

// Input is used to send data to workers
type Input struct {
	Idx     int
	Invoice orm.Invoice
}

// Record is used to receive data from workers
type Record struct {
	Idx  int
	Data string
}

// GetCSVReport return csv report by duration
func (inv InvoiceService) GetCSVReport(duration model.Duration) ([]byte, error) {
	csvData, err := inv.generateCSVReportByDuration(duration)
	if err != nil {
		// TODO: Load admin email from config and mail template
		inv.MailService.Send([]string{"admin@example.com"}, []byte("Error when create CSV file"))
		return nil, err
	}

	inv.MailService.Send([]string{"admin@example.com"}, []byte("Export csv successfully!"))
	return csvData, nil
}

// GenerateCSVInvoice generate csv invoice

func (inv InvoiceService) generateCSVReportByDuration(duration model.Duration) ([]byte, error) {
	// STEP 1.
	// Get all invoices by duration sorted by created_at
	invoices, repoErr := inv.InvoiceRepository.GetInvoicesByDuration(duration)
	if repoErr != nil {
		return nil, repoErr
	}

	// STEP 1.2
	// Return error ERR_NO_DATA if no invoices found
	invoiceLen := len(invoices)
	if invoiceLen == 0 {
		return nil, fmt.Errorf(errors.ERR_NO_DATA)
	}

	// STEP 2.
	// Define worker number and buffer size
	workerNum := 4 // TODO: Choose worker number based on invoiceLen later
	buffSize := workerNum

	// STEP 3.
	//Create channels for sending data to workers and receiving data from workers
	invCh := make(chan Input, buffSize)
	recordCh := make(chan Record, buffSize)

	// STEP 4.
	// Create workers
	var wg sync.WaitGroup
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go writeCSVWorker(i, invCh, recordCh, &wg)
	}

	// STEP 5.
	// Send data to workers
	go func() {
		for i := 0; i < invoiceLen; i++ {
			invCh <- Input{i, invoices[i]}
		}
		// Close the channel when all invoices are sent
		close(invCh)
	}()

	// STEP 6.
	// Receive data from workers
	sortedRecords := make([]string, invoiceLen, invoiceLen)
	for i := 0; i < invoiceLen; i++ {
		record := <-recordCh
		sortedRecords[record.Idx] = record.Data
	}

	// STEP 7.
	// Wait for all workers to finish
	wg.Wait()

	// Be careful big-O of strings.Join
	return []byte(strings.Join(sortedRecords, "\n")), nil
}

// toString converts Invoice to csv record as string
func toString(invoice orm.Invoice) string {
	return invoice.ID + "," + invoice.StorekeeperID + "," + invoice.CreatedAt.Format("2006-01-02 15:04:05") + "," + invoice.UpdatedAt.Format("2006-01-02 15:04:05")
}

// writeCSVWorker is a worker that
// receives data from invCh and converts Invoice to csv record
// then sends data to recordCh
func writeCSVWorker(id int, inpCh <-chan Input, recordCh chan<- Record, wg *sync.WaitGroup) {
	for in := range inpCh {
		recordCh <- Record{in.Idx, toString(in.Invoice)}
	}
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func (inv InvoiceService) GetAllByPaginate(offset, limit int) ([]orm.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) GetByID(id string) (orm.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) Create(invoice orm.Invoice) (orm.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) Update(id string, invoice orm.Invoice) (orm.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (inv InvoiceService) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
