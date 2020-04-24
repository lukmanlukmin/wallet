package transaction

import (
	"fmt"

	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	httpRequestEntity "github.com/lukmanlukmin/wallet/entity/http/request"
	repository "github.com/lukmanlukmin/wallet/repository/database"
)

type TransactionService struct {
	balanceRepository        repository.UserBalanceRepositoryInterface
	balanceHistoryRepository repository.UserBalanceHistoryRepositoryInterface
}

func TransactionServiceHandler() *TransactionService {
	return &TransactionService{
		balanceRepository:        repository.UserBalanceRepositoryHandler(),
		balanceHistoryRepository: repository.UserBalanceHistoryRepositoryHandler(),
	}
}

type TransactionServiceInterface interface {
	TopUp(obligorId int, ip string, agent string, payload httpRequestEntity.TopUpRequest) error
	Transfer(obligorId int, ip string, agent string, payload httpRequestEntity.TransferRequest) error
}

func (service *TransactionService) TopUp(benificiaryID int, ip string, agent string, payload httpRequestEntity.TopUpRequest) error {
	balanceData := &dbEntity.UserBalances{}
	err := service.balanceRepository.GetCurrentBalance(benificiaryID, balanceData)
	curBalance := 0
	if err == nil {
		curBalance = int(balanceData.BalanceAchieves)
	}
	newBalanceData := &dbEntity.UserBalances{}
	newBalanceData.Balance = uint(curBalance)
	newBalanceData.BalanceAchieves = uint(payload.Amount) + uint(curBalance)
	newBalanceData.UserID = uint(benificiaryID)
	errr := service.balanceRepository.InsertBalanceData(newBalanceData)
	if errr != nil {
		return errr
	}
	balanceDataHistory := &dbEntity.UserBalanceHistories{}
	balanceDataHistory.UserBalanceID = uint(benificiaryID)
	balanceDataHistory.BalanceBefore = uint(curBalance)
	balanceDataHistory.BalanceAfter = uint(newBalanceData.BalanceAchieves)
	balanceDataHistory.Activity = "TopUp"
	balanceDataHistory.TypeActivity = "debit"
	balanceDataHistory.IP = ip
	balanceDataHistory.UserAgent = agent
	errrr := service.balanceHistoryRepository.InsertBalanceHistoryData(balanceDataHistory)
	return errrr
}

func (service *TransactionService) Transfer(obligorId int, ip string, agent string, payload httpRequestEntity.TransferRequest) error {
	balanceData := &dbEntity.UserBalances{}
	err1 := service.balanceRepository.GetCurrentBalance(obligorId, balanceData)
	curBalance := 0
	if err1 == nil {
		curBalance = int(balanceData.BalanceAchieves)
	}
	if curBalance < payload.Amount {
		return fmt.Errorf("Insuficent Balance")
	}
	newBalanceData := &dbEntity.UserBalances{}
	newBalanceData.Balance = uint(curBalance)
	newBalanceData.BalanceAchieves = uint(payload.Amount) + uint(curBalance)
	newBalanceData.UserID = uint(obligorId)
	err2 := service.balanceRepository.InsertBalanceData(newBalanceData)
	if err2 != nil {
		return err2
	}
	balanceDataHistory := &dbEntity.UserBalanceHistories{}
	balanceDataHistory.UserBalanceID = uint(obligorId)
	balanceDataHistory.BalanceBefore = uint(curBalance)
	balanceDataHistory.BalanceAfter = uint(newBalanceData.BalanceAchieves)
	balanceDataHistory.Activity = "Transfer"
	balanceDataHistory.TypeActivity = "debit"
	balanceDataHistory.IP = ip
	balanceDataHistory.UserAgent = agent
	errrr := service.balanceHistoryRepository.InsertBalanceHistoryData(balanceDataHistory)
	return errrr
}
