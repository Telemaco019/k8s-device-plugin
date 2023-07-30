package main

import "github.com/NVIDIA/k8s-device-plugin/internal/rm"

type MPSDevice struct {
	AnnotatedID rm.MPSAnnotatedID
	Index       string
}

type MPSDeviceList []MPSDevice

// Subset returns the subset of MPS devices in MPSDeviceList matching the provided ids.
func (m MPSDeviceList) Subset(ids []string) MPSDeviceList {
	res := make(MPSDeviceList, 0)
	for _, device := range m {
		for _, id := range ids {
			if id == device.AnnotatedID.String() {
				res = append(res, device)
			}
		}
	}
	return res
}
