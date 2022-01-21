// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    addToCart, err := UnmarshalAddToCart(bytes)
//    bytes, err = addToCart.Marshal()
//
//    addToList, err := UnmarshalAddToList(bytes)
//    bytes, err = addToList.Marshal()
//
//    aPISettings, err := UnmarshalAPISettings(bytes)
//    bytes, err = aPISettings.Marshal()
//
//    builderFn1, err := UnmarshalBuilderFn1(bytes)
//    bytes, err = builderFn1.Marshal()
//
//    builderFn2, err := UnmarshalBuilderFn2(bytes)
//    bytes, err = builderFn2.Marshal()
//
//    builderVariable, err := UnmarshalBuilderVariable(bytes)
//    bytes, err = builderVariable.Marshal()
//
//    changeItemStockState, err := UnmarshalChangeItemStockState(bytes)
//    bytes, err = changeItemStockState.Marshal()
//
//    checkoutStart, err := UnmarshalCheckoutStart(bytes)
//    bytes, err = checkoutStart.Marshal()
//
//    common, err := UnmarshalCommon(bytes)
//    bytes, err = common.Marshal()
//
//    detailItemView, err := UnmarshalDetailItemView(bytes)
//    bytes, err = detailItemView.Marshal()
//
//    homePageView, err := UnmarshalHomePageView(bytes)
//    bytes, err = homePageView.Marshal()
//
//    imageInteraction, err := UnmarshalImageInteraction(bytes)
//    bytes, err = imageInteraction.Marshal()
//
//    itemAttributesSelection, err := UnmarshalItemAttributesSelection(bytes)
//    bytes, err = itemAttributesSelection.Marshal()
//
//    itemRemove, err := UnmarshalItemRemove(bytes)
//    bytes, err = itemRemove.Marshal()
//
//    itemsView, err := UnmarshalItemsView(bytes)
//    bytes, err = itemsView.Marshal()
//
//    itemUpsert, err := UnmarshalItemUpsert(bytes)
//    bytes, err = itemUpsert.Marshal()
//
//    listView, err := UnmarshalListView(bytes)
//    bytes, err = listView.Marshal()
//
//    offlineRecommendationsRemove, err := UnmarshalOfflineRecommendationsRemove(bytes)
//    bytes, err = offlineRecommendationsRemove.Marshal()
//
//    offlineRecommendationsUpsert, err := UnmarshalOfflineRecommendationsUpsert(bytes)
//    bytes, err = offlineRecommendationsUpsert.Marshal()
//
//    otherInteraction, err := UnmarshalOtherInteraction(bytes)
//    bytes, err = otherInteraction.Marshal()
//
//    pageVisit, err := UnmarshalPageVisit(bytes)
//    bytes, err = pageVisit.Marshal()
//
//    placementRemove, err := UnmarshalPlacementRemove(bytes)
//    bytes, err = placementRemove.Marshal()
//
//    placementStatisticsJSONReady, err := UnmarshalPlacementStatisticsJSONReady(bytes)
//    bytes, err = placementStatisticsJSONReady.Marshal()
//
//    placementUpsert, err := UnmarshalPlacementUpsert(bytes)
//    bytes, err = placementUpsert.Marshal()
//
//    purchaseComplete, err := UnmarshalPurchaseComplete(bytes)
//    bytes, err = purchaseComplete.Marshal()
//
//    rateItem, err := UnmarshalRateItem(bytes)
//    bytes, err = rateItem.Marshal()
//
//    recoACK, err := UnmarshalRecoACK(bytes)
//    bytes, err = recoACK.Marshal()
//
//    recoRequest, err := UnmarshalRecoRequest(bytes)
//    bytes, err = recoRequest.Marshal()
//
//    recoShow, err := UnmarshalRecoShow(bytes)
//    bytes, err = recoShow.Marshal()
//
//    removeFromCart, err := UnmarshalRemoveFromCart(bytes)
//    bytes, err = removeFromCart.Marshal()
//
//    removeFromList, err := UnmarshalRemoveFromList(bytes)
//    bytes, err = removeFromList.Marshal()
//
//    searchItems, err := UnmarshalSearchItems(bytes)
//    bytes, err = searchItems.Marshal()
//
//    cartPageView, err := UnmarshalCartPageView(bytes)
//    bytes, err = cartPageView.Marshal()
//
//    smartSearchRequest, err := UnmarshalSmartSearchRequest(bytes)
//    bytes, err = smartSearchRequest.Marshal()
//
//    smartSearchShow, err := UnmarshalSmartSearchShow(bytes)
//    bytes, err = smartSearchShow.Marshal()
//
//    sortItems, err := UnmarshalSortItems(bytes)
//    bytes, err = sortItems.Marshal()
//
//    strategyParametersTypes, err := UnmarshalStrategyParametersTypes(bytes)
//    bytes, err = strategyParametersTypes.Marshal()
//
//    unknownEvent, err := UnmarshalUnknownEvent(bytes)
//    bytes, err = unknownEvent.Marshal()
//
//    videoInteraction, err := UnmarshalVideoInteraction(bytes)
//    bytes, err = videoInteraction.Marshal()

