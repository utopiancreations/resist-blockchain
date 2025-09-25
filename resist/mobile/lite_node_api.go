package mobile

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MobileLiteNodeAPI provides REST API endpoints optimized for mobile devices
type MobileLiteNodeAPI struct {
	authService    *AuthenticationService
	contentService *ContentService
	syncService    *SynchronizationService
	signalService  *SignalService
	nodeService    *NodeService
}

// NewMobileLiteNodeAPI creates a new mobile API instance
func NewMobileLiteNodeAPI() *MobileLiteNodeAPI {
	return &MobileLiteNodeAPI{
		authService:    NewAuthenticationService(),
		contentService: NewContentService(),
		syncService:    NewSynchronizationService(),
		signalService:  NewSignalService(),
		nodeService:    NewNodeService(),
	}
}

// RegisterRoutes sets up all mobile API routes
func (api *MobileLiteNodeAPI) RegisterRoutes(router *mux.Router) {
	// Authentication routes
	router.HandleFunc("/api/v1/auth/challenge", api.RequestChallenge).Methods("POST")
	router.HandleFunc("/api/v1/auth/verify", api.VerifyChallenge).Methods("POST")

	// Content routes
	router.HandleFunc("/api/v1/posts/feed", api.GetFeed).Methods("GET")
	router.HandleFunc("/api/v1/posts/create", api.CreatePost).Methods("POST")
	router.HandleFunc("/api/v1/posts/{id}", api.GetPost).Methods("GET")

	// Node management routes
	router.HandleFunc("/api/v1/node/status", api.GetNodeStatus).Methods("GET")
	router.HandleFunc("/api/v1/node/resource-offer", api.CreateResourceOffer).Methods("POST")

	// Synchronization routes
	router.HandleFunc("/api/v1/sync/selective", api.SelectiveSync).Methods("POST")
	router.HandleFunc("/api/v1/sync/status", api.GetSyncStatus).Methods("GET")

	// Signal protocol routes
	router.HandleFunc("/api/v1/signal/establish-channel", api.EstablishSignalChannel).Methods("POST")
	router.HandleFunc("/api/v1/signal/send", api.SendSignalMessage).Methods("POST")
}

// Authentication Handlers

