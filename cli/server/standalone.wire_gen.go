// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject && !harness
// +build !wireinject,!harness

package server

import (
	"context"
	"github.com/harness/gitness/events"
	"github.com/harness/gitness/gitrpc"
	server2 "github.com/harness/gitness/gitrpc/server"
	"github.com/harness/gitness/gitrpc/server/cron"
	"github.com/harness/gitness/internal/api/controller/githook"
	"github.com/harness/gitness/internal/api/controller/principal"
	"github.com/harness/gitness/internal/api/controller/pullreq"
	"github.com/harness/gitness/internal/api/controller/repo"
	"github.com/harness/gitness/internal/api/controller/service"
	"github.com/harness/gitness/internal/api/controller/serviceaccount"
	"github.com/harness/gitness/internal/api/controller/space"
	"github.com/harness/gitness/internal/api/controller/user"
	webhook2 "github.com/harness/gitness/internal/api/controller/webhook"
	"github.com/harness/gitness/internal/auth/authn"
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/bootstrap"
	events3 "github.com/harness/gitness/internal/events/git"
	events2 "github.com/harness/gitness/internal/events/pullreq"
	"github.com/harness/gitness/internal/router"
	"github.com/harness/gitness/internal/server"
	"github.com/harness/gitness/internal/services"
	"github.com/harness/gitness/internal/services/codecomments"
	pullreq2 "github.com/harness/gitness/internal/services/pullreq"
	"github.com/harness/gitness/internal/services/webhook"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/internal/store/cache"
	"github.com/harness/gitness/internal/store/database"
	"github.com/harness/gitness/internal/url"
	"github.com/harness/gitness/lock"
	"github.com/harness/gitness/pubsub"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"
)

// Injectors from standalone.wire.go:

func initSystem(ctx context.Context, config *types.Config) (*system, error) {
	principalUID := check.ProvidePrincipalUIDCheck()
	authorizer := authz.ProvideAuthorizer()
	db, err := database.ProvideDatabase(ctx, config)
	if err != nil {
		return nil, err
	}
	principalUIDTransformation := store.ProvidePrincipalUIDTransformation()
	principalStore := database.ProvidePrincipalStore(db, principalUIDTransformation)
	tokenStore := database.ProvideTokenStore(db)
	controller := user.NewController(principalUID, authorizer, principalStore, tokenStore)
	serviceController := service.NewController(principalUID, authorizer, principalStore)
	bootstrapBootstrap := bootstrap.ProvideBootstrap(config, controller, serviceController)
	authenticator := authn.ProvideAuthenticator(principalStore, tokenStore)
	provider, err := url.ProvideURLProvider(config)
	if err != nil {
		return nil, err
	}
	pathUID := check.ProvidePathUIDCheck()
	pathTransformation := store.ProvidePathTransformation()
	pathStore := database.ProvidePathStore(db, pathTransformation)
	pathCache := cache.ProvidePathCache(pathStore, pathTransformation)
	repoStore := database.ProvideRepoStore(db, pathCache)
	spaceStore := database.ProvideSpaceStore(db, pathCache)
	gitrpcConfig, err := ProvideGitRPCClientConfig()
	if err != nil {
		return nil, err
	}
	gitrpcInterface, err := gitrpc.ProvideClient(gitrpcConfig)
	if err != nil {
		return nil, err
	}
	repoController := repo.ProvideController(config, db, provider, pathUID, authorizer, pathStore, repoStore, spaceStore, principalStore, gitrpcInterface)
	spaceController := space.ProvideController(db, provider, pathUID, authorizer, pathStore, spaceStore, repoStore, principalStore)
	principalInfoView := database.ProvidePrincipalInfoView(db)
	principalInfoCache := cache.ProvidePrincipalInfoCache(principalInfoView)
	pullReqStore := database.ProvidePullReqStore(db, principalInfoCache)
	pullReqActivityStore := database.ProvidePullReqActivityStore(db, principalInfoCache)
	codeCommentView := database.ProvideCodeCommentView(db)
	pullReqReviewStore := database.ProvidePullReqReviewStore(db)
	pullReqReviewerStore := database.ProvidePullReqReviewerStore(db, principalInfoCache)
	eventsConfig, err := ProvideEventsConfig()
	if err != nil {
		return nil, err
	}
	universalClient, err := ProvideRedis(config)
	if err != nil {
		return nil, err
	}
	eventsSystem, err := events.ProvideSystem(eventsConfig, universalClient)
	if err != nil {
		return nil, err
	}
	reporter, err := events2.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	lockConfig := lock.ProvideConfig(config)
	mutexManager := lock.ProvideMutexManager(lockConfig, universalClient)
	migrator := codecomments.ProvideMigrator(gitrpcInterface)
	pullreqController := pullreq.ProvideController(db, provider, authorizer, pullReqStore, pullReqActivityStore, codeCommentView, pullReqReviewStore, pullReqReviewerStore, repoStore, principalStore, gitrpcInterface, reporter, mutexManager, migrator)
	webhookConfig, err := ProvideWebhookConfig()
	if err != nil {
		return nil, err
	}
	webhookStore := database.ProvideWebhookStore(db)
	webhookExecutionStore := database.ProvideWebhookExecutionStore(db)
	readerFactory, err := events3.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	eventsReaderFactory, err := events2.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	webhookService, err := webhook.ProvideService(ctx, webhookConfig, readerFactory, eventsReaderFactory, webhookStore, webhookExecutionStore, repoStore, pullReqStore, provider, principalStore, gitrpcInterface)
	if err != nil {
		return nil, err
	}
	webhookController := webhook2.ProvideController(webhookConfig, db, authorizer, webhookStore, webhookExecutionStore, repoStore, webhookService)
	eventsReporter, err := events3.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	githookController := githook.ProvideController(db, authorizer, principalStore, repoStore, eventsReporter)
	serviceaccountController := serviceaccount.NewController(principalUID, authorizer, principalStore, spaceStore, repoStore, tokenStore)
	principalController := principal.NewController(principalStore)
	apiHandler := router.ProvideAPIHandler(config, authenticator, repoController, spaceController, pullreqController, webhookController, githookController, serviceaccountController, controller, principalController)
	gitHandler := router.ProvideGitHandler(config, provider, repoStore, authenticator, authorizer, gitrpcInterface)
	webHandler := router.ProvideWebHandler(config)
	routerRouter := router.ProvideRouter(config, apiHandler, gitHandler, webHandler)
	serverServer := server.ProvideServer(config, routerRouter)
	serverConfig, err := ProvideGitRPCServerConfig()
	if err != nil {
		return nil, err
	}
	server3, err := server2.ProvideServer(serverConfig)
	if err != nil {
		return nil, err
	}
	cronManager := cron.ProvideCronManager(serverConfig)
	repoGitInfoView := database.ProvideRepoGitInfoView(db)
	repoGitInfoCache := cache.ProvideRepoGitInfoCache(repoGitInfoView)
	pubsubConfig := pubsub.ProvideConfig(config)
	pubSub := pubsub.ProvidePubSub(pubsubConfig, universalClient)
	pullreqService, err := pullreq2.ProvideService(ctx, config, readerFactory, eventsReaderFactory, reporter, gitrpcInterface, db, repoGitInfoCache, principalInfoCache, repoStore, pullReqStore, pullReqActivityStore, codeCommentView, migrator, pubSub)
	if err != nil {
		return nil, err
	}
	servicesServices := services.ProvideServices(webhookService, pullreqService)
	serverSystem := newSystem(bootstrapBootstrap, serverServer, server3, cronManager, servicesServices)
	return serverSystem, nil
}
