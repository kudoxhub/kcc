// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main KCC network.
var MainnetBootnodes = []string{
	"enode://d16cf20214ae987c0b200252df331633a680866e75fec511ae25573c06585a8ce2d7379526c9ca186501fbcf1af92443a178f7a800049947a178d5a7ba130566@13.230.226.55:30303",  // kcc-mainnet-node-boot-01
	"enode://0a3e227c2babb0f63e7fbb942f6c1abf00e47dc2e8b970b8c10b6c99dba7ac5eb223a11259eae50d195f88d27d3779a9b7b97618ca2afa5932c7ffdfabc91c31@35.74.148.173:30303",  //kcc-mainnet-node-boot-02
	"enode://0be483c73c2ce90374da59cf0126e364ab1d2a8c96b02e6f2aaadc4d4ee15e50c3962034435fcc9490d0acd3050cc119f50860bc53a5042e02fe9d0586609f44@18.142.129.189:30303", // d-01
	"enode://067ad547a7663ab378d5fdc8d0a6fef5e9e687e07223c38146525d83b7f72d69279a52b4deb1b362f497406c47547be1f5625bbc7366d3f5d5019b287dbe1330@52.47.152.245:30303",  // d-02
	"enode://0ef8c8675a417ec2456e1e2b54f6d382a8c9925062367ec9a381408dd252c9c157e8bbf3561720f8e93814aafd2df7ebce9cdc4f1a51c9d55c56fc5bb6e9ebcd@13.43.80.116:30303",   // d-03
	"enode://d6ff3a55f335d7b82d3fd18e2e2df77c91f635f256fc30080ed24c723bbd5a1614f9adb930b82247b473fea8f640596453e1c1b0041b92e670fa6dc50d4c7498@35.165.249.173:30303", // d-04
	"enode://0cbc9ec289e5f6a5d98e375110b476e141527053b28634ae6d83e9372fde9de94963f12f2c94ad74575dc161241dc48bf074c2d466ee0fda7cd935b257923174@35.175.52.53:30303",   // d-05
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://b2aa772d077311e827228e65cf0cd8fccf8e0a2d7d1933debde45737124140253469e47faa46fb2cc7014eaaa05828c599855cc4d7803fa76a336f44b6e67c05@54.168.253.129:30303", // kcc-testnet-node-boot-01
	"enode://223d6e031bc7ae5911dc377451641ad5807bfdcf46f0809c642f044751dfbf89de52b36ac1c379048f30a566de8bc4908d85c244b9f541599a419be97439bab1@3.216.120.238:30303",  // kcc-testnet-node-boot-02
	"enode://a12eb76fd1c7c2ae9a237b86e9357762a15a8e46c66c6f9b668acb349e937e27fd676275fe73a9e08ce1b667da59fb3ae6f8c29016d55bef17a415d65f767924@54.254.194.123:30303", // kcc-testnet-node-boot-03
	"enode://68ed7bf65b937eefc5bc9d8c08d83d2c8e7d3a5f628f0f5cc53b055d66793afcd12829854c8d9dd85754eb6badb8c05bbe2e25bdc1a004b8122571184e01540e@13.208.138.241:30303", // kcc-testnet-node-boot-04
	"enode://71bea53f03654e1c9bfdf9488c88087492a207c30129e93b5b5c89d79fcd4cbd432eee55ed12ab2cacb7b84aa24bd03c1b76ca9d95164f33b1088d64756e17a7@15.152.3.151:53370",   // kcc-testnet-node-sync-01
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
