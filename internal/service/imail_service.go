package service

type IMailService interface {
	Send(to []string, msg []byte) error
}
