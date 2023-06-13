package controllers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Upgrade interface {
	upgradeScenario(context.Context, client.Client, interface{}) bool
	updateStatefulSet(context.Context, client.Client, metav1.ObjectMeta, string, interface{}) error
}
