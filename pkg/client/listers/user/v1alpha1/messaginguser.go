/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/enmasseproject/enmasse/pkg/apis/user/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MessagingUserLister helps list MessagingUsers.
type MessagingUserLister interface {
	// List lists all MessagingUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MessagingUser, err error)
	// MessagingUsers returns an object that can list and get MessagingUsers.
	MessagingUsers(namespace string) MessagingUserNamespaceLister
	MessagingUserListerExpansion
}

// messagingUserLister implements the MessagingUserLister interface.
type messagingUserLister struct {
	indexer cache.Indexer
}

// NewMessagingUserLister returns a new MessagingUserLister.
func NewMessagingUserLister(indexer cache.Indexer) MessagingUserLister {
	return &messagingUserLister{indexer: indexer}
}

// List lists all MessagingUsers in the indexer.
func (s *messagingUserLister) List(selector labels.Selector) (ret []*v1alpha1.MessagingUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MessagingUser))
	})
	return ret, err
}

// MessagingUsers returns an object that can list and get MessagingUsers.
func (s *messagingUserLister) MessagingUsers(namespace string) MessagingUserNamespaceLister {
	return messagingUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MessagingUserNamespaceLister helps list and get MessagingUsers.
type MessagingUserNamespaceLister interface {
	// List lists all MessagingUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MessagingUser, err error)
	// Get retrieves the MessagingUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MessagingUser, error)
	MessagingUserNamespaceListerExpansion
}

// messagingUserNamespaceLister implements the MessagingUserNamespaceLister
// interface.
type messagingUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MessagingUsers in the indexer for a given namespace.
func (s messagingUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MessagingUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MessagingUser))
	})
	return ret, err
}

// Get retrieves the MessagingUser from the indexer for a given namespace and name.
func (s messagingUserNamespaceLister) Get(name string) (*v1alpha1.MessagingUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("messaginguser"), name)
	}
	return obj.(*v1alpha1.MessagingUser), nil
}
