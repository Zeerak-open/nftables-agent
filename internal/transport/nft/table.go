package nft

import (
	"github.com/google/nftables"
)

func (ns *NftService) GetTables() ([]*nftables.Table, error) {
	return ns.conn.ListTables()
}

func (ns *NftService) GetTable(name string) (*nftables.Table, error) {
	tables, err := ns.GetTables()
	if err != nil {
		return nil, err
	}

	for _, t := range tables {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, nil
}
