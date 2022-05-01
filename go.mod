module github.com/inspirit941/convert-jenkinsfile

go 1.17

require (
	github.com/alecthomas/participle v0.7.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	k8s.io/api v0.23.6
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.23.6 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

//replace k8s.io/api => k8s.io/api v0.0.0-20181128191700-6db15a15d2d3
//
//replace k8s.io/metrics => k8s.io/metrics v0.0.0-20181128195641-3954d62a524d
//
//replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190122181752-bebe27e40fb7
//
//replace k8s.io/client-go => k8s.io/client-go v2.0.0-alpha.0.0.20190115164855-701b91367003+incompatible
//
//replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181128195303-1f84094d7e8e
