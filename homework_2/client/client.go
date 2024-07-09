package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"homework_2/models"
	"io"
	"net/http"
	"os"
)

type Query struct {
	Host    string
	Port    int
	Command string
	Name    string
	Amount  int
	NewName string
}

func main() {
	hostVal := flag.String("host", "0.0.0.0", "server host")
	portVal := flag.Int("port", 8080, "server port")
	commandVal := flag.String("cmd", "", "command to execute: 'create', 'delete', 'update-amount',"+
		" 'update-name', 'get'")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account; required amount for query 'update-amount'")
	newNameVal := flag.String("new-name", "", "new name of account for 'update-name'")

	flag.Usage = func() {
		_, err := fmt.Fprint(os.Stderr, "Usage:\n")
		if err != nil {
			panic(err)
		}
		flag.PrintDefaults()
	}

	flag.Parse()

	query := Query{
		Host:    *hostVal,
		Port:    *portVal,
		Command: *commandVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newNameVal,
	}

	if err := Do(&query); err != nil {
		panic(err)
	}
}

func Do(q *Query) error {
	switch q.Command {
	case "create":
		if err := CliCreate(q); err != nil {
			return fmt.Errorf("create failed: %w", err)
		}
		return nil
	case "delete":
		if err := CliDelete(q); err != nil {
			return fmt.Errorf("delete failed: %w", err)
		}
		return nil
	case "update-amount":
		if err := CliUpdateAmount(q); err != nil {
			return fmt.Errorf("update-amount failed: %w", err)
		}
		return nil
	case "update-name":
		if err := CliUpdateName(q); err != nil {
			return fmt.Errorf("update-name failed: %w", err)
		}
		return nil
	case "get":
		if err := CliGetAccount(q); err != nil {
			return fmt.Errorf("get failed: %w", err)
		}
		return nil
	default:
		flag.Usage()
		return fmt.Errorf("unknown command: %s", q.Command)
	}
}

func CliCreate(q *Query) error {
	request := models.CreateAccountRequest{
		Name: q.Name,
	}

	req, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	response, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", q.Host, q.Port),
		"application/json",
		bytes.NewReader(req),
	)
	if err != nil {
		return fmt.Errorf("server create failed: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("response error: %s", string(body))
}

func CliDelete(q *Query) error {
	request := models.DeleteAccountRequest{
		Name: q.Name,
	}

	_, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s:%d/account/delete?name=%s", q.Host, q.Port, request.Name),
		nil)
	if err != nil {
		return fmt.Errorf("http DELETE creation failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("server delete failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("response error: %s", string(body))
}

func CliUpdateAmount(q *Query) error {
	request := models.UpdateAmountRequest{
		Name:   q.Name,
		Amount: q.Amount,
	}

	req, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/update/amount", q.Host, q.Port),
		"application/json",
		bytes.NewReader(req),
	)
	if err != nil {
		return fmt.Errorf("server update-amount failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("response error: %s", string(body))
}

func CliUpdateName(q *Query) error {
	request := models.UpdateNameRequest{
		Name:    q.Name,
		NewName: q.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("http://%s:%d/account/update/name", q.Host, q.Port),
		bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("http PUT creation failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("server update-name failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return fmt.Errorf("response error: %s", string(body))
}

func CliGetAccount(q *Query) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", q.Host, q.Port, q.Name),
	)
	if err != nil {
		return fmt.Errorf("server get failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("response error: %s", string(body))
	}

	var response models.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("Name: %s   Amount: %d", response.Name, response.Amount)
	return nil
}
