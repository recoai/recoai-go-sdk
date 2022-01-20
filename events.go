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
//    categoryPageView, err := UnmarshalCategoryPageView(bytes)
//    bytes, err = categoryPageView.Marshal()
//
//    checkoutStart, err := UnmarshalCheckoutStart(bytes)
//    bytes, err = checkoutStart.Marshal()
//
//    common, err := UnmarshalCommon(bytes)
//    bytes, err = common.Marshal()
//
//    detailProductView, err := UnmarshalDetailProductView(bytes)
//    bytes, err = detailProductView.Marshal()
//
//    homePageView, err := UnmarshalHomePageView(bytes)
//    bytes, err = homePageView.Marshal()
//
//    imageInteraction, err := UnmarshalImageInteraction(bytes)
//    bytes, err = imageInteraction.Marshal()
//
//    listView, err := UnmarshalListView(bytes)
//    bytes, err = listView.Marshal()
//
//    otherInteraction, err := UnmarshalOtherInteraction(bytes)
//    bytes, err = otherInteraction.Marshal()
//
//    pageVisit, err := UnmarshalPageVisit(bytes)
//    bytes, err = pageVisit.Marshal()
//
//    purchaseComplete, err := UnmarshalPurchaseComplete(bytes)
//    bytes, err = purchaseComplete.Marshal()
//
//    rateProduct, err := UnmarshalRateProduct(bytes)
//    bytes, err = rateProduct.Marshal()
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
//    removeItem, err := UnmarshalRemoveItem(bytes)
//    bytes, err = removeItem.Marshal()
//
//    cartPageView, err := UnmarshalCartPageView(bytes)
//    bytes, err = cartPageView.Marshal()
//
//    sortItems, err := UnmarshalSortItems(bytes)
//    bytes, err = sortItems.Marshal()
//
//    unknownEvent, err := UnmarshalUnknownEvent(bytes)
//    bytes, err = unknownEvent.Marshal()
//
//    upsertItem, err := UnmarshalUpsertItem(bytes)
//    bytes, err = upsertItem.Marshal()

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

