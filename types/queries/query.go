package queries

import (
	"fmt"
	"net/url"
)

type (
	Action interface {
		String() string
	}

	Query struct {
		url.Values
	}
)

func NewQuery(token string) *Query {
	return &Query{
		url.Values{
			"apikey": []string{token},
		},
	}
}

func (q *Query) SetTarget(module string, action Action) *Query {
	q.Add("module", module)
	q.Add("action", action.String())
	return q
}

func (q *Query) Paginate(page, maxRecords uint64) *Query {
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("offset", fmt.Sprintf("%d", maxRecords))
	return q
}

func (q *Query) SetAddress(address string) *Query {
	q.Add("address", address)
	return q
}

func (q *Query) SetAddresses(addresses []string) *Query {
	for _, addr := range addresses {
		q.Add("address", addr)
	}
	return q
}

func (q *Query) SetContractAddress(address string) *Query {
	q.Add("contractaddress", address)
	return q
}

func (q *Query) SetTxHash(txHash string) *Query {
	q.Add("txhash", txHash)
	return q
}

func (q *Query) SetBlockNo(block uint64) *Query {
	q.Add("blockno", fmt.Sprintf("%d", block))
	return q
}

//TODO: support time.duration alongside string
func (q *Query) SetTimestamp(ts string) *Query {
	q.Add("timestamp", ts)
	return q
}

func (q *Query) SetBlockRange(begin, end uint64) *Query {
	q.Add("startblock", fmt.Sprintf("%d", begin))
	q.Add("endblock", fmt.Sprintf("%d", end))
	return q
}

func (q *Query) SortAsc() *Query {
	q.Add("sort", "asc")
	return q
}

func (q *Query) SortDesc() *Query {
	q.Add("sort", "desc")
	return q
}
