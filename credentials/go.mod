module github.com/aws/aws-sdk-go-v2/credentials

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.23.1
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.5
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.3
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.20.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.4
	github.com/aws/smithy-go v1.17.0
	github.com/google/go-cmp v0.5.8
)

require (
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.4 // indirect
)

replace github.com/aws/aws-sdk-go-v2 => ../

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../service/sso/

replace github.com/aws/aws-sdk-go-v2/service/ssooidc => ../service/ssooidc/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../service/sts/