func UnmarshalCategoryPageView(data []byte) (CategoryPageView, error) {
	var r CategoryPageView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CategoryPageView) Marshal() ([]byte, error) {
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

func UnmarshalDetailProductView(data []byte) (DetailProductView, error) {
	var r DetailProductView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DetailProductView) Marshal() ([]byte, error) {
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

func UnmarshalListView(data []byte) (ListView, error) {
	var r ListView
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListView) Marshal() ([]byte, error) {
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

func UnmarshalPurchaseComplete(data []byte) (PurchaseComplete, error) {
	var r PurchaseComplete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PurchaseComplete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRateProduct(data []byte) (RateProduct, error) {
	var r RateProduct
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RateProduct) Marshal() ([]byte, error) {
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

func UnmarshalRemoveItem(data []byte) (RemoveItem, error) {
	var r RemoveItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemoveItem) Marshal() ([]byte, error) {
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

func UnmarshalSortItems(data []byte) (SortItems, error) {
	var r SortItems
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SortItems) Marshal() ([]byte, error) {
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

func UnmarshalUpsertItem(data []byte) (UpsertItem, error) {
	var r UpsertItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertItem) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AddToCart struct {
	CartID      *string       `json:"cart_id"`     
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	RecID       *string       `json:"rec_id"`      
	UserInfo    UserInfo      `json:"user_info"`   
}

type EventDetail struct {
	EventAttributes map[string]string `json:"event_attributes"`
	ExperimentIDS   *int64            `json:"experiment_ids"`  
	RecID           *string           `json:"rec_id"`          
	URL             string            `json:"url"`             
}

type ItemDetails struct {
	CurrencyCode  *Currency `json:"currency_code"` 
	DisplayPrice  *float64  `json:"display_price"` 
	ItemID        string    `json:"item_id"`       
	OriginalPrice *float64  `json:"original_price"`
	Quantity      *int64    `json:"quantity"`      
}

type UserInfo struct {
	AdditionalInfo *UserAdditionalInfo `json:"additional_info"`
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
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	ListID      string        `json:"list_id"`     
	UserInfo    UserInfo      `json:"user_info"`   
}

type APISettings struct {
	URLAPI string `json:"url_api"`
}

type CategoryPageView struct {
	EventDetail    *EventDetail  `json:"event_detail"`   
	EventTime      int64         `json:"event_time"`     
	EventType      EventType     `json:"event_type"`     
	Items          []ItemDetails `json:"items"`          
	OnScreen       bool          `json:"on_screen"`      
	PageCategories []string      `json:"page_categories"`
	UserInfo       UserInfo      `json:"user_info"`      
}

type CheckoutStart struct {
	CartID              *string             `json:"cart_id"`             
	EventDetail         *EventDetail        `json:"event_detail"`        
	EventTime           int64               `json:"event_time"`          
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

type Costs struct {
	Cost          *float64 `json:"cost"`         
	Manufacturing *float64 `json:"manufacturing"`
}

type Taxes struct {
	Local *float64 `json:"local"`
	State *float64 `json:"state"`
}

type DetailProductView struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	RecID       *string      `json:"rec_id"`      
	UserInfo    UserInfo     `json:"user_info"`   
}

type HomePageView struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type ImageInteraction struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	UserInfo    UserInfo     `json:"user_info"`   
}

type ListView struct {
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	ListID      string        `json:"list_id"`     
	UserInfo    UserInfo      `json:"user_info"`   
}

type OtherInteraction struct {
	EventDetail     *EventDetail `json:"event_detail"`    
	EventTime       int64        `json:"event_time"`      
	EventType       EventType    `json:"event_type"`      
	InteractionName string       `json:"interaction_name"`
	Item            ItemDetails  `json:"item"`            
	UserInfo        UserInfo     `json:"user_info"`       
}

type PageVisit struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type PurchaseComplete struct {
	CartID              *string             `json:"cart_id"`             
	EventDetail         *EventDetail        `json:"event_detail"`        
	EventTime           int64               `json:"event_time"`          
	EventType           EventType           `json:"event_type"`          
	Items               []ItemDetails       `json:"items"`               
	PurchaseTransaction PurchaseTransaction `json:"purchase_transaction"`
	UserInfo            UserInfo            `json:"user_info"`           
}

type RateProduct struct {
	Comment     *string      `json:"comment"`     
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	Item        ItemDetails  `json:"item"`        
	Rating      *float64     `json:"rating"`      
	UserInfo    UserInfo     `json:"user_info"`   
}

type RecoRequest struct {
	Content       *string      `json:"content"`       
	EventDetail   *EventDetail `json:"event_detail"`  
	EventTime     int64        `json:"event_time"`    
	EventType     EventType    `json:"event_type"`    
	Location      *Location    `json:"location"`      
	NItems        int64        `json:"n_items"`       
	PlacementName *string      `json:"placement_name"`
	UserInfo      UserInfo     `json:"user_info"`     
}

type LocationClass struct {
	ProductPage  *string     `json:"ProductPage,omitempty"` 
	AddToCart    *string     `json:"AddToCart,omitempty"`   
	CategoryPage *string     `json:"CategoryPage,omitempty"`
	SearchPage   *SearchInfo `json:"SearchPage,omitempty"`  
	OtherPage    *PageInfo   `json:"OtherPage,omitempty"`   
	UnknownPage  *PageInfo   `json:"UnknownPage,omitempty"` 
}

type PageInfo struct {
	Content *string `json:"content"`
	URL     *string `json:"url"`    
}

type SearchInfo struct {
	Query *string `json:"query"`
}

type RecoShow struct {
	EventDetail   *EventDetail             `json:"event_detail"`  
	EventTime     int64                    `json:"event_time"`    
	EventType     EventType                `json:"event_type"`    
	ExperimentID  *string                  `json:"experiment_id"` 
	Items         []ProductDetailsRecoShow `json:"items"`         
	Location      *Location                `json:"location"`      
	PlacementName *string                  `json:"placement_name"`
	RecID         string                   `json:"rec_id"`        
	UserInfo      UserInfo                 `json:"user_info"`     
}

type ProductDetailsRecoShow struct {
	CanonicalProductURI *string    `json:"canonical_product_uri"`
	CurrencyCode        Currency   `json:"currency_code"`        
	ExactPrice          ExactPrice `json:"exact_price"`          
	ID                  string     `json:"id"`                   
	Images              []Image    `json:"images"`               
	Title               string     `json:"title"`                
}

type ExactPrice struct {
	DisplayPrice  float64 `json:"display_price"` 
	OriginalPrice float64 `json:"original_price"`
}

type Image struct {
	Height string `json:"height"`
	URI    string `json:"uri"`   
	Width  string `json:"width"` 
}

type RemoveFromCart struct {
	CartID      *string       `json:"cart_id"`     
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	UserInfo    UserInfo      `json:"user_info"`   
}

type RemoveFromList struct {
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	ListID      *string       `json:"list_id"`     
	UserInfo    UserInfo      `json:"user_info"`   
}

type RemoveItem struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	ItemID      string       `json:"item_id"`     
	UserInfo    UserInfo     `json:"user_info"`   
}

type CartPageView struct {
	CartID      *string       `json:"cart_id"`     
	EventDetail *EventDetail  `json:"event_detail"`
	EventTime   int64         `json:"event_time"`  
	EventType   EventType     `json:"event_type"`  
	Items       []ItemDetails `json:"items"`       
	UserInfo    UserInfo      `json:"user_info"`   
}

type SortItems struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	SortOrder   *SortOrder   `json:"sort_order"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type UnknownEvent struct {
	EventDetail *EventDetail `json:"event_detail"`
	EventTime   int64        `json:"event_time"`  
	EventType   EventType    `json:"event_type"`  
	UserInfo    UserInfo     `json:"user_info"`   
}

type UpsertItem struct {
	EventDetail    *EventDetail   `json:"event_detail"`   
	EventTime      int64          `json:"event_time"`     
	EventType      EventType      `json:"event_type"`     
	ProductDetails ProductDetails `json:"product_details"`
	UserInfo       UserInfo       `json:"user_info"`      
}

type ProductDetails struct {
	Attributes            map[string]string  `json:"attributes"`            
	AvailableQuantity     *int64             `json:"available_quantity"`    
	CanonicalProductURI   *string            `json:"canonical_product_uri"` 
	CategoricalAttributes map[string]string  `json:"categorical_attributes"`
	Categories            [][]string         `json:"categories"`            
	Costs                 map[string]float64 `json:"costs"`                 
	CurrencyCode          Currency           `json:"currency_code"`         
	Description           *string            `json:"description"`           
	ExactPrice            ExactPrice         `json:"exact_price"`           
	ID                    string             `json:"id"`                    
	Images                []Image            `json:"images"`                
	ItemGroupID           *string            `json:"item_group_id"`         
	LanguageCode          *string            `json:"language_code"`         
	NumericAttributes     map[string]float64 `json:"numeric_attributes"`    
	StockState            StockState         `json:"stock_state"`           
	Tags                  []string           `json:"tags"`                  
	Title                 string             `json:"title"`                 
}

type EventType string
const (
	EventTypeAddToCart EventType = "AddToCart"
	EventTypeAddToList EventType = "AddToList"
	EventTypeCartPageView EventType = "CartPageView"
	EventTypeCategoryPageView EventType = "CategoryPageView"
	EventTypeCheckoutStart EventType = "CheckoutStart"
	EventTypeDetailProductView EventType = "DetailProductView"
	EventTypeHomePageView EventType = "HomePageView"
	EventTypeImageInteraction EventType = "ImageInteraction"
	EventTypeListView EventType = "ListView"
	EventTypeOtherInteraction EventType = "OtherInteraction"
	EventTypePageVisit EventType = "PageVisit"
	EventTypePurchaseComplete EventType = "PurchaseComplete"
	EventTypeRateProduct EventType = "RateProduct"
	EventTypeRecoRequest EventType = "RecoRequest"
	EventTypeRecoShow EventType = "RecoShow"
	EventTypeRemoveFromCart EventType = "RemoveFromCart"
	EventTypeRemoveFromList EventType = "RemoveFromList"
	EventTypeRemoveItem EventType = "RemoveItem"
	EventTypeSortItems EventType = "SortItems"
	EventTypeUnknownEvent EventType = "UnknownEvent"
	EventTypeUpsertItem EventType = "UpsertItem"
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

type Gender string
const (
	Female Gender = "Female"
	Male Gender = "Male"
	Other Gender = "Other"
)

type LocationEnum string
const (
	CheckoutPage LocationEnum = "CheckoutPage"
	Error404 LocationEnum = "Error404"
	HomePage LocationEnum = "HomePage"
)

type SortOrder string
const (
	NameAsc SortOrder = "NameAsc"
	NameDesc SortOrder = "NameDesc"
	PopularityAsc SortOrder = "PopularityAsc"
	PopularityDesc SortOrder = "PopularityDesc"
	PriceAsc SortOrder = "PriceAsc"
	PriceDesc SortOrder = "PriceDesc"
	RatingAsc SortOrder = "RatingAsc"
	RatingDesc SortOrder = "RatingDesc"
)

type StockState string
const (
	BackOrder StockState = "BackOrder"
	InStock StockState = "InStock"
	OutOfStock StockState = "OutOfStock"
	PreOrder StockState = "PreOrder"
)

type Location struct {
	Enum          *LocationEnum
	LocationClass *LocationClass
}

func (x *Location) UnmarshalJSON(data []byte) error {
	x.LocationClass = nil
	x.Enum = nil
	var c LocationClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.LocationClass = &c
	}
	return nil
}

func (x *Location) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.LocationClass != nil, x.LocationClass, false, nil, x.Enum != nil, x.Enum, false)
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
