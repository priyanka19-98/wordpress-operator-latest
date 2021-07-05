package controllers

import (
	"context"

	v1 "github.com/priyanka19-98/wordpress-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *WordpressReconciler) deploymentForMysql(cr *v1.Wordpress) *appsv1.Deployment {
	labels := map[string]string{
		"app": cr.Name,
	}
	matchlabels := map[string]string{
		"app":  cr.Name,
		"tier": "mysql",
	}

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "wordpress-mysql",
			Namespace: cr.Namespace,
			Labels:    labels,
		},

		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: matchlabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: matchlabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "mysql:5.6",
						Name:  "mysql",

						Env: []corev1.EnvVar{
							{
								Name:  "MYSQL_ROOT_PASSWORD",
								Value: cr.Spec.SQLRootPassword,
							},
						},

						Ports: []corev1.ContainerPort{{
							ContainerPort: 3306,
							Name:          "mysql",
						}},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      "mysql-persistent-storage",
								MountPath: "/var/lib/mysql",
							},
						},
					},
					},

					Volumes: []corev1.Volume{

						{
							Name: "mysql-persistent-storage",
							VolumeSource: corev1.VolumeSource{

								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mysql-pv-claim",
								},
							},
						},
					},
				},
			},
		},
	}

	controllerutil.SetControllerReference(cr, dep, r.Scheme)
	return dep
}

func (r *WordpressReconciler) pvcForMysql(cr *v1.Wordpress) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app": cr.Name,
	}

	pvc := &corev1.PersistentVolumeClaim{

		ObjectMeta: metav1.ObjectMeta{
			Name:      "mysql-pv-claim",
			Namespace: cr.Namespace,
			Labels:    labels,
		},

		Spec: corev1.PersistentVolumeClaimSpec{

			AccessModes: []corev1.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},

			Resources: corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse("10Gi"),
				},
			},
		},
	}

	controllerutil.SetControllerReference(cr, pvc, r.Scheme)
	return pvc

}

func (r *WordpressReconciler) serviceForMysql(cr *v1.Wordpress) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name,
	}
	matchlabels := map[string]string{
		"app":  cr.Name,
		"tier": "mysql",
	}

	ser := &corev1.Service{

		ObjectMeta: metav1.ObjectMeta{
			Name:      "wordpress-mysql",
			Namespace: cr.Namespace,
			Labels:    labels,
		},

		Spec: corev1.ServiceSpec{
			Selector: matchlabels,

			Ports: []corev1.ServicePort{
				{
					Port: 3306,
					Name: cr.Name,
				},
			},
			ClusterIP: "None",
		},
	}

	controllerutil.SetControllerReference(cr, ser, r.Scheme)
	return ser

}

func (r *WordpressReconciler) isMysqlUp(v *v1.Wordpress) bool {
	deployment := &appsv1.Deployment{}

	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name:      "wordpress-mysql",
		Namespace: v.Namespace,
	}, deployment)

	if err != nil {
		r.Log.Error(err, "Deployment mysql not found")
		return false
	}
	if deployment.Status.ReadyReplicas == 1 {
		return true
	}

	return false

}
