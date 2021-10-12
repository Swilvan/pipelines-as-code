package repository

import (
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis/duck/v1beta1"
)

type RepoTestcreationOpts struct {
	Name             string
	URL              string
	Branch           string
	InstallNamespace string
	EventType        string
	SecretName       string
	VcsURL           string
	CreateTime       metav1.Time
}

func NewRepo(opts RepoTestcreationOpts) *v1alpha1.Repository {
	cw := clockwork.NewFakeClock()
	repo := &v1alpha1.Repository{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:              opts.Name,
			Namespace:         opts.InstallNamespace,
			CreationTimestamp: opts.CreateTime,
		},
		Spec: v1alpha1.RepositorySpec{
			URL:       opts.URL,
			Branch:    opts.Branch,
			EventType: opts.EventType,
		},
		Status: []v1alpha1.RepositoryRunStatus{
			{
				Status:          v1beta1.Status{},
				PipelineRunName: "pipelinerun5",
				StartTime:       &metav1.Time{Time: cw.Now().Add(-56 * time.Minute)},
				CompletionTime:  &metav1.Time{Time: cw.Now().Add(-55 * time.Minute)},
			},
			{
				Status:          v1beta1.Status{},
				PipelineRunName: "pipelinerun4",
				StartTime:       &metav1.Time{Time: cw.Now().Add(-46 * time.Minute)},
				CompletionTime:  &metav1.Time{Time: cw.Now().Add(-45 * time.Minute)},
			},
			{
				Status:          v1beta1.Status{},
				PipelineRunName: "pipelinerun3",
				StartTime:       &metav1.Time{Time: cw.Now().Add(-36 * time.Minute)},
				CompletionTime:  &metav1.Time{Time: cw.Now().Add(-35 * time.Minute)},
			},
			{
				Status:          v1beta1.Status{},
				PipelineRunName: "pipelinerun2",
				StartTime:       &metav1.Time{Time: cw.Now().Add(-26 * time.Minute)},
				CompletionTime:  &metav1.Time{Time: cw.Now().Add(-25 * time.Minute)},
			},
			{
				Status:          v1beta1.Status{},
				PipelineRunName: "pipelinerun1",
				StartTime:       &metav1.Time{Time: cw.Now().Add(-16 * time.Minute)},
				CompletionTime:  &metav1.Time{Time: cw.Now().Add(-15 * time.Minute)},
			},
		},
	}
	if opts.SecretName != "" {
		repo.Spec.WebvcsSecret = &v1alpha1.WebvcsSecretSpec{
			Name: opts.SecretName,
		}
	}
	if opts.VcsURL != "" {
		repo.Spec.WebvcsAPIURL = opts.VcsURL
	}
	return repo
}
