package handlers

import (
    "jwt-auth-crud/internal/dto"
    "jwt-auth-crud/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ProductHandler interface {
    CreateProduct(c *gin.Context)
    GetProducts(c *gin.Context)
    GetProduct(c *gin.Context)
    UpdateProduct(c *gin.Context)
    DeleteProduct(c *gin.Context)
    GetAllProducts(c *gin.Context) // Uncomment if needed

}

type productHandler struct {
    productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
    return &productHandler{
        productService: productService,
    }
}

func (h *productHandler) CreateProduct(c *gin.Context) {
    userID := c.GetUint("user_id")

    var req dto.CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product, err := h.productService.CreateProduct(userID, &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Product created successfully",
        "data":    product,
    })
}

func (h *productHandler) GetProducts(c *gin.Context) {
    userID := c.GetUint("user_id")

    products, err := h.productService.GetProducts(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Products retrieved successfully",
        "data":    products,
    })
}

func (h *productHandler) GetProduct(c *gin.Context) {
    // Get the product ID from URL params
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    // Get the requesting user's ID (for authorization if needed)
    requestingUserID := c.GetUint("user_id")
    
    // Get the product
    product, err := h.productService.GetProduct(uint(id), requestingUserID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    // Return the product with user data
    c.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var req dto.UpdateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.productService.UpdateProduct(uint(id), userID, &req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Product updated successfully",
    })
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    if err := h.productService.DeleteProduct(uint(id), userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Product deleted successfully",
    })
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
    products, err := h.productService.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "All products retrieved successfully",
        "data":    products,
    })
}

