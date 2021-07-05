package controllers

import (
	"context"

	v1 "github.com/priyanka19-98/wordpress-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (r *WordpressReconciler) ensureDeployment(request reconcile.Request,
	instance *v1.Wordpress,
	dep *appsv1.Deployment,
) (*reconcile.Result, error) {

	found := &appsv1.Deployment{}

	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name:      dep.Name,
		Namespace: instance.Namespace,
	}, found)

	if err != nil && errors.IsNotFound(err) {

		// Create the deployment
		r.Log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		err = r.Client.Create(context.TODO(), dep)

		if err != nil {
			// Deployment failed
			r.Log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return &reconcile.Result{}, err
		}
		// Deployment was successful
		return nil, nil

	} else if err != nil {
		// Error that isn't due to the deployment not existing
		r.Log.Error(err, "Failed to get Deployment")
		return &ctrl.Result{}, err
	}

	return nil, nil

}

func (r *WordpressReconciler) ensureService(request reconcile.Request,
	instance *v1.Wordpress,
	s *corev1.Service,
) (*reconcile.Result, error) {
	found := &corev1.Service{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name:      s.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create the service
		r.Log.Info("Creating a new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
		err = r.Client.Create(context.TODO(), s)

		if err != nil {
			// Creation failed
			r.Log.Error(err, "Failed to create new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
			return &ctrl.Result{}, err
		}
		// Creation was successful
		return nil, nil

	} else if err != nil {
		// Error that isn't due to the service not existing
		r.Log.Error(err, "Failed to get Service")
		return &ctrl.Result{}, err
	}

	return nil, nil
}

func (r *WordpressReconciler) ensurePVC(request reconcile.Request,
	instance *v1.Wordpress,
	s *corev1.PersistentVolumeClaim,
) (*reconcile.Result, error) {

	found := &corev1.PersistentVolumeClaim{}

	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name:      s.Name,
		Namespace: instance.Namespace,
	}, found)

	if err != nil && errors.IsNotFound(err) {
		// Create the PVC
		r.Log.Info("Creating a new PVC", "PVC.Namespace", s.Namespace, "PVC.Name", s.Name)
		err = r.Client.Create(context.TODO(), s)

		if err != nil {
			// Creation failed
			r.Log.Error(err, "Failed to create new PVC", "PVC.Namespace", s.Namespace, "PVC.Name", s.Name)
			return &ctrl.Result{}, err
		}
		// Creation was successful
		return nil, nil

	} else if err != nil {
		// Error that isn't due to the pvc not existing
		r.Log.Error(err, "Failed to get PVC")
		return &ctrl.Result{}, err
	}

	return nil, nil

}
