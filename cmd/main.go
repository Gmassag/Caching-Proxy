package main

import (
	"caching-proxy/internal"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	redisClient := internal.NewRedisClient()

	port := os.Getenv("PORT")
	origin := os.Getenv("ORIGIN")

	// Debug: print environment variables
	log.Printf("PORT: %s", port)
	log.Printf("ORIGIN: %s", origin)
	log.Printf("REDIS_URL: %s", os.Getenv("REDIS_URL"))

	if origin == "" {
		log.Fatal("Devi specificare ORIGIN environment variable")
	}

	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cacheKey := r.URL.String()

		val, err := redisClient.Get(internal.Ctx, cacheKey).Result()
		if err == nil {
			log.Printf("CACHE HIT %s", cacheKey)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(val))
			return
		}

		url := origin + r.URL.String()
		log.Printf("CACHE MISS %s - fetching from %s", cacheKey, url)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Errore fetching origin: %v", err)
			http.Error(w, "Errore fetching origin", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Errore reading response body: %v", err)
			http.Error(w, "Errore reading response", http.StatusInternalServerError)
			return
		}

		err = redisClient.Set(internal.Ctx, cacheKey, body, 60*time.Second).Err()
		if err != nil {
			log.Printf("Errore salvataggio cache: %v", err)
		}

		for k, v := range resp.Header {
			for _, vv := range v {
				w.Header().Add(k, vv)
			}
		}
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})

	log.Printf("Server avviato su :%s, origin: %s", port, origin)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
