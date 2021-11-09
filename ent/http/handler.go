// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/open-farms/inventory/ent"
	"go.uber.org/zap"
)

// NewHandler returns a ready to use handler with all generated endpoints mounted.
func NewHandler(c *ent.Client, l *zap.Logger) chi.Router {
	r := chi.NewRouter()
	MountRoutes(c, l, r)
	return r
}

// MountRoutes mounts all generated routes on the given router.
func MountRoutes(c *ent.Client, l *zap.Logger, r chi.Router) {
	NewCategoryHandler(c, l).MountRoutes(r)
	NewEquipmentHandler(c, l).MountRoutes(r)
	NewImplementHandler(c, l).MountRoutes(r)
	NewLocationHandler(c, l).MountRoutes(r)
	NewToolHandler(c, l).MountRoutes(r)
	NewVehicleHandler(c, l).MountRoutes(r)
}

// CategoryHandler handles http crud operations on ent.Category.
type CategoryHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewCategoryHandler(c *ent.Client, l *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		client: c,
		log:    l.With(zap.String("handler", "CategoryHandler")),
	}
}
func (h *CategoryHandler) MountCreateRoute(r chi.Router) *CategoryHandler {
	r.Post("/categories", h.Create)
	return h
}
func (h *CategoryHandler) MountReadRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}", h.Read)
	return h
}
func (h *CategoryHandler) MountUpdateRoute(r chi.Router) *CategoryHandler {
	r.Patch("/categories/{id}", h.Update)
	return h
}
func (h *CategoryHandler) MountDeleteRoute(r chi.Router) *CategoryHandler {
	r.Delete("/categories/{id}", h.Delete)
	return h
}
func (h *CategoryHandler) MountListRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories", h.List)
	return h
}
func (h *CategoryHandler) MountVehicleRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/vehicle", h.Vehicle)
	return h
}
func (h *CategoryHandler) MountToolRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/tool", h.Tool)
	return h
}
func (h *CategoryHandler) MountImplementRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/implement", h.Implement)
	return h
}
func (h *CategoryHandler) MountEquipmentRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/equipment", h.Equipment)
	return h
}
func (h *CategoryHandler) MountLocationRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/location", h.Location)
	return h
}
func (h *CategoryHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountVehicleRoute(r).MountToolRoute(r).MountImplementRoute(r).MountEquipmentRoute(r).MountLocationRoute(r)
}

// EquipmentHandler handles http crud operations on ent.Equipment.
type EquipmentHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewEquipmentHandler(c *ent.Client, l *zap.Logger) *EquipmentHandler {
	return &EquipmentHandler{
		client: c,
		log:    l.With(zap.String("handler", "EquipmentHandler")),
	}
}
func (h *EquipmentHandler) MountCreateRoute(r chi.Router) *EquipmentHandler {
	r.Post("/equipment", h.Create)
	return h
}
func (h *EquipmentHandler) MountReadRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment/{id}", h.Read)
	return h
}
func (h *EquipmentHandler) MountUpdateRoute(r chi.Router) *EquipmentHandler {
	r.Patch("/equipment/{id}", h.Update)
	return h
}
func (h *EquipmentHandler) MountDeleteRoute(r chi.Router) *EquipmentHandler {
	r.Delete("/equipment/{id}", h.Delete)
	return h
}
func (h *EquipmentHandler) MountListRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment", h.List)
	return h
}
func (h *EquipmentHandler) MountLocationRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment/{id}/location", h.Location)
	return h
}
func (h *EquipmentHandler) MountCategoryRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment/{id}/category", h.Category)
	return h
}
func (h *EquipmentHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountLocationRoute(r).MountCategoryRoute(r)
}

// ImplementHandler handles http crud operations on ent.Implement.
type ImplementHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewImplementHandler(c *ent.Client, l *zap.Logger) *ImplementHandler {
	return &ImplementHandler{
		client: c,
		log:    l.With(zap.String("handler", "ImplementHandler")),
	}
}
func (h *ImplementHandler) MountCreateRoute(r chi.Router) *ImplementHandler {
	r.Post("/implements", h.Create)
	return h
}
func (h *ImplementHandler) MountReadRoute(r chi.Router) *ImplementHandler {
	r.Get("/implements/{id}", h.Read)
	return h
}
func (h *ImplementHandler) MountUpdateRoute(r chi.Router) *ImplementHandler {
	r.Patch("/implements/{id}", h.Update)
	return h
}
func (h *ImplementHandler) MountDeleteRoute(r chi.Router) *ImplementHandler {
	r.Delete("/implements/{id}", h.Delete)
	return h
}
func (h *ImplementHandler) MountListRoute(r chi.Router) *ImplementHandler {
	r.Get("/implements", h.List)
	return h
}
func (h *ImplementHandler) MountLocationRoute(r chi.Router) *ImplementHandler {
	r.Get("/implements/{id}/location", h.Location)
	return h
}
func (h *ImplementHandler) MountCategoryRoute(r chi.Router) *ImplementHandler {
	r.Get("/implements/{id}/category", h.Category)
	return h
}
func (h *ImplementHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountLocationRoute(r).MountCategoryRoute(r)
}

