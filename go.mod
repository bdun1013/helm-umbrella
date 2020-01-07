module github.com/bpdunni/helm-umbrella

go 1.13

require (
	9fans.net/go v0.0.2
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	helm.sh/helm v2.16.1+incompatible
	helm.sh/helm/v3 v3.0.2
	k8s.io/helm v2.16.1+incompatible // indirect
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
