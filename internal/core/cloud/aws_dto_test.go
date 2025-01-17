package cloud

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AWSMetadataDtoTestSuite struct {
	suite.Suite
}

func TestAWSMetadataDtoTestSuite(t *testing.T) {
	suite.Run(t, new(AWSMetadataDtoTestSuite))
}

func (suite *AWSMetadataDtoTestSuite) TestNewAWSMetadataDto() {

	awsMetadata := &AWSMetadata{ //nolint
		AmiID: "some-ami",
		BlockDeviceMapping: map[string]string{
			"ami":  "sda",
			"ebs1": "/dev/sdb1",
			"ebs2": "/dev/sdb2",
			"root": "/dev/sda",
		},
		InstanceID:   "some-instance",
		InstanceType: "some-instance-type",
		Placement: Placement{
			AvailabilityZone: "some-availability-zone",
			Region:           "some-region",
		},
	}

	awsMetadata.IdentityCredentials.EC2.Info.AccountID = "some-account"
	awsMetadata.Network.Interfaces.Macs = make(map[string]MacEntry)
	macEntry := MacEntry{VpcID: "some-vpc-id"}
	awsMetadata.Network.Interfaces.Macs["eth1"] = macEntry

	awsMetadataDto := NewAWSMetadataDto(awsMetadata)

	expectedDto := &AWSMetadataDto{
		AccountID:        "some-account",
		AmiID:            "some-ami",
		AvailabilityZone: "some-availability-zone",
		DataDiskNumber:   2,
		InstanceID:       "some-instance",
		InstanceType:     "some-instance-type",
		Region:           "some-region",
		VpcID:            "some-vpc-id",
	}
	suite.Equal(expectedDto, awsMetadataDto)
}
