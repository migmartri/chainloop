//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import (
	"context"
	"fmt"
	"slices"

	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/referrer"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type ReferrerRepo struct {
	data         *Data
	log          *log.Helper
	workflowRepo biz.WorkflowRepo
}

func NewReferrerRepo(data *Data, wfRepo biz.WorkflowRepo, logger log.Logger) biz.ReferrerRepo {
	return &ReferrerRepo{
		data:         data,
		log:          log.NewHelper(logger),
		workflowRepo: wfRepo,
	}
}

type storedReferrerMap map[string]uuid.UUID

func (r *ReferrerRepo) Save(ctx context.Context, referrers []*biz.Referrer, workflowID uuid.UUID) (err error) {
	// find the workflow
	wf, err := r.workflowRepo.FindByID(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to find workflow: %w", err)
	} else if wf == nil {
		return biz.NewErrNotFound("workflow")
	}

	storedMap := make(storedReferrerMap)

	for _, ref := range referrers {
		// Check if it exists already, if not create it
		storedID, err := r.data.DB.Referrer.Create().
			SetDigest(ref.Digest).SetKind(ref.Kind).SetDownloadable(ref.Downloadable).
			SetMetadata(ref.Metadata).SetAnnotations(ref.Annotations).
			AddWorkflowIDs(workflowID).
			OnConflictColumns(
				referrer.FieldDigest, referrer.FieldKind,
			).UpdateNewValues().ID(ctx)
		if err != nil {
			return fmt.Errorf("failed to create referrer: %w", err)
		}

		storedRef, err := r.data.DB.Referrer.Query().Select(referrer.FieldID).Where(referrer.ID(storedID)).First(ctx)
		if err != nil {
			return fmt.Errorf("failed to load referrer: %w", err)
		} else if storedRef == nil {
			return fmt.Errorf("failed to load referrer: %w", err)
		}

		// Store it in the map
		storedMap[ref.MapID()] = storedRef.ID
	}

	// 2 - define the relationship between referrers
	for _, parentRef := range referrers {
		// This is the current item stored in DB
		storedReferrer := storedMap[parentRef.MapID()]
		// Iterate on the items it refer to (references)
		var references []uuid.UUID
		for _, ref := range parentRef.References {
			// amd find it in the DB
			storedReference, ok := storedMap[ref.MapID()]
			if !ok {
				return fmt.Errorf("referrer %v not found", ref)
			}

			references = append(references, storedReference)
		}

		if len(references) == 0 {
			continue
		}

		// Create the relationship
		if err := r.data.DB.Referrer.UpdateOneID(storedReferrer).AddReferenceIDs(references...).Exec(ctx); err != nil {
			return fmt.Errorf("failed to create referrer relationship: %w", err)
		}
	}

	return nil
}

// Check if a given referrer by digest exist. The query can be scoped further down if needed by providing the kind or visibility status
func (r *ReferrerRepo) Exist(ctx context.Context, digest string, filters ...biz.GetFromRootFilter) (bool, error) {
	opts := &biz.GetFromRootFilters{}
	for _, f := range filters {
		f(opts)
	}

	query := r.data.DB.Referrer.Query().Where(referrer.DigestEQ(digest))
	// We might be filtering by the rootKind, this will prevent ambiguity
	if opts.RootKind != nil {
		query = query.Where(referrer.Kind(*opts.RootKind))
	}

	if opts.Public != nil {
		query = query.WithWorkflows(func(q *ent.WorkflowQuery) { q.Where(workflow.PublicEQ(*opts.Public)) })
	}

	return query.Exist(ctx)
}

func (r *ReferrerRepo) GetFromRoot(ctx context.Context, digest string, orgIDs []uuid.UUID, filters ...biz.GetFromRootFilter) (*biz.StoredReferrer, error) {
	opts := &biz.GetFromRootFilters{}
	for _, f := range filters {
		f(opts)
	}

	// Find the referrer from its digest + artifactType (optional)
	// if there is more than 1 item we return ReferrerAmbiguous error
	// filter by the allowed organizations and by the visibility of the attached workflows if needed
	predicateReferrer := []predicate.Referrer{
		referrer.Digest(digest),
	}

	// We might be filtering by the rootKind, this will prevent ambiguity
	if opts.RootKind != nil {
		predicateReferrer = append(predicateReferrer, referrer.Kind(*opts.RootKind))
	}

	// Prepare the workflow query predicate
	predicateWF := []predicate.Workflow{
		workflow.DeletedAtIsNil(), workflow.HasOrganizationWith(organization.IDIn(orgIDs...)),
	}

	// optionally attaching its visibility
	if opts.Public != nil {
		predicateWF = append(predicateWF, workflow.Public(*opts.Public))
	}

	// Attach the workflow predicate
	predicateReferrer = append(predicateReferrer, referrer.HasWorkflowsWith(predicateWF...))

	refs, err := r.data.DB.Referrer.Query().Where(predicateReferrer...).WithWorkflows().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query referrer: %w", err)
	}

	// No items found
	if numrefs := len(refs); numrefs == 0 {
		return nil, nil
	} else if numrefs > 1 {
		// if there is more than 1 item with the same digest+artifactType we will fail
		var kinds []string
		for _, r := range refs {
			kinds = append(kinds, r.Kind)
		}
		return nil, biz.NewErrReferrerAmbiguous(digest, kinds)
	}

	// Find the referrer recursively starting from the root
	res, err := r.doGet(ctx, refs[0], orgIDs, opts.ProjectIDs, opts.Public, 0)
	if err != nil && !biz.IsErrUnauthorized(err) {
		return nil, fmt.Errorf("failed to get referrer: %w", err)
	}

	return res, nil
}

