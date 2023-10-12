package pki

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	PrefixRoot = "root/"
)

type Storage struct {
	etcd *clientv3.Client
}

// NewStorage
func NewStorage(endpoints []string) (*Storage, error) {
	var err error
	manage := &Storage{}
	// connect etcd
	if manage.etcd, err = clientv3.New(clientv3.Config{Endpoints: endpoints}); err != nil {
		return nil, err
	}
	return manage, nil
}

// Put put Certificate into etcd
func (s *Storage) Put(CertificateId string, Certificate string) error {
	_, err := s.etcd.Put(context.TODO(), CertificateId, Certificate)
	return err
}

// Get get Certificate from etcd
func (s *Storage) Get(CertificateId string) (string, error) {
	rsp, err := s.etcd.Get(context.TODO(), CertificateId)
	if err != nil {
		return "", err
	}
	if len(rsp.Kvs) == 0 {
		return "", nil
	}
	return string(rsp.Kvs[0].Value), err
}

// Delete delete Certificate from etcd
func (s *Storage) Delete(CertificateId string) error {
	_, err := s.etcd.Delete(context.TODO(), CertificateId)
	return err
}

// Close close etcd connection
func (s *Storage) Close() error {
	return s.etcd.Close()
}
