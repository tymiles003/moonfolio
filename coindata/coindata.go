package coindata

import (
	"sync"

	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

// CoinData contains information
type CoinData struct {
	Name         string
	Symbol       string
	Rank         int
	PriceUSD     float64
	HistoHour    []*cryptocompare.Histo
	HistoDay     []*cryptocompare.Histo
	transactions []*transactions.Transaction
	lock         sync.Mutex
}

// SetName to set the name of the coin
func (c *CoinData) SetName(name string) {
	c.Name = name
}

// SetSymbol to set the symbol for the coin
func (c *CoinData) SetSymbol(symbol string) {
	c.Symbol = symbol
}

// SetRank to set the current rank of the coin
func (c *CoinData) SetRank(rank int) {
	c.Rank = rank
}

// SetPriceUSD set the price in USD
func (c *CoinData) SetPriceUSD(price float64) {
	c.PriceUSD = price
}

// SetTransactions sets a slice of transactions for a coin
func (c *CoinData) SetTransactions(transactions []*transactions.Transaction) {
	c.lock.Lock()
	c.transactions = transactions
	c.lock.Unlock()
}

// AddTransaction adds an user transaction to the coin
func (c *CoinData) AddTransaction(transaction *transactions.Transaction) {
	c.lock.Lock()

	if c.transactions == nil {
		c.transactions = make([]*transactions.Transaction, 0)
	}

	c.transactions = append(c.transactions, transaction)
	c.lock.Unlock()
}

// GetTransactions returns all user transactions for a coin
func (c *CoinData) GetTransactions() []*transactions.Transaction {
	if c.transactions == nil {
		return make([]*transactions.Transaction, 0)
	}

	return c.transactions
}

// GetCoinAmount returns the total amount of coins the user is hodling
func (c *CoinData) GetCoinAmount() float64 {
	c.lock.Lock()

	amount := 0.0
	for _, t := range c.GetTransactions() {
		if t.Type == transactions.TransactionBuy {
			amount += t.CoinAmount
		} else if t.Type == transactions.TransactionSell {
			amount -= t.CoinAmount
		}
	}

	c.lock.Unlock()
	return amount
}

// GetBalance returns the Balance in USD for this coin
func (c *CoinData) GetBalance() float64 {
	coinAmount := c.GetCoinAmount()
	return coinAmount * c.PriceUSD
}

func (c *CoinData) GetChange1H() string {
	c.lock.Lock()
	c.lock.Unlock()

	return ""
}

func (c *CoinData) GetChange1D() string {
	c.lock.Lock()
	c.lock.Unlock()

	return ""
}

func (c *CoinData) GetChange7D() string {
	c.lock.Lock()
	c.lock.Unlock()

	return ""
}
