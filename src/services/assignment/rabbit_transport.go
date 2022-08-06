package assignment

import (
	"context"
	"encoding/json"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/logger"
)

type rabbitTransport struct {
	service IService
}

func NewRabbitTransport(service IService) *rabbitTransport {
	return &rabbitTransport{service: service}
}

func (t *rabbitTransport) Bootstrap(ctx context.Context) {
	conn := config.GetConfig().RabbitMQ

	ch, err := conn.Channel()
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "direct", true, false, false, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	err = ch.QueueBind(q.Name, "assignment", "logs", false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Sugar().Error(err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			content := d.Body
			var assignment createAssignmentLogInput
			err := json.Unmarshal(content, &assignment)
			if err != nil {
				logger.Sugar().Error(err)
			} else {
				logger.Sugar().Info(assignment)
				err = t.service.Create(ctx, assignment)
				if err != nil {
					logger.Sugar().Error(err)
				}
			}

		}
	}()
	<-forever
}
