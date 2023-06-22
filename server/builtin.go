package server

import (
	"github.com/mauricio-b-r/polygon-edge/consensus"
	consensusDev "github.com/mauricio-b-r/polygon-edge/consensus/dev"
	consensusDummy "github.com/mauricio-b-r/polygon-edge/consensus/dummy"
	consensusIBFT "github.com/mauricio-b-r/polygon-edge/consensus/ibft"
	"github.com/mauricio-b-r/polygon-edge/secrets"
	"github.com/mauricio-b-r/polygon-edge/secrets/awsssm"
	"github.com/mauricio-b-r/polygon-edge/secrets/gcpssm"
	"github.com/mauricio-b-r/polygon-edge/secrets/hashicorpvault"
	"github.com/mauricio-b-r/polygon-edge/secrets/local"
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
