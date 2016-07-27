/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"fmt"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = framework.KubeDescribe("Downward API volume", func() {
	// How long to wait for a log pod to be displayed
	const podLogTimeout = 45 * time.Second

	f := framework.NewDefaultFramework("downward-api")
	It("should provide podname only [Conformance]", func() {
		podName := "downwardapi-volume-" + string(util.NewUUID())
		pod := downwardAPIVolumePodForSimpleTest(podName, "/etc/podname")

		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("%s\n", podName),
		}, f.Namespace.Name)
	})

	It("should provide podname as non-root with fsgroup [Feature:FSGroup]", func() {
		podName := "metadata-volume-" + string(util.NewUUID())
		uid := int64(1001)
		gid := int64(1234)
		pod := downwardAPIVolumePodForSimpleTest(podName, "/etc/podname")
		pod.Spec.SecurityContext = &api.PodSecurityContext{
			RunAsUser: &uid,
			FSGroup:   &gid,
		}
		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("%s\n", podName),
		}, f.Namespace.Name)
	})

	It("should update labels on modification [Conformance]", func() {
		labels := map[string]string{}
		labels["key1"] = "value1"
		labels["key2"] = "value2"

		podName := "labelsupdate" + string(util.NewUUID())
		pod := downwardAPIVolumePodForUpdateTest(podName, labels, map[string]string{}, "/etc/labels")
		containerName := "client-container"
		defer func() {
			By("Deleting the pod")
			f.Client.Pods(f.Namespace.Name).Delete(pod.Name, api.NewDeleteOptions(0))
		}()
		By("Creating the pod")
		_, err := f.Client.Pods(f.Namespace.Name).Create(pod)
		Expect(err).NotTo(HaveOccurred())

		framework.ExpectNoError(framework.WaitForPodRunningInNamespace(f.Client, pod.Name, f.Namespace.Name))

		pod, err = f.Client.Pods(f.Namespace.Name).Get(pod.Name)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() (string, error) {
			return framework.GetPodLogs(f.Client, f.Namespace.Name, pod.Name, containerName)
		},
			podLogTimeout, framework.Poll).Should(ContainSubstring("key1=\"value1\"\n"))

		//modify labels
		pod.Labels["key3"] = "value3"
		pod.ResourceVersion = "" // to force update
		_, err = f.Client.Pods(f.Namespace.Name).Update(pod)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() (string, error) {
			return framework.GetPodLogs(f.Client, f.Namespace.Name, pod.Name, containerName)
		},
			podLogTimeout, framework.Poll).Should(ContainSubstring("key3=\"value3\"\n"))

	})

	It("should update annotations on modification [Conformance]", func() {
		annotations := map[string]string{}
		annotations["builder"] = "bar"
		podName := "annotationupdate" + string(util.NewUUID())
		pod := downwardAPIVolumePodForUpdateTest(podName, map[string]string{}, annotations, "/etc/annotations")

		containerName := "client-container"
		defer func() {
			By("Deleting the pod")
			f.Client.Pods(f.Namespace.Name).Delete(pod.Name, api.NewDeleteOptions(0))
		}()
		By("Creating the pod")
		_, err := f.Client.Pods(f.Namespace.Name).Create(pod)
		Expect(err).NotTo(HaveOccurred())
		framework.ExpectNoError(framework.WaitForPodRunningInNamespace(f.Client, pod.Name, f.Namespace.Name))

		pod, err = f.Client.Pods(f.Namespace.Name).Get(pod.Name)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() (string, error) {
			return framework.GetPodLogs(f.Client, f.Namespace.Name, pod.Name, containerName)
		},
			podLogTimeout, framework.Poll).Should(ContainSubstring("builder=\"bar\"\n"))

		//modify annotations
		pod.Annotations["builder"] = "foo"
		pod.ResourceVersion = "" // to force update
		_, err = f.Client.Pods(f.Namespace.Name).Update(pod)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() (string, error) {
			return framework.GetPodLogs(f.Client, f.Namespace.Name, pod.Name, containerName)
		},
			podLogTimeout, framework.Poll).Should(ContainSubstring("builder=\"foo\"\n"))

	})

	It("should provide container's cpu limit", func() {
		podName := "downwardapi-volume-" + string(util.NewUUID())
		pod := downwardAPIVolumeForContainerResources(podName, "/etc/cpu_limit")

		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("2\n"),
		}, f.Namespace.Name)
	})

	It("should provide container's memory limit", func() {
		podName := "downwardapi-volume-" + string(util.NewUUID())
		pod := downwardAPIVolumeForContainerResources(podName, "/etc/memory_limit")

		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("67108864\n"),
		}, f.Namespace.Name)
	})

	It("should provide container's cpu request", func() {
		podName := "downwardapi-volume-" + string(util.NewUUID())
		pod := downwardAPIVolumeForContainerResources(podName, "/etc/cpu_request")

		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("1\n"),
		}, f.Namespace.Name)
	})

	It("should provide container's memory request", func() {
		podName := "downwardapi-volume-" + string(util.NewUUID())
		pod := downwardAPIVolumeForContainerResources(podName, "/etc/memory_request")

		framework.TestContainerOutput("downward API volume plugin", f.Client, pod, 0, []string{
			fmt.Sprintf("33554432\n"),
		}, f.Namespace.Name)
	})

})

