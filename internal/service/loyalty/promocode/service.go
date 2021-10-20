package promocode

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
	"log"
)

type PromocodeService interface {
	Describe(promocodeId uint64) (*loyalty.Promocode, error)
	List(cursor uint64, limit uint64) ([]loyalty.Promocode, error)
	Create(loyalty.Promocode) (uint64, error)
	Update(promocodeId uint64, promocode loyalty.Promocode) error
	Remove(promocodeId uint64) (bool, error)
}

type DummyPromocodeService struct{}

func (d *DummyPromocodeService) Describe(promocodeId uint64) (*loyalty.Promocode, error) {
	if uint64(len(loyalty.AllPromocodes)) < promocodeId {
		return nil, fmt.Errorf("invalid promocode id")
	}
	return &loyalty.AllPromocodes[int(promocodeId)], nil
}

func (d *DummyPromocodeService) List(cursor uint64, limit uint64) ([]loyalty.Promocode, error) {
	if cursor+1 > uint64(len(loyalty.AllPromocodes)) {
		return nil, fmt.Errorf("invalid cursor value")
	}
	if cursor+limit > uint64(len(loyalty.AllPromocodes))+1 {
		return loyalty.AllPromocodes[cursor:], nil
	}
	return loyalty.AllPromocodes[cursor : cursor+limit], nil
}

func (d DummyPromocodeService) Create(promocode loyalty.Promocode) (uint64, error) {
	loyalty.AllPromocodes = append(loyalty.AllPromocodes, promocode)
	log.Printf("Promocode %s was created", promocode.Promocode)
	return uint64(len(loyalty.AllPromocodes) - 1), nil
}

func (d DummyPromocodeService) Update(promocodeId uint64, promocode loyalty.Promocode) error {
	if uint64(len(loyalty.AllPromocodes)) < promocodeId || len(loyalty.AllPromocodes) == 0 {
		return fmt.Errorf("invalid promocode id")
	}

	loyalty.AllPromocodes[promocodeId] = promocode
	log.Printf("Promocode with id %d was updated", promocodeId)
	return nil
}

func (d DummyPromocodeService) Remove(promocodeId uint64) (bool, error) {
	if uint64(len(loyalty.AllPromocodes)) < promocodeId || len(loyalty.AllPromocodes) == 0 {

		return false, fmt.Errorf("invalid promocode id")
	}
	loyalty.AllPromocodes = append(loyalty.AllPromocodes[:promocodeId], loyalty.AllPromocodes[promocodeId+1:]...)
	log.Printf("Promocode with id %d was removed", promocodeId)
	return true, nil
}

func NewDummyPromocodeService() *DummyPromocodeService {
	return &DummyPromocodeService{}
}
