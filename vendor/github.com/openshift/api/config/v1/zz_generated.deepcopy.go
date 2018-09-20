// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionPluginConfig) DeepCopyInto(out *AdmissionPluginConfig) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionPluginConfig.
func (in *AdmissionPluginConfig) DeepCopy() *AdmissionPluginConfig {
	if in == nil {
		return nil
	}
	out := new(AdmissionPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditConfig) DeepCopyInto(out *AuditConfig) {
	*out = *in
	in.PolicyConfiguration.DeepCopyInto(&out.PolicyConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditConfig.
func (in *AuditConfig) DeepCopy() *AuditConfig {
	if in == nil {
		return nil
	}
	out := new(AuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertInfo) DeepCopyInto(out *CertInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertInfo.
func (in *CertInfo) DeepCopy() *CertInfo {
	if in == nil {
		return nil
	}
	out := new(CertInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionOverrides) DeepCopyInto(out *ClientConnectionOverrides) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionOverrides.
func (in *ClientConnectionOverrides) DeepCopy() *ClientConnectionOverrides {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConnectionInfo) DeepCopyInto(out *EtcdConnectionInfo) {
	*out = *in
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CertInfo = in.CertInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConnectionInfo.
func (in *EtcdConnectionInfo) DeepCopy() *EtcdConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(EtcdConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdStorageConfig) DeepCopyInto(out *EtcdStorageConfig) {
	*out = *in
	in.EtcdConnectionInfo.DeepCopyInto(&out.EtcdConnectionInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdStorageConfig.
func (in *EtcdStorageConfig) DeepCopy() *EtcdStorageConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericAPIServerConfig) DeepCopyInto(out *GenericAPIServerConfig) {
	*out = *in
	in.ServingInfo.DeepCopyInto(&out.ServingInfo)
	if in.CORSAllowedOrigins != nil {
		in, out := &in.CORSAllowedOrigins, &out.CORSAllowedOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AuditConfig.DeepCopyInto(&out.AuditConfig)
	in.StorageConfig.DeepCopyInto(&out.StorageConfig)
	if in.AdmissionPluginConfig != nil {
		in, out := &in.AdmissionPluginConfig, &out.AdmissionPluginConfig
		*out = make(map[string]AdmissionPluginConfig, len(*in))
		for key, val := range *in {
			newVal := new(AdmissionPluginConfig)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	out.KubeClientConfig = in.KubeClientConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericAPIServerConfig.
func (in *GenericAPIServerConfig) DeepCopy() *GenericAPIServerConfig {
	if in == nil {
		return nil
	}
	out := new(GenericAPIServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPServingInfo) DeepCopyInto(out *HTTPServingInfo) {
	*out = *in
	in.ServingInfo.DeepCopyInto(&out.ServingInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPServingInfo.
func (in *HTTPServingInfo) DeepCopy() *HTTPServingInfo {
	if in == nil {
		return nil
	}
	out := new(HTTPServingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeClientConfig) DeepCopyInto(out *KubeClientConfig) {
	*out = *in
	out.ConnectionOverrides = in.ConnectionOverrides
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeClientConfig.
func (in *KubeClientConfig) DeepCopy() *KubeClientConfig {
	if in == nil {
		return nil
	}
	out := new(KubeClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElection) DeepCopyInto(out *LeaderElection) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElection.
func (in *LeaderElection) DeepCopy() *LeaderElection {
	if in == nil {
		return nil
	}
	out := new(LeaderElection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedCertificate) DeepCopyInto(out *NamedCertificate) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CertInfo = in.CertInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedCertificate.
func (in *NamedCertificate) DeepCopy() *NamedCertificate {
	if in == nil {
		return nil
	}
	out := new(NamedCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteConnectionInfo) DeepCopyInto(out *RemoteConnectionInfo) {
	*out = *in
	out.CertInfo = in.CertInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteConnectionInfo.
func (in *RemoteConnectionInfo) DeepCopy() *RemoteConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(RemoteConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingInfo) DeepCopyInto(out *ServingInfo) {
	*out = *in
	out.CertInfo = in.CertInfo
	if in.NamedCertificates != nil {
		in, out := &in.NamedCertificates, &out.NamedCertificates
		*out = make([]NamedCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CipherSuites != nil {
		in, out := &in.CipherSuites, &out.CipherSuites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingInfo.
func (in *ServingInfo) DeepCopy() *ServingInfo {
	if in == nil {
		return nil
	}
	out := new(ServingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringSource) DeepCopyInto(out *StringSource) {
	*out = *in
	out.StringSourceSpec = in.StringSourceSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringSource.
func (in *StringSource) DeepCopy() *StringSource {
	if in == nil {
		return nil
	}
	out := new(StringSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringSourceSpec) DeepCopyInto(out *StringSourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringSourceSpec.
func (in *StringSourceSpec) DeepCopy() *StringSourceSpec {
	if in == nil {
		return nil
	}
	out := new(StringSourceSpec)
	in.DeepCopyInto(out)
	return out
}
