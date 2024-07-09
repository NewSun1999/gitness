// Copyright 2023 Harness, Inc.
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

package gitspace

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	apiauth "github.com/harness/gitness/app/api/auth"
	"github.com/harness/gitness/app/auth"
	events "github.com/harness/gitness/app/events/gitspace"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"
	"github.com/harness/gitness/types/enum"

	gonanoid "github.com/matoous/go-nanoid"
)

const defaultAccessKey = "Harness@123"
const defaultMachineUser = "harness"

type ActionInput struct {
	Action     enum.GitspaceActionType `json:"action"`
	Identifier string                  `json:"-"`
	SpaceRef   string                  `json:"-"` // Ref of the parent space
}

func (c *Controller) Action(
	ctx context.Context,
	session *auth.Session,
	in *ActionInput,
) (*types.GitspaceConfig, error) {
	if err := c.sanitizeActionInput(in); err != nil {
		return nil, fmt.Errorf("failed to sanitize input: %w", err)
	}
	space, err := c.spaceStore.FindByRef(ctx, in.SpaceRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find space: %w", err)
	}
	err = apiauth.CheckGitspace(ctx, c.authorizer, session, space.Path, in.Identifier, enum.PermissionGitspaceAccess)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize: %w", err)
	}
	gitspaceConfig, err := c.gitspaceConfigStore.FindByIdentifier(ctx, space.ID, in.Identifier)
	gitspaceConfig.SpacePath = space.Path
	gitspaceConfig.SpaceID = space.ID
	if err != nil {
		return nil, fmt.Errorf("failed to find gitspace config: %w", err)
	}
	// All the actions should be idempotent.
	switch in.Action {
	case enum.GitspaceActionTypeStart:
		c.emitGitspaceConfigEvent(ctx, gitspaceConfig, enum.GitspaceEventTypeGitspaceActionStart)
		gitspace, err := c.startGitspace(ctx, gitspaceConfig)
		if err != nil {
			c.emitGitspaceConfigEvent(ctx, gitspaceConfig, enum.GitspaceEventTypeGitspaceActionStartFailed)
		}
		return gitspace, err
	case enum.GitspaceActionTypeStop:
		c.emitGitspaceConfigEvent(ctx, gitspaceConfig, enum.GitspaceEventTypeGitspaceActionStop)
		gitspace, err := c.stopGitspace(ctx, gitspaceConfig)
		if err != nil {
			c.emitGitspaceConfigEvent(ctx, gitspaceConfig, enum.GitspaceEventTypeGitspaceActionStopFailed)
		}
		return gitspace, err
	default:
		return nil, fmt.Errorf("unknown action %s on gitspace : %s", string(in.Action), gitspaceConfig.Identifier)
	}
}

