package logging

import (
	"context"
	"log/slog"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/queries"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type Application struct {
	application.App
	logger *slog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger *slog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) FindAll(ctx context.Context) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindAll")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindAll", "error", err)
		} else {
			a.logger.Info("<-- Notifications.FindAll")
		}
	}()

	return a.App.FindAll(ctx)
}

func (a Application) FindById(ctx context.Context, query queries.FindById) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindById")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindById", "error", err)
		} else {
			a.logger.Info("<-- Notifications.FindById")
		}
	}()

	return a.App.FindById(ctx, query)
}

func (a Application) FindByStatus(ctx context.Context, query queries.FindByStatus) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindByStatus")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindByStatus", "error", err)
		} else {
			a.logger.Info("<-- Notifications.FindByStatus")
		}
	}()

	return a.App.FindByStatus(ctx, query)
}

func (a Application) FindByType(ctx context.Context, query queries.FindByType) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindByType")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindByType", "error", err)
		} else {
			a.logger.Info("<-- Notifications.FindByType")
		}
	}()

	return a.App.FindByType(ctx, query)
}

func (a Application) CreateNotification(ctx context.Context, cmd commands.CreateNotificationCommand) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.CreateNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.CreateNotification", "error", err)
		} else {
			a.logger.Info("<-- Notifications.CreateNotification")
		}
	}()

	return a.App.CreateNotification(ctx, cmd)
}

func (a Application) UpdateNotification(ctx context.Context, cmd commands.UpdateNotificationCommand) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.UpdateNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.UpdateNotification", "error", err)
		} else {
			a.logger.Info("<-- Notifications.UpdateNotification")
		}
	}()

	return a.App.UpdateNotification(ctx, cmd)
}

func (a Application) DeleteNotification(ctx context.Context, cmd commands.DeleteNotificationCommand) (err error) {
	a.logger.Info("--> Notifications.DeleteNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.DeleteNotification", "error", err)
		} else {
			a.logger.Info("<-- Notifications.DeleteNotification")
		}
	}()

	return a.App.DeleteNotification(ctx, cmd)
}
