package alirds

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/util"
	"github.com/teamssix/cf/pkg/util/cmdutil"
)

func RDSClient(region string) *rds.Client {
	aliconfig := cmdutil.GetAliCredential()
	if aliconfig.AccessKeyId == "" {
		log.Warnln("需要先配置访问凭证 (Access Key need to be configured first)")
		os.Exit(0)
		return nil
	} else {
		config := sdk.NewConfig()
		if aliconfig.STSToken == "" {
			credential := credentials.NewAccessKeyCredential(aliconfig.AccessKeyId, aliconfig.AccessKeySecret)
			client, err := rds.NewClientWithOptions(region, config, credential)
			util.HandleErr(err)
			if err == nil {
				log.Traceln("RDS Client 连接成功 (RDS Client connection successful)")
			}
			return client
		} else {
			credential := credentials.NewStsTokenCredential(aliconfig.AccessKeyId, aliconfig.AccessKeySecret, aliconfig.STSToken)
			client, err := rds.NewClientWithOptions(region, config, credential)
			util.HandleErr(err)
			if err == nil {
				log.Traceln("RDS Client 连接成功 (RDS Client connection successful)")
			}
			return client
		}
	}
}

func GetRDSRegions() []rds.RDSRegion {
	client := RDSClient("cn-hangzhou")
	request := rds.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := client.DescribeRegions(request)
	util.HandleErr(err)
	return response.Regions.RDSRegion
}
