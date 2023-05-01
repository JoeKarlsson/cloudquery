// Code generated by MockGen. DO NOT EDIT.
// Source: cognitoidentityprovider.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cognitoidentityprovider "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	gomock "github.com/golang/mock/gomock"
)

// MockCognitoidentityproviderClient is a mock of CognitoidentityproviderClient interface.
type MockCognitoidentityproviderClient struct {
	ctrl     *gomock.Controller
	recorder *MockCognitoidentityproviderClientMockRecorder
}

// MockCognitoidentityproviderClientMockRecorder is the mock recorder for MockCognitoidentityproviderClient.
type MockCognitoidentityproviderClientMockRecorder struct {
	mock *MockCognitoidentityproviderClient
}

// NewMockCognitoidentityproviderClient creates a new mock instance.
func NewMockCognitoidentityproviderClient(ctrl *gomock.Controller) *MockCognitoidentityproviderClient {
	mock := &MockCognitoidentityproviderClient{ctrl: ctrl}
	mock.recorder = &MockCognitoidentityproviderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCognitoidentityproviderClient) EXPECT() *MockCognitoidentityproviderClientMockRecorder {
	return m.recorder
}

// DescribeIdentityProvider mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeIdentityProvider(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeIdentityProviderInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeIdentityProvider")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentityProvider", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeIdentityProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIdentityProvider indicates an expected call of DescribeIdentityProvider.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeIdentityProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentityProvider", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeIdentityProvider), varargs...)
}

// DescribeResourceServer mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeResourceServer(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeResourceServerInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeResourceServer")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeResourceServer", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeResourceServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeResourceServer indicates an expected call of DescribeResourceServer.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeResourceServer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeResourceServer", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeResourceServer), varargs...)
}

// DescribeRiskConfiguration mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeRiskConfiguration(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeRiskConfigurationInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeRiskConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRiskConfiguration", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeRiskConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRiskConfiguration indicates an expected call of DescribeRiskConfiguration.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeRiskConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRiskConfiguration", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeRiskConfiguration), varargs...)
}

// DescribeUserImportJob mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeUserImportJob(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserImportJobInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUserImportJob")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserImportJob", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserImportJob indicates an expected call of DescribeUserImportJob.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserImportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserImportJob", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserImportJob), varargs...)
}

// DescribeUserPool mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeUserPool(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUserPool")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserPool", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserPool indicates an expected call of DescribeUserPool.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserPool", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPool), varargs...)
}

// DescribeUserPoolClient mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeUserPoolClient(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolClientInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUserPoolClient")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserPoolClient", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolClientOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserPoolClient indicates an expected call of DescribeUserPoolClient.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPoolClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserPoolClient", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPoolClient), varargs...)
}

// DescribeUserPoolDomain mocks base method.
func (m *MockCognitoidentityproviderClient) DescribeUserPoolDomain(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolDomainInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUserPoolDomain")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserPoolDomain", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserPoolDomain indicates an expected call of DescribeUserPoolDomain.
func (mr *MockCognitoidentityproviderClientMockRecorder) DescribeUserPoolDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserPoolDomain", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).DescribeUserPoolDomain), varargs...)
}

// GetCSVHeader mocks base method.
func (m *MockCognitoidentityproviderClient) GetCSVHeader(arg0 context.Context, arg1 *cognitoidentityprovider.GetCSVHeaderInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCSVHeader")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCSVHeader", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetCSVHeaderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCSVHeader indicates an expected call of GetCSVHeader.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetCSVHeader(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCSVHeader", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetCSVHeader), varargs...)
}

// GetDevice mocks base method.
func (m *MockCognitoidentityproviderClient) GetDevice(arg0 context.Context, arg1 *cognitoidentityprovider.GetDeviceInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDevice")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevice", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetDeviceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetDevice(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetDevice), varargs...)
}

// GetGroup mocks base method.
func (m *MockCognitoidentityproviderClient) GetGroup(arg0 context.Context, arg1 *cognitoidentityprovider.GetGroupInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetGroup")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroup", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetGroup), varargs...)
}

// GetIdentityProviderByIdentifier mocks base method.
func (m *MockCognitoidentityproviderClient) GetIdentityProviderByIdentifier(arg0 context.Context, arg1 *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetIdentityProviderByIdentifier")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIdentityProviderByIdentifier", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentityProviderByIdentifier indicates an expected call of GetIdentityProviderByIdentifier.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetIdentityProviderByIdentifier(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentityProviderByIdentifier", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetIdentityProviderByIdentifier), varargs...)
}

// GetSigningCertificate mocks base method.
func (m *MockCognitoidentityproviderClient) GetSigningCertificate(arg0 context.Context, arg1 *cognitoidentityprovider.GetSigningCertificateInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSigningCertificate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSigningCertificate", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetSigningCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSigningCertificate indicates an expected call of GetSigningCertificate.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetSigningCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigningCertificate", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetSigningCertificate), varargs...)
}

