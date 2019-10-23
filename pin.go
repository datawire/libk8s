// Code generated by [go run ./cmd/generate-libk8s "refs/heads/release-1.15"] (for k8s.io/cli-runtime commit cb4016b5ceb7a3ca63779ce26a8328c1a2c7208c). DO NOT EDIT.

package libk8s

import (
	_ "cloud.google.com/go/compute/metadata"
	_ "github.com/Azure/go-autorest/autorest"
	_ "github.com/Azure/go-autorest/autorest/adal"
	_ "github.com/Azure/go-autorest/autorest/azure"
	_ "github.com/Azure/go-autorest/autorest/date"
	_ "github.com/Azure/go-autorest/logger"
	_ "github.com/Azure/go-autorest/version"
	_ "github.com/PuerkitoBio/purell"
	_ "github.com/PuerkitoBio/urlesc"
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/docker/spdystream"
	_ "github.com/docker/spdystream/spdy"
	_ "github.com/emicklei/go-restful"
	_ "github.com/emicklei/go-restful/log"
	_ "github.com/evanphx/json-patch"
	_ "github.com/ghodss/yaml"
	_ "github.com/go-openapi/jsonpointer"
	_ "github.com/go-openapi/jsonreference"
	_ "github.com/go-openapi/spec"
	_ "github.com/go-openapi/swag"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/sortkeys"
	_ "github.com/golang/groupcache/lru"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/google/btree"
	_ "github.com/google/go-cmp/cmp"
	_ "github.com/google/gofuzz"
	_ "github.com/googleapis/gnostic/OpenAPIv2"
	_ "github.com/googleapis/gnostic/compiler"
	_ "github.com/googleapis/gnostic/extensions"
	_ "github.com/gophercloud/gophercloud"
	_ "github.com/gophercloud/gophercloud/openstack"
	_ "github.com/gophercloud/gophercloud/openstack/identity/v2/tenants"
	_ "github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	_ "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	_ "github.com/gophercloud/gophercloud/openstack/utils"
	_ "github.com/gophercloud/gophercloud/pagination"
	_ "github.com/gregjones/httpcache"
	_ "github.com/gregjones/httpcache/diskcache"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/hashicorp/golang-lru/simplelru"
	_ "github.com/imdario/mergo"
	_ "github.com/json-iterator/go"
	_ "github.com/mailru/easyjson/buffer"
	_ "github.com/mailru/easyjson/jlexer"
	_ "github.com/mailru/easyjson/jwriter"
	_ "github.com/modern-go/concurrent"
	_ "github.com/modern-go/reflect2"
	_ "github.com/peterbourgon/diskv"
	_ "github.com/pkg/errors"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	_ "golang.org/x/crypto/ssh/terminal"
	_ "golang.org/x/net/context"
	_ "golang.org/x/net/context/ctxhttp"
	_ "golang.org/x/net/http/httpguts"
	_ "golang.org/x/net/http2"
	_ "golang.org/x/net/http2/hpack"
	_ "golang.org/x/net/idna"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/google"
	_ "golang.org/x/oauth2/jws"
	_ "golang.org/x/oauth2/jwt"
	_ "golang.org/x/sys/unix"
	_ "golang.org/x/text/encoding"
	_ "golang.org/x/text/encoding/unicode"
	_ "golang.org/x/text/runes"
	_ "golang.org/x/text/secure/bidirule"
	_ "golang.org/x/text/transform"
	_ "golang.org/x/text/unicode/bidi"
	_ "golang.org/x/text/unicode/norm"
	_ "golang.org/x/text/width"
	_ "golang.org/x/time/rate"
	_ "gopkg.in/inf.v0"
	_ "gopkg.in/yaml.v2"
	_ "k8s.io/api/admissionregistration/v1beta1"
	_ "k8s.io/api/apps/v1"
	_ "k8s.io/api/apps/v1beta1"
	_ "k8s.io/api/apps/v1beta2"
	_ "k8s.io/api/auditregistration/v1alpha1"
	_ "k8s.io/api/authentication/v1"
	_ "k8s.io/api/authentication/v1beta1"
	_ "k8s.io/api/authorization/v1"
	_ "k8s.io/api/authorization/v1beta1"
	_ "k8s.io/api/autoscaling/v1"
	_ "k8s.io/api/autoscaling/v2beta1"
	_ "k8s.io/api/autoscaling/v2beta2"
	_ "k8s.io/api/batch/v1"
	_ "k8s.io/api/batch/v1beta1"
	_ "k8s.io/api/batch/v2alpha1"
	_ "k8s.io/api/certificates/v1beta1"
	_ "k8s.io/api/coordination/v1"
	_ "k8s.io/api/coordination/v1beta1"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/api/events/v1beta1"
	_ "k8s.io/api/extensions/v1beta1"
	_ "k8s.io/api/imagepolicy/v1alpha1"
	_ "k8s.io/api/networking/v1"
	_ "k8s.io/api/networking/v1beta1"
	_ "k8s.io/api/node/v1alpha1"
	_ "k8s.io/api/node/v1beta1"
	_ "k8s.io/api/policy/v1beta1"
	_ "k8s.io/api/rbac/v1"
	_ "k8s.io/api/rbac/v1alpha1"
	_ "k8s.io/api/rbac/v1beta1"
	_ "k8s.io/api/scheduling/v1"
	_ "k8s.io/api/scheduling/v1alpha1"
	_ "k8s.io/api/scheduling/v1beta1"
	_ "k8s.io/api/settings/v1alpha1"
	_ "k8s.io/api/storage/v1"
	_ "k8s.io/api/storage/v1alpha1"
	_ "k8s.io/api/storage/v1beta1"
	_ "k8s.io/apimachinery/pkg/api/equality"
	_ "k8s.io/apimachinery/pkg/api/errors"
	_ "k8s.io/apimachinery/pkg/api/meta"
	_ "k8s.io/apimachinery/pkg/api/resource"
	_ "k8s.io/apimachinery/pkg/api/validation"
	_ "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured/unstructuredscheme"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	_ "k8s.io/apimachinery/pkg/conversion"
	_ "k8s.io/apimachinery/pkg/conversion/queryparams"
	_ "k8s.io/apimachinery/pkg/fields"
	_ "k8s.io/apimachinery/pkg/labels"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/apimachinery/pkg/runtime/serializer"
	_ "k8s.io/apimachinery/pkg/runtime/serializer/json"
	_ "k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
	_ "k8s.io/apimachinery/pkg/runtime/serializer/recognizer"
	_ "k8s.io/apimachinery/pkg/runtime/serializer/streaming"
	_ "k8s.io/apimachinery/pkg/runtime/serializer/versioning"
	_ "k8s.io/apimachinery/pkg/selection"
	_ "k8s.io/apimachinery/pkg/types"
	_ "k8s.io/apimachinery/pkg/util/cache"
	_ "k8s.io/apimachinery/pkg/util/clock"
	_ "k8s.io/apimachinery/pkg/util/diff"
	_ "k8s.io/apimachinery/pkg/util/errors"
	_ "k8s.io/apimachinery/pkg/util/framer"
	_ "k8s.io/apimachinery/pkg/util/httpstream"
	_ "k8s.io/apimachinery/pkg/util/httpstream/spdy"
	_ "k8s.io/apimachinery/pkg/util/intstr"
	_ "k8s.io/apimachinery/pkg/util/json"
	_ "k8s.io/apimachinery/pkg/util/mergepatch"
	_ "k8s.io/apimachinery/pkg/util/naming"
	_ "k8s.io/apimachinery/pkg/util/net"
	_ "k8s.io/apimachinery/pkg/util/remotecommand"
	_ "k8s.io/apimachinery/pkg/util/runtime"
	_ "k8s.io/apimachinery/pkg/util/sets"
	_ "k8s.io/apimachinery/pkg/util/strategicpatch"
	_ "k8s.io/apimachinery/pkg/util/validation"
	_ "k8s.io/apimachinery/pkg/util/validation/field"
	_ "k8s.io/apimachinery/pkg/util/wait"
	_ "k8s.io/apimachinery/pkg/util/yaml"
	_ "k8s.io/apimachinery/pkg/version"
	_ "k8s.io/apimachinery/pkg/watch"
	_ "k8s.io/apimachinery/third_party/forked/golang/json"
	_ "k8s.io/apimachinery/third_party/forked/golang/netutil"
	_ "k8s.io/apimachinery/third_party/forked/golang/reflect"
	_ "k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/cli-runtime/pkg/kustomize"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/configmapandsecret"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/kunstruct"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/kv"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/transformer"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/transformer/hash"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/transformer/patch"
	_ "k8s.io/cli-runtime/pkg/kustomize/k8sdeps/validator"
	_ "k8s.io/cli-runtime/pkg/printers"
	_ "k8s.io/cli-runtime/pkg/resource"
	_ "k8s.io/client-go/deprecated-dynamic"
	_ "k8s.io/client-go/discovery"
	_ "k8s.io/client-go/discovery/cached"
	_ "k8s.io/client-go/discovery/cached/disk"
	_ "k8s.io/client-go/discovery/cached/memory"
	_ "k8s.io/client-go/discovery/fake"
	_ "k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/dynamic/dynamicinformer"
	_ "k8s.io/client-go/dynamic/dynamiclister"
	_ "k8s.io/client-go/dynamic/fake"
	_ "k8s.io/client-go/examples/fake-client"
	_ "k8s.io/client-go/informers"
	_ "k8s.io/client-go/informers/admissionregistration"
	_ "k8s.io/client-go/informers/admissionregistration/v1beta1"
	_ "k8s.io/client-go/informers/apps"
	_ "k8s.io/client-go/informers/apps/v1"
	_ "k8s.io/client-go/informers/apps/v1beta1"
	_ "k8s.io/client-go/informers/apps/v1beta2"
	_ "k8s.io/client-go/informers/auditregistration"
	_ "k8s.io/client-go/informers/auditregistration/v1alpha1"
	_ "k8s.io/client-go/informers/autoscaling"
	_ "k8s.io/client-go/informers/autoscaling/v1"
	_ "k8s.io/client-go/informers/autoscaling/v2beta1"
	_ "k8s.io/client-go/informers/autoscaling/v2beta2"
	_ "k8s.io/client-go/informers/batch"
	_ "k8s.io/client-go/informers/batch/v1"
	_ "k8s.io/client-go/informers/batch/v1beta1"
	_ "k8s.io/client-go/informers/batch/v2alpha1"
	_ "k8s.io/client-go/informers/certificates"
	_ "k8s.io/client-go/informers/certificates/v1beta1"
	_ "k8s.io/client-go/informers/coordination"
	_ "k8s.io/client-go/informers/coordination/v1"
	_ "k8s.io/client-go/informers/coordination/v1beta1"
	_ "k8s.io/client-go/informers/core"
	_ "k8s.io/client-go/informers/core/v1"
	_ "k8s.io/client-go/informers/events"
	_ "k8s.io/client-go/informers/events/v1beta1"
	_ "k8s.io/client-go/informers/extensions"
	_ "k8s.io/client-go/informers/extensions/v1beta1"
	_ "k8s.io/client-go/informers/internalinterfaces"
	_ "k8s.io/client-go/informers/networking"
	_ "k8s.io/client-go/informers/networking/v1"
	_ "k8s.io/client-go/informers/networking/v1beta1"
	_ "k8s.io/client-go/informers/node"
	_ "k8s.io/client-go/informers/node/v1alpha1"
	_ "k8s.io/client-go/informers/node/v1beta1"
	_ "k8s.io/client-go/informers/policy"
	_ "k8s.io/client-go/informers/policy/v1beta1"
	_ "k8s.io/client-go/informers/rbac"
	_ "k8s.io/client-go/informers/rbac/v1"
	_ "k8s.io/client-go/informers/rbac/v1alpha1"
	_ "k8s.io/client-go/informers/rbac/v1beta1"
	_ "k8s.io/client-go/informers/scheduling"
	_ "k8s.io/client-go/informers/scheduling/v1"
	_ "k8s.io/client-go/informers/scheduling/v1alpha1"
	_ "k8s.io/client-go/informers/scheduling/v1beta1"
	_ "k8s.io/client-go/informers/settings"
	_ "k8s.io/client-go/informers/settings/v1alpha1"
	_ "k8s.io/client-go/informers/storage"
	_ "k8s.io/client-go/informers/storage/v1"
	_ "k8s.io/client-go/informers/storage/v1alpha1"
	_ "k8s.io/client-go/informers/storage/v1beta1"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1beta2/fake"
	_ "k8s.io/client-go/kubernetes/typed/auditregistration/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/auditregistration/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/authentication/v1"
	_ "k8s.io/client-go/kubernetes/typed/authentication/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/authentication/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/authorization/v1"
	_ "k8s.io/client-go/kubernetes/typed/authorization/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/authorization/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	_ "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2/fake"
	_ "k8s.io/client-go/kubernetes/typed/batch/v1"
	_ "k8s.io/client-go/kubernetes/typed/batch/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/batch/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	_ "k8s.io/client-go/kubernetes/typed/batch/v2alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/certificates/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/coordination/v1"
	_ "k8s.io/client-go/kubernetes/typed/coordination/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/coordination/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/events/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/extensions/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/networking/v1"
	_ "k8s.io/client-go/kubernetes/typed/networking/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/networking/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/node/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/node/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/policy/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/rbac/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1/fake"
	_ "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/settings/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1/fake"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1alpha1/fake"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	_ "k8s.io/client-go/kubernetes/typed/storage/v1beta1/fake"
	_ "k8s.io/client-go/listers/admissionregistration/v1beta1"
	_ "k8s.io/client-go/listers/apps/v1"
	_ "k8s.io/client-go/listers/apps/v1beta1"
	_ "k8s.io/client-go/listers/apps/v1beta2"
	_ "k8s.io/client-go/listers/auditregistration/v1alpha1"
	_ "k8s.io/client-go/listers/authentication/v1"
	_ "k8s.io/client-go/listers/authentication/v1beta1"
	_ "k8s.io/client-go/listers/authorization/v1"
	_ "k8s.io/client-go/listers/authorization/v1beta1"
	_ "k8s.io/client-go/listers/autoscaling/v1"
	_ "k8s.io/client-go/listers/autoscaling/v2beta1"
	_ "k8s.io/client-go/listers/autoscaling/v2beta2"
	_ "k8s.io/client-go/listers/batch/v1"
	_ "k8s.io/client-go/listers/batch/v1beta1"
	_ "k8s.io/client-go/listers/batch/v2alpha1"
	_ "k8s.io/client-go/listers/certificates/v1beta1"
	_ "k8s.io/client-go/listers/coordination/v1"
	_ "k8s.io/client-go/listers/coordination/v1beta1"
	_ "k8s.io/client-go/listers/core/v1"
	_ "k8s.io/client-go/listers/events/v1beta1"
	_ "k8s.io/client-go/listers/extensions/v1beta1"
	_ "k8s.io/client-go/listers/imagepolicy/v1alpha1"
	_ "k8s.io/client-go/listers/networking/v1"
	_ "k8s.io/client-go/listers/networking/v1beta1"
	_ "k8s.io/client-go/listers/node/v1alpha1"
	_ "k8s.io/client-go/listers/node/v1beta1"
	_ "k8s.io/client-go/listers/policy/v1beta1"
	_ "k8s.io/client-go/listers/rbac/v1"
	_ "k8s.io/client-go/listers/rbac/v1alpha1"
	_ "k8s.io/client-go/listers/rbac/v1beta1"
	_ "k8s.io/client-go/listers/scheduling/v1"
	_ "k8s.io/client-go/listers/scheduling/v1alpha1"
	_ "k8s.io/client-go/listers/scheduling/v1beta1"
	_ "k8s.io/client-go/listers/settings/v1alpha1"
	_ "k8s.io/client-go/listers/storage/v1"
	_ "k8s.io/client-go/listers/storage/v1alpha1"
	_ "k8s.io/client-go/listers/storage/v1beta1"
	_ "k8s.io/client-go/pkg/apis/clientauthentication"
	_ "k8s.io/client-go/pkg/apis/clientauthentication/install"
	_ "k8s.io/client-go/pkg/apis/clientauthentication/v1alpha1"
	_ "k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"
	_ "k8s.io/client-go/pkg/version"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/exec"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
	_ "k8s.io/client-go/rest"
	_ "k8s.io/client-go/rest/fake"
	_ "k8s.io/client-go/rest/watch"
	_ "k8s.io/client-go/restmapper"
	_ "k8s.io/client-go/scale"
	_ "k8s.io/client-go/scale/fake"
	_ "k8s.io/client-go/scale/scheme"
	_ "k8s.io/client-go/scale/scheme/appsint"
	_ "k8s.io/client-go/scale/scheme/appsv1beta1"
	_ "k8s.io/client-go/scale/scheme/appsv1beta2"
	_ "k8s.io/client-go/scale/scheme/autoscalingv1"
	_ "k8s.io/client-go/scale/scheme/extensionsint"
	_ "k8s.io/client-go/scale/scheme/extensionsv1beta1"
	_ "k8s.io/client-go/testing"
	_ "k8s.io/client-go/third_party/forked/golang/template"
	_ "k8s.io/client-go/tools/auth"
	_ "k8s.io/client-go/tools/cache"
	_ "k8s.io/client-go/tools/cache/testing"
	_ "k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/tools/clientcmd/api"
	_ "k8s.io/client-go/tools/clientcmd/api/latest"
	_ "k8s.io/client-go/tools/clientcmd/api/v1"
	_ "k8s.io/client-go/tools/events"
	_ "k8s.io/client-go/tools/leaderelection"
	_ "k8s.io/client-go/tools/leaderelection/resourcelock"
	_ "k8s.io/client-go/tools/metrics"
	_ "k8s.io/client-go/tools/pager"
	_ "k8s.io/client-go/tools/portforward"
	_ "k8s.io/client-go/tools/record"
	_ "k8s.io/client-go/tools/record/util"
	_ "k8s.io/client-go/tools/reference"
	_ "k8s.io/client-go/tools/remotecommand"
	_ "k8s.io/client-go/tools/watch"
	_ "k8s.io/client-go/transport"
	_ "k8s.io/client-go/transport/spdy"
	_ "k8s.io/client-go/util/cert"
	_ "k8s.io/client-go/util/certificate"
	_ "k8s.io/client-go/util/certificate/csr"
	_ "k8s.io/client-go/util/connrotation"
	_ "k8s.io/client-go/util/exec"
	_ "k8s.io/client-go/util/flowcontrol"
	_ "k8s.io/client-go/util/homedir"
	_ "k8s.io/client-go/util/jsonpath"
	_ "k8s.io/client-go/util/keyutil"
	_ "k8s.io/client-go/util/retry"
	_ "k8s.io/client-go/util/testing"
	_ "k8s.io/client-go/util/workqueue"
	_ "k8s.io/klog"
	_ "k8s.io/kube-openapi/pkg/common"
	_ "k8s.io/kube-openapi/pkg/util/proto"
	_ "k8s.io/utils/buffer"
	_ "k8s.io/utils/integer"
	_ "k8s.io/utils/trace"
	_ "sigs.k8s.io/kustomize/pkg/commands/build"
	_ "sigs.k8s.io/kustomize/pkg/constants"
	_ "sigs.k8s.io/kustomize/pkg/expansion"
	_ "sigs.k8s.io/kustomize/pkg/factory"
	_ "sigs.k8s.io/kustomize/pkg/fs"
	_ "sigs.k8s.io/kustomize/pkg/git"
	_ "sigs.k8s.io/kustomize/pkg/gvk"
	_ "sigs.k8s.io/kustomize/pkg/ifc"
	_ "sigs.k8s.io/kustomize/pkg/ifc/transformer"
	_ "sigs.k8s.io/kustomize/pkg/image"
	_ "sigs.k8s.io/kustomize/pkg/loader"
	_ "sigs.k8s.io/kustomize/pkg/patch"
	_ "sigs.k8s.io/kustomize/pkg/patch/transformer"
	_ "sigs.k8s.io/kustomize/pkg/resid"
	_ "sigs.k8s.io/kustomize/pkg/resmap"
	_ "sigs.k8s.io/kustomize/pkg/resource"
	_ "sigs.k8s.io/kustomize/pkg/target"
	_ "sigs.k8s.io/kustomize/pkg/transformers"
	_ "sigs.k8s.io/kustomize/pkg/transformers/config"
	_ "sigs.k8s.io/kustomize/pkg/transformers/config/defaultconfig"
	_ "sigs.k8s.io/kustomize/pkg/types"
	_ "sigs.k8s.io/yaml"
)
