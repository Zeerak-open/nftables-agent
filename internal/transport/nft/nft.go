package nft

import (
	"github.com/google/nftables"
)

type NftService struct {
	conn *nftables.Conn
}

func New() *NftService {
	return &NftService{
		conn: &nftables.Conn{},
	}
}
