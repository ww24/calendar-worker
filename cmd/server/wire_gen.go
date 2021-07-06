// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/ww24/calendar-notifier/domain/repository"
	"github.com/ww24/calendar-notifier/domain/service"
	"github.com/ww24/calendar-notifier/interface/action"
	"github.com/ww24/calendar-notifier/interface/calendar"
	"github.com/ww24/calendar-notifier/interface/http/handler"
	"github.com/ww24/calendar-notifier/usecase"
)

// Injectors from wire.go:

func initialize(ctx context.Context, cnf repository.Config) (*app, error) {
	config := service.NewConfig(cnf)
	calendarCalendar := calendar.New(cnf)
	actionAction, err := action.New(ctx)
	if err != nil {
		return nil, err
	}
	synchronizer := service.NewSynchronizer(cnf, calendarCalendar, actionAction)
	usecaseSynchronizer := usecase.NewSynchronizer(config, synchronizer)
	httpHandler := handler.New(usecaseSynchronizer)
	mainApp := newApp(httpHandler, usecaseSynchronizer)
	return mainApp, nil
}