func downwardAPIVolumePodForSimpleTest(name string, filePath string) *api.Pod {
	pod := downwardAPIVolumeBasePod(name, nil, nil)

	pod.Spec.Containers = []api.Container{
		{
			Name:    "client-container",
			Image:   "gcr.io/google_containers/mounttest:0.6",
			Command: []string{"/mt", "--file_content=" + filePath},
			VolumeMounts: []api.VolumeMount{
				{
					Name:      "podinfo",
					MountPath: "/etc",
					ReadOnly:  false,
				},
			},
		},
	}

	return pod
}

func downwardAPIVolumeForContainerResources(name string, filePath string) *api.Pod {
	pod := downwardAPIVolumeBasePod(name, nil, nil)
	pod.Spec.Containers = downwardAPIVolumeBaseContainers("client-container", filePath)
	return pod
}

func downwardAPIVolumeBaseContainers(name, filePath string) []api.Container {
	return []api.Container{
		{
			Name:    name,
			Image:   "gcr.io/google_containers/mounttest:0.6",
			Command: []string{"/mt", "--file_content=" + filePath},
			Resources: api.ResourceRequirements{
				Requests: api.ResourceList{
					api.ResourceCPU:    resource.MustParse("250m"),
					api.ResourceMemory: resource.MustParse("32Mi"),
				},
				Limits: api.ResourceList{
					api.ResourceCPU:    resource.MustParse("1250m"),
					api.ResourceMemory: resource.MustParse("64Mi"),
				},
			},
			VolumeMounts: []api.VolumeMount{
				{
					Name:      "podinfo",
					MountPath: "/etc",
					ReadOnly:  false,
				},
			},
		},
	}

}

func downwardAPIVolumePodForUpdateTest(name string, labels, annotations map[string]string, filePath string) *api.Pod {
	pod := downwardAPIVolumeBasePod(name, labels, annotations)

	pod.Spec.Containers = []api.Container{
		{
			Name:    "client-container",
			Image:   "gcr.io/google_containers/mounttest:0.6",
			Command: []string{"/mt", "--break_on_expected_content=false", "--retry_time=120", "--file_content_in_loop=" + filePath},
			VolumeMounts: []api.VolumeMount{
				{
					Name:      "podinfo",
					MountPath: "/etc",
					ReadOnly:  false,
				},
			},
		},
	}

	applyLabelsAndAnnotationsToDownwardAPIPod(labels, annotations, pod)
	return pod
}

func downwardAPIVolumeBasePod(name string, labels, annotations map[string]string) *api.Pod {
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:        name,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: api.PodSpec{
			Volumes: []api.Volume{
				{
					Name: "podinfo",
					VolumeSource: api.VolumeSource{
						DownwardAPI: &api.DownwardAPIVolumeSource{
							Items: []api.DownwardAPIVolumeFile{
								{
									Path: "podname",
									FieldRef: &api.ObjectFieldSelector{
										APIVersion: "v1",
										FieldPath:  "metadata.name",
									},
								},
								{
									Path: "cpu_limit",
									ResourceFieldRef: &api.ResourceFieldSelector{
										ContainerName: "client-container",
										Resource:      "limits.cpu",
									},
								},
								{
									Path: "cpu_request",
									ResourceFieldRef: &api.ResourceFieldSelector{
										ContainerName: "client-container",
										Resource:      "requests.cpu",
									},
								},
								{
									Path: "memory_limit",
									ResourceFieldRef: &api.ResourceFieldSelector{
										ContainerName: "client-container",
										Resource:      "limits.memory",
									},
								},
								{
									Path: "memory_request",
									ResourceFieldRef: &api.ResourceFieldSelector{
										ContainerName: "client-container",
										Resource:      "requests.memory",
									},
								},
							},
						},
					},
				},
			},
			RestartPolicy: api.RestartPolicyNever,
		},
	}

	return pod
}

func applyLabelsAndAnnotationsToDownwardAPIPod(labels, annotations map[string]string, pod *api.Pod) {
	if len(labels) > 0 {
		pod.Spec.Volumes[0].DownwardAPI.Items = append(pod.Spec.Volumes[0].DownwardAPI.Items, api.DownwardAPIVolumeFile{
			Path: "labels",
			FieldRef: &api.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "metadata.labels",
			},
		})
	}

	if len(annotations) > 0 {
		pod.Spec.Volumes[0].DownwardAPI.Items = append(pod.Spec.Volumes[0].DownwardAPI.Items, api.DownwardAPIVolumeFile{
			Path: "annotations",
			FieldRef: &api.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "metadata.annotations",
			},
		})
	}
}

// TODO: add test-webserver example as pointed out in https://github.com/kubernetes/kubernetes/pull/5093#discussion-diff-37606771
