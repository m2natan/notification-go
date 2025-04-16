package application

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/queries"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CreateNotification(ctx context.Context, cmd commands.CreateNotificationCommand) (*domain.Notification, error)
		UpdateNotification(ctx context.Context, cmd commands.UpdateNotificationCommand) (*domain.Notification, error)
		DeleteNotification(ctx context.Context, cmd commands.DeleteNotificationCommand) error
	}

	Queries interface {
		FindAll(ctx context.Context) ([]domain.Notification, error)
		FindById(ctx context.Context, cmd queries.FindById) (*domain.Notification, error)
		FindByStatus(ctx context.Context, cmd queries.FindByStatus) ([]domain.Notification, error)
		FindByType(ctx context.Context, cmd queries.FindByType) ([]domain.Notification, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CreateNotificationHandler
		commands.UpdateNotificationHandler
		commands.DeleteNotificationHandler
	}
	appQueries struct {
		queries.FindAllHandler
		queries.FindByIdHandler
		queries.FindByStatusHandler
		queries.FindByTypeHandler
	}
)

var _ App = (*Application)(nil)

func New(
	notification domain.NotificationRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			commands.NewCreateNotificationHandler(notification),
			commands.NewUpdateNotificationHandler(notification),
			commands.NewDeleteNotificationHandler(notification),
		},
		appQueries: appQueries{
			queries.NewFindAllHandler(notification),
			queries.NewFindByIdHandler(notification),
			queries.NewFindByStatusHandler(notification),
			queries.NewFindByTypeHandler(notification),
		},
	}
}
