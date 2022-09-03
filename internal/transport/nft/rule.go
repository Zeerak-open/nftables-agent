package nft

import (
	"github.com/google/nftables"
)

func (ns *NftService) GetRules(table *nftables.Table, chain *nftables.Chain) ([]*nftables.Rule, error) {
	return ns.conn.GetRule(table, chain)
}

func (ns *NftService) GetRule(table *nftables.Table, chain *nftables.Chain, id uint64) (*nftables.Rule, error) {
	rules, err := ns.GetRules(table, chain)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if rule.Position == id {
			return rule, nil
		}
	}

	return nil, nil
}
