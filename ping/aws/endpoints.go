package aws

var AWSEndpoints = map[string]string{
	"us-east-1":      "https://ec2.us-east-1.amazonaws.com/ping",
	"us-east-2":      "https://ec2.us-east-2.amazonaws.com/ping",
	"us-west-1":      "https://ec2.us-west-1.amazonaws.com/ping",
	"us-west-2":      "https://ec2.us-west-2.amazonaws.com/ping",
	"ca-central-1":   "https://ec2.ca-central-1.amazonaws.com/ping",
	"eu-central-1":   "https://ec2.eu-central-1.amazonaws.com/ping",
	"eu-west-1":      "https://ec2.eu-west-1.amazonaws.com/ping",
	"eu-west-2":      "https://ec2.eu-west-2.amazonaws.com/ping",
	"eu-west-3":      "https://ec2.eu-west-3.amazonaws.com/ping",
	"eu-north-1":     "https://ec2.eu-north-1.amazonaws.com/ping",
	"ap-east-1":      "https://ec2.ap-east-1.amazonaws.com/ping",
	"ap-northeast-1": "https://ec2.ap-northeast-1.amazonaws.com/ping",
	"ap-northeast-2": "https://ec2.ap-northeast-2.amazonaws.com/ping",
	"ap-northeast-3": "https://ec2.ap-northeast-3.amazonaws.com/ping",
	"ap-southeast-1": "https://ec2.ap-southeast-1.amazonaws.com/ping",
	"ap-southeast-2": "https://ec2.ap-southeast-2.amazonaws.com/ping",
	"ap-south-1":     "https://ec2.ap-south-1.amazonaws.com/ping",
	"me-south-1":     "https://ec2.me-south-1.amazonaws.com/ping",
	"sa-east-1":      "https://ec2.sa-east-1.amazonaws.com/ping",
}

var AWSRegions = map[string]string{
	"us-east-1":      "US East (N. Virginia)",
	"us-east-2":      "US East (Ohio)",
	"us-west-1":      "US West (N. California)",
	"us-west-2":      "US West (Oregon)",
	"ca-central-1":   "Canada (Central)",
	"eu-central-1":   "EU (Frankfurt)",
	"eu-west-1":      "EU (Ireland)",
	"eu-west-2":      "EU (London)",
	"eu-west-3":      "EU (Paris)",
	"eu-north-1":     "EU (Stockholm)",
	"ap-east-1":      "Asia Pacific (Hong Kong)",
	"ap-northeast-1": "Asia Pacific (Tokyo)",
	"ap-northeast-2": "Asia Pacific (Seoul)",
	"ap-northeast-3": "Asia Pacific (Osaka-Local)",
	"ap-southeast-1": "Asia Pacific (Singapore)",
	"ap-southeast-2": "Asia Pacific (Sydney)",
	"ap-south-1":     "Asia Pacific (Mumbai)",
	"me-south-1":     "Middle East (Bahrain)",
	"sa-east-1":      "South America (São Paulo)",
}
