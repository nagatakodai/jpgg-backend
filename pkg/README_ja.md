作成日時: 2025-07-29 04:41:15 UTC - 本ファイルはAIによって翻訳されました。

# `/pkg`

外部アプリケーションから利用されても問題ないライブラリコードを置きます（例: `/pkg/mypubliclib`）。ここに置いたライブラリは他のプロジェクトからインポートされることを想定しているため、配置する際は慎重に考えてください。プライベートなパッケージをインポート不可にしたい場合は、Go が強制してくれる `internal` ディレクトリを使う方が確実です。それでも `/pkg` ディレクトリは、そこにあるコードが他者の利用を想定していることを明示するという点で有用です。Travis Jeffery による [I'll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) では、`pkg` と `internal` ディレクトリの概要や使い分けについて説明されています。

また、ルートディレクトリに多くの非 Go コンポーネントやディレクトリがある場合でも、Go コードを1か所にまとめることで各種 Go ツールを実行しやすくなるという利点があります（以下のトークで触れられています: [Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018、[GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)、[GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

もっともこのパターンが常に受け入れられるわけではなく、採用している人気リポジトリもあれば、していないものも多数あります。採用するかどうかはあなた次第ですが、良いパターンかどうかはともかく、こうしておけば意図が伝わりやすいでしょう。新しい Go 開発者には少し混乱を与えるかもしれませんが、理解は簡単で、それがこのプロジェクトレイアウトの目的のひとつです。

アプリケーションが非常に小さく、階層を増やすメリットが少ない場合は無理に使う必要はありません。ただ、規模が大きくなりルートディレクトリが煩雑になってきたら検討してみてください（特に非 Go コンポーネントが多い場合）。

`pkg` ディレクトリの由来: 以前の Go ソースコードはパッケージを `pkg` に配置しており、コミュニティの様々なプロジェクトもこのパターンをコピーし始めました（詳しくは [こちら](https://twitter.com/bradfitz/status/1039512487538970624) の Brad Fitzpatrick のツイートを参照）。

例:

* https://github.com/containerd/containerd/tree/main/pkg
* https://github.com/slimtoolkit/slim/tree/master/pkg
* https://github.com/telepresenceio/telepresence/tree/release/v2/pkg
* https://github.com/jaegertracing/jaeger/tree/master/pkg
* https://github.com/istio/istio/tree/master/pkg
* https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
* https://github.com/google/gvisor/tree/master/pkg
* https://github.com/google/syzkaller/tree/master/pkg
* https://github.com/perkeep/perkeep/tree/master/pkg
* https://github.com/heptio/ark/tree/master/pkg
* https://github.com/argoproj/argo-workflows/tree/master/pkg
* https://github.com/argoproj/argo-cd/tree/master/pkg
* https://github.com/heptio/sonobuoy/tree/master/pkg
* https://github.com/helm/helm/tree/master/pkg
* https://github.com/k3s-io/k3s/tree/master/pkg
* https://github.com/kubernetes/kubernetes/tree/master/pkg
* https://github.com/kubernetes/kops/tree/master/pkg
* https://github.com/moby/moby/tree/master/pkg
* https://github.com/grafana/grafana/tree/master/pkg
* https://github.com/influxdata/influxdb/tree/master/pkg
* https://github.com/cockroachdb/cockroach/tree/master/pkg
* https://github.com/derekparker/delve/tree/master/pkg
* https://github.com/etcd-io/etcd/tree/master/pkg
* https://github.com/oklog/oklog/tree/master/pkg
* https://github.com/flynn/flynn/tree/master/pkg
* https://github.com/jesseduffield/lazygit/tree/master/pkg
* https://github.com/gopasspw/gopass/tree/master/pkg
* https://github.com/sosedoff/pgweb/tree/master/pkg
* https://github.com/GoogleContainerTools/skaffold/tree/master/pkg
* https://github.com/knative/serving/tree/master/pkg
* https://github.com/grafana/loki/tree/master/pkg
* https://github.com/bloomberg/goldpinger/tree/master/pkg
* https://github.com/Ne0nd0g/merlin/tree/master/pkg
* https://github.com/jenkins-x/jx/tree/master/pkg
* https://github.com/DataDog/datadog-agent/tree/master/pkg
* https://github.com/dapr/dapr/tree/master/pkg
* https://github.com/cortexproject/cortex/tree/master/pkg
* https://github.com/dexidp/dex/tree/master/pkg
* https://github.com/pusher/oauth2_proxy/tree/master/pkg
* https://github.com/pdfcpu/pdfcpu/tree/master/pkg
* https://github.com/weaveworks/kured/tree/master/pkg
* https://github.com/weaveworks/footloose/tree/master/pkg
* https://github.com/weaveworks/ignite/tree/master/pkg
* https://github.com/tmrts/boilr/tree/master/pkg
* https://github.com/kata-containers/runtime/tree/master/pkg
* https://github.com/okteto/okteto/tree/master/pkg
* https://github.com/solo-io/squash/tree/master/pkg
* https://github.com/google/exposure-notifications-server/tree/main/pkg
* https://github.com/spiffe/spire/tree/main/pkg
* https://github.com/rook/rook/tree/master/pkg
* https://github.com/buildpacks/pack/tree/main/pkg
* https://github.com/cilium/cilium/tree/main/pkg
* https://github.com/containernetworking/cni/tree/main/pkg
* https://github.com/crossplane/crossplane/tree/master/pkg
* https://github.com/dragonflyoss/Dragonfly2/tree/main/pkg
* https://github.com/kubeedge/kubeedge/tree/master/pkg
* https://github.com/kubevela/kubevela/tree/master/pkg
* https://github.com/kubevirt/kubevirt/tree/main/pkg
* https://github.com/kyverno/kyverno/tree/main/pkg
* https://github.com/thanos-io/thanos/tree/main/pkg
* https://github.com/cri-o/cri-o/tree/main/pkg
* https://github.com/fluxcd/flux2/tree/main/pkg
* https://github.com/kedacore/keda/tree/main/pkg
* https://github.com/linkerd/linkerd2/tree/main/pkg
* https://github.com/opencost/opencost/tree/develop/pkg
* https://github.com/antrea-io/antrea/tree/main/pkg
* https://github.com/karmada-io/karmada/tree/master/pkg
* https://github.com/kuberhealthy/kuberhealthy/tree/master/pkg
* https://github.com/submariner-io/submariner/tree/devel/pkg
* https://github.com/trickstercache/trickster/tree/main/pkg
* https://github.com/tellerops/teller/tree/master/pkg
* https://github.com/OpenFunction/OpenFunction/tree/main/pkg
* https://github.com/external-secrets/external-secrets/tree/main/pkg
* https://github.com/ko-build/ko/tree/main/pkg
* https://github.com/lima-vm/lima/tree/master/pkg
* https://github.com/clastix/capsule/tree/master/pkg
* https://github.com/carvel-dev/ytt/tree/develop/pkg
* https://github.com/clusternet/clusternet/tree/main/pkg
* https://github.com/fluid-cloudnative/fluid/tree/master/pkg
* https://github.com/inspektor-gadget/inspektor-gadget/tree/main/pkg
* https://github.com/sustainable-computing-io/kepler/tree/main/pkg
* https://github.com/GoogleContainerTools/kpt/tree/main/pkg
* https://github.com/guacsec/guac/tree/main/pkg
* https://github.com/kubeovn/kube-ovn/tree/master/pkg
* https://github.com/kube-vip/kube-vip/tree/main/pkg
* https://github.com/kubescape/kubescape/tree/master/pkg
* https://github.com/kudobuilder/kudo/tree/main/pkg
* https://github.com/kumahq/kuma/tree/master/pkg
* https://github.com/kubereboot/kured/tree/main/pkg
* https://github.com/nocalhost/nocalhost/tree/main/pkg
* https://github.com/openelb/openelb/tree/master/pkg
* https://github.com/openfga/openfga/tree/main/pkg
* https://github.com/openyurtio/openyurt/tree/master/pkg
* https://github.com/getporter/porter/tree/main/pkg
* https://github.com/sealerio/sealer/tree/main/pkg
* https://github.com/werf/werf/tree/main/pkg
