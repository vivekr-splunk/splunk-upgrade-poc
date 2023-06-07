package controllers

import (
	"context"
	"fmt"

	enterprisev1 "github.com/vivekrsplunk/splunk-upgrade-poc/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

/*
apiVersion: v1
kind: Service
metadata:

	name: nginx
	labels:
	  app: nginx

spec:

	ports:
	- port: 80
	  name: web
	clusterIP: None
	selector:
	  app: nginx

---
apiVersion: apps/v1
kind: StatefulSet
metadata:

	name: web

spec:

	selector:
	  matchLabels:
	    app: nginx # has to match .spec.template.metadata.labels
	serviceName: "nginx"
	replicas: 3 # by default is 1
	minReadySeconds: 10 # by default is 0
	template:
	  metadata:
	    labels:
	      app: nginx # has to match .spec.selector.matchLabels
	  spec:
	    terminationGracePeriodSeconds: 10
	    containers:
	    - name: nginx
	      image: registry.k8s.io/nginx-slim:0.8
	      ports:
	      - containerPort: 80
	        name: web
	      volumeMounts:
	      - name: www
	        mountPath: /usr/share/nginx/html
	volumeClaimTemplates:
	- metadata:
	    name: www
	  spec:
	    accessModes: [ "ReadWriteOnce" ]
	    storageClassName: "my-storage-class"
	    resources:
	      requests:
	        storage: 1Gi
*/
func updateLMStatefulSet(ctx context.Context, c client.Client, meta metav1.ObjectMeta, image string, request *enterprisev1.LicenseManager) error {
	replicas := int32(3)
	name := fmt.Sprintf("%s-%s", meta.GetName(), "st")
	namespacedName := types.NamespacedName{Namespace: meta.GetNamespace(), Name: name}
	statefulSet := appsv1.StatefulSet{}
	err := c.Get(ctx, namespacedName, &statefulSet)
	if err != nil && k8serrors.IsNotFound(err) {
		// create new statefulset
		statefulSet.Name = name
		statefulSet.Namespace = meta.GetNamespace()
		statefulSet.Spec.Replicas = &replicas
		statefulSet.Spec.Template.Spec.Containers = []corev1.Container{
			{
				Name:  "nginx",
				Image: image,
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: 80,
						Name:          "web",
					},
				},
			},
		}
		err = c.Create(ctx, &statefulSet)
		if err != nil {
			return err
		}

	}
	// else update the statefulset image
	if statefulSet.Spec.Template.Spec.Containers[0].Image == image {
		return fmt.Errorf("image tag is same")
	}
	if statefulSet.Status.AvailableReplicas != 3 {
		return fmt.Errorf("replicas not 3")
	}

	statefulSet.Spec.Template.Spec.Containers[0].Image = image
	err = c.Update(ctx, &statefulSet)
	if err != nil {
		return err
	}
	request.Status.Image = image
	request.Status.Phase = "Ready"

	return nil
}
func updateCMStatefulSet(ctx context.Context, c client.Client, meta metav1.ObjectMeta, image string, request *enterprisev1.ClusterManager) error {
	replicas := int32(3)
	name := fmt.Sprintf("%s-%s", meta.GetName(), "st")
	namespacedName := types.NamespacedName{Namespace: meta.GetNamespace(), Name: name}
	statefulSet := appsv1.StatefulSet{}
	err := c.Get(ctx, namespacedName, &statefulSet)
	if err != nil && k8serrors.IsNotFound(err) {
		// create new statefulset
		statefulSet.Name = name
		statefulSet.Namespace = meta.GetNamespace()
		statefulSet.Spec.Replicas = &replicas
		statefulSet.Spec.Template.Spec.Containers = []corev1.Container{
			{
				Name:  "nginx",
				Image: image,
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: 80,
						Name:          "web",
					},
				},
			},
		}
		err = c.Create(ctx, &statefulSet)
		if err != nil {
			return err
		}

	}
	// else update the statefulset image
	if statefulSet.Spec.Template.Spec.Containers[0].Image == image {
		return fmt.Errorf("image tag is same")
	}
	if statefulSet.Status.AvailableReplicas != 3 {
		return fmt.Errorf("replicas not 3")
	}

	statefulSet.Spec.Template.Spec.Containers[0].Image = image
	err = c.Update(ctx, &statefulSet)
	if err != nil {
		return err
	}
	request.Status.Image = image
	request.Status.Phase = "Ready"

	return nil
}
func updateStatefulSet(ctx context.Context, c client.Client, meta metav1.ObjectMeta, image string) error {
	replicas := int32(3)
	name := fmt.Sprintf("%s-%s", meta.GetName(), "st")
	namespacedName := types.NamespacedName{Namespace: meta.GetNamespace(), Name: name}
	statefulSet := appsv1.StatefulSet{}
	err := c.Get(ctx, namespacedName, &statefulSet)
	if err != nil && k8serrors.IsNotFound(err) {
		// create new statefulset
		statefulSet.Name = name
		statefulSet.Namespace = meta.GetNamespace()
		statefulSet.Spec.Replicas = &replicas
		statefulSet.Spec.Template.Spec.Containers = []corev1.Container{
			{
				Name:  "nginx",
				Image: image,
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: 80,
						Name:          "web",
					},
				},
			},
		}
		err = c.Create(ctx, &statefulSet)
		if err != nil {
			return err
		}

	}
	// else update the statefulset image
	if statefulSet.Spec.Template.Spec.Containers[0].Image == image {
		return fmt.Errorf("image tag is same")
	}
	if statefulSet.Status.AvailableReplicas != 3 {
		return fmt.Errorf("replicas not 3")
	}

	statefulSet.Spec.Template.Spec.Containers[0].Image = image
	err = c.Update(ctx, &statefulSet)
	if err != nil {
		return err
	}

	return nil
}

func upgradeScenarioForLicenseManager(ctx context.Context, c client.Client, request *enterprisev1.LicenseManager) bool {
	return false
}

func upgradeScenarioForClusterManager(ctx context.Context, c client.Client, request *enterprisev1.ClusterManager) bool {

	reqLogger := log.FromContext(ctx)
	// read license manager reference

	licenseManagerRef := request.Spec.LicenseManagerRef
	namespacedName := types.NamespacedName{Namespace: request.GetNamespace(), Name: licenseManagerRef.Name}

	// create new object
	licenseManager := &enterprisev1.LicenseManager{}

	// get the license manager referred in cluster manager
	err := c.Get(ctx, namespacedName, licenseManager)
	if err != nil {
		reqLogger.Error(err, "unable to find license manager", "name", licenseManagerRef.Name, "namespace", licenseManagerRef.Namespace)
		return false
	}
	if request.Status.Image != request.Spec.Image && licenseManager.Status.Image == request.Spec.Image && licenseManager.Status.Phase == "Ready" {
		// update possible
		fmt.Println("1")
		return true
	}
	if request.Status.Image != request.Spec.Image {
		request.Status.Phase = "Pending"
	}
	fmt.Println("0")

	return false
}

func upgradeScenarioForMonitoringConsole(ctx context.Context, c client.Client, request *enterprisev1.MonitoringConsole) bool {
	return false
}

func upgradeScenarioForIndexerCluster(ctx context.Context, c client.Client, request *enterprisev1.IndexerCluster) bool {
	return false
}

func upgradeScenarioForSearchHeadCluster(ctx context.Context, c client.Client, request *enterprisev1.SearchHeadCluster) bool {
	return false
}

func upgradeScenarioForStandalone(ctx context.Context, c client.Client, request *enterprisev1.Standalone) bool {
	return false
}