// GetUICustomization mocks base method.
func (m *MockCognitoidentityproviderClient) GetUICustomization(arg0 context.Context, arg1 *cognitoidentityprovider.GetUICustomizationInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetUICustomization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUICustomization", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUICustomizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUICustomization indicates an expected call of GetUICustomization.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetUICustomization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUICustomization", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUICustomization), varargs...)
}

// GetUser mocks base method.
func (m *MockCognitoidentityproviderClient) GetUser(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetUser")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUser), varargs...)
}

// GetUserAttributeVerificationCode mocks base method.
func (m *MockCognitoidentityproviderClient) GetUserAttributeVerificationCode(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetUserAttributeVerificationCode")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserAttributeVerificationCode", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttributeVerificationCode indicates an expected call of GetUserAttributeVerificationCode.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetUserAttributeVerificationCode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttributeVerificationCode", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUserAttributeVerificationCode), varargs...)
}

// GetUserPoolMfaConfig mocks base method.
func (m *MockCognitoidentityproviderClient) GetUserPoolMfaConfig(arg0 context.Context, arg1 *cognitoidentityprovider.GetUserPoolMfaConfigInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetUserPoolMfaConfig")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserPoolMfaConfig", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.GetUserPoolMfaConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPoolMfaConfig indicates an expected call of GetUserPoolMfaConfig.
func (mr *MockCognitoidentityproviderClientMockRecorder) GetUserPoolMfaConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPoolMfaConfig", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).GetUserPoolMfaConfig), varargs...)
}

// ListDevices mocks base method.
func (m *MockCognitoidentityproviderClient) ListDevices(arg0 context.Context, arg1 *cognitoidentityprovider.ListDevicesInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDevices")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDevices", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDevices indicates an expected call of ListDevices.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListDevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevices", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListDevices), varargs...)
}

// ListGroups mocks base method.
func (m *MockCognitoidentityproviderClient) ListGroups(arg0 context.Context, arg1 *cognitoidentityprovider.ListGroupsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListGroups), varargs...)
}

// ListIdentityProviders mocks base method.
func (m *MockCognitoidentityproviderClient) ListIdentityProviders(arg0 context.Context, arg1 *cognitoidentityprovider.ListIdentityProvidersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListIdentityProviders")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentityProviders", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListIdentityProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdentityProviders indicates an expected call of ListIdentityProviders.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListIdentityProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentityProviders", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListIdentityProviders), varargs...)
}

// ListResourceServers mocks base method.
func (m *MockCognitoidentityproviderClient) ListResourceServers(arg0 context.Context, arg1 *cognitoidentityprovider.ListResourceServersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResourceServers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceServers", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListResourceServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceServers indicates an expected call of ListResourceServers.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListResourceServers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceServers", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListResourceServers), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCognitoidentityproviderClient) ListTagsForResource(arg0 context.Context, arg1 *cognitoidentityprovider.ListTagsForResourceInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListTagsForResource), varargs...)
}

// ListUserImportJobs mocks base method.
func (m *MockCognitoidentityproviderClient) ListUserImportJobs(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserImportJobsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUserImportJobs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserImportJobs", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserImportJobs indicates an expected call of ListUserImportJobs.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserImportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserImportJobs", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserImportJobs), varargs...)
}

// ListUserPoolClients mocks base method.
func (m *MockCognitoidentityproviderClient) ListUserPoolClients(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserPoolClientsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUserPoolClients")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserPoolClients", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserPoolClientsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPoolClients indicates an expected call of ListUserPoolClients.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserPoolClients(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPoolClients", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserPoolClients), varargs...)
}

// ListUserPools mocks base method.
func (m *MockCognitoidentityproviderClient) ListUserPools(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserPoolsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUserPools")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserPools", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPools indicates an expected call of ListUserPools.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListUserPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPools", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUserPools), varargs...)
}

// ListUsers mocks base method.
func (m *MockCognitoidentityproviderClient) ListUsers(arg0 context.Context, arg1 *cognitoidentityprovider.ListUsersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUsers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUsers), varargs...)
}

// ListUsersInGroup mocks base method.
func (m *MockCognitoidentityproviderClient) ListUsersInGroup(arg0 context.Context, arg1 *cognitoidentityprovider.ListUsersInGroupInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentityprovider.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUsersInGroup")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsersInGroup", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUsersInGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersInGroup indicates an expected call of ListUsersInGroup.
func (mr *MockCognitoidentityproviderClientMockRecorder) ListUsersInGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersInGroup", reflect.TypeOf((*MockCognitoidentityproviderClient)(nil).ListUsersInGroup), varargs...)
}