// max number of recursive levels to traverse
// we just care about 1 level, i.e att -> commit, or commit -> attestation
// we also need to limit this because there might be cycles
const maxTraverseLevels = 1

func (r *ReferrerRepo) doGet(ctx context.Context, root *ent.Referrer, allowedOrgs []uuid.UUID, visibleProjectsMap map[uuid.UUID][]uuid.UUID, public *bool, level int) (*biz.StoredReferrer, error) {
	// Assemble the referrer to return
	res := &biz.StoredReferrer{
		ID:        root.ID,
		CreatedAt: toTimePtr(root.CreatedAt),
		Referrer: &biz.Referrer{
			Digest:       root.Digest,
			Kind:         root.Kind,
			Downloadable: root.Downloadable,
			Metadata:     root.Metadata,
			Annotations:  root.Annotations,
		},
	}

	// add additional information related to the workflows
	hydrateWorkflowsInfo(root, res)

	// check that, if RBAC is required, the user has visibility on the artifact in at least 1 org/project
	if visible := isReferrerVisible(res, allowedOrgs, visibleProjectsMap); !visible {
		return nil, biz.NewErrUnauthorizedStr("referrer not allowed")
	}

	// Next: We'll find the references recursively up to a max of maxTraverseLevels levels
	if level >= maxTraverseLevels {
		return res, nil
	}

	// Find the references and call recursively filtered out by the allowed organizations
	// and by the visibility if needed
	predicateReferrer := []predicate.Referrer{}

	predicateWF := []predicate.Workflow{
		workflow.DeletedAtIsNil(), workflow.HasOrganizationWith(organization.IDIn(allowedOrgs...)),
	}

	// optionally attaching its visibility
	if public != nil {
		predicateWF = append(predicateWF, workflow.Public(*public))
	}

	// Attach the workflow predicate
	predicateReferrer = append(predicateReferrer, referrer.HasWorkflowsWith(predicateWF...))

	// sort the references by creation date in descending order
	// so whenever we add pagination we'll get the latest x references
	refs, err := root.QueryReferences().Where(predicateReferrer...).WithWorkflows().Order(referrer.ByCreatedAt(), ent.Desc()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query references: %w", err)
	}

	// Add the references to the result
	for _, reference := range refs {
		// Call recursively the function
		// we return all the references
		ref, err := r.doGet(ctx, reference, allowedOrgs, visibleProjectsMap, public, level+1)
		if err != nil && !biz.IsErrUnauthorized(err) {
			return nil, fmt.Errorf("failed to get referrer: %w", err)
		}

		if ref != nil {
			res.References = append(res.References, ref)
		}
	}

	return res, nil
}

func isReferrerVisible(ref *biz.StoredReferrer, allowedOrgs []uuid.UUID, visibleProjectsMap map[uuid.UUID][]uuid.UUID) bool {
	if ref.InPublicWorkflow {
		return true
	}

	for _, oid := range ref.OrgIDs {
		if !slices.Contains(allowedOrgs, oid) {
			// skip check in organizations where the user doesn't have access
			continue
		}
		if visibleProjects, ok := visibleProjectsMap[oid]; ok {
			// if entry is present, it means we need to apply RBAC
			// check if visible projects and referrer projects match
			// by checking if any project is visible by the user
			for _, pid := range ref.ProjectIDs {
				if slices.Contains(visibleProjects, pid) {
					return true
				}
			}
		} else {
			// if entry is not found in the map, it means that RBAC is not needed for this org, we have finished
			return true
		}
	}

	return false
}

// hydrate the referrer with the following information:
// - isPublic: if it has a public workflow associated
// - workflowIDs: the list of associated workflows
// - orgIDs: the list of associated organizations
func hydrateWorkflowsInfo(root *ent.Referrer, out *biz.StoredReferrer) {
	isPublic := false
	workflowIDs := make([]uuid.UUID, 0, len(root.Edges.Workflows))
	projectIDs := make(map[uuid.UUID]bool, 0)
	orgIDs := make([]uuid.UUID, 0)
	orgsMap := make(map[uuid.UUID]struct{}, 0)
	for _, wf := range root.Edges.Workflows {
		if wf.Public {
			isPublic = true
		}
		workflowIDs = append(workflowIDs, wf.ID)
		if _, ok := orgsMap[wf.OrganizationID]; !ok {
			orgIDs = append(orgIDs, wf.OrganizationID)
		}
		if _, ok := projectIDs[wf.ProjectID]; !ok {
			projectIDs[wf.ProjectID] = true
		}
		orgsMap[wf.OrganizationID] = struct{}{}
	}

	out.ProjectIDs = maps.Keys(projectIDs)
	out.InPublicWorkflow = isPublic
	out.WorkflowIDs = workflowIDs
	out.OrgIDs = orgIDs
}
