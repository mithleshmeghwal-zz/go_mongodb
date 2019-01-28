package album

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	albums := c.Repository.GetAlbums()
	log.Println(albums)
	data, _ := json.Marshal(albums)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origins", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) AddAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Fatalln("Error AddAlbum", err)
		respondWithError(w, http.StatusInternalServerError, "Body Size Error")
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddAlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddAlbum unmarshalling data", err)
			respondWithError(w, http.StatusInternalServerError, "json encoding error")
			return
		}
		log.Fatalln("Error AddAlbum unmarshalling data", err)
		respondWithError(w, http.StatusUnprocessableEntity, "Params Error")
		return
	}

	success := c.Repository.AddAlbum(album)
	if !success {
		respondWithError(w, http.StatusUnprocessableEntity, "Params Error")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respondWithJSON(w, http.StatusCreated, "created")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
