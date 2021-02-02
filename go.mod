module github.com/bpdunni/helm-umbrella

go 1.13

require (
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	helm.sh/helm/v3 v3.5.1
	rsc.io/letsencrypt v0.0.3 // indirect
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
