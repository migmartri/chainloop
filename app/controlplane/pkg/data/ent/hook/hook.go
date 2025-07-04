// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent"
)

// The APITokenFunc type is an adapter to allow the use of ordinary
// function as APIToken mutator.
type APITokenFunc func(context.Context, *ent.APITokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f APITokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.APITokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.APITokenMutation", m)
}

// The AttestationFunc type is an adapter to allow the use of ordinary
// function as Attestation mutator.
type AttestationFunc func(context.Context, *ent.AttestationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttestationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AttestationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttestationMutation", m)
}

// The CASBackendFunc type is an adapter to allow the use of ordinary
// function as CASBackend mutator.
type CASBackendFunc func(context.Context, *ent.CASBackendMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CASBackendFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CASBackendMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CASBackendMutation", m)
}

// The CASMappingFunc type is an adapter to allow the use of ordinary
// function as CASMapping mutator.
type CASMappingFunc func(context.Context, *ent.CASMappingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CASMappingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CASMappingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CASMappingMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
}

// The GroupMembershipFunc type is an adapter to allow the use of ordinary
// function as GroupMembership mutator.
type GroupMembershipFunc func(context.Context, *ent.GroupMembershipMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GroupMembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMembershipMutation", m)
}

// The IntegrationFunc type is an adapter to allow the use of ordinary
// function as Integration mutator.
type IntegrationFunc func(context.Context, *ent.IntegrationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IntegrationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.IntegrationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IntegrationMutation", m)
}

// The IntegrationAttachmentFunc type is an adapter to allow the use of ordinary
// function as IntegrationAttachment mutator.
type IntegrationAttachmentFunc func(context.Context, *ent.IntegrationAttachmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IntegrationAttachmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.IntegrationAttachmentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IntegrationAttachmentMutation", m)
}

// The MembershipFunc type is an adapter to allow the use of ordinary
// function as Membership mutator.
type MembershipFunc func(context.Context, *ent.MembershipMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MembershipFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MembershipMutation", m)
}

// The OrgInvitationFunc type is an adapter to allow the use of ordinary
// function as OrgInvitation mutator.
type OrgInvitationFunc func(context.Context, *ent.OrgInvitationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgInvitationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrgInvitationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrgInvitationMutation", m)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary
// function as Organization mutator.
type OrganizationFunc func(context.Context, *ent.OrganizationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationMutation", m)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *ent.ProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectMutation", m)
}

// The ProjectVersionFunc type is an adapter to allow the use of ordinary
// function as ProjectVersion mutator.
type ProjectVersionFunc func(context.Context, *ent.ProjectVersionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectVersionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProjectVersionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectVersionMutation", m)
}

// The ReferrerFunc type is an adapter to allow the use of ordinary
// function as Referrer mutator.
type ReferrerFunc func(context.Context, *ent.ReferrerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReferrerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReferrerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReferrerMutation", m)
}

// The RobotAccountFunc type is an adapter to allow the use of ordinary
// function as RobotAccount mutator.
type RobotAccountFunc func(context.Context, *ent.RobotAccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RobotAccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RobotAccountMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RobotAccountMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The WorkflowFunc type is an adapter to allow the use of ordinary
// function as Workflow mutator.
type WorkflowFunc func(context.Context, *ent.WorkflowMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkflowMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowMutation", m)
}

// The WorkflowContractFunc type is an adapter to allow the use of ordinary
// function as WorkflowContract mutator.
type WorkflowContractFunc func(context.Context, *ent.WorkflowContractMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowContractFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkflowContractMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowContractMutation", m)
}

// The WorkflowContractVersionFunc type is an adapter to allow the use of ordinary
// function as WorkflowContractVersion mutator.
type WorkflowContractVersionFunc func(context.Context, *ent.WorkflowContractVersionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowContractVersionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkflowContractVersionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowContractVersionMutation", m)
}

// The WorkflowRunFunc type is an adapter to allow the use of ordinary
// function as WorkflowRun mutator.
type WorkflowRunFunc func(context.Context, *ent.WorkflowRunMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowRunFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WorkflowRunMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowRunMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
