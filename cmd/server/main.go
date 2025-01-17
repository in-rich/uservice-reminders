package main

import (
	"fmt"
	"github.com/in-rich/lib-go/deploy"
	"github.com/in-rich/lib-go/monitor"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/config"
	"github.com/in-rich/uservice-reminders/migrations"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/handlers"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"github.com/rs/zerolog"
	"os"
)

func getLogger() monitor.GRPCLogger {
	if deploy.IsReleaseEnv() {
		return monitor.NewGCPGRPCLogger(zerolog.New(os.Stdout), "uservice-reminders")
	}

	return monitor.NewConsoleGRPCLogger()
}

func main() {
	logger := getLogger()

	logger.Info("Starting server")
	db, closeDB, err := deploy.OpenDB(config.App.Postgres.DSN)
	if err != nil {
		logger.Fatal(err, "failed to connect to database")
	}
	defer closeDB()

	logger.Info("Running migrations")
	if err := migrations.Migrate(db); err != nil {
		logger.Fatal(err, "failed to migrate")
	}

	depCheck := deploy.DepsCheck{
		Dependencies: func() map[string]error {
			return map[string]error{
				"Postgres": db.Ping(),
			}
		},
		Services: deploy.DepCheckServices{
			"GetReminder":    {"Postgres"},
			"UpsertReminder": {"Postgres"},
		},
	}

	createReminderDAO := dao.NewCreateReminderRepository(db)
	deleteReminderDAO := dao.NewDeleteReminderRepository(db)
	getReminderDAO := dao.NewGetReminderRepository(db)
	updateReminderDAO := dao.NewUpdateReminderRepository(db)
	getReminderByIDDAO := dao.NewGetReminderByIDRepository(db)

	getReminderService := services.NewGetReminderService(getReminderDAO)
	upsertReminderService := services.NewUpsertReminderService(
		updateReminderDAO,
		createReminderDAO,
		deleteReminderDAO,
	)
	getReminderByIDService := services.NewGetReminderByIDService(getReminderByIDDAO)

	getReminderHandler := handlers.NewGetReminderHandler(getReminderService, logger)
	upsertReminderHandler := handlers.NewUpsertReminderHandler(upsertReminderService, logger)
	getReminderByIDHandler := handlers.NewGetReminderByIDHandler(getReminderByIDService)

	logger.Info(fmt.Sprintf("Starting to listen on port %v", config.App.Server.Port))
	listener, server, health := deploy.StartGRPCServer(logger, config.App.Server.Port, depCheck)
	defer deploy.CloseGRPCServer(listener, server)
	go health()

	reminders_pb.RegisterGetReminderServer(server, getReminderHandler)
	reminders_pb.RegisterUpsertReminderServer(server, upsertReminderHandler)
	reminders_pb.RegisterGetReminderByIDServer(server, getReminderByIDHandler)

	logger.Info("Server started")
	if err := server.Serve(listener); err != nil {
		logger.Fatal(err, "failed to serve")
	}
}