package main

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalAddToCart(data []byte) (AddToCart, error) {
	var r AddToCart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddToCart) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAddToList(data []byte) (AddToList, error) {
	var r AddToList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddToList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAPISettings(data []byte) (APISettings, error) {
	var r APISettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *APISettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBuilderFn1(data []byte) (BuilderFn1, error) {
	var r BuilderFn1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuilderFn1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBuilderFn2(data []byte) (BuilderFn2, error) {
	var r BuilderFn2
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuilderFn2) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBuilderVariable(data []byte) (BuilderVariable, error) {
	var r BuilderVariable
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuilderVariable) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChangeItemStockState(data []byte) (ChangeItemStockState, error) {
	var r ChangeItemStockState
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChangeItemStockState) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCheckoutStart(data []byte) (CheckoutStart, error) {
	var r CheckoutStart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CheckoutStart) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Common interface{}

func UnmarshalCommon(data []byte) (Common, error) {
	var r Common
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Common) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDetailItemView(data []byte) (DetailItemView, error) {
	var r DetailItemView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DetailItemView) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalHomePageView(data []byte) (HomePageView, error) {
	var r HomePageView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *HomePageView) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalImageInteraction(data []byte) (ImageInteraction, error) {
	var r ImageInteraction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ImageInteraction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalItemAttributesSelection(data []byte) (ItemAttributesSelection, error) {
	var r ItemAttributesSelection
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ItemAttributesSelection) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalItemRemove(data []byte) (ItemRemove, error) {
	var r ItemRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ItemRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalItemsView(data []byte) (ItemsView, error) {
	var r ItemsView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ItemsView) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalItemUpsert(data []byte) (ItemUpsert, error) {
	var r ItemUpsert
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ItemUpsert) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListView(data []byte) (ListView, error) {
	var r ListView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListView) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOfflineRecommendationsRemove(data []byte) (OfflineRecommendationsRemove, error) {
	var r OfflineRecommendationsRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OfflineRecommendationsRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOfflineRecommendationsUpsert(data []byte) (OfflineRecommendationsUpsert, error) {
	var r OfflineRecommendationsUpsert
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OfflineRecommendationsUpsert) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOtherInteraction(data []byte) (OtherInteraction, error) {
	var r OtherInteraction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OtherInteraction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPageVisit(data []byte) (PageVisit, error) {
	var r PageVisit
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PageVisit) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlacementRemove(data []byte) (PlacementRemove, error) {
	var r PlacementRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlacementRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlacementStatisticsJSONReady(data []byte) (PlacementStatisticsJSONReady, error) {
	var r PlacementStatisticsJSONReady
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlacementStatisticsJSONReady) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlacementUpsert(data []byte) (PlacementUpsert, error) {
	var r PlacementUpsert
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlacementUpsert) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPurchaseComplete(data []byte) (PurchaseComplete, error) {
	var r PurchaseComplete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PurchaseComplete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRateItem(data []byte) (RateItem, error) {
	var r RateItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RateItem) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRecoACK(data []byte) (RecoACK, error) {
	var r RecoACK
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RecoACK) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRecoRequest(data []byte) (RecoRequest, error) {
	var r RecoRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RecoRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRecoShow(data []byte) (RecoShow, error) {
	var r RecoShow
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RecoShow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRemoveFromCart(data []byte) (RemoveFromCart, error) {
	var r RemoveFromCart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemoveFromCart) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRemoveFromList(data []byte) (RemoveFromList, error) {
	var r RemoveFromList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemoveFromList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchItems(data []byte) (SearchItems, error) {
	var r SearchItems
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchItems) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCartPageView(data []byte) (CartPageView, error) {
	var r CartPageView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CartPageView) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSmartSearchRequest(data []byte) (SmartSearchRequest, error) {
	var r SmartSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SmartSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSmartSearchShow(data []byte) (SmartSearchShow, error) {
	var r SmartSearchShow
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SmartSearchShow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSortItems(data []byte) (SortItems, error) {
	var r SortItems
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SortItems) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStrategyParametersTypes(data []byte) (StrategyParametersTypes, error) {
	var r StrategyParametersTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StrategyParametersTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUnknownEvent(data []byte) (UnknownEvent, error) {
	var r UnknownEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UnknownEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalVideoInteraction(data []byte) (VideoInteraction, error) {
	var r VideoInteraction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VideoInteraction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AddToCart struct {
	CartID      *string      `json:"cart_id"`     
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

type EventDetail struct {
	EventAttributes map[string]string `json:"event_attributes"`
	ExperimentIDS   *int64            `json:"experiment_ids"`  
	RecID           *string           `json:"rec_id"`          
	URL             string            `json:"url"`             
}

type ItemDetails struct {
	Attributes *Attributes `json:"attributes"`
	Item       Item        `json:"item"`      
}

// This attribute structure is inspired by ECS dense components also know as table-based
// component list.
type Attributes struct {
	Article      *Article           `json:"article"`      
	Categories   *Categories        `json:"categories"`   
	Costs        *Costs             `json:"costs"`        
	Description  *Description       `json:"description"`  
	Ecommerce    *ItemEcommerceSpec `json:"ecommerce"`    
	Images       *Images            `json:"images"`       
	Price        *ExactPrice        `json:"price"`        
	RelatedItems []Item             `json:"related_items"`
	Stock        *Stock             `json:"stock"`        
	Tags         *Tags              `json:"tags"`         
	URL          ItemURL            `json:"url"`          
	Video        *Video             `json:"video"`        
}

type Article struct {
	Author         string `json:"author"`         
	Snippet        string `json:"snippet"`        
	TimestampAdded int64  `json:"timestamp_added"`
}

type Categories struct {
	Categories [][]TextContent `json:"categories"`
}

type Costs struct {
	Costs map[string]float64 `json:"costs"`
}

type Description struct {
	Content      *ContentUnion `json:"content"`      
	LanguageCode *string       `json:"language_code"`
	Title        *TextContent  `json:"title"`        
}

type ItemEcommerceSpec struct {
	ItemCode    *string `json:"item_code"`    
	ItemGroupID *string `json:"item_group_id"`
}

type Images struct {
	Images []Image `json:"images"`
}

type Image struct {
	Height string `json:"height"`
	URI    string `json:"uri"`   
	Width  string `json:"width"` 
}

type ExactPrice struct {
	CurrencyCode  *Currency `json:"currency_code"` 
	DisplayPrice  float64   `json:"display_price"` 
	OriginalPrice float64   `json:"original_price"`
}

type Item struct {
	ItemID   string   `json:"item_id"`  
	ItemType ItemType `json:"item_type"`
}

type Stock struct {
	AvailableQuantity *int64         `json:"available_quantity"`
	Quantity          *int64         `json:"quantity"`          
	StockState        StockStateEnum `json:"stock_state"`       
}

type Tags struct {
	Tags []string `json:"tags"`
}

type ItemURL struct {
	CanonicalURI           string  `json:"canonical_uri"`            
	CanonicalURIWithParams *string `json:"canonical_uri_with_params"`
	URLParams              *string `json:"url_params"`               
}

type Video struct {
	DurationSecs int64   `json:"duration_secs"`
	URI          *string `json:"uri"`          
}

type UserInfo struct {
	AdditionalInfo *UserAdditionalInfo `json:"additional_info"`
	PrivacySetting *PrivacySetting     `json:"privacy_setting"`
	SessionID      *string             `json:"session_id"`     
	UserID         *string             `json:"user_id"`        
	VisitorID      string              `json:"visitor_id"`     
}

type UserAdditionalInfo struct {
	BirthYear *int64  `json:"birth_year"`
	Gender    *Gender `json:"gender"`    
	IP        *string `json:"ip"`        
	Location  *string `json:"location"`  
	SessionID *string `json:"session_id"`
	UserAgent *string `json:"user_agent"`
}

type AddToList struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	ListID      string       `json:"list_id"`     
	UserInfo    UserInfo     `json:"user_info"`   
}

type APISettings struct {
	URLAPI string `json:"url_api"`
}

type ChangeItemStockState struct {
	EventDetail *EventDetail   `json:"event_detail"`
	EventTime   *int64         `json:"event_time"`  
	EventType   EventType      `json:"event_type"`  
	Item        Item           `json:"item"`        
	StockState  StockStateEnum `json:"stock_state"` 
	UserInfo    UserInfo       `json:"user_info"`   
}

type CheckoutStart struct {
	CartID              *string             `json:"cart_id"`             
	EventDetail         *EventDetail        `json:"event_detail"`        
	EventTime           *int64              `json:"event_time"`          
	EventType           EventType           `json:"event_type"`          
	Items               []ItemDetails       `json:"items"`               
	PurchaseTransaction PurchaseTransaction `json:"purchase_transaction"`
	UserInfo            UserInfo            `json:"user_info"`           
}

type PurchaseTransaction struct {
	Costs        *Costs   `json:"costs"`        
	CurrencyCode Currency `json:"currency_code"`
	ID           *string  `json:"id"`           
	Revenue      float64  `json:"revenue"`      
	Taxes        *Taxes   `json:"taxes"`        
}

type Taxes struct {
	Local *float64 `json:"local"`
	State *float64 `json:"state"`
}

type DetailItemView struct {
	EventDetail    *EventDetail              `json:"event_detail"`   
	EventTime      *int64                    `json:"event_time"`     
	EventType      EventType                 `json:"event_type"`     
	Item           ItemDetails               `json:"item"`           
	RecID          *string                   `json:"rec_id"`         
	UserInfo       UserInfo                  `json:"user_info"`      
	ViewAttributes *DetailItemViewAttributes `json:"view_attributes"`
}

type DetailItemViewAttributes struct {
	Completed    *bool  `json:"completed"`     // Completed can mean for - video - whether someone watched the whole video - article -; whether someone read the whole article
	ViewTimeSecs *int64 `json:"view_time_secs"`// This attribute can be used for different item types. For - Video - means watch time -; Article - reading time - Ecommerce - time spend on viewing the product
}

type HomePageView struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type ImageInteraction struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

type ItemAttributesSelectionClass struct {
	SelectedAttributes []ItemAttributesFieldNames `json:"SelectedAttributes"`
}

type ItemAttributesFieldNames struct {
	URL           *ItemURLFieldName           `json:"Url,omitempty"`          
	Price         *ExactPriceFieldName        `json:"Price,omitempty"`        
	Description   *DescriptionFieldName       `json:"Description,omitempty"`  
	Categories    *CategoriesFieldName        `json:"Categories,omitempty"`   
	Images        *ImagesFieldName            `json:"Images,omitempty"`       
	Video         *VideoFieldName             `json:"Video,omitempty"`        
	Tags          *TagsFieldName              `json:"Tags,omitempty"`         
	Article       *ArticleFieldName           `json:"Article,omitempty"`      
	EcommerceSpec *ItemEcommerceSpecFieldName `json:"EcommerceSpec,omitempty"`
	Stock         *StockFieldName             `json:"Stock,omitempty"`        
	Costs         *CostsFieldName             `json:"Costs,omitempty"`        
}

type ItemRemove struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        Item         `json:"item"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

type ItemsView struct {
	EventDetail    *EventDetail  `json:"event_detail"`   
	EventTime      *int64        `json:"event_time"`     
	EventType      EventType     `json:"event_type"`     
	Items          []ItemDetails `json:"items"`          
	OnScreen       bool          `json:"on_screen"`      
	PageCategories []string      `json:"page_categories"`
	UserInfo       UserInfo      `json:"user_info"`      
}

type ItemUpsert struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	ItemDetails ItemDetails  `json:"item_details"`
	UserInfo    UserInfo     `json:"user_info"`   
}

type ListView struct {
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   *int64        `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	ListID      string        `json:"list_id"`     
	UserInfo    UserInfo      `json:"user_info"`   
}

type OfflineRecommendationsRemove struct {
	EventDetail *EventDetail                `json:"event_detail"`
	EventTime   *int64                      `json:"event_time"`  
	EventType   EventType                   `json:"event_type"`  
	Name        *OfflineRecommendationsType `json:"name"`        
	UserInfo    UserInfo                    `json:"user_info"`   
}

type OfflineRecommendationsTypeClass struct {
	OtherSimilarity string `json:"OtherSimilarity"`
}

type OfflineRecommendationsUpsert struct {
	EventDetail *EventDetail                `json:"event_detail"`
	EventTime   *int64                      `json:"event_time"`  
	EventType   EventType                   `json:"event_type"`  
	Matrix      map[string]map[string]int64 `json:"matrix"`      
	Name        *OfflineRecommendationsType `json:"name"`        
	UserInfo    UserInfo                    `json:"user_info"`   
}

type OtherInteraction struct {
	EventDetail     *EventDetail `json:"event_detail"`    
	EventTime       *int64       `json:"event_time"`      
	EventType       EventType    `json:"event_type"`      
	InteractionName string       `json:"interaction_name"`
	Item            ItemDetails  `json:"item"`            
	UserInfo        UserInfo     `json:"user_info"`       
}

type PageVisit struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type PlacementRemove struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Name        string       `json:"name"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

// Json safe version of PlacementStatistics, unfortunately serde_with doesn't work because
// of conflicts with JsonSchema
type PlacementStatisticsJSONReady struct {
	LoadingTimesMicroseconds map[string]LIFOVecForUint128              `json:"loading_times_microseconds"`
	PlacementsStatistics     map[string][][]PlacementsStatisticElement `json:"placements_statistics"`     
}

type LIFOVecForUint128 struct {
	Base     []int64 `json:"base"`    
	Capacity int64   `json:"capacity"`
}

// Parametrized strategies
//
// Pre-defined strategies
//
// Build your custom strategies
type SuccessTriesClass struct {
	Parametrized    *ParametrizedStrategy `json:"Parametrized,omitempty"`   
	Generic         *GenericStrategy      `json:"Generic,omitempty"`        
	StrategyBuilder *string               `json:"StrategyBuilder,omitempty"`
	NImpressions    *int64                `json:"n_impressions,omitempty"`  
	NSuccess        *int64                `json:"n_success,omitempty"`      
}

// Symbolic reference to accumulators instances in the context
type ParametrizedStrategy struct {
	VisitorItemCounter            *EventVisitorItemCounterStaticParams       `json:"VisitorItemCounter,omitempty"`           
	ItemCounter                   *EventTypeItemCounterStaticParams          `json:"ItemCounter,omitempty"`                  
	ItemCooccurences              *EventItemTypeItemCooccurenceStaticParams  `json:"ItemCooccurences,omitempty"`             
	ContentItemMatcher            *ContentItemMatcherStaticParams            `json:"ContentItemMatcher,omitempty"`           
	OfflineRecommendationsStorage *OfflineRecommendationsStorageStaticParams `json:"OfflineRecommendationsStorage,omitempty"`
	SessionBasedCooccurence       *SessionItemsCooccurenceStaticParams       `json:"SessionBasedCooccurence,omitempty"`      
}

type ContentItemMatcherStaticParams struct {
	ItemType *ItemType `json:"item_type"`
}

type EventItemTypeItemCooccurenceStaticParams struct {
	EventTypeA EventType `json:"event_type_a"`
	EventTypeB EventType `json:"event_type_b"`
	ItemTypeA  ItemType  `json:"item_type_a"` 
	ItemTypeB  ItemType  `json:"item_type_b"` 
}

type EventTypeItemCounterStaticParams struct {
	EventType EventType `json:"event_type"`
	ItemType  ItemType  `json:"item_type"` 
}

type OfflineRecommendationsStorageStaticParams struct {
	RecType *OfflineRecommendationsType `json:"rec_type"`
}

// These are strategies that are just using other accumulators to generate candidates. They
// don't have any internal state
//
// The use case is to generate candidates based on the recent user history and event type
// and item type coocurrences
type SessionItemsCooccurenceStaticParams struct {
	Cooccurence   EventItemTypeItemCooccurenceStaticParams `json:"cooccurence"`   
	ItemGenerator EventVisitorItemCounterStaticParams      `json:"item_generator"`
}

type EventVisitorItemCounterStaticParams struct {
	EventType    EventType    `json:"event_type"`    
	ItemType     ItemType     `json:"item_type"`     
	UserInfoType UserInfoType `json:"user_info_type"`
}

type PlacementUpsert struct {
	Enabled      *bool                             `json:"enabled,omitempty"`
	EventDetail  *EventDetail                      `json:"event_detail"`     
	EventTime    *int64                            `json:"event_time"`       
	EventType    EventType                         `json:"event_type"`       
	HTMLTemplate *string                           `json:"html_template"`    
	ItemType     ItemType                          `json:"item_type"`        
	Location     *Location                         `json:"location"`         
	Name         string                            `json:"name"`             // Lowercase no spaces allowed
	Ranking      StrategySelectorStrategyChooseOne `json:"ranking"`          // How the strategies are selected
	Strategies   []WeightedGenericCandidateRec     `json:"strategies"`       
	URLParams    map[string]string                 `json:"url_params"`       
	UserInfo     UserInfo                          `json:"user_info"`        
}

type WeightedGenericCandidateRec struct {
	Strategy *Strategy `json:"strategy"`
	Weight   *float64  `json:"weight"`  
}

// Parametrized strategies
//
// Pre-defined strategies
//
// Build your custom strategies
type StrategyDynamicStrategies struct {
	Parametrized    *ParametrizedStrategy `json:"Parametrized,omitempty"`   
	Generic         *GenericStrategy      `json:"Generic,omitempty"`        
	StrategyBuilder *string               `json:"StrategyBuilder,omitempty"`
}

type PurchaseComplete struct {
	CartID              *string             `json:"cart_id"`             
	EventDetail         *EventDetail        `json:"event_detail"`        
	EventTime           *int64              `json:"event_time"`          
	EventType           EventType           `json:"event_type"`          
	Items               []ItemDetails       `json:"items"`               
	PurchaseTransaction PurchaseTransaction `json:"purchase_transaction"`
	UserInfo            UserInfo            `json:"user_info"`           
}

type RateItem struct {
	Comment     *string      `json:"comment"`     
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	Rating      *float64     `json:"rating"`      
	UserInfo    UserInfo     `json:"user_info"`   
}

type RecoACK struct {
	EventDetail     *EventDetail          `json:"event_detail"`    
	EventTime       *int64                `json:"event_time"`      
	EventType       EventType             `json:"event_type"`      
	Items           []ItemDetailsRecoShow `json:"items"`           
	PlacementConfig PlacementConfig       `json:"placement_config"`
	UserInfo        UserInfo              `json:"user_info"`       
}

// Note that ItemDetailsRecoShow is already translated
type ItemDetailsRecoShow struct {
	Attributes       Attributes                `json:"attributes"`       
	Item             Item                      `json:"item"`             
	RecID            string                    `json:"rec_id"`           
	Score            *float64                  `json:"score"`            
	StrategiesUsed   [][]StrategiesUsedElement `json:"strategies_used"`  
	StrategySelected *Strategy                 `json:"strategy_selected"`
}

// Parametrized strategies
//
// Pre-defined strategies
//
// Build your custom strategies
type StrategiesUsedDynamicStrategies struct {
	Parametrized    *ParametrizedStrategy `json:"Parametrized,omitempty"`   
	Generic         *GenericStrategy      `json:"Generic,omitempty"`        
	StrategyBuilder *string               `json:"StrategyBuilder,omitempty"`
}

type PlacementConfig struct {
	HTMLTemplate            *string                       `json:"html_template"`            
	ItemAttributesSelection *ItemAttributesSelectionUnion `json:"item_attributes_selection"`
	ItemType                ItemType                      `json:"item_type"`                
	Location                *Location                     `json:"location"`                 
	Name                    string                        `json:"name"`                     
	Parameters              *StrategyParametersValues     `json:"parameters"`               
}

type ItemAttributesSelectionItemAttributesSelectionClass struct {
	SelectedAttributes []ItemAttributesFieldNames `json:"SelectedAttributes"`
}

type StrategyParametersValues struct {
	AdditionalIntParams map[string]int64  `json:"additional_int_params"`
	AdditionalStrParams map[string]string `json:"additional_str_params"`
	CategoryID          []string          `json:"category_id"`          
	CategoryIDS         [][]string        `json:"category_ids"`         
	Item                *Item             `json:"item"`                 
	Items               []Item            `json:"items"`                
	PageInfo            *PageInfo         `json:"page_info"`            
	SearchInfo          *SearchInfo       `json:"search_info"`          
}

type PageInfo struct {
	Content  *string   `json:"content"`  // Set page content
	ItemType *ItemType `json:"item_type"`// Item type
	URL      *string   `json:"url"`      // Set page URL
}

type SearchInfo struct {
	Query *string `json:"query"`// Set the search query
}

type RecoRequest struct {
	AdditionalURIParams map[string]string `json:"additional_uri_params"`
	EventDetail         *EventDetail      `json:"event_detail"`         
	EventTime           *int64            `json:"event_time"`           
	EventType           EventType         `json:"event_type"`           
	Locale              *Locale           `json:"locale"`               
	Location            Location          `json:"location"`             
	NItems              int64             `json:"n_items"`              
	PlacementConfig     PlacementConfig   `json:"placement_config"`     
	UserInfo            UserInfo          `json:"user_info"`            
}

type RecoShow struct {
	AdditionalURIParams map[string]string     `json:"additional_uri_params"`
	EventDetail         *EventDetail          `json:"event_detail"`         
	EventTime           *int64                `json:"event_time"`           
	EventType           EventType             `json:"event_type"`           
	ExperimentID        *string               `json:"experiment_id"`        
	Items               []ItemDetailsRecoShow `json:"items"`                
	PlacementConfig     PlacementConfig       `json:"placement_config"`     
	UserInfo            UserInfo              `json:"user_info"`            
}

type RemoveFromCart struct {
	CartID      *string      `json:"cart_id"`     
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

type RemoveFromList struct {
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   *int64        `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	ListID      *string       `json:"list_id"`     
	UserInfo    UserInfo      `json:"user_info"`   
}

type SearchItems struct {
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   *int64        `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	Query       string        `json:"query"`       
	UserInfo    UserInfo      `json:"user_info"`   
}

type CartPageView struct {
	CartID      *string       `json:"cart_id"`     
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   *int64        `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	UserInfo    UserInfo      `json:"user_info"`   
}

type SmartSearchRequest struct {
	EventDetail *EventDetail      `json:"event_detail"`
	EventTime   *int64            `json:"event_time"`  
	EventType   EventType         `json:"event_type"`  
	Filter      map[string]string `json:"filter"`      
	NItems      int64             `json:"n_items"`     
	Page        int64             `json:"page"`        
	Query       string            `json:"query"`       
	SearchOrder SearchOrder       `json:"search_order"`
	UserInfo    UserInfo          `json:"user_info"`   
}

type SmartSearchShow struct {
	EventDetail *EventDetail          `json:"event_detail"`
	EventTime   *int64                `json:"event_time"`  
	EventType   EventType             `json:"event_type"`  
	Items       []ItemDetailsRecoShow `json:"items"`       
	UserInfo    UserInfo              `json:"user_info"`   
}

type SortItems struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	SortOrder   *SortOrder   `json:"sort_order"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type StrategyParametersTypes struct {
	AdditionalIntParams []string `json:"additional_int_params"`
	AdditionalStrParams []string `json:"additional_str_params"`
	CategoryID          bool     `json:"category_id"`          
	CategoryIDS         bool     `json:"category_ids"`         
	Item                bool     `json:"item"`                 
	ItemsInfo           bool     `json:"items_info"`           
	PageInfo            bool     `json:"page_info"`            
	SearchInfo          bool     `json:"search_info"`          
}

type UnknownEvent struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   *int64       `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type VideoInteraction struct {
	Completed   *bool         `json:"completed"`   
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   *int64        `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	UserInfo    UserInfo      `json:"user_info"`   
	VideoItem   *Item         `json:"video_item"`  
	WatchedSecs *int64        `json:"watched_secs"`
}

type EventType string
const (
	EventTypeAddToCart EventType = "AddToCart"
	EventTypeAddToList EventType = "AddToList"
	EventTypeCartPageView EventType = "CartPageView"
	EventTypeChangeItemStockState EventType = "ChangeItemStockState"
	EventTypeCheckoutStart EventType = "CheckoutStart"
	EventTypeDetailItemView EventType = "DetailItemView"
	EventTypeHomePageView EventType = "HomePageView"
	EventTypeImageInteraction EventType = "ImageInteraction"
	EventTypeItemRemove EventType = "ItemRemove"
	EventTypeItemUpsert EventType = "ItemUpsert"
	EventTypeItemsView EventType = "ItemsView"
	EventTypeListView EventType = "ListView"
	EventTypeOfflineRecommendationsRemove EventType = "OfflineRecommendationsRemove"
	EventTypeOfflineRecommendationsUpsert EventType = "OfflineRecommendationsUpsert"
	EventTypeOtherInteraction EventType = "OtherInteraction"
	EventTypePageVisit EventType = "PageVisit"
	EventTypePlacementRemove EventType = "PlacementRemove"
	EventTypePlacementUpsert EventType = "PlacementUpsert"
	EventTypePurchaseComplete EventType = "PurchaseComplete"
	EventTypeRateItem EventType = "RateItem"
	EventTypeRecoACK EventType = "RecoAck"
	EventTypeRecoRequest EventType = "RecoRequest"
	EventTypeRecoShow EventType = "RecoShow"
	EventTypeRemoveFromCart EventType = "RemoveFromCart"
	EventTypeRemoveFromList EventType = "RemoveFromList"
	EventTypeSearchItems EventType = "SearchItems"
	EventTypeSmartSearchRequest EventType = "SmartSearchRequest"
	EventTypeSmartSearchShow EventType = "SmartSearchShow"
	EventTypeSortItems EventType = "SortItems"
	EventTypeUnknownEvent EventType = "UnknownEvent"
	EventTypeVideoInteraction EventType = "VideoInteraction"
)

type Currency string
const (
	AMD Currency = "AMD"
	Aed Currency = "AED"
	Afn Currency = "AFN"
	All Currency = "ALL"
	Ang Currency = "ANG"
	Aoa Currency = "AOA"
	Ars Currency = "ARS"
	Aud Currency = "AUD"
	Awg Currency = "AWG"
	Azn Currency = "AZN"
	BAM Currency = "BAM"
	BSD Currency = "BSD"
	Bbd Currency = "BBD"
	Bdt Currency = "BDT"
	Bgn Currency = "BGN"
	Bhd Currency = "BHD"
	Bif Currency = "BIF"
	Bmd Currency = "BMD"
	Bnd Currency = "BND"
	Bob Currency = "BOB"
	Bov Currency = "BOV"
	Brl Currency = "BRL"
	Btn Currency = "BTN"
	Bwp Currency = "BWP"
	Byn Currency = "BYN"
	Bzd Currency = "BZD"
	CAD Currency = "CAD"
	CRC Currency = "CRC"
	Cdf Currency = "CDF"
	Che Currency = "CHE"
	Chf Currency = "CHF"
	Chw Currency = "CHW"
	Clf Currency = "CLF"
	Clp Currency = "CLP"
	Cny Currency = "CNY"
	Cop Currency = "COP"
	Cou Currency = "COU"
	Cuc Currency = "CUC"
	Cup Currency = "CUP"
	Cve Currency = "CVE"
	Czk Currency = "CZK"
	Djf Currency = "DJF"
	Dkk Currency = "DKK"
	Dop Currency = "DOP"
	Dzd Currency = "DZD"
	EGP Currency = "EGP"
	Ern Currency = "ERN"
	Etb Currency = "ETB"
	Eur Currency = "EUR"
	Fjd Currency = "FJD"
	Fkp Currency = "FKP"
	Gbp Currency = "GBP"
	Gel Currency = "GEL"
	Ghs Currency = "GHS"
	Gip Currency = "GIP"
	Gmd Currency = "GMD"
	Gnf Currency = "GNF"
	Gtq Currency = "GTQ"
	Gyd Currency = "GYD"
	Hkd Currency = "HKD"
	Hnl Currency = "HNL"
	Hrk Currency = "HRK"
	Htg Currency = "HTG"
	Huf Currency = "HUF"
	Idr Currency = "IDR"
	Ils Currency = "ILS"
	Inr Currency = "INR"
	Iqd Currency = "IQD"
	Irr Currency = "IRR"
	Isk Currency = "ISK"
	Jmd Currency = "JMD"
	Jod Currency = "JOD"
	Jpy Currency = "JPY"
	Kes Currency = "KES"
	Kgs Currency = "KGS"
	Khr Currency = "KHR"
	Kmf Currency = "KMF"
	Kpw Currency = "KPW"
	Krw Currency = "KRW"
	Kwd Currency = "KWD"
	Kyd Currency = "KYD"
	Kzt Currency = "KZT"
	Lak Currency = "LAK"
	Lbp Currency = "LBP"
	Lkr Currency = "LKR"
	Lrd Currency = "LRD"
	Lsl Currency = "LSL"
	Lyd Currency = "LYD"
	Mad Currency = "MAD"
	Mdl Currency = "MDL"
	Mga Currency = "MGA"
	Mkd Currency = "MKD"
	Mmk Currency = "MMK"
	Mnt Currency = "MNT"
	Mop Currency = "MOP"
	Mru Currency = "MRU"
	Mur Currency = "MUR"
	Mvr Currency = "MVR"
	Mwk Currency = "MWK"
	Mxn Currency = "MXN"
	Mxv Currency = "MXV"
	Myr Currency = "MYR"
	Mzn Currency = "MZN"
	NIO Currency = "NIO"
	Nad Currency = "NAD"
	Ngn Currency = "NGN"
	Nok Currency = "NOK"
	Npr Currency = "NPR"
	Nzd Currency = "NZD"
	OMR Currency = "OMR"
	PHP Currency = "PHP"
	Pab Currency = "PAB"
	Pen Currency = "PEN"
	Pgk Currency = "PGK"
	Pkr Currency = "PKR"
	Pln Currency = "PLN"
	Pyg Currency = "PYG"
	Qar Currency = "QAR"
	Ron Currency = "RON"
	Rsd Currency = "RSD"
	Rub Currency = "RUB"
	Rwf Currency = "RWF"
	SSP Currency = "SSP"
	SVC Currency = "SVC"
	Sar Currency = "SAR"
	Sbd Currency = "SBD"
	Scr Currency = "SCR"
	Sdg Currency = "SDG"
	Sek Currency = "SEK"
	Sgd Currency = "SGD"
	Shp Currency = "SHP"
	Sll Currency = "SLL"
	Sos Currency = "SOS"
	Srd Currency = "SRD"
	Stn Currency = "STN"
	Syp Currency = "SYP"
	Szl Currency = "SZL"
	Thb Currency = "THB"
	Tjs Currency = "TJS"
	Tmt Currency = "TMT"
	Tnd Currency = "TND"
	Top Currency = "TOP"
	Try Currency = "TRY"
	Ttd Currency = "TTD"
	Twd Currency = "TWD"
	Tzs Currency = "TZS"
	Uah Currency = "UAH"
	Ugx Currency = "UGX"
	Unk Currency = "UNK"
	Usd Currency = "USD"
	Usn Currency = "USN"
	Uyi Currency = "UYI"
	Uyu Currency = "UYU"
	Uzs Currency = "UZS"
	Vef Currency = "VEF"
	Vnd Currency = "VND"
	Vuv Currency = "VUV"
	Wst Currency = "WST"
	Xaf Currency = "XAF"
	Xcd Currency = "XCD"
	Xdr Currency = "XDR"
	Xof Currency = "XOF"
	Xpf Currency = "XPF"
	Xsu Currency = "XSU"
	Xua Currency = "XUA"
	Yer Currency = "YER"
	Zar Currency = "ZAR"
	Zmw Currency = "ZMW"
	Zwl Currency = "ZWL"
)

type ItemType string
const (
	Ecommerce ItemType = "Ecommerce"
	ItemTypeArticle ItemType = "Article"
	ItemTypeUnknown ItemType = "Unknown"
	ItemTypeVideo ItemType = "Video"
)

type StockStateEnum string
const (
	BackOrder StockStateEnum = "BackOrder"
	InStock StockStateEnum = "InStock"
	OutOfStock StockStateEnum = "OutOfStock"
	PreOrder StockStateEnum = "PreOrder"
)

type Gender string
const (
	Female Gender = "Female"
	Male Gender = "Male"
	Other Gender = "Other"
)

type PrivacySetting string
const (
	NonPersonalized PrivacySetting = "NonPersonalized"
	PrivacySettingPersonalized PrivacySetting = "Personalized"
)

type BuilderFn1 string
const (
	ArgMax BuilderFn1 = "ArgMax"
	ArgMin BuilderFn1 = "ArgMin"
	ConvertToPlacementItemType BuilderFn1 = "ConvertToPlacementItemType"
	Invert BuilderFn1 = "Invert"
	Ranking BuilderFn1 = "Ranking"
)

type BuilderFn2 string
const (
	Expand BuilderFn2 = "Expand"
	Highest BuilderFn2 = "Highest"
	Intersect BuilderFn2 = "Intersect"
	Less BuilderFn2 = "Less"
	LessEq BuilderFn2 = "LessEq"
	Lookup BuilderFn2 = "Lookup"
	Lowest BuilderFn2 = "Lowest"
	More BuilderFn2 = "More"
	MoreEq BuilderFn2 = "MoreEq"
	Remove BuilderFn2 = "Remove"
	Union BuilderFn2 = "Union"
)

type BuilderVariable string
const (
	AllItems BuilderVariable = "AllItems"
	ItemCurrent BuilderVariable = "ItemCurrent"
	ItemCurrentType BuilderVariable = "ItemCurrentType"
	ItemsAlsoAddedToCartInSession BuilderVariable = "ItemsAlsoAddedToCartInSession"
	ItemsAlsoBoughtInSession BuilderVariable = "ItemsAlsoBoughtInSession"
	ItemsAlsoSeenInSession BuilderVariable = "ItemsAlsoSeenInSession"
	ItemsInCart BuilderVariable = "ItemsInCart"
	ItemsRecommendedSessionCounter BuilderVariable = "ItemsRecommendedSessionCounter"
	ItemsSeenInSession BuilderVariable = "ItemsSeenInSession"
	ItemsVisitedCounter BuilderVariable = "ItemsVisitedCounter"
)

type ArticleFieldName string
const (
	Author ArticleFieldName = "Author"
	Snippet ArticleFieldName = "Snippet"
	TimestampAdded ArticleFieldName = "TimestampAdded"
)

type CategoriesFieldName string
const (
	CategoriesFieldNameCategories CategoriesFieldName = "Categories"
)

type CostsFieldName string
const (
	CostsFieldNameCosts CostsFieldName = "Costs"
)

type DescriptionFieldName string
const (
	Content DescriptionFieldName = "Content"
	LanguageCode DescriptionFieldName = "LanguageCode"
	Title DescriptionFieldName = "Title"
)

type ItemEcommerceSpecFieldName string
const (
	ItemCode ItemEcommerceSpecFieldName = "ItemCode"
	ItemGroupID ItemEcommerceSpecFieldName = "ItemGroupId"
)

type ImagesFieldName string
const (
	ImagesFieldNameImages ImagesFieldName = "Images"
)

type ExactPriceFieldName string
const (
	CurrencyCode ExactPriceFieldName = "CurrencyCode"
	DisplayPrice ExactPriceFieldName = "DisplayPrice"
	OriginalPrice ExactPriceFieldName = "OriginalPrice"
)

type StockFieldName string
const (
	AvailableQuantity StockFieldName = "AvailableQuantity"
	Quantity StockFieldName = "Quantity"
	StockState StockFieldName = "StockState"
)

type TagsFieldName string
const (
	TagsFieldNameTags TagsFieldName = "Tags"
)

type ItemURLFieldName string
const (
	CanonicalURI ItemURLFieldName = "CanonicalUri"
	CanonicalURIWithParams ItemURLFieldName = "CanonicalUriWithParams"
	URLParams ItemURLFieldName = "UrlParams"
)

type VideoFieldName string
const (
	DurationSecs VideoFieldName = "DurationSecs"
	URI VideoFieldName = "Uri"
)

type ItemAttributesSelectionEnum string
const (
	AllAttributes ItemAttributesSelectionEnum = "AllAttributes"
)

type OfflineRecommendationsTypeEnum string
const (
	AttributesSimilarity OfflineRecommendationsTypeEnum = "AttributesSimilarity"
	ImageSimilarity OfflineRecommendationsTypeEnum = "ImageSimilarity"
	TextSimilarity OfflineRecommendationsTypeEnum = "TextSimilarity"
)

type GenericStrategy string
const (
	BestsellerCategory GenericStrategy = "BestsellerCategory"
	SearchMatching GenericStrategy = "SearchMatching"
	SeenInSessionCoccurAddedToCart GenericStrategy = "SeenInSessionCoccurAddedToCart"
	SeenInSessionCoccurSeen GenericStrategy = "SeenInSessionCoccurSeen"
)

type UserInfoType string
const (
	Session UserInfoType = "SESSION"
	User UserInfoType = "USER"
	Visitor UserInfoType = "VISITOR"
)

type StrategyEnum string
const (
	StrategyUnknown StrategyEnum = "Unknown"
)

// Choose location
type Location string
const (
	CategoryPage Location = "CategoryPage"
	CheckoutPage Location = "CheckoutPage"
	Error404 Location = "Error404"
	HomePage Location = "HomePage"
	ItemPage Location = "ItemPage"
	LocationAddToCart Location = "AddToCart"
	OtherPage Location = "OtherPage"
	SearchPage Location = "SearchPage"
	UnknownPage Location = "UnknownPage"
)

// How the strategies are selected
type StrategySelectorStrategyChooseOne string
const (
	RankingModel StrategySelectorStrategyChooseOne = "RankingModel"
	ThompsonSampling StrategySelectorStrategyChooseOne = "ThompsonSampling"
	WeightedSample StrategySelectorStrategyChooseOne = "WeightedSample"
)

type Locale string
const (
	AFZA Locale = "af_ZA"
	ArAR Locale = "ar_AR"
	AsIN Locale = "as_IN"
	AzAZ Locale = "az_AZ"
	BeBY Locale = "be_BY"
	BgBG Locale = "bg_BG"
	BnIN Locale = "bn_IN"
	BrFR Locale = "br_FR"
	BsBA Locale = "bs_BA"
	CAES Locale = "ca_ES"
	CSCZ Locale = "cs_CZ"
	CbIQ Locale = "cb_IQ"
	CoFR Locale = "co_FR"
	CxPH Locale = "cx_PH"
	CyGB Locale = "cy_GB"
	DaDK Locale = "da_DK"
	DeDE Locale = "de_DE"
	ElGR Locale = "el_GR"
	EnGB Locale = "en_GB"
	EnUD Locale = "en_UD"
	EnUS Locale = "en_US"
	EsES Locale = "es_ES"
	EsLA Locale = "es_LA"
	EtEE Locale = "et_EE"
	EuES Locale = "eu_ES"
	FaIR Locale = "fa_IR"
	FfNG Locale = "ff_NG"
	FiFI Locale = "fi_FI"
	FoFO Locale = "fo_FO"
	FrCA Locale = "fr_CA"
	FrFR Locale = "fr_FR"
	FyNL Locale = "fy_NL"
	GaIE Locale = "ga_IE"
	GlES Locale = "gl_ES"
	GnPY Locale = "gn_PY"
	GuIN Locale = "gu_IN"
	HaNG Locale = "ha_NG"
	HeIL Locale = "he_IL"
	HiIN Locale = "hi_IN"
	HrHR Locale = "hr_HR"
	HuHU Locale = "hu_HU"
	HyAM Locale = "hy_AM"
	IDID Locale = "id_ID"
	IsIS Locale = "is_IS"
	ItIT Locale = "it_IT"
	JaJP Locale = "ja_JP"
	JaKS Locale = "ja_KS"
	JvID Locale = "jv_ID"
	KMKH Locale = "km_KH"
	KaGE Locale = "ka_GE"
	KkKZ Locale = "kk_KZ"
	KnIN Locale = "kn_IN"
	KoKR Locale = "ko_KR"
	KuTR Locale = "ku_TR"
	LVLV Locale = "lv_LV"
	LtLT Locale = "lt_LT"
	MSMY Locale = "ms_MY"
	MTMT Locale = "mt_MT"
	MgMG Locale = "mg_MG"
	MkMK Locale = "mk_MK"
	MlIN Locale = "ml_IN"
	MnMN Locale = "mn_MN"
	MrIN Locale = "mr_IN"
	MyMM Locale = "my_MM"
	NbNO Locale = "nb_NO"
	NeNP Locale = "ne_NP"
	NlBE Locale = "nl_BE"
	NlNL Locale = "nl_NL"
	NnNO Locale = "nn_NO"
	OrIN Locale = "or_IN"
	PSAF Locale = "ps_AF"
	PaIN Locale = "pa_IN"
	PlPL Locale = "pl_PL"
	PtBR Locale = "pt_BR"
	PtPT Locale = "pt_PT"
	QzMM Locale = "qz_MM"
	RoRO Locale = "ro_RO"
	RuRU Locale = "ru_RU"
	RwRW Locale = "rw_RW"
	ScIT Locale = "sc_IT"
	SiLK Locale = "si_LK"
	SkSK Locale = "sk_SK"
	SlSI Locale = "sl_SI"
	SoSO Locale = "so_SO"
	SqAL Locale = "sq_AL"
	SrRS Locale = "sr_RS"
	SvSE Locale = "sv_SE"
	SwKE Locale = "sw_KE"
	SzPL Locale = "sz_PL"
	TaIN Locale = "ta_IN"
	TeIN Locale = "te_IN"
	TgTJ Locale = "tg_TJ"
	ThTH Locale = "th_TH"
	TlPH Locale = "tl_PH"
	TrTR Locale = "tr_TR"
	TzMA Locale = "tz_MA"
	UkUA Locale = "uk_UA"
	UrPK Locale = "ur_PK"
	UzUZ Locale = "uz_UZ"
	ViVN Locale = "vi_VN"
	ZhCN Locale = "zh_CN"
	ZhHK Locale = "zh_HK"
	ZhTW Locale = "zh_TW"
)

type SearchOrder string
const (
	Newest SearchOrder = "Newest"
	Oldest SearchOrder = "Oldest"
	RelevanceAsc SearchOrder = "RelevanceAsc"
	RelevanceDesc SearchOrder = "RelevanceDesc"
	SearchOrderPersonalized SearchOrder = "Personalized"
	SearchOrderPopularityAsc SearchOrder = "PopularityAsc"
	SearchOrderPopularityDesc SearchOrder = "PopularityDesc"
	SearchOrderPriceAsc SearchOrder = "PriceAsc"
	SearchOrderPriceDesc SearchOrder = "PriceDesc"
	SearchOrderRatingAsc SearchOrder = "RatingAsc"
	SearchOrderRatingDesc SearchOrder = "RatingDesc"
)

type SortOrder string
const (
	NameAsc SortOrder = "NameAsc"
	NameDesc SortOrder = "NameDesc"
	SortOrderPopularityAsc SortOrder = "PopularityAsc"
	SortOrderPopularityDesc SortOrder = "PopularityDesc"
	SortOrderPriceAsc SortOrder = "PriceAsc"
	SortOrderPriceDesc SortOrder = "PriceDesc"
	SortOrderRatingAsc SortOrder = "RatingAsc"
	SortOrderRatingDesc SortOrder = "RatingDesc"
)

type TextContent struct {
	String    *string
	StringMap map[string]string
}

func (x *TextContent) UnmarshalJSON(data []byte) error {
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.StringMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *TextContent) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.StringMap != nil, x.StringMap, false, nil, false)
}

type ContentUnion struct {
	String    *string
	StringMap map[string]string
}

func (x *ContentUnion) UnmarshalJSON(data []byte) error {
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.StringMap, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ContentUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.StringMap != nil, x.StringMap, false, nil, true)
}

type ItemAttributesSelection struct {
	Enum                         *ItemAttributesSelectionEnum
	ItemAttributesSelectionClass *ItemAttributesSelectionClass
}

func (x *ItemAttributesSelection) UnmarshalJSON(data []byte) error {
	x.ItemAttributesSelectionClass = nil
	x.Enum = nil
	var c ItemAttributesSelectionClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.ItemAttributesSelectionClass = &c
	}
	return nil
}

func (x *ItemAttributesSelection) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.ItemAttributesSelectionClass != nil, x.ItemAttributesSelectionClass, false, nil, x.Enum != nil, x.Enum, false)
}

type OfflineRecommendationsType struct {
	Enum                            *OfflineRecommendationsTypeEnum
	OfflineRecommendationsTypeClass *OfflineRecommendationsTypeClass
}

func (x *OfflineRecommendationsType) UnmarshalJSON(data []byte) error {
	x.OfflineRecommendationsTypeClass = nil
	x.Enum = nil
	var c OfflineRecommendationsTypeClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.OfflineRecommendationsTypeClass = &c
	}
	return nil
}

func (x *OfflineRecommendationsType) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.OfflineRecommendationsTypeClass != nil, x.OfflineRecommendationsTypeClass, false, nil, x.Enum != nil, x.Enum, false)
}

type PlacementsStatisticElement struct {
	Enum              *StrategyEnum
	SuccessTriesClass *SuccessTriesClass
}

func (x *PlacementsStatisticElement) UnmarshalJSON(data []byte) error {
	x.SuccessTriesClass = nil
	x.Enum = nil
	var c SuccessTriesClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.SuccessTriesClass = &c
	}
	return nil
}

func (x *PlacementsStatisticElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.SuccessTriesClass != nil, x.SuccessTriesClass, false, nil, x.Enum != nil, x.Enum, false)
}

type Strategy struct {
	Enum                      *StrategyEnum
	StrategyDynamicStrategies *StrategyDynamicStrategies
}

func (x *Strategy) UnmarshalJSON(data []byte) error {
	x.StrategyDynamicStrategies = nil
	x.Enum = nil
	var c StrategyDynamicStrategies
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.StrategyDynamicStrategies = &c
	}
	return nil
}

func (x *Strategy) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.StrategyDynamicStrategies != nil, x.StrategyDynamicStrategies, false, nil, x.Enum != nil, x.Enum, false)
}

type StrategiesUsedElement struct {
	Double                          *float64
	Enum                            *StrategyEnum
	StrategiesUsedDynamicStrategies *StrategiesUsedDynamicStrategies
}

func (x *StrategiesUsedElement) UnmarshalJSON(data []byte) error {
	x.StrategiesUsedDynamicStrategies = nil
	x.Enum = nil
	var c StrategiesUsedDynamicStrategies
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.StrategiesUsedDynamicStrategies = &c
	}
	return nil
}

func (x *StrategiesUsedElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, false, nil, x.StrategiesUsedDynamicStrategies != nil, x.StrategiesUsedDynamicStrategies, false, nil, x.Enum != nil, x.Enum, false)
}

type ItemAttributesSelectionUnion struct {
	Enum                                                *ItemAttributesSelectionEnum
	ItemAttributesSelectionItemAttributesSelectionClass *ItemAttributesSelectionItemAttributesSelectionClass
}

func (x *ItemAttributesSelectionUnion) UnmarshalJSON(data []byte) error {
	x.ItemAttributesSelectionItemAttributesSelectionClass = nil
	x.Enum = nil
	var c ItemAttributesSelectionItemAttributesSelectionClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, true)
	if err != nil {
		return err
	}
	if object {
		x.ItemAttributesSelectionItemAttributesSelectionClass = &c
	}
	return nil
}

func (x *ItemAttributesSelectionUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.ItemAttributesSelectionItemAttributesSelectionClass != nil, x.ItemAttributesSelectionItemAttributesSelectionClass, false, nil, x.Enum != nil, x.Enum, true)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
