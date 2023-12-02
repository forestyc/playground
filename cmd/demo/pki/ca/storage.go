package ca

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	PrefixRoot     = "root/"
	PrefixMiddle   = "middle/"
	PrefixTerminal = "terminal/"
)

type Storage struct {
	etcd *clientv3.Client
}

func NewStorage(endpoints []string) (*Storage, error) {
	var err error
	manage := &Storage{}
	// connect etcd
	if manage.etcd, err = clientv3.New(clientv3.Config{Endpoints: endpoints}); err != nil {
		return nil, err
	}
	return manage, nil
}

// Put Certificate into etcd
func (s *Storage) Put(CertificateId string, Certificate string, option ...clientv3.OpOption) error {
	_, err := s.etcd.Put(context.TODO(), CertificateId, Certificate, option...)
	return err
}

// Get Certificate from etcd
func (s *Storage) Get(CertificateId string, option ...clientv3.OpOption) ([]*mvccpb.KeyValue, error) {
	rsp, err := s.etcd.Get(context.TODO(), CertificateId, option...)
	if err != nil {
		return nil, err
	}
	return rsp.Kvs, nil
}

// Delete delete Certificate from etcd
func (s *Storage) Delete(CertificateId string, option ...clientv3.OpOption) error {
	_, err := s.etcd.Delete(context.TODO(), CertificateId, option...)
	return err
}

// Close etcd connection
func (s *Storage) Close() error {
	return s.etcd.Close()
}
