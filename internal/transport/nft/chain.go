package nft

import (
	"github.com/google/nftables"
	"github.com/thoas/go-funk"
)

func (ns *NftService) GetChains(table *nftables.Table) ([]*nftables.Chain, error) {
	chains, err := ns.conn.ListChains()
	if err != nil {
		return nil, err
	}

	return funk.Filter(chains, func(c *nftables.Chain) bool {
		return c.Table.Name == table.Name
	}).([]*nftables.Chain), nil
}

func (ns *NftService) GetChain(table *nftables.Table, name string) (*nftables.Chain, error) {
	chains, err := ns.GetChains(table)
	if err != nil {
		return nil, err
	}

	for _, chain := range chains {
		if chain.Name == name {
			return chain, nil
		}
	}

	return nil, nil
}