// LocationHandler handles http crud operations on ent.Location.
type LocationHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewLocationHandler(c *ent.Client, l *zap.Logger) *LocationHandler {
	return &LocationHandler{
		client: c,
		log:    l.With(zap.String("handler", "LocationHandler")),
	}
}
func (h *LocationHandler) MountCreateRoute(r chi.Router) *LocationHandler {
	r.Post("/locations", h.Create)
	return h
}
func (h *LocationHandler) MountReadRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}", h.Read)
	return h
}
func (h *LocationHandler) MountUpdateRoute(r chi.Router) *LocationHandler {
	r.Patch("/locations/{id}", h.Update)
	return h
}
func (h *LocationHandler) MountDeleteRoute(r chi.Router) *LocationHandler {
	r.Delete("/locations/{id}", h.Delete)
	return h
}
func (h *LocationHandler) MountListRoute(r chi.Router) *LocationHandler {
	r.Get("/locations", h.List)
	return h
}
func (h *LocationHandler) MountVehicleRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}/vehicle", h.Vehicle)
	return h
}
func (h *LocationHandler) MountToolRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}/tool", h.Tool)
	return h
}
func (h *LocationHandler) MountImplementRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}/implement", h.Implement)
	return h
}
func (h *LocationHandler) MountEquipmentRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}/equipment", h.Equipment)
	return h
}
func (h *LocationHandler) MountCategoryRoute(r chi.Router) *LocationHandler {
	r.Get("/locations/{id}/category", h.Category)
	return h
}
func (h *LocationHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountVehicleRoute(r).MountToolRoute(r).MountImplementRoute(r).MountEquipmentRoute(r).MountCategoryRoute(r)
}

// ToolHandler handles http crud operations on ent.Tool.
type ToolHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewToolHandler(c *ent.Client, l *zap.Logger) *ToolHandler {
	return &ToolHandler{
		client: c,
		log:    l.With(zap.String("handler", "ToolHandler")),
	}
}
func (h *ToolHandler) MountCreateRoute(r chi.Router) *ToolHandler {
	r.Post("/tools", h.Create)
	return h
}
func (h *ToolHandler) MountReadRoute(r chi.Router) *ToolHandler {
	r.Get("/tools/{id}", h.Read)
	return h
}
func (h *ToolHandler) MountUpdateRoute(r chi.Router) *ToolHandler {
	r.Patch("/tools/{id}", h.Update)
	return h
}
func (h *ToolHandler) MountDeleteRoute(r chi.Router) *ToolHandler {
	r.Delete("/tools/{id}", h.Delete)
	return h
}
func (h *ToolHandler) MountListRoute(r chi.Router) *ToolHandler {
	r.Get("/tools", h.List)
	return h
}
func (h *ToolHandler) MountLocationRoute(r chi.Router) *ToolHandler {
	r.Get("/tools/{id}/location", h.Location)
	return h
}
func (h *ToolHandler) MountCategoryRoute(r chi.Router) *ToolHandler {
	r.Get("/tools/{id}/category", h.Category)
	return h
}
func (h *ToolHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountLocationRoute(r).MountCategoryRoute(r)
}

// VehicleHandler handles http crud operations on ent.Vehicle.
type VehicleHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewVehicleHandler(c *ent.Client, l *zap.Logger) *VehicleHandler {
	return &VehicleHandler{
		client: c,
		log:    l.With(zap.String("handler", "VehicleHandler")),
	}
}
func (h *VehicleHandler) MountCreateRoute(r chi.Router) *VehicleHandler {
	r.Post("/vehicles", h.Create)
	return h
}
func (h *VehicleHandler) MountReadRoute(r chi.Router) *VehicleHandler {
	r.Get("/vehicles/{id}", h.Read)
	return h
}
func (h *VehicleHandler) MountUpdateRoute(r chi.Router) *VehicleHandler {
	r.Patch("/vehicles/{id}", h.Update)
	return h
}
func (h *VehicleHandler) MountDeleteRoute(r chi.Router) *VehicleHandler {
	r.Delete("/vehicles/{id}", h.Delete)
	return h
}
func (h *VehicleHandler) MountListRoute(r chi.Router) *VehicleHandler {
	r.Get("/vehicles", h.List)
	return h
}
func (h *VehicleHandler) MountLocationRoute(r chi.Router) *VehicleHandler {
	r.Get("/vehicles/{id}/location", h.Location)
	return h
}
func (h *VehicleHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountLocationRoute(r)
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
