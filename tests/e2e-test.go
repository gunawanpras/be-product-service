package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// baseURL is the endpoint of the live API service.
const baseURL = "http://localhost:8080"

// APIClient represents a client for the product API.
type APIClient struct {
	BaseURL string
	Client  *http.Client
}

// newAPIClient creates a new APIClient with the provided base URL.
func newAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// doRequest executes the given HTTP request and returns the status code and response body.
func (c *APIClient) doRequest(req *http.Request) (int, []byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, body, nil
}

// GetProducts sends a GET request to the /products endpoint.
func (c *APIClient) GetProducts(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// GetProductByID sends a GET request to the /products/:id endpoint.
func (c *APIClient) GetProductByID(ctx context.Context, productID string) (int, []byte, error) {
	url := fmt.Sprintf("%s/products/%s", c.BaseURL, productID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// SearchAProductByName sends a GET request to the /products?product_name=kangkung endpoint.
func (c *APIClient) SearchAProductByName(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products?product_name=kangkung", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// SearchAProductByNameAndFilterByCategoryType sends a GET request to the /products?product_name=kangkung&category_type=Sayuran endpoint.
func (c *APIClient) SearchAProductByNameAndFilterByCategoryType(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products?product_name=kangkung&category_type=Sayuran", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// GetProductsAndFilteredByCategoryType sends a GET request to the /products?category_type=Sayuran endpoint.
func (c *APIClient) GetProductsAndFilteredByCategoryType(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products?category_type=Sayuran", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// GetProductsAndSortedByNameAsc sends a GET request to the /products?sort=name&directive=asc endpoint.
func (c *APIClient) GetProductsAndSortedByNameAsc(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products?sort=name&directive=asc", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// GetProductsAndSortedByBasePriceDesc sends a GET request to the /products?sort=base_price&directive=desc endpoint.
func (c *APIClient) GetProductsAndSortedByBasePriceDesc(ctx context.Context) (int, []byte, error) {
	url := fmt.Sprintf("%s/products?sort=base_price&directive=desc", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}
	return c.doRequest(req)
}

// CreateProduct sends a POST request to create a new product.
// The productPayload is expected to be a map representing the JSON payload.
func (c *APIClient) CreateProduct(ctx context.Context, productPayload map[string]any) (int, []byte, error) {
	payloadBytes, err := json.Marshal(productPayload)
	if err != nil {
		return 0, nil, err
	}

	url := fmt.Sprintf("%s/products", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func execCreateProductTest(ctx context.Context, client *APIClient) {
	productPayload := []map[string]any{
		{
			"name":        "Kangkung Potong 1",
			"description": "A sample product for E2E testing",
			"base_price":  3000.00,
			"stock":       100,
			"category_id": "00000000-0000-0000-0000-000000000001",
			"supplier_id": "00000000-0000-0000-0000-000000000011",
			"unit_id":     "00000000-0000-0000-0000-000000000021",
		},
		{
			"name":        "Kangkung Potong 2",
			"description": "A sample product for E2E testing",
			"base_price":  3500.00,
			"stock":       100,
			"category_id": "00000000-0000-0000-0000-000000000003",
			"supplier_id": "00000000-0000-0000-0000-000000000011",
			"unit_id":     "00000000-0000-0000-0000-000000000021",
		},
	}

	fmt.Printf("1) POST /products\n")
	for _, payload := range productPayload {
		status, _, err := client.CreateProduct(ctx, payload)
		if err != nil {
			log.Fatalf("error in CreateProduct: %v", err)
		}
		fmt.Printf("  - Create product '%s' - Status: %d\n", payload["name"], status)
	}
}

func execGetAllProductsTest(ctx context.Context, client *APIClient) {
	status, _, err := client.GetProducts(ctx)
	if err != nil {
		log.Fatalf("Error in GetProducts: %v", err)
	}
	fmt.Printf("2) GET /products - Status: %d\n", status)
}

func execGetProductByIDTest(ctx context.Context, client *APIClient) {
	sampleID := "00000000-0000-0000-0000-000000000031"
	status, _, err := client.GetProductByID(ctx, sampleID)
	if err != nil {
		log.Fatalf("Error in GetProductByID: %v", err)
	}
	fmt.Printf("3) GET /products/%s - Status: %d\n", sampleID, status)
}

func execSearchProductByNameTest(ctx context.Context, client *APIClient) {
	productName := "kangkung"
	searchQuery := fmt.Sprintf("?product_name=%s", productName)

	status, _, err := client.SearchAProductByName(ctx)
	if err != nil {
		log.Fatalf("Error in SearchAProductByName: %v", err)
	}
	fmt.Printf("4) GET /products%s - Status: %d\n", searchQuery, status)
}

func execSearchProductByCategoryAndProductNameTest(ctx context.Context, client *APIClient) {
	productName := "kangkung"
	categoryType := "Sayuran"

	searchQuery := fmt.Sprintf("?product_name=%s&category_type=%s", productName, categoryType)
	status, _, err := client.SearchAProductByNameAndFilterByCategoryType(ctx)
	if err != nil {
		log.Fatalf("Error in SearchAProductByNameAndFilterByCategoryType: %v", err)
	}
	fmt.Printf("5) GET /products%s - Status: %d\n", searchQuery, status)
}

func execGetAllProductsAndSortedByNameAscTest(ctx context.Context, client *APIClient) {
	sort := "name"
	directive := "asc"
	searchQuery := fmt.Sprintf("?sort=%s&directive=%s", sort, directive)

	status, _, err := client.GetProductsAndSortedByNameAsc(ctx)
	if err != nil {
		log.Fatalf("Error in GetProductsAndSortedByNameAsc: %v", err)
	}
	fmt.Printf("6) GET /products%s - Status: %d\n", searchQuery, status)
}

func execGetAllProductsAndSortedByBasePriceDescTest(ctx context.Context, client *APIClient) {
	sort := "base_price"
	directive := "desc"
	searchQuery := fmt.Sprintf("?sort=%s&directive=%s", sort, directive)

	status, _, err := client.GetProductsAndSortedByBasePriceDesc(ctx)
	if err != nil {
		log.Fatalf("Error in GetProductsAndSortedByBasePriceDesc: %v", err)
	}
	fmt.Printf("7) GET /products%s - Status: %d\n\n", searchQuery, status)
}

func main() {
	client := newAPIClient(baseURL)
	// Create a context with timeout for all API calls.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Run tests
	execCreateProductTest(ctx, client)
	execGetAllProductsTest(ctx, client)
	execGetProductByIDTest(ctx, client)
	execSearchProductByNameTest(ctx, client)
	execSearchProductByCategoryAndProductNameTest(ctx, client)
	execGetAllProductsAndSortedByNameAscTest(ctx, client)
	execGetAllProductsAndSortedByBasePriceDescTest(ctx, client)
}
