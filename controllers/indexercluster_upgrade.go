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

func (r *IndexerClusterReconciler) updateStatefulSet(ctx context.Context, c client.Client, meta metav1.ObjectMeta, image string, instance interface{}) error {
	replicas := int32(3)
	matchlabels := map[string]string{
		"app":  meta.GetName(), //changed from request.name
		"tier": "splunk",
	}
	annotations := map[string]string{
		"test": "a",
	}
	request := instance.(*enterprisev1.IndexerCluster)
	site := request.Spec.Site
	namePrefix := fmt.Sprintf("%s-%s", meta.GetName(), "st")
	for s := 1; s <= site; s++ {
		name := fmt.Sprintf("%s-%s-%d", namePrefix, "site", s)
		namespacedName := types.NamespacedName{Namespace: meta.GetNamespace(), Name: name}
		statefulSet := appsv1.StatefulSet{}
		err := c.Get(ctx, namespacedName, &statefulSet)

		if err != nil && k8serrors.IsNotFound(err) {
			// create new statefulset
			statefulSet.Name = name
			statefulSet.Namespace = meta.GetNamespace()
			statefulSet.Spec.Replicas = &replicas
			statefulSet.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: matchlabels,
			}
			statefulSet.Spec.Template = corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      matchlabels,
					Annotations: annotations,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: image,
						Name:  "nginx",
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: 80,
								Name:          "web",
							},
						},
					},
					},
				},
			}
			statefulSet.OwnerReferences = []metav1.OwnerReference{
				{
					Name:       request.Name,
					Kind:       request.Kind,
					UID:        request.UID,
					APIVersion: request.APIVersion,
				},
			}
			err = c.Create(ctx, &statefulSet)
			if err != nil {
				return err
			}
			request.Status.Image = request.Spec.Image
			request.Status.Phase = "Ready"
			c.Status().Update(context.Background(), request)
			fmt.Println("IC Image Status Create", request.Status.Image)
			fmt.Println("IC Image Spec Create", request.Spec.Image)

			continue

		}
		// else update the statefulset image
		fmt.Println("IC Image Status Before Upgrade", request.Status.Image)
		fmt.Println("IC Image Spec Before Upgrade", request.Spec.Image)
		if statefulSet.Spec.Template.Spec.Containers[0].Image == image {
			return fmt.Errorf("image tag is same")
		}
		if statefulSet.Status.AvailableReplicas != 3 {
			return fmt.Errorf("replicas not 3")
		}

		statefulSet.Spec.Template.Spec.Containers[0].Image = image
		err = c.Update(ctx, &statefulSet)
		// if err != nil {
		// 	return err
		// }

		request.Status.Image = request.Spec.Image
		request.Status.Phase = "Ready"
		fmt.Println("IC Image Status 2", request.Status.Image)
		fmt.Println("IC Image Status 2", request.Status.Phase)
		fmt.Println("IC Upgrade Done!")
		c.Status().Update(context.Background(), request)

	}
	return nil
}
func (r *IndexerClusterReconciler) upgradeScenario(ctx context.Context, c client.Client, instance interface{}) bool {

	reqLogger := log.FromContext(ctx)
	request := instance.(*enterprisev1.IndexerCluster)
	// read license manager reference

	clusterManagerRef := request.Spec.ClusterManagerRef
	namespacedName := types.NamespacedName{Namespace: request.GetNamespace(), Name: clusterManagerRef.Name}

	// create new object
	clusterManager := &enterprisev1.ClusterManager{}

	// get the license manager referred in cluster manager
	err := c.Get(ctx, namespacedName, clusterManager)
	if err != nil {
		reqLogger.Error(err, "unable to find license manager", "name", clusterManagerRef.Name, "namespace", clusterManagerRef.Namespace)
		return false
	}
	fmt.Println("IC Status Image in IC", request.Status.Image)
	fmt.Println("IC Spec Image in IC", request.Spec.Image)
	fmt.Println("CM Status Image in IC", clusterManager.Status.Image)
	fmt.Println("CM spec Image in IC", clusterManager.Spec.Image)
	fmt.Println("CM Status Phase in IC", clusterManager.Status.Phase)

	if request.Status.Image != request.Spec.Image && clusterManager.Status.Image == request.Spec.Image && clusterManager.Status.Phase == "Ready" {
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
