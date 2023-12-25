package dialogs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/server/dialogs/gen/api"
)

func (h *Handler) DialogUserIDListGet(ctx context.Context, params api.DialogUserIDListGetParams) (api.DialogUserIDListGetRes, error) {
	toUserID := uuid.MustParse(string(params.UserID))
	fromUserID := uuid.MustParse(params.XScUserID)
	messages, err := h.app.ListMessages(ctx, fromUserID, toUserID, time.Time{})
	if err != nil {
		return nil, err
	}

	return convertMessagesToAPI(messages), nil
}

func convertMessagesToAPI(messages []entity.Message) *api.DialogUserIDListGetOKApplicationJSON {
	dialogMessages := lo.Map(messages, func(m entity.Message, _ int) api.DialogMessage {
		return api.DialogMessage{
			From: api.UserId(m.Author.String()),
			To:   api.UserId(m.Recipient.String()),
			Text: api.DialogMessageText(m.Text),
		}
	})
	return lo.ToPtr(api.DialogUserIDListGetOKApplicationJSON(dialogMessages))
}

func (h *Handler) DialogUserIDSendPost(ctx context.Context, req *api.DialogUserIDSendPostReq, params api.DialogUserIDSendPostParams) (api.DialogUserIDSendPostRes, error) {
	toUserID := uuid.MustParse(string(params.UserID))
	fromUserID := uuid.MustParse(params.XScUserID)

	err := h.app.SendMessage(ctx, fromUserID, toUserID, string(req.Text), time.Now())
	if err != nil {
		return nil, err
	}

	return &api.DialogUserIDSendPostOK{}, nil
}
