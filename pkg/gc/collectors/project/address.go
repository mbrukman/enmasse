/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

package project

import (
	corev1alpha1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1beta1"
	"github.com/enmasseproject/enmasse/pkg/apis/iot/v1alpha1"
	"github.com/enmasseproject/enmasse/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *projectCollector) collectAddresses() error {

	log.Info("Collect Addresses")

	opts := metav1.ListOptions{}

	list, err := p.client.EnmasseV1beta1().
		Addresses(p.namespace).
		List(opts)

	if err != nil {
		return err
	}

	mt := util.MultiTool{}

	for _, addr := range list.Items {
		mt.Ran(p.checkAddresss(&addr))
	}

	return mt.Error
}

func (p *projectCollector) checkAddresss(addr *corev1alpha1.Address) error {
	log.Info("Checking address", "Address", addr)

	found, notFound, err := p.findOwningProjects(addr, true)
	if err != nil {
		return err
	}

	if len(found) <= 0 && len(notFound) > 0 {
		// we were owned, but now everyone is gone
		return p.deleteAddress(addr)
	}

	for _, proj := range found {
		if p.needAddress(addr, &proj) {
			return nil
		}
	}

	return p.deleteAddress(addr)

	return nil
}

func (p *projectCollector) needAddress(addr *corev1alpha1.Address, proj *v1alpha1.IoTProject) bool {
	if proj.Spec.DownstreamStrategy.ManagedDownstreamStrategy == nil {
		return false
	}

	return true
}

func (p *projectCollector) deleteAddress(addr *corev1alpha1.Address) error {

	log.Info("Deleting Address", "Address", addr, "UID", addr.UID)

	return p.client.EnmasseV1beta1().
		Addresses(addr.Namespace).
		Delete(addr.Name, &metav1.DeleteOptions{
			Preconditions: &metav1.Preconditions{UID: &addr.UID},
		})
}
