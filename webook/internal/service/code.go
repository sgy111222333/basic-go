package service

import (
	"context"
	"fmt"
	"github.com/sgy111222333/basic-go/webook/internal/repository"
	"github.com/sgy111222333/basic-go/webook/internal/service/sms"
	"math/rand"
)

var ErrCodeSendTooMany = repository.ErrCodeVerifyTooMany

type CodeService struct {
	repo *repository.CodeRepository
	sms  sms.Service
}

func NewCodeService(repo *repository.CodeRepository, smsSvc sms.Service) *CodeService {
	return &CodeService{
		repo: repo,
		sms:  smsSvc,
	}
}

func (svc *CodeService) Send(ctx context.Context, biz string, phone string) error {
	code := svc.generate()
	err := svc.repo.Set(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	const codeTplId = "1877556"
	return svc.sms.Send(ctx, codeTplId, []string{code}, phone)
}

func (svc *CodeService) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	ok, err := svc.repo.Verify(ctx, biz, phone, inputCode)
	if err == repository.ErrCodeVerifyTooMany {
		// 相当于对外面屏蔽了验证次数过多的错误, 就告诉调用者不对
		return false, nil
	}
	return ok, err
}

func (svc *CodeService) generate() string {
	// 0~999999
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code) // 不足六位数的前面补零
}
