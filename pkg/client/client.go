package client

import (
	"golanglearning/new_project/multi_cluster_client/pkg/config"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	"k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1alpha1 "k8s.io/client-go/kubernetes/typed/authentication/v1alpha1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	"k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	certificatesv1alpha1 "k8s.io/client-go/kubernetes/typed/certificates/v1alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	flowcontrolv1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	flowcontrolv1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	resourcev1alpha2 "k8s.io/client-go/kubernetes/typed/resource/v1alpha2"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

// MultiClientSet 多集群client
type MultiClientSet struct {
	// clientSets 存储所有的k8s clientSet
	clientSets  map[string]kubernetes.Interface
	// 目前是冗余字段
	clusterList []string
	// selectedCluster 指定时传入的集群名
	selectedCluster string
}

// Interface 嵌套接口
type Interface interface {
	// Cluster 选择集群
	Cluster(cluster string) Interface
	kubernetes.Interface
}

// NewForConfig 初始化multi client
func NewForConfig(c *config.Config) (*MultiClientSet, error) {

	mc := &MultiClientSet{
		clientSets: map[string]kubernetes.Interface{},
		clusterList: make([]string, 0),
	}

	for _, v := range c.Clusters {
		if v.MetaData.ConfigPath != "" {
			ccg, err := clientcmd.BuildConfigFromFlags("", v.MetaData.ConfigPath)
			if err != nil {
				return nil, err
			}
			ccg.Insecure = v.MetaData.Insecure
			clientSet, err := kubernetes.NewForConfig(ccg)
			if err != nil {
				return nil, err
			}

			mc.clientSets[v.MetaData.ClusterName] = clientSet
			mc.clusterList = append(mc.clusterList, v.MetaData.ClusterName)
		}
	}

	return mc, nil
}

var _ Interface = &MultiClientSet{}

func (m *MultiClientSet) Cluster(cluster string) Interface {
	// TODO: 当没指定集群时，获取所有集群
	if cluster == "" {
		panic("not specified cluster name")
	}
	m.selectedCluster = cluster
	return m
}

func (m *MultiClientSet) Discovery() discovery.DiscoveryInterface {
	return m.clientSets[m.selectedCluster].Discovery()
}

func (m *MultiClientSet) AdmissionregistrationV1() v1.AdmissionregistrationV1Interface {
	return m.clientSets[m.selectedCluster].AdmissionregistrationV1()
}

func (m *MultiClientSet) AdmissionregistrationV1alpha1() v1alpha1.AdmissionregistrationV1alpha1Interface {
	return m.clientSets[m.selectedCluster].AdmissionregistrationV1alpha1()
}

func (m *MultiClientSet) AdmissionregistrationV1beta1() v1beta1.AdmissionregistrationV1beta1Interface {
	return m.clientSets[m.selectedCluster].AdmissionregistrationV1beta1()
}

func (m *MultiClientSet) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return m.clientSets[m.selectedCluster].InternalV1alpha1()
}

func (m *MultiClientSet) AppsV1() appsv1.AppsV1Interface {
	return m.clientSets[m.selectedCluster].AppsV1()
}

func (m *MultiClientSet) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return m.clientSets[m.selectedCluster].AppsV1beta1()
}

func (m *MultiClientSet) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return m.clientSets[m.selectedCluster].AppsV1beta2()
}

func (m *MultiClientSet) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return m.clientSets[m.selectedCluster].AuthenticationV1()
}

func (m *MultiClientSet) AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return m.clientSets[m.selectedCluster].AuthenticationV1alpha1()
}

func (m *MultiClientSet) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	return m.clientSets[m.selectedCluster].AuthenticationV1beta1()
}

func (m *MultiClientSet) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return m.clientSets[m.selectedCluster].AuthorizationV1()
}

func (m *MultiClientSet) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	return m.clientSets[m.selectedCluster].AuthorizationV1beta1()
}

func (m *MultiClientSet) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return m.clientSets[m.selectedCluster].AutoscalingV1()
}

func (m *MultiClientSet) AutoscalingV2() autoscalingv2.AutoscalingV2Interface {
	return m.clientSets[m.selectedCluster].AutoscalingV2()
}

func (m *MultiClientSet) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return m.clientSets[m.selectedCluster].AutoscalingV2beta1()
}

func (m *MultiClientSet) AutoscalingV2beta2() v2beta2.AutoscalingV2beta2Interface {
	return m.clientSets[m.selectedCluster].AutoscalingV2beta2()
}