func (c *Controller) startGitspace(ctx context.Context, config *types.GitspaceConfig) (*types.GitspaceConfig, error) {
	savedGitspaceInstance, err := c.gitspaceInstanceStore.FindLatestByGitspaceConfigID(ctx, config.ID, config.SpaceID)
	const resourceNotFoundErr = "Failed to find gitspace: resource not found"
	if err != nil && err.Error() != resourceNotFoundErr { // TODO fix this
		return nil, fmt.Errorf("failed to find gitspace with config ID : %s %w", config.Identifier, err)
	}
	config.GitspaceInstance = savedGitspaceInstance
	if savedGitspaceInstance == nil || savedGitspaceInstance.State.IsFinalStatus() {
		codeServerPassword := defaultAccessKey
		gitspaceMachineUser := defaultMachineUser
		now := time.Now().UnixMilli()
		suffixUID, err := gonanoid.Generate(allowedUIDAlphabet, 6)
		if err != nil {
			return nil, fmt.Errorf("could not generate UID for gitspace config : %q %w", config.Identifier, err)
		}
		identifier := strings.ToLower(config.Identifier + "-" + suffixUID)
		var gitspaceInstance = &types.GitspaceInstance{
			GitSpaceConfigID: config.ID,
			Identifier:       identifier,
			State:            enum.GitspaceInstanceStateUninitialized,
			UserID:           config.UserID,
			SpaceID:          config.SpaceID,
			SpacePath:        config.SpacePath,
			Created:          now,
			Updated:          now,
			TotalTimeUsed:    0,
		}
		if config.IDE == enum.IDETypeVSCodeWeb || config.IDE == enum.IDETypeVSCode {
			gitspaceInstance.AccessKey = &codeServerPassword
			gitspaceInstance.AccessType = enum.GitspaceAccessTypeUserCredentials
			gitspaceInstance.MachineUser = &gitspaceMachineUser
		}
		if err = c.gitspaceInstanceStore.Create(ctx, gitspaceInstance); err != nil {
			return nil, fmt.Errorf("failed to create gitspace instance for %s %w", config.Identifier, err)
		}
		newGitspaceInstance, err := c.gitspaceInstanceStore.FindLatestByGitspaceConfigID(ctx, config.ID, config.SpaceID)
		newGitspaceInstance.SpacePath = config.SpacePath
		if err != nil {
			return nil, fmt.Errorf("failed to find gitspace with config ID : %s %w", config.Identifier, err)
		}
		config.GitspaceInstance = newGitspaceInstance
	}

	updatedGitspace, err := c.orchestrator.StartGitspace(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to find start gitspace : %s %w", config.Identifier, err)
	}
	if err = c.gitspaceInstanceStore.Update(ctx, updatedGitspace); err != nil {
		return nil, fmt.Errorf("failed to update gitspace %w", err)
	}
	config.GitspaceInstance = updatedGitspace
	config.State, _ = enum.GetGitspaceStateFromInstance(updatedGitspace.State)
	return config, nil
}

func (c *Controller) stopGitspace(ctx context.Context, config *types.GitspaceConfig) (*types.GitspaceConfig, error) {
	savedGitspace, err := c.gitspaceInstanceStore.FindLatestByGitspaceConfigID(ctx, config.ID, config.SpaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to find gitspace with config ID : %s %w", config.Identifier, err)
	}
	if savedGitspace.State.IsFinalStatus() {
		return nil, fmt.Errorf(
			"gitspace Instance cannot be stopped with ID %s %w", savedGitspace.Identifier, err)
	}
	config.GitspaceInstance = savedGitspace
	if updatedGitspace, stopErr := c.orchestrator.StopGitspace(ctx, config); stopErr != nil {
		return nil, fmt.Errorf(
			"failed to stop gitspace instance with ID %s %w", savedGitspace.Identifier, stopErr)
	} else if updatedGitspace != nil {
		if stopErr = c.gitspaceInstanceStore.Update(ctx, updatedGitspace); stopErr != nil {
			return nil, fmt.Errorf(
				"unable to update the gitspace with config id %s %w", savedGitspace.Identifier, stopErr)
		}
		config.GitspaceInstance = updatedGitspace
		config.State, _ = enum.GetGitspaceStateFromInstance(updatedGitspace.State)
	}
	return config, nil
}

func (c *Controller) sanitizeActionInput(in *ActionInput) error {
	if err := check.Identifier(in.Identifier); err != nil {
		return err
	}
	parentRefAsID, err := strconv.ParseInt(in.SpaceRef, 10, 64)
	if (err == nil && parentRefAsID <= 0) || (len(strings.TrimSpace(in.SpaceRef)) == 0) {
		return ErrGitspaceRequiresParent
	}
	return nil
}

func (c *Controller) emitGitspaceConfigEvent(
	ctx context.Context,
	config *types.GitspaceConfig,
	eventType enum.GitspaceEventType) {
	c.eventReporter.EmitGitspaceEvent(ctx, events.GitspaceEvent, &events.GitspaceEventPayload{
		QueryKey:   config.Identifier,
		EntityID:   config.ID,
		EntityType: enum.GitspaceEntityTypeGitspaceConfig,
		EventType:  eventType,
		Created:    time.Now().UnixMilli(),
	})
}