package jobs

import (
	"context"
	"github.com/filedrive-team/filplus-info/models"
	"github.com/filedrive-team/filplus-info/rpc"
	"github.com/filedrive-team/filplus-info/utils"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Syncer struct {
}

func NewSyncer() *Syncer {
	s := &Syncer{}
	return s
}

func (s *Syncer) Run(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Minute)
	s.syncNotaryAllowance(ctx)
	s.syncClientAllowance(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.syncNotaryAllowance(ctx)
			s.syncClientAllowance(ctx)
		}
	}
}

func (s *Syncer) syncNotaryAllowance(ctx context.Context) {
	epoch, err := models.GetLastEpochFromNotaryAllowance()
	if err != nil {
		logger.Errorf("GetLastEpochFromNotaryAllowance failed, %v", err)
		return
	}
	currentEpoch := utils.GetEpochByTime(time.Now())
	data, err := rpc.NotaryAllowanceList(epoch, currentEpoch)
	if err != nil {
		logger.Errorf("NotaryAllowanceList failed, %v", err)
		return
	}
	for _, item := range data {
		err = models.UpsertNotaryAllowance(item)
		if err != nil {
			logger.Errorf("UpsertNotaryAllowance failed, %v", err)
			return
		}
	}
}

func (s *Syncer) syncClientAllowance(ctx context.Context) {
	blockTime, err := models.GetLastBlockTimeFromNotaryAllowance()
	if err != nil {
		logger.Errorf("GetLastBlockTimeFromNotaryAllowance failed, %v", err)
		return
	}
	now := time.Now().Unix()
	data, err := rpc.ClientAllowanceList(blockTime, now)
	if err != nil {
		logger.Errorf("ClientAllowanceList failed, %v", err)
		return
	}
	for _, item := range data {
		err = models.UpsertClientAllowance(item)
		if err != nil {
			logger.Errorf("UpsertClientAllowance failed, %v", err)
			return
		}
	}
}