func (m *MultiClientSet) BatchV1() batchv1.BatchV1Interface {
	return m.clientSets[m.selectedCluster].BatchV1()
}

func (m *MultiClientSet) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return m.clientSets[m.selectedCluster].BatchV1beta1()
}

func (m *MultiClientSet) CertificatesV1() certificatesv1.CertificatesV1Interface {
	return m.clientSets[m.selectedCluster].CertificatesV1()
}

func (m *MultiClientSet) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return m.clientSets[m.selectedCluster].CertificatesV1beta1()
}

func (m *MultiClientSet) CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface {
	return m.clientSets[m.selectedCluster].CertificatesV1alpha1()
}

func (m *MultiClientSet) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	return m.clientSets[m.selectedCluster].CoordinationV1beta1()
}

func (m *MultiClientSet) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return m.clientSets[m.selectedCluster].CoordinationV1()
}

func (m *MultiClientSet) CoreV1() corev1.CoreV1Interface {
	return m.clientSets[m.selectedCluster].CoreV1()
}

func (m *MultiClientSet) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	return m.clientSets[m.selectedCluster].DiscoveryV1()
}

func (m *MultiClientSet) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return m.clientSets[m.selectedCluster].DiscoveryV1beta1()
}

func (m *MultiClientSet) EventsV1() eventsv1.EventsV1Interface {
	return m.clientSets[m.selectedCluster].EventsV1()
}

func (m *MultiClientSet) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return m.clientSets[m.selectedCluster].EventsV1beta1()
}

func (m *MultiClientSet) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return m.clientSets[m.selectedCluster].ExtensionsV1beta1()
}

func (m *MultiClientSet) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1Interface {
	return m.clientSets[m.selectedCluster].FlowcontrolV1alpha1()
}

func (m *MultiClientSet) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	return m.clientSets[m.selectedCluster].FlowcontrolV1beta1()
}

func (m *MultiClientSet) FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2Interface {
	return m.clientSets[m.selectedCluster].FlowcontrolV1beta2()
}

func (m *MultiClientSet) FlowcontrolV1beta3() flowcontrolv1beta3.FlowcontrolV1beta3Interface {
	return m.clientSets[m.selectedCluster].FlowcontrolV1beta3()
}

func (m *MultiClientSet) NetworkingV1() networkingv1.NetworkingV1Interface {
	return m.clientSets[m.selectedCluster].NetworkingV1()
}

func (m *MultiClientSet) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return m.clientSets[m.selectedCluster].NetworkingV1alpha1()
}

func (m *MultiClientSet) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return m.clientSets[m.selectedCluster].NetworkingV1beta1()
}

func (m *MultiClientSet) NodeV1() nodev1.NodeV1Interface {
	return m.clientSets[m.selectedCluster].NodeV1()
}

func (m *MultiClientSet) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return m.clientSets[m.selectedCluster].NodeV1alpha1()
}

func (m *MultiClientSet) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	return m.clientSets[m.selectedCluster].NodeV1beta1()
}

func (m *MultiClientSet) PolicyV1() policyv1.PolicyV1Interface {
	return m.clientSets[m.selectedCluster].PolicyV1()
}

func (m *MultiClientSet) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return m.clientSets[m.selectedCluster].PolicyV1beta1()
}

func (m *MultiClientSet) RbacV1() rbacv1.RbacV1Interface {
	return m.clientSets[m.selectedCluster].RbacV1()
}

func (m *MultiClientSet) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return m.clientSets[m.selectedCluster].RbacV1beta1()
}

func (m *MultiClientSet) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return m.clientSets[m.selectedCluster].RbacV1alpha1()
}

func (m *MultiClientSet) ResourceV1alpha2() resourcev1alpha2.ResourceV1alpha2Interface {
	return m.clientSets[m.selectedCluster].ResourceV1alpha2()
}

func (m *MultiClientSet) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return m.clientSets[m.selectedCluster].SchedulingV1alpha1()
}

func (m *MultiClientSet) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return m.clientSets[m.selectedCluster].SchedulingV1beta1()
}

func (m *MultiClientSet) SchedulingV1() schedulingv1.SchedulingV1Interface {
	return m.clientSets[m.selectedCluster].SchedulingV1()
}

func (m *MultiClientSet) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return m.clientSets[m.selectedCluster].StorageV1beta1()
}

func (m *MultiClientSet) StorageV1() storagev1.StorageV1Interface {
	return m.clientSets[m.selectedCluster].StorageV1()
}

func (m *MultiClientSet) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return m.clientSets[m.selectedCluster].StorageV1alpha1()
}
