/*
Copyright 2025.

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

package v1alpha1

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GatherProviderSecrets selects unique secrets based on ProviderSelectors
func (s *SopsProvider) GatherProviderSecrets(ctx context.Context, client client.Client) ([]corev1.Secret, error) {
	// 1️⃣ Fetch all Secrets in a single API call
	secretList := &corev1.SecretList{}
	if err := client.List(ctx, secretList); err != nil {
		return nil, fmt.Errorf("failed to list secrets: %w", err)
	}

	// Prepare a map to store unique secrets
	uniqueSecrets := make(map[string]*corev1.Secret)

	// 2️⃣ Loop through ProviderSecrets and filter relevant secrets
	for _, selector := range s.Spec.ProviderSecrets {
		if selector == nil || selector.NamespacedSelector == nil {
			continue
		}

		// ✅ Use optimized `MatchObjects()` to filter secrets
		matchingSecrets, err := selector.NamespacedSelector.MatchObjects(ctx, client, toObjectList(secretList.Items))
		if err != nil {
			return nil, fmt.Errorf("error matching secrets: %w", err)
		}

		// ✅ Store unique secrets in a map
		for _, sec := range matchingSecrets {
			secret, ok := sec.(*corev1.Secret)
			if ok {
				uniqueSecrets[secret.Namespace+"/"+secret.Name] = secret
			}
		}
	}

	// 3️⃣ Convert map to a list of unique secrets
	finalSecrets := make([]corev1.Secret, 0, len(uniqueSecrets))
	for _, sec := range uniqueSecrets {
		finalSecrets = append(finalSecrets, *sec)
	}

	return finalSecrets, nil
}

// Helper function to convert []corev1.Secret to []metav1.Object
func toObjectList(secrets []corev1.Secret) []metav1.Object {
	objectList := make([]metav1.Object, len(secrets))
	for i := range secrets {
		objectList[i] = &secrets[i]
	}
	return objectList
}
