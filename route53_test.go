package main

import (
	"net"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/stretchr/testify/assert"
)

func TestConvertR53RecordsToHosts(t *testing.T) {
	input := []*route53.ResourceRecordSet{
		&route53.ResourceRecordSet{
			Name: aws.String("test1.test.com"),
			Type: aws.String("A"),
			TTL:  aws.Int64(300),
			ResourceRecords: []*route53.ResourceRecord{
				&route53.ResourceRecord{Value: aws.String("1.2.3.4")},
			},
		},
		&route53.ResourceRecordSet{
			Name: aws.String("test2.test.com"),
			Type: aws.String("CNAME"),
			TTL:  aws.Int64(300),
			ResourceRecords: []*route53.ResourceRecord{
				&route53.ResourceRecord{Value: aws.String("test1.test.com")},
			},
		},
		&route53.ResourceRecordSet{
			Name: aws.String("test3.test.com"),
			Type: aws.String("A"),
			TTL:  aws.Int64(300),
			ResourceRecords: []*route53.ResourceRecord{
				&route53.ResourceRecord{Value: aws.String("1.2.3.4")},
				&route53.ResourceRecord{Value: aws.String("1.2.3.5")},
			},
		},
		&route53.ResourceRecordSet{
			Name: aws.String("test4.test.com"),
			Type: aws.String("A"),
			TTL:  aws.Int64(300),
			ResourceRecords: []*route53.ResourceRecord{
				&route53.ResourceRecord{Value: aws.String("abc")},
			},
		},
	}

	expected := []hostEntry{
		{hostname: "test1.test.com", ip: net.ParseIP("1.2.3.4")},
	}

	output := convertR53RecordsToHosts(input)
	assert.Equal(t, output, expected)

}
