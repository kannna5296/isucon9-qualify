package scenario

import (
	"sync"

	"github.com/isucon/isucon9-qualify/bench/asset"
	"github.com/isucon/isucon9-qualify/bench/fails"
)

func Initialize() *fails.Critical {
	critical := fails.NewCritical()

	_, err := initialize("", "")
	if err != nil {
		critical.Add(err)
	}

	return critical
}

func Verify() *fails.Critical {
	var wg sync.WaitGroup

	critical := fails.NewCritical()

	user1, user2 := asset.GetRandomUser(), asset.GetRandomUser()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := sellAndBuy(user1, user2)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1, user2 := asset.GetRandomUser(), asset.GetRandomUser()
		err := bumpAndNewItems(user1, user2)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1 := asset.GetRandomUser()
		err := newCategoryItems(user1)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1 := asset.GetRandomUser()
		err := itemEdit(user1)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1 := asset.GetRandomUser()
		err := transactionEvidence(user1)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1 := asset.GetRandomUser()
		user2 := asset.GetRandomUser()
		err := userItemsAndItem(user1, user2)
		if err != nil {
			critical.Add(err)
		}
	}()

	user3 := asset.GetRandomUser()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := irregularLoginWrongPassword(user3)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := irregularSellAndBuy(user2, user1, user3)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Wait()

	return critical
}

func Validation(critical *fails.Critical) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		check(critical)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		load(critical)
	}()

	wg.Wait()
}

func check(critical *fails.Critical) {
	var wg sync.WaitGroup

	user1, user2, user3 := asset.GetRandomUser(), asset.GetRandomUser(), asset.GetRandomUser()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := irregularLoginWrongPassword(user3)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := irregularSellAndBuy(user2, user1, user3)
		if err != nil {
			critical.Add(err)
		}
	}()

	wg.Wait()
}

func load(critical *fails.Critical) {}

func FinalCheck(critical *fails.Critical) {}