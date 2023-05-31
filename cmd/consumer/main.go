package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flpgst/go-reportgen/configs"
	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/infra/database/mongodb"
	"github.com/flpgst/go-reportgen/internal/infra/queue/interfaces"
	"github.com/flpgst/go-reportgen/internal/infra/queue/rabbitmq"
	"github.com/flpgst/go-reportgen/internal/infra/web"
	"github.com/flpgst/go-reportgen/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/streadway/amqp"
)

type QueueConn struct {
	queue interfaces.RabbitMQInterface
}

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := mongodb.MongoConnection(configs.DBDriver, configs.DBName, configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort)
	if err != nil {
		panic(err)
	}
	defer db.Client().Disconnect(context.TODO())

	reportRepository := mongodb.NewReportRepository(db)
	createReportUseCase := usecase.NewCreateReportUseCase(reportRepository)

	reportHandler := web.NewWebReportHandler(reportRepository)
	router := chi.NewRouter()
	router.Route("/report", func(r chi.Router) {
		r.Get("/", reportHandler.Get)
	})
	go http.ListenAndServe(":"+configs.WebServerPort, router)

	queueConn := QueueConn{
		queue: rabbitmq.NewRabbitMQConn(configs.RABBITMQ_USER, configs.RABBITMQ_PASSWORD, configs.RABBITMQ_HOST, configs.RABBITMQ_PORT),
	}

	ch := queueConn.queue.OpenChannel()
	defer ch.Close()
	msgs := make(chan amqp.Delivery)

	go queueConn.queue.Consume(ch, msgs, configs.RABBITMQ_QUEUE_NAME)

	for msg := range msgs {
		var message dto.ReportInputDTO
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			fmt.Println("Error decoding message:", err)
			msg.Nack(false, false)
			continue
		}
		_, err = createReportUseCase.Execute(message)
		if err != nil {
			fmt.Println(err)
			msg.Nack(false, false)
		} else {
			msg.Ack(false)
		}
	}

}