func (api *MobileLiteNodeAPI) RequestChallenge(w http.ResponseWriter, r *http.Request) {
	var req ChallengeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	challenge, err := api.authService.GenerateChallenge(req.DeviceID, req.PublicKey, req.DeviceInfo)
	if err != nil {
		http.Error(w, "Failed to generate challenge", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challenge)
}

func (api *MobileLiteNodeAPI) VerifyChallenge(w http.ResponseWriter, r *http.Request) {
	var req VerificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session, err := api.authService.VerifyChallenge(req.Challenge, req.Signature, req.SessionID)
	if err != nil {
		http.Error(w, "Challenge verification failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

// Content Handlers

func (api *MobileLiteNodeAPI) GetFeed(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 || limit > 100 {
		limit = 20 // Default limit
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	maxSize, _ := strconv.Atoi(r.URL.Query().Get("max_size"))
	if maxSize == 0 {
		maxSize = 1024 * 1024 // 1MB default
	}

	since, _ := strconv.ParseInt(r.URL.Query().Get("since"), 10, 64)
	contentTypes := r.URL.Query().Get("content_types")
	topics := r.URL.Query().Get("topics")

	feedRequest := FeedRequest{
		Limit:        limit,
		Offset:       offset,
		MaxSize:      maxSize,
		Since:        since,
		ContentTypes: contentTypes,
		Topics:       topics,
	}

	feed, err := api.contentService.GetPersonalizedFeed(feedRequest)
	if err != nil {
		http.Error(w, "Failed to get feed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

func (api *MobileLiteNodeAPI) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	post, err := api.contentService.CreatePost(req)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (api *MobileLiteNodeAPI) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, err := api.contentService.GetPost(postID)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Node Management Handlers

func (api *MobileLiteNodeAPI) GetNodeStatus(w http.ResponseWriter, r *http.Request) {
	status := api.nodeService.GetNodeStatus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (api *MobileLiteNodeAPI) CreateResourceOffer(w http.ResponseWriter, r *http.Request) {
	var req ResourceOfferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	offer, err := api.nodeService.CreateResourceOffer(req)
	if err != nil {
		http.Error(w, "Failed to create resource offer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(offer)
}

// Synchronization Handlers

func (api *MobileLiteNodeAPI) SelectiveSync(w http.ResponseWriter, r *http.Request) {
	var req SelectiveSyncRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	syncJob, err := api.syncService.StartSelectiveSync(req)
	if err != nil {
		http.Error(w, "Failed to start sync", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(syncJob)
}

func (api *MobileLiteNodeAPI) GetSyncStatus(w http.ResponseWriter, r *http.Request) {
	status := api.syncService.GetSyncStatus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// Signal Protocol Handlers

func (api *MobileLiteNodeAPI) EstablishSignalChannel(w http.ResponseWriter, r *http.Request) {
	var req EstablishChannelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	channel, err := api.signalService.EstablishChannel(req.TargetNode, req.Purpose, req.AutoRotateKeys)
	if err != nil {
		http.Error(w, "Failed to establish channel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(channel)
}

func (api *MobileLiteNodeAPI) SendSignalMessage(w http.ResponseWriter, r *http.Request) {
	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := api.signalService.SendMessage(req.ChannelID, req.MessageType, req.Payload)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Request/Response Types

type ChallengeRequest struct {
	DeviceID   string     `json:"device_id"`
	PublicKey  string     `json:"public_key"`
	DeviceInfo DeviceInfo `json:"device_info"`
}

type DeviceInfo struct {
	Platform     string   `json:"platform"`
	Version      string   `json:"version"`
	Capabilities []string `json:"capabilities"`
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
	ExpiresAt int64  `json:"expires_at"`
	SessionID string `json:"session_id"`
}

type VerificationRequest struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
	SessionID string `json:"session_id"`
}

type FeedRequest struct {
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	MaxSize      int    `json:"max_size"`
	Since        int64  `json:"since"`
	ContentTypes string `json:"content_types"`
	Topics       string `json:"topics"`
}

type CreatePostRequest struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	MediaData       string `json:"media_data,omitempty"`
	MediaType       string `json:"media_type,omitempty"`
	Sources         []string `json:"sources"`
	Intent          string `json:"intent"`
	ContextType     string `json:"context_type"`
	OfflineCreated  bool   `json:"offline_created"`
	LocalTimestamp  int64  `json:"local_timestamp"`
}

type ResourceOfferRequest struct {
	StorageGB     float64              `json:"storage_gb"`
	BandwidthMbps float64              `json:"bandwidth_mbps"`
	DurationHours int                  `json:"duration_hours"`
	PricePerHour  int64                `json:"price_per_hour"`
	Conditions    ResourceConditions   `json:"conditions"`
}

type ResourceConditions struct {
	WiFiOnly         bool     `json:"wifi_only"`
	ChargingRequired bool     `json:"charging_required"`
	IdleHours        []string `json:"idle_hours"`
}

type SelectiveSyncRequest struct {
	SyncProfile        string             `json:"sync_profile"`
	ContentFilters     ContentFilters     `json:"content_filters"`
	NetworkConstraints NetworkConstraints `json:"network_constraints"`
}

type ContentFilters struct {
	Topics              []string `json:"topics"`
	Authors             []string `json:"authors"`
	EngagementThreshold int      `json:"engagement_threshold"`
	MaxAgeDays          int      `json:"max_age_days"`
}

type NetworkConstraints struct {
	WiFiOnly      bool `json:"wifi_only"`
	MaxBandwidth  int  `json:"max_bandwidth"`
	MaxDuration   int  `json:"max_duration"`
}

type EstablishChannelRequest struct {
	TargetNode      string `json:"target_node"`
	Purpose         string `json:"purpose"`
	AutoRotateKeys  bool   `json:"auto_rotate_keys"`
}

type SendMessageRequest struct {
	ChannelID   string      `json:"channel_id"`
	MessageType string      `json:"message_type"`
	Payload     interface{} `json:"payload"`
}