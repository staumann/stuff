package ui

import (
	"github.com/staumann/caluclation/model"
	"log"
	"net/http"
	"strconv"
)

func HandleShowShop(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("shops.html")
	id := r.URL.Query().Get("id")
	start := int64(0)
	idParsed, e := strconv.ParseInt(id, 10, 64)
	if e == nil {
		log.Printf("setting page to %d", idParsed)
		start = idParsed
	}
	data := map[string]interface{}{
		"Shops": shopRepository.GetShops(start),
	}
	err := t.Execute(w, data)

	if err != nil {
		log.Printf("error showing all shops: %v", err)
		handleError(w, r, err)
	}
}

func HandleNewShop(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("shops_new.html")

	err := t.Execute(w, nil)

	if err != nil {
		log.Printf("error showing create new shop form, %v", err)
		handleError(w, r, err)
	}
}

func HandleCreateShop(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil {
		log.Printf("error parsing form: %v", e)
		handleError(w, r, e)
		return
	}
	shop := &model.Shop{
		Name:        r.Form.Get("name"),
		Street:      r.Form.Get("street"),
		HouseNumber: r.Form.Get("houseNumber"),
		City:        r.Form.Get("city"),
		PostalCode:  r.Form.Get("postCode"),
		Infos:       r.Form.Get("infos"),
	}

	if shop.Name == "" {
		t := templates.Lookup("shops_new.html")
		w.WriteHeader(http.StatusOK)
		if e := t.Execute(w, map[string]interface{}{
			"Name":        shop.Name,
			"Street":      shop.Street,
			"HouseNumber": shop.HouseNumber,
			"City":        shop.City,
			"PostalCode":  shop.PostalCode,
			"Infos":       shop.Infos,
			"Errors": []string{
				"Name was not defined",
			},
		}); e != nil {
			log.Printf("error exectuing template with error: %v", e)
			handleError(w, r, e)
		}
		return
	}

	if err := shopRepository.SaveShop(shop); err != nil {
		log.Printf("error saving shop %v", err)
		handleError(w, r, err)
	}
}
