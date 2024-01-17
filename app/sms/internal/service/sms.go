package service

import (
	"akita/quantum_cat/app/sms/internal/biz"
	v1 "akita/quantum_cat/proto/sms/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type SmsService struct {
	v1.UnimplementedSmsServer
	uc  *biz.SmsUseCase
	log *log.Helper
}

// NewUserService new a User Service
func NewUserService(uc *biz.SmsUseCase, logger log.Logger) *SmsService {
	return &SmsService{
		UnimplementedSmsServer: v1.UnimplementedSmsServer{},
		uc:                     uc,
		log:                    log.NewHelper(log.With(logger, "module", "service/sms")),
	}
}
