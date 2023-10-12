package pki

type Pki struct {
	manage *CertificateManage
}

// NewPki
func NewPki(endpoints []string) (*Pki, error) {
	var err error
	pki := &Pki{}
	if pki.manage, err = NewCertificateManage(endpoints); err != nil {
		return nil, err
	}
	return pki, nil
}

func (pki Pki) NewRootCertificate(name string) {

}
