package pki

import (
	"context"
	"encoding/json"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type CertificateManage struct {
	etcd *clientv3.Client
}

// NewCertificateManage
func NewCertificateManage(endpoints []string) (*CertificateManage, error) {
	var err error
	manage := &CertificateManage{}
	// connect etcd
	if manage.etcd, err = clientv3.New(clientv3.Config{Endpoints: endpoints}); err != nil {
		return nil, err
	}
	return manage, nil
}

// Put put Certificate into etcd
func (crtm *CertificateManage) Put(CertificateId string, Certificate Certificate) error {

	jsonBytes, err := json.Marshal(&Certificate)
	if err != nil {
		return err
	}
	_, err = crtm.etcd.Put(context.TODO(), CertificateId, string(jsonBytes))
	return err
}

// Get get Certificate from etcd
func (crtm *CertificateManage) Get(CertificateId string) (Certificate, error) {
	var Certificate Certificate
	rsp, err := crtm.etcd.Get(context.TODO(), CertificateId)
	if err != nil {
		return Certificate, err
	}
	if len(rsp.Kvs) == 0 {
		return Certificate, nil
	}
	err = json.Unmarshal(rsp.Kvs[0].Value, &Certificate)
	return Certificate, err
}
