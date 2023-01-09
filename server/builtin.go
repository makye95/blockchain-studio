package server

import (
	"github.com/gitshock-labs/go-client/consensus"
	consensusDev "github.com/gitshock-labs/go-client/consensus/dev"
	consensusDummy "github.com/gitshock-labs/go-client/consensus/dummy"
	consensusIBFT "github.com/gitshock-labs/go-client/consensus/ibft"
	"github.com/gitshock-labs/go-client/secrets"
	"github.com/gitshock-labs/go-client/secrets/awsssm"
	"github.com/gitshock-labs/go-client/secrets/gcpssm"
	"github.com/gitshock-labs/go-client/secrets/hashicorpvault"
	"github.com/gitshock-labs/go-client/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
