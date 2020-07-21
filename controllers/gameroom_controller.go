/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/k0kubun/pp"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
)

// GameroomReconciler reconciles a Gameroom object
type GameroomReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=pod,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=pods/status,verbs=get
func (r *GameroomReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("gameroom", req.NamespacedName)
	log.Info("Start ...")

	var instance corev1.Pod

	// Get record from kubernetes api
	if err := r.Client.Get(ctx, req.NamespacedName, &instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// // If object hasn't been deleted and doesn't have a finalizer,
	// // Add a finalizer to newly created objects.
	// if instance.ObjectMeta.DeletionTimestamp.IsZero() && !util.Contains(instance.ObjectMeta.Finalizers, sqlv1alpha2.Finalizer) {
	// 	instance.Finalizers = append(instance.Finalizers, sqlv1alpha2.Finalizer)
	// 	if err := r.Patch(context.Background(), &instance); err != nil {
	// 		log.Error(err, "failed to add finalizer to sql instance")
	// 		return ctrl.Result{}, err
	// 	}

	// 	// Since adding the finalizer updates the object return to avoid later update issues
	// 	return ctrl.Result{}, nil
	// }

	pp.Println(instance)

	// your logic here

	return ctrl.Result{}, nil
}

// https://github.com/packyzbq/CRDSample/blob/master/controllers/demo_controller.go
func (r *GameroomReconciler) SetupWithManager(mgr ctrl.Manager) error {
	mgr.GetWebhookServer()
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		Complete(r)
}
