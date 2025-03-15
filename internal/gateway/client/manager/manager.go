package manager

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/gateway/client/manager/convert"
	"go-micro-service-template/internal/gateway/client/manager/model"
	er "go-micro-service-template/pkg/error"
)

type Manager struct {
	options options
	client  *http.Client
}

func New(in ...Option) *Manager {
	cfg := options{
		host:    "localhost",
		port:    8080,
		timeout: time.Second * 30,
	}

	for _, o := range in {
		o(&cfg)
	}

	client := &http.Client{}
	return &Manager{
		options: cfg,
		client:  client,
	}
}

func (m *Manager) CreateGame(ctx context.Context) (string, error) {
	req, err := http.NewRequest(http.MethodPost, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/game", bytes.NewReader([]byte("{}")))
	if err != nil {
		return "", er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return "", er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", er.Wrap(err, "io.ReadAll")

	}

	var gamesResp model.PostCreateGameResponse
	if err = json.Unmarshal(body, &gamesResp); err != nil {
		return "", er.Wrap(err, "json.Unmarshal")
	}

	return gamesResp.Payload.ID, nil
}

func (m *Manager) StartGame(ctx context.Context, id string) error {
	req, err := http.NewRequest(http.MethodPut, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/game/"+id, bytes.NewReader([]byte(`{"action": "run"}`)))
	if err != nil {
		return er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	return nil
}

func (m *Manager) StopGame(ctx context.Context, id string) error {
	req, err := http.NewRequest(http.MethodPut, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/game/"+id, bytes.NewReader([]byte(`{"action": "stop"}`)))
	if err != nil {
		return er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	return nil
}

func (m *Manager) Game(ctx context.Context, id string) (*entity.Game, error) {
	req, err := http.NewRequest(http.MethodGet, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/game/"+id, nil)
	if err != nil {
		return nil, er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, er.Wrap(err, "io.ReadAll")

	}

	var gamesResp model.GetNodeGameResponse
	if err = json.Unmarshal(body, &gamesResp); err != nil {
		return nil, er.Wrap(err, "json.Unmarshal")
	}

	return convert.ModelToGame(gamesResp.Payload), nil
}

func (m *Manager) Games(ctx context.Context) ([]*entity.Game, error) {
	req, err := http.NewRequest(http.MethodGet, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/games", nil)
	if err != nil {
		return nil, er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, er.Wrap(err, "io.ReadAll")

	}

	var gamesResp model.GetNodeGamesResponse
	if err = json.Unmarshal(body, &gamesResp); err != nil {
		return nil, er.Wrap(err, "json.Unmarshal")
	}

	return convert.ModelToGames(gamesResp.Payload), nil
}

func (m *Manager) AddBot(ctx context.Context, id string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/bot/"+id, nil)
	if err != nil {
		return "", er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return "", er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", er.Wrap(err, "io.ReadAll")

	}

	var gamesResp model.PostBotResponse
	if err = json.Unmarshal(body, &gamesResp); err != nil {
		return "", er.Wrap(err, "json.Unmarshal")
	}

	return "", nil
}

func (m *Manager) View(ctx context.Context, id string) ([][]*entity.Cell, error) {
	req, err := http.NewRequest(http.MethodGet, "http://"+m.options.host+":"+strconv.Itoa(m.options.port)+"/game/view/"+id, nil)
	if err != nil {
		return nil, er.Wrap(err, "http.NewRequest")
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, er.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, er.Wrap(err, "io.ReadAll")

	}

	var gamesResp model.GetNodeViewResponse
	if err = json.Unmarshal(body, &gamesResp); err != nil {
		return nil, er.Wrap(err, "json.Unmarshal")
	}

	return convert.ModelToCells(gamesResp.Payload), nil
}
